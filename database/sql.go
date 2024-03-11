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
	INSERT INTO Sources (name, list)
	VALUES ($1, $2)
	ON CONFLICT(name) 
	DO UPDATE SET
		list = EXCLUDED.list;
`
const UPDATE_SOURCES = `
	UPDATE Sources (Name, List)
	SET List = $2
	WHERE Name = $1;
`
const INSERT_SOURCES = `
	INSERT INTO Sources (Name, List) 
	VALUES($1, $2)
`
const DROP_TABLE_SOURCES = `
	DROP TABLE Sources;
`
const DROP_TABLE_PACKAGE = `
	DROP TABLE Package;
`
