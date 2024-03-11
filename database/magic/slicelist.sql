# gpt wizardry
# select x packages from list

# 1
SELECT json_build_object('packages', json_agg(value))
FROM (
  SELECT value
  FROM json_array_elements((SELECT list FROM sources) -> 'packages')
  LIMIT 20
) subquery;

# 3
SELECT json_build_object('packages', json_agg(value))
FROM (
  SELECT value
  FROM sources, json_array_elements(sources.list::json -> 'packages')
  LIMIT 20
) subquery;

# custom
SELECT json_build_object('packages', json_agg(value))
FROM (
  SELECT value
  FROM sources, json_array_elements(sources.list::json -> 'packages')
  LIMIT 20 OFFSET 20
) subquery;

# custom
SELECT json_build_object('packages', json_agg(value))
FROM (
  SELECT value
  FROM sources, json_array_elements(list::json -> 'packages') AS package(value)
  WHERE package.value->>'name' LIKE 'lib%'
  LIMIT 20
) subquery;