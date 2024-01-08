package repository

import (
	"errors"
	"log"
	"time"
	"xsis-code-test/models/model"
)

func (ar *AppRepository) CreateMovie(movie model.Movie) error {
	if err := ar.DB.Create(&movie).Error; err != nil {
		log.Println(err.Error())
		return errors.New("Cannot Perform DB Creation")
	}
	return nil
}
func (ar *AppRepository) ListMovie() (*[]model.Movie, error) {
	movies := make([]model.Movie, 0)

	if err := ar.DB.Where("deleted_at is null").Find(&movies).Error; err != nil {
		log.Println(err.Error())
		return nil, errors.New("Cannot Perform DB Query")
	}

	return &movies, nil
}
func (ar *AppRepository) GetMovie(id int64) (*model.Movie, error) {
	var movie model.Movie

	if err := ar.DB.Where("id = ?", id).Find(&movie).Error; err != nil {
		log.Println(err.Error())
		return nil, errors.New("Cannot Perform DB Query")
	}

	return &movie, nil
}
func (ar *AppRepository) UpdateMovie(id int64, movie model.Movie) error {
	if err := ar.DB.Model(&model.Movie{}).Where("id = ?", id).Updates(&movie).Error; err != nil {
		log.Println(err.Error())
		return errors.New("Cannot Perform DB Update")
	}
	return nil
}

func (ar *AppRepository) DeleteMovie(id int64) error {
	if err := ar.DB.Model(&model.Movie{}).Where("id = ?", id).Update("deleted_at", time.Now()).Error; err != nil {
		log.Println(err.Error())
		return errors.New("Cannot Perform DB Delete")
	}
	return nil
}
