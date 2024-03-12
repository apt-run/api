package database

const CREATE_SOURCE_TABLE = `
	CREATE TABLE IF NOT EXISTS Sources (
		SourceID 	SERIAL,
		Name 		TEXT NOT NULL UNIQUE,
		List		JSON NOT NULL,
		PRIMARY 	KEY(SourceID)
	);
`
const CREATE_PACKAGE_TABLE = `
	CREATE TABLE IF NOT EXISTS Packages (
		PackageID 		SERIAL,
		Name 			TEXT NOT NULL UNIQUE,
		Description 	TEXT NOT NULL,
		Installs 	    TEXT NOT NULL,
		Versions 	    TEXT NOT NULL,
		URL 			TEXT NOT NULL UNIQUE,
		PRIMARY 		KEY(PackageID)
	);
`
const UPSERT_SOURCES = `
	INSERT INTO sources (name, list)
	VALUES ($1, $2)
	ON CONFLICT(name) 
	DO UPDATE SET
		list = EXCLUDED.list;
`
const DROP_TABLE_SOURCES = `
	DROP TABLE Sources;
`
const DROP_TABLE_PACKAGE = `
	DROP TABLE Package;
`

const SEARCH_PACKAGES = `
	SELECT json_build_object('packages', json_agg(value))
	FROM (
		SELECT value 
		FROM sources, json_array_elements(sources.list::json -> 'packages') AS package(value)
		WHERE package.value->>'name' LIKE ($1 || '%')
		LIMIT $2
	) subquery;
`

const SEARCH_PACKAGES_PAGINATE = `
	SELECT json_build_object('packages', json_agg(value))
	FROM (
		SELECT value 
		FROM sources, json_array_elements(list::json -> 'packages') AS package(value)
		WHERE package.value->>'name' LIKE ($1 || '%')
		LIMIT $1 OFFSET $1	
	) subquery;
`

const PAGINATE_PACKAGES_START = `
	SELECT json_build_object('packages', json_agg(value))
	FROM (
		SELECT value
		FROM sources, json_array_elements(sources.list::json -> 'packages')
		LIMIT 20
	) subquery;
`

const PAGINATE_PACKAGES = `
	SELECT json_build_object('packages', json_agg(value))
	FROM (
		SELECT value
		FROM sources, json_array_elements(sources.list::json -> 'packages')
		LIMIT $1 OFFSET $1
	) subquery;
`
