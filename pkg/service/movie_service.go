package service

import (
	"errors"
	"fmt"
	"web-service/pkg/db"
	"web-service/pkg/model"
)

func FindMovies(page int, pageSize int) ([]model.Movie, error) {

	movies, err := db.GetMovies(page*pageSize, pageSize)

	if err == nil {
		fmt.Printf("movies: %v\n", movies)
		return movies, nil
	}

	return nil, errors.New("problems with database: " + err.Error())
}
