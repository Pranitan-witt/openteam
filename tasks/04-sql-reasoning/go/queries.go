// tasks/04‑sql‑reasoning/go/queries.go
package queries

// Task A
const SQLA = `
	select c.id as campaign_id, sum(p.amount_thb) as total_thb, round(sum(p.amount_thb) / cast(c.target_thb as REAL), 4) as pct_of_target  from campaign c
	inner join pledge p on c.id = p.campaign_id
	group by c.id
	order by pct_of_target desc, c.id;
`

// Task B
const SQLB = `
	WITH all_data_thailand AS (
		SELECT
			amount_thb
		FROM pledge p
		JOIN donor d ON p.donor_id = d.id
		where d.country = 'Thailand'
		order by amount_thb
	),
	all_data_global AS (
		SELECT
			amount_thb
		FROM pledge p
		JOIN donor d ON p.donor_id = d.id
		order by amount_thb
	),
	counts_global AS (
		SELECT 
			COUNT(*) AS cnt
		FROM all_data_global
	),
	counts_thailand AS (
		SELECT 
			COUNT(*) AS cnt
		FROM all_data_thailand
	),
	ranked_global AS (
		SELECT 
			a.amount_thb,
			ROW_NUMBER() OVER (ORDER BY amount_thb) AS rn
		FROM all_data_global a
	),
	ranked_thailand AS (
		SELECT 
			a.amount_thb,
			ROW_NUMBER() OVER (ORDER BY amount_thb) AS rn
		FROM all_data_thailand a
	),
	target_ranks_global AS (
		SELECT 
			CAST(ROUND(0.9 * cnt) AS INT) AS target_rank
		FROM counts_global
	),
	target_ranks_thailand AS (
		SELECT 
			CAST(ROUND(0.9 * cnt) AS INT) AS target_rank
		FROM counts_thailand
	)
	SELECT 
		'global' as scope,
		r.amount_thb AS p90_thb
	FROM ranked_global r
	JOIN target_ranks_global t 
	ON r.rn = t.target_rank
	union all
	SELECT 
		'thailand' as scope,
		r.amount_thb AS p90_thb
	FROM ranked_thailand r
	JOIN target_ranks_thailand t 
	ON r.rn = t.target_rank
	;
`

var Indexes = []string{`
	CREATE INDEX idx_pledge_campaign_id
	ON pledge (campaign_id)
`,
`
	CREATE INDEX idx_pledge_donor_id
	ON pledge (campaign_id);
`,

} // skipped
