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
	"low_name" Text NOT NULL );

CREATE INDEX "authorIdx" ON "author"( "id" );
CREATE INDEX "authorIdx1" ON "author"( "name" );
CREATE INDEX "authorIdx2" ON "author"( "low_name" );

-- ------------------------------------------
-- Dump of "book"
-- ------------------------------------------

DROP TABLE IF EXISTS "book";

CREATE TABLE "book"(
	"author_id" Integer,
	"date" Date,
	"id" Integer NOT NULL PRIMARY KEY AUTOINCREMENT,
	"name" Text,
	"station_id" Integer,
	"url" Text,
	"low_name" Text,
	CONSTRAINT "link_author_book_2" FOREIGN KEY ( "author_id" ) REFERENCES "author"( "id" )
		ON DELETE Cascade
		ON UPDATE Cascade
		DEFERRABLE INITIALLY DEFERRED
,
	CONSTRAINT "link_station_book_3" FOREIGN KEY ( "station_id" ) REFERENCES "station"( "id" )
		ON DELETE Cascade
		ON UPDATE Cascade
,
CONSTRAINT "Unique_2" UNIQUE ( "name" ),
CONSTRAINT "unique_id" UNIQUE ( "id" ) );

CREATE INDEX "bookIdx" ON "book"( "name" );
CREATE INDEX "bookIdx1" ON "book"( "date" );
CREATE INDEX "bookIdx2" ON "book"( "id" );
CREATE INDEX "bookIdx3" ON "book"( "station_id" );
CREATE INDEX "bookIdx4" ON "book"( "author_id" );
CREATE INDEX "bookIdx5" ON "book"( "low_name" );

-- ------------------------------------------
-- Dump of "file"
-- ------------------------------------------

DROP TABLE IF EXISTS "file";

CREATE TABLE "file"(
	"book_id" Integer NOT NULL,
	"id" Integer NOT NULL PRIMARY KEY,
	"name" Text NOT NULL,
	"size" Integer NOT NULL,
	"url" Text NOT NULL,
	CONSTRAINT "link_book_file_0" FOREIGN KEY ( "book_id" ) REFERENCES "book"( "id" )
		ON DELETE Cascade
		ON UPDATE Cascade
		DEFERRABLE INITIALLY DEFERRED
,
CONSTRAINT "Unique_1" UNIQUE ( "id" ),
CONSTRAINT "Unique_2" UNIQUE ( "url" ) );

CREATE INDEX "fileIdx" ON "file"( "id", "book_id", "name", "size", "url" );

-- ------------------------------------------
-- Dump of "history"
-- ------------------------------------------

DROP TABLE IF EXISTS "history";

CREATE TABLE "history"(
	"id" Integer PRIMARY KEY AUTOINCREMENT,
	"book_id" Integer,
	"date" DateTime,
	CONSTRAINT "lnk_history_book" FOREIGN KEY ( "book_id" ) REFERENCES "book"( "id" )
		DEFERRABLE INITIALLY DEFERRED
,
CONSTRAINT "unique_id" UNIQUE ( "id" ) );

CREATE INDEX "index_book_id" ON "history"( "book_id" );
CREATE INDEX "index_data" ON "history"( "date" );
CREATE INDEX "index_id" ON "history"( "id" );

-- ------------------------------------------
-- Dump of "station"
-- ------------------------------------------

DROP TABLE IF EXISTS "station";

CREATE TABLE "station"(
	"id" Integer PRIMARY KEY,
	"name" Text,
CONSTRAINT "stationUnique" UNIQUE ( "name" ) );

CREATE INDEX "stationIdx" ON "station"( "id" );
CREATE INDEX "stationIdx1" ON "station"( "name" );


`

func ResetDb() (err error) {
	_, err = db.Exec(dump)
	return
}