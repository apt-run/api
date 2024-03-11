package database

const CREATE_SOURCE_TABLE = `
	CREATE TABLE IF NOT EXISTS Sources (
		SourceID 	SERIAL,
		Name 		TEXT NOT NULL UNIQUE,
		URL 		TEXT NOT NULL UNIQUE,
		BaseURL 	TEXT NOT NULL UNIQUE,
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

const DROP_TABLE_SOURCES = `
	DROP TABLE Sourec;
`

const DROP_TABLE_PACKAGE = `
	DROP TABLE Package;
`

const INSERT_SOURCE = `
	INSERT INTO Users (Username, Password)
	VALUES ($1, $2); 
`

const SELECT_SOURCE = `
	SELECT * 
	FROM users
	WHERE username = $1
	ORDER BY username ASC
	LIMIT 1
`

const UPDATE_SOURCE = `
	UPDATE Users
	SET Username = $1, Password = $2
	WHERE Username = $3; 
`

const DELETE_SOURCE = `
	DELETE FROM Users 
	WHERE Username = $1;
`

const INSERT_PACKAGE = `
	INSERT INTO Users (Username, Password)
	VALUES ($1, $2); 
`

const SELECT_PACKAGE = `
	SELECT * 
	FROM users
	WHERE username = $1
	ORDER BY username ASC
	LIMIT 1
`

const UPDATE_PACKAGE = `
	UPDATE Users
	SET Username = $1, Password = $2
	WHERE Username = $3; 
`

const DELETE_PACKAGE = `
	DELETE FROM Users 
	WHERE Username = $1;
`
