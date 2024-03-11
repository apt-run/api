SELECT json_build_object('packages', json_agg(value))
FROM (
  SELECT value
  FROM sources, json_array_elements(list::json -> 'packages') AS j(value)
  WHERE j.value->>'name' LIKE 'lib%'
  LIMIT 20
) subquery;


SELECT json_build_object('packages', json_agg(value))
FROM (
  SELECT value
  FROM sources, json_array_elements(list::json -> 'packages') AS j(value)
  WHERE j.value->>'name' LIKE 'lib%'
) subquery;