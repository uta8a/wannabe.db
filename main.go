package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type City struct {
	ID          int    `json:"id,omitempty" db:"ID"`
	Name        string `json:"name,omitempty" db:"Name"`
	CountryCode string `json:"countryCode,omitempty" db:"CountryCode"`
	District    string `json:"district,omitempty" db:"District"`
	Population  int    `json:"population,omitempty" db:"Population"`
}

type RequestedCity struct {
	Name        string `json:"name,omitempty" db:"Name"`
	CountryCode string `json:"countryCode,omitempty" db:"CountryCode"`
	District    string `json:"district,omitempty" db:"District"`
	Population  int    `json:"population,omitempty" db:"Population"`
}

var (
	db *sqlx.DB
)

func main() {
	_db, err := sqlx.Connect("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOSTNAME"), os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE")))
	if err != nil {
		log.Fatalf("Cannot Connect to Database: %s", err)
	}

	db = _db
	e := echo.New()

	e.GET("/cities/:cityName", getCityInfoHandler)
	e.POST("/cities", addCityHandler)

	e.Start(":4000")
}

func getCityInfoHandler(c echo.Context) error {
	cityName := c.Param("cityName")
	fmt.Println(cityName)

	city := City{}
	db.Get(&city, "SELECT * FROM city WHERE Name=?", cityName)
	if city.Name == "" {
		return c.NoContent(http.StatusNotFound)
	}
	return c.JSON(http.StatusOK, city)
}

func addCityHandler(c echo.Context) error {
	var cityData RequestedCity // Cityでもいける。IDは0が入る。
	if err := c.Bind(&cityData); err != nil {
		return c.JSON(http.StatusBadRequest, cityData)
	}
	fmt.Println(cityData)
	if _, err := db.NamedExec("INSERT INTO city (Name, CountryCode, District, Population) VALUES (:Name,:CountryCode,:District,:Population)", &cityData); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, cityData)
}

// $ curl localhost:4000/cities/Osaka
// {"id":1534,"name":"Osaka","countryCode":"JPN","district":"Osaka","population":2595674}

// $ curl -X POST -H "Content-Type: application/json" -d '{"name":"NekoNekoClub","countryCode":"JPN","district":"Osaka","population":100}' localhost:4000/cities
// {"name":"NekoNekoClub","countryCode":"JPN","district":"Osaka","population":100}

// $ curl localhost:4000/cities/NekoNekoClub
// {"id":4081,"name":"NekoNekoClub","countryCode":"JPN","district":"Osaka","population":100}
