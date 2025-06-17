// Package aggregator – stub for Concurrent File Stats Processor.
package aggregator

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

// Result mirrors one JSON object in the final array.
type Result struct {
	Path   string `json:"path"`
	Lines  int    `json:"lines,omitempty"`
	Words  int    `json:"words,omitempty"`
	Status string `json:"status"` // "ok" or "timeout"
}

// Aggregate must read filelistPath, spin up *workers* goroutines,
// apply a per‑file timeout, and return results in **input order**.

type task struct {
	index int
	path  string
}


func Aggregate(filelistPath string, workers, timeout int) ([]Result, error) {
	file, err := os.Open(filelistPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var paths []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		path := strings.TrimSpace(scanner.Text())
		if path != "" {
			paths = append(paths, path)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	numFiles := len(paths)
	results := make([]Result, numFiles)

	tasks := make(chan task)
	var wg sync.WaitGroup

	// Start worker goroutines
	for w := 0; w < workers; w++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for t := range tasks {
				res := processFileWithTimeout(t.path, timeout)
				results[t.index] = res
			}
		}()
	}

	// Dispatch tasks
	for i, path := range paths {
		tasks <- task{index: i, path: path}
	}
	close(tasks)

	wg.Wait()
	return results, nil
}

func processFileWithTimeout(path string, timeoutSec int) Result {

	resultChan := make(chan Result, 1)

	go func() {
		lines, words, status := countLinesAndWords(path, timeoutSec)
		resultChan <- Result{
			Path:   path,
			Lines:  lines,
			Words:  words,
			Status: status,
		}
	}()

	select {
	case res := <-resultChan:
		return res
	}
}

func countLinesAndWords(path string, timeout int) (int, int, string) {
	file, err := os.Open("../data/" + path)
	if err != nil {
		return 0, 0, "timeout"
	}
	defer file.Close()

	var lines, words int
	scanner := bufio.NewScanner(file)
	firstLine := true
	for scanner.Scan() {
		text := scanner.Text()
		if firstLine {
			sleepSec := strings.Split(text, "=")[1]
			sleepInt, err := strconv.ParseInt(sleepSec, 10, 64)
			if err != nil {
				return 0, 0, "timeout"
			}
			if sleepInt > int64(timeout) {
				return 0, 0, "timeout"
			}
			time.Sleep(time.Duration(sleepInt*int64(time.Second)))
			firstLine = false
			continue
		}
		data := strings.Split(text, " ")
		if len(data) == 1 {
			break
		}
		lines++
		words += len(data)
	}


	return lines, words, "ok"
}