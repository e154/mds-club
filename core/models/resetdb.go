package models

import (
)

const dump = `
------------- SQLite3 Dump File -------------

-- ------------------------------------------
-- Dump of "author"
-- ------------------------------------------

DROP TABLE IF EXISTS "author";

CREATE TABLE "author"(
	"id" Integer NOT NULL PRIMARY KEY AUTOINCREMENT,
	"name" Text NOT NULL,
CONSTRAINT "Unique_1" UNIQUE ( "id" ),
CONSTRAINT "Unique_2" UNIQUE ( "name" ) );

CREATE INDEX "authorIdx" ON "author"( "id", "name" );

-- ------------------------------------------
-- Dump of "book"
-- ------------------------------------------

DROP TABLE IF EXISTS "book";

CREATE TABLE "book"(
	"id" Integer NOT NULL PRIMARY KEY,
	"author_id" Integer,
	"name" Text,
	"datetime" DateTime,
	"station_id" Integer,
	CONSTRAINT "lnk_book_station" FOREIGN KEY ( "station_id" ) REFERENCES "station"( "id" )
		ON DELETE Cascade
		ON UPDATE Cascade
,
	CONSTRAINT "lnk_book_author" FOREIGN KEY ( "author_id" ) REFERENCES "author"( "id" )
		ON DELETE Cascade
		ON UPDATE Cascade
		DEFERRABLE INITIALLY DEFERRED
,
CONSTRAINT "Unique_1" UNIQUE ( "id" ) );


-- ------------------------------------------
-- Dump of "station"
-- ------------------------------------------

DROP TABLE IF EXISTS "station";

CREATE TABLE "station"(
	"id" Integer PRIMARY KEY,
	"name" Text );


-- ------------------------------------------
-- Dump of "file"
-- ------------------------------------------

DROP TABLE IF EXISTS "file";

CREATE TABLE "file"(
	"id" Integer PRIMARY KEY,
	"book_id" Integer,
	"name" Text,
	"url" Text,
	"size" Integer,
	CONSTRAINT "lnk_file_book" FOREIGN KEY ( "book_id" ) REFERENCES "book"( "id" )
		ON DELETE Cascade
		ON UPDATE Cascade
		DEFERRABLE INITIALLY DEFERRED
 );

CREATE INDEX "fileIdx" ON "file"( "id", "book_id", "name", "size", "url" );


`

func ResetDb() (err error) {
	_, err = db.Exec(dump)
	return
}