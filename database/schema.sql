-- Postgres 15 

-- Database Schema
---------------------------------------

DROP TABLE Source;
DROP TABLE Sources;

DROP TABLE Package;
DROP TABLE Packages;

---------------------------------------

CREATE TABLE IF NOT EXISTS Sources (
    UserID SERIAL PRIMARY KEY,
    Username TEXT NOT NULL UNIQUE,
    Password TEXT NOT NULL UNIQUE,
);

CREATE TABLE IF NOT EXISTS Sources (
    UserID SERIAL PRIMARY KEY,
    Username TEXT NOT NULL UNIQUE,
    Password TEXT NOT NULL UNIQUE,
);

CREATE TABLE IF NOT EXISTS Package (
    UserID SERIAL PRIMARY KEY,
    Username TEXT NOT NULL UNIQUE,
    Password TEXT NOT NULL UNIQUE,
); 

CREATE TABLE IF NOT EXISTS Packages (
    UserID SERIAL PRIMARY KEY,
    Username TEXT NOT NULL UNIQUE,
    Password TEXT NOT NULL UNIQUE,
); 


---------------------------------------

SELECT * FROM Package
SELECT * FROM Packages

SELECT * FROM Source
SELECT * FROM Sources

---------------------------------------

SELECT ? FROM Package
SELECT ? FROM Packages

SELECT ? FROM Source
SELECT ? FROM Sources

---------------------------------------

INSERT INTO Users (Username, Password)
VALUES (?, ?); 

INSERT INTO Users (Username, Password)
VALUES (?, ?); 

INSERT INTO Users (Username, Password)
VALUES (?, ?); 

INSERT INTO Users (Username, Password)
VALUES (?, ?); 

---------------------------------------

UPDATE Users
SET Username = ?, Password = ?
WHERE Username = ?; 

UPDATE Users
SET Username = ?, Password = ?
WHERE Username = ?; 

UPDATE Users
SET Username = ?, Password = ?
WHERE Username = ?; 

UPDATE Users
SET Username = ?, Password = ?
WHERE Username = ?; 


---------------------------------------

DELETE FROM Users 
WHERE Username = ?;

DELETE FROM Users 
WHERE Username = ?;

DELETE FROM Users 
WHERE Username = ?;

DELETE FROM Users 
WHERE Username = ?;

---------------------------------------

