package database

import (
	"database/sql"

	"log"

	_ "github.com/mattn/go-sqlite3" // Sqlite3 driver
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("sqlite", "crmlite.db")
	if err != nil {
		log.Fatal("Could not open/create database!")
		return
	}
}

func buildDb() error {
	query := `
    PRAGMA foreign_keys = ON;
    Create Table if not exists contact (
        id INTEGER PRIMARY KEY,
        name text null,
        phone text null,
        email text null,
        website text null,
        created_date text not null,
        updated_date text null,
        deleted_date text null
    );

    Create Table if not exists note (
        id INTEGER PRIMARY KEY,
        body text null,
        created_date text not null,
        updated_date text null,
        deleted_date text null
    );

    Create Table if not exists contact_note (
        contact_id int,
        note_id int,
        created_date text not null,
        updated_date text null,
        deleted_date text null,
        Foreign Key (contact_id) references conctact(id),
        Foreign Key (note_id) references Note(id)
    );
    `
	_, err := db.Exec(query)
	return err
}
