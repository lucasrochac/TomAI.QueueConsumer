package dataaccess

import (
	"database/sql"
	"fmt"
	"log"

	"TomAI.QueueConsumer/domain"

	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	Conn *sql.DB
}

func (db *Database) Init() {
	connectionString := fmt.Sprintf("%s:%s@/%s", "root", "Crvg#2108", "TomAI")

	var err error
	db.Conn, err = sql.Open("mysql", connectionString)

	if err != nil {
		log.Fatal(err)
	}

	if err = db.Conn.Ping(); err != nil {
		log.Fatal(err)
	}
}

func (db *Database) CheckIfBeerExists(beername string) (int, error) {
	var id int
	row := db.Conn.QueryRow("Select Id from Beer where Name = ?", beername)
	err := row.Scan(&id)

	if err != nil {
		return id, err
	} else {
		return 0, err
	}
}

func (db *Database) CheckIfBreweryExists(breweryid int) (int, error) {
	var id int
	row := db.Conn.QueryRow("Select Id from Brewery where Id = ?", breweryid)
	err := row.Scan(&id)

	if err != nil {
		return id, err
	} else {
		return 0, err
	}
}

func (db *Database) CheckIfStyleExists(stylename string) (int, error) {
	var id int
	row := db.Conn.QueryRow("Select Id from BeerStyle where Name = ?", stylename)
	err := row.Scan(&id)

	if err != nil {
		return id, err
	} else {
		return 0, err
	}
}

func (db *Database) InsertBeer(beer domain.Beer) error {
	var brewery, err = db.CheckIfBreweryExists(beer.BrewerId)
	if brewery == 0 {
		brewery, err = db.InsertBrewery(beer.BrewerId)

		if err != nil {
			log.Fatal("Erro ao salvar novo estilo. Beer: %s | Style: %s", beer.Name, beer.Style)
		}
	}

	var style, stylecheckerr = db.CheckIfStyleExists(beer.Style)
	if style == 0 {
		style, stylecheckerr = db.InsertStyle(beer.Style)

		if stylecheckerr != nil {
			log.Fatal("Erro ao salvar novo estilo. Beer: %s | Style: %s", beer.Name, beer.Style)
		}
	}

	insert, err := db.Conn.Query("Insert into Beer (Name, BreweryId, ABV, StyleId) Values  (?, ?, ?)", beer.Name, brewery, beer.ABV, style)
	if err != nil {
		return err
	}

	defer insert.Close()
	return nil
}

func (db *Database) InsertBrewery(breweryid int) (int, error) {
	insert, err := db.Conn.Exec("Insert into Brewery (Id) Values  (?)", breweryid)

	if err != nil {
		return 0, err
	}

	bid, err := insert.LastInsertId()

	if err != nil {
		return 0, err
	}

	return int(bid), err
}

func (db *Database) InsertStyle(stylename string) (int, error) {
	insert, err := db.Conn.Exec("Insert into BeerStyle (Name) Values  (?)", stylename)

	if err != nil {
		return 0, err
	}

	bid, err := insert.LastInsertId()

	if err != nil {
		return 0, err
	}

	return int(bid), err
}

func (db *Database) Close() {
	db.Conn.Close()
}
