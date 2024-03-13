SELECT jsonb_agg(to_jsonb(stats)) AS stats
FROM stats
WHERE name LIKE 'golang%';


SELECT jsonb_agg(to_jsonb(stats.name)) AS stats
FROM stats
WHERE name LIKE 'golang%';


SELECT jsonb_agg(to_jsonb(stats.name) || to_jsonb(stats.installs)) AS stats
FROM stats
WHERE name LIKE 'golang%';