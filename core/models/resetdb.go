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
	"date" Date,
	"station_id" Integer,
	CONSTRAINT "link_author_book_2" FOREIGN KEY ( "author_id" ) REFERENCES "author"( "id" )
		ON DELETE Cascade
		ON UPDATE Cascade
		DEFERRABLE INITIALLY DEFERRED
,
	CONSTRAINT "link_station_book_3" FOREIGN KEY ( "station_id" ) REFERENCES "station"( "id" )
		ON DELETE Cascade
		ON UPDATE Cascade
,
CONSTRAINT "Unique_1" UNIQUE ( "id", "name" ) );

CREATE INDEX "bookIdx" ON "book"( "date", "author_id", "name", "id", "station_id" );

-- ------------------------------------------
-- Dump of "file"
-- ------------------------------------------

DROP TABLE IF EXISTS "file";

CREATE TABLE "file"(
	"id" Integer NOT NULL PRIMARY KEY,
	"book_id" Integer NOT NULL,
	"name" Text NOT NULL,
	"url" Text NOT NULL,
	"size" Integer NOT NULL,
	CONSTRAINT "link_book_file_0" FOREIGN KEY ( "book_id" ) REFERENCES "book"( "id" )
		ON DELETE Cascade
		ON UPDATE Cascade
		DEFERRABLE INITIALLY DEFERRED
,
CONSTRAINT "Unique_1" UNIQUE ( "url", "id" ) );

CREATE INDEX "fileIdx" ON "file"( "id", "book_id", "name", "size", "url" );

-- ------------------------------------------
-- Dump of "station"
-- ------------------------------------------

DROP TABLE IF EXISTS "station";

CREATE TABLE "station"(
	"id" Integer PRIMARY KEY,
	"name" Text );

CREATE INDEX "stationIdx" ON "station"( "id", "name" );
`

func ResetDb() (err error) {
	_, err = db.Exec(dump)
	return
}