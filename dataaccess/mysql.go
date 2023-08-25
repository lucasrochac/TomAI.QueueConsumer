package dataaccess

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"TomAI.QueueConsumer/domain"
	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	Conn *sql.DB
}

func InsertReview(review domain.BeerReview) (int, error) {
	var beerId, err = CheckIfBeerExists(review.Beer.Name)

	if err != nil {
		log.Printf(err.Error())
		return 0, err
	} else {
		if beerId == -1 {
			var breweryId = GetBreweryId(review.Beer.BrewerId)
			var styleId = GetStyleId(review.Beer.Style)

			beerId, err = InsertBeer(review.Beer.Name, styleId, review.Beer.ABV, breweryId)
			if err != nil {
				fmt.Printf(err.Error())
			}

			return beerId, nil
		}

		sucesso, err := InsertBeerReview(review.Review, beerId)

		if err != nil {
			fmt.Printf(err.Error())
		} else if sucesso == false {
			fmt.Printf("Erro ao inserir Revisão")
		}

		return beerId, nil
	}

	/**/
}

//-------- Review..

func InsertBeerReview(review domain.Review, beerId int) (bool, error) {
	dsn := "root:Crvg#2108@tcp(localhost:3306)/tomai"
	conn, err := sql.Open("mysql", dsn)

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	insert, err := conn.Exec("Insert into Review (Appearence, Aroma, Palate, Taste, Overall, ReviewDate, Profile, Text, BeerId) Values  (?,?,?,?,?,?,?,?,?)",
		review.Appearence, review.Aroma, review.Palate, review.Taste, review.Overall, time.Unix(int64(review.Time), 0), review.ProfileName, review.Text, beerId)

	if err != nil {
		fmt.Printf(err.Error())
		return false, err
	}

	insertedRows, err := insert.RowsAffected()

	if insertedRows > 0 {
		return true, nil
	} else {

		fmt.Printf(err.Error())
		return false, err
	}
}

//-------- Style..

func GetStyleId(styleName string) int {
	var styleId, err = CheckIfStyleExists(styleName)

	if err != nil {
		log.Printf(err.Error())
		styleId = -1
	} else {
		if styleId == -1 {
			styleId, err = InsertStyle(styleName)
		}
	}
	return styleId
}

func CheckIfStyleExists(stylename string) (int, error) {
	dsn := "root:Crvg#2108@tcp(localhost:3306)/tomai"
	conn, err := sql.Open("mysql", dsn)

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	var id int
	row := conn.QueryRow("Select Id from BeerStyle where Name = ?", stylename)
	erro := row.Scan(&id)

	if erro != nil {
		return -1, err
	} else {
		return id, nil
	}
}

func InsertStyle(stylename string) (int, error) {
	dsn := "root:Crvg#2108@tcp(localhost:3306)/tomai"
	conn, err := sql.Open("mysql", dsn)

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	insert, err := conn.Exec("Insert into BeerStyle (Name) Values  (?)", stylename)

	if err != nil {
		return 0, err
	}

	bid, err := insert.LastInsertId()

	if err != nil {
		return 0, err
	}

	return int(bid), err
}

//-------- Brewery..

func GetBreweryId(breweryId int) int {
	var bid, err = CheckIfBreweryExists(breweryId)

	if err != nil {
		log.Printf(err.Error())
	} else {
		if bid == -1 {
			bid, err = InsertBrewery(breweryId)
		}
	}

	return bid
}

func CheckIfBreweryExists(breweryId int) (int, error) {
	dsn := "root:Crvg#2108@tcp(localhost:3306)/tomai"
	conn, err := sql.Open("mysql", dsn)

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	var id int
	row := conn.QueryRow("Select Id from Brewery where Id = ?", breweryId)
	erro := row.Scan(&id)

	if erro != nil {
		return -1, err
	} else {
		return id, nil
	}
}

func InsertBrewery(breweryId int) (int, error) {
	dsn := "root:Crvg#2108@tcp(localhost:3306)/tomai"
	conn, openerr := sql.Open("mysql", dsn)

	if openerr != nil {
		panic(openerr)
	}

	defer conn.Close()

	insertCommand, inserterr := conn.Exec("insert into brewery (Id) values  (?)", breweryId)

	if inserterr != nil {
		fmt.Printf(" Error: %s", inserterr.Error())
		return 0, inserterr
	}

	rowsAffected, err := insertCommand.RowsAffected()

	if err != nil {
		fmt.Printf("Erro ao inserir Brewery: %d", breweryId)
	} else if rowsAffected > 0 {
		return int(breweryId), inserterr
	}

	panic(err)
}

//-------- Beer..

func CheckIfBeerExists(beername string) (int, error) {

	dsn := "root:Crvg#2108@tcp(localhost:3306)/tomai"
	conn, err := sql.Open("mysql", dsn)

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	var id int
	row := conn.QueryRow("Select Id from Beer where Name = ?", beername)
	erro := row.Scan(&id)

	if erro != nil {
		return -1, err
	} else {
		return id, nil
	}
}

func InsertBeer(beerName string, styleId int, abv float64, breweryId int) (int, error) {
	dsn := "root:Crvg#2108@tcp(localhost:3306)/tomai"
	conn, err := sql.Open("mysql", dsn)

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	insert, err := conn.Exec("Insert into Beer (Name, StyleId, ABV, BreweryId) Values (?,?,?,?)", beerName, styleId, abv, breweryId)

	if err != nil {
		return 0, err
	}

	bid, err := insert.LastInsertId()

	if err != nil {
		return 0, err
	}

	return int(bid), err
}
