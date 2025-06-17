## Solution notes

### Task 01 – Run‑Length Encoder

- Language: Go
- Approach: Using map string to keep tracking number of character and use order for order integrity that map cannot provided
- Why: a bit hard to read if does not know why use order instead of loop through map
- Time spent: ~10 min
- AI tools used: No

### Task 02 – Fix‑the‑Bug

- Language: Go
- Approach: ~~Output is only want to avoid the duplication so I decided to keep it simple for reduce time for other problem~~ After trying to not use mutex so I already got it that I misunderstand the problem
- Why: Should use realible library to generate the uuid instead if id not necessary to order and need to use mutex to lock the accessibility of worker
- Time spent: ~7 min

### Task 03 – Sync-Aggregator

- Language: Go
- Approach: Using Go routine for concurrency read the file and use wait group to wait for all the concurrent complete and send the result back
- Why: This case code will be complexity and low readability should refactor for easier to maintain
- Time spent: ~40 min

### Task 04 – SQL-Reasoning

- Language: Go
- Approach: Use concept of join to get interesting data to identify the record or make them aggregation
- Why: If this case happend much so we need to identify and looking for good performance of database and built-in function that support about data transformation
- Time spent: ~2 hr
