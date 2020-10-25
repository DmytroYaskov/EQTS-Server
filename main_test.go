package main

import (
	"database/sql"
	"os"
	"testing"
)

func TestInsertRecord(t *testing.T) {

	file, err := os.Create("test.db")
	if err != nil {
		t.Error()
	}
	file.Close()

	sqliteDatabase, _ := sql.Open("sqlite3", "./test.db")
	defer sqliteDatabase.Close()

	initDatabase(sqliteDatabase)

	var testDeviceID int = 1
	var testTemperature float32 = 40.1

	insertRecord(sqliteDatabase, testDeviceID, testTemperature)

	testRecords := getRecords(sqliteDatabase)

	if testRecords[0].device != testDeviceID && testRecords[0].temperature != testTemperature {
		t.Error()
	}

	os.Remove("test.db")
}
