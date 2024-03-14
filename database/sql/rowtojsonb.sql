SELECT jsonb_agg(to_jsonb(stats)) AS stats
FROM stats
WHERE name LIKE 'golang%';


SELECT jsonb_agg(to_jsonb(stats.name)) AS stats
FROM stats
WHERE name LIKE 'golang%';


SELECT jsonb_agg(to_jsonb(stats.name) || to_jsonb(stats.installs)) AS stats
FROM stats
WHERE name LIKE 'golang%';

-- array
SELECT jsonb_agg(results)
FROM (
    SELECT to_jsonb(stats.name) || 
    to_jsonb(stats.installs) ||
    to_jsonb(stats.maintainer) as results
    FROM stats
    LIMIT 20
)as subquery

-- object
SELECT json_build_object('results', 
    jsonb_agg(json_build_object(
        'name', subquery.name,
        'installs', subquery.installs,
        'maintainer', subquery.maintainer
))) AS results
FROM (
    SELECT 
        to_jsonb(stats.name) as name,
        to_jsonb(stats.installs) as installs,
        to_jsonb(stats.maintainer) as maintainer
    FROM stats
    LIMIT 20
) AS subquery;


SELECT json_build_object('results', 
    jsonb_agg(json_build_object(
        'name', subquery.name,
        'installs', subquery.installs,
        'maintainer', subquery.maintainer
))) AS results
FROM (
    SELECT 
        to_jsonb(stats.name) as name,
        to_jsonb(stats.installs) as installs,
        to_jsonb(stats.maintainer) as maintainer
    FROM stats
    WHERE name LIKE '%golang%'
    LIMIT 20
) AS subquery;