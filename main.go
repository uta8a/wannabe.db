package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type City struct {
	ID          int    `json:"id,omitempty" db:"ID"`
	Name        string `json:"name,omitempty" db:"Name"`
	CountryCode string `json:"countryCode,omitempty" db:"CountryCode"`
	District    string `json:"district,omitempty" db:"District"`
	Population  int    `json:"population,omitempty" db:"Population"`
}

func main() {
	// if len(os.Args) < 2 {
	// 	log.Fatalln("usage: go run main.go <city_name>")
	// }
	// cityName := os.Args[1]

	db, err := sqlx.Connect("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOSTNAME"), os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE")))
	if err != nil {
		log.Fatalf("Cannot Connect to Database: %s", err)
	}

	fmt.Println("Connected!")
	// city := City{}
	// db.Get(&city, "SELECT * FROM city WHERE Name = ?", cityName)

	// fmt.Printf("%sの人口は%d人です\n", cityName, city.Population)
	db.Exec("INSERT INTO city (Name, CountryCode, District, Population) VALUES ('wakuwaku-ookoku', 'JPN', 'wakuwaku-island', 100)")
}
