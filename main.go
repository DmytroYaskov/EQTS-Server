package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type record struct {
	id          int
	device      int
	temperature float32
}

func main() {
	log.Println("App started")
	os.Remove("app.db")

	log.Println("Creating database")
	file, err := os.Create("app.db")
	if err != nil {
		log.Fatal("Error during file creating")
		log.Fatal(err.Error())
	}
	file.Close()

	sqliteDatabase, _ := sql.Open("sqlite3", "./app.db")
	defer sqliteDatabase.Close()

	initDatabase(sqliteDatabase)

	for i := 0; i < 10; i++ {
		insertRecord(sqliteDatabase, rand.Int(), rand.Float32())
	}

	// displayRecords(sqliteDatabase)
	log.Printf("%d", len(getRecords(sqliteDatabase)))
}

func initDatabase(db *sql.DB) {
	createRecordsTableStatement := `CREATE TABLE records (
		"idRecord" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
		"idDevice" integer,
		"temperature" real
	);`

	statement, err := db.Prepare(createRecordsTableStatement)
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec()
	log.Println("Database initialized")
}

func insertRecord(db *sql.DB, deviceID int, temperature float32) {
	log.Println(fmt.Sprintf("Inserting record with device id:%d, temperature:%f", deviceID, temperature))
	insertRecordStatement := `INSERT INTO records (idDevice, temperature) VALUES (?, ?)`
	statement, err := db.Prepare(insertRecordStatement)
	if err != nil {
		log.Print("Error in preparation")
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(deviceID, temperature)
	if err != nil {
		log.Print("Error while executing query")
		log.Fatalln(err.Error())
	}
}

func getRecords(db *sql.DB) []record {
	row, err := db.Query("SELECT * FROM records ORDER BY idRecord DESC")
	if err != nil {
		log.Print("Error while getting")
		log.Fatal(err.Error())
	}
	defer row.Close()
	var records []record
	for row.Next() {
		var newRecord record
		row.Scan(&newRecord.id, &newRecord.device, &newRecord.temperature)
		records = append(records, newRecord)
	}
	return records
}

func displayRecords(db *sql.DB) {
	row, err := db.Query("SELECT * FROM records ORDER BY idRecord DESC")
	if err != nil {
		log.Print("Error while displaying")
		log.Fatal(err.Error())
	}
	defer row.Close()
	for row.Next() {
		var record int
		var device int
		var temperature float32
		row.Scan(&record, &device, &temperature)
		log.Println("Record:", record, " Device:", device, " Temperature:", temperature)
	}
}

// djfkjgdhskjng
