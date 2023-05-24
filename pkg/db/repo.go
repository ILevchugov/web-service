package db

import (
	"database/sql"
	"fmt"
	"web-service/pkg/model"

	_ "github.com/lib/pq"
)

var dbConnection *sql.DB

func InitRepo() {
	var err error
	dbConnection, err = GetConnection(
		"localhost",
		5432,
		"postgres",
		"postgres",
		"postgres")

	if err != nil {
		fmt.Print(err)
	}
	err = dbConnection.Ping()
	if err != nil {
		panic("SHIT")
	}
}

func GetMovies(offset int, limit int) ([]model.Movie, error) {
	var err error

	if dbConnection == nil {
		fmt.Println("There is no db connection, init new one")
		InitRepo()
	}

	var movies []model.Movie

	query := "select * from movies offset $1 limit $2"
	rows, err := dbConnection.Query(query, offset, limit)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	for rows.Next() {
		var movie model.Movie

		if err := rows.Scan(&movie.Id, &movie.Title, &movie.Image, &movie.Description); err != nil {
			return nil, fmt.Errorf("movies  offser %q: limit: %q,  %v", offset, limit, err)
		}
		movies = append(movies, movie)
	}

	return movies, err

}
