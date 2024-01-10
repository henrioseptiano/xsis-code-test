package usecase

import (
	"errors"
	"time"
	"xsis-code-test/models/model"
	"xsis-code-test/models/request"
	"xsis-code-test/models/response"
)

func (au *AppUsecase) CreateMovie(req request.CreateMovie) error {
	if req.Title == "" {
		return errors.New("Movie Title Cannot Be Empty")
	}
	if req.Description == "" {
		return errors.New("Movie Description Cannot Be Empty")
	}
	if req.Image == "" {
		return errors.New("Image Cannot Be Empty")
	}
	if req.Rating < 0 || req.Rating > 10 {
		return errors.New("Rating between 0 to 10")
	}
	movie := model.Movie{
		Title:       req.Title,
		Description: req.Description,
		Image:       req.Image,
		Rating:      req.Rating,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err := au.AppRepository.CreateMovie(movie)
	if err != nil {
		return err
	}

	return nil
}

func (au *AppUsecase) ListMovie() (*[]response.ListMovie, error) {
	movies, err := au.AppRepository.ListMovie()
	if err != nil {
		return nil, err
	}

	listMovies := make([]response.ListMovie, 0)
	for _, movie := range *movies {
		getCreatedAt := movie.CreatedAt.Format("2006-01-02 15:04:05")
		getUpdatedAt := movie.UpdatedAt.Format("2006-01-02 15:04:05")
		listMovie := response.ListMovie{
			ID:          movie.ID,
			Title:       movie.Title,
			Description: movie.Description,
			Rating:      movie.Rating,
			Image:       movie.Image,
			CreatedAt:   getCreatedAt,
			UpdatedAt:   getUpdatedAt,
		}
		listMovies = append(listMovies, listMovie)
	}

	return &listMovies, nil
}

func (au *AppUsecase) GetMovie(id int64) (*response.GetMovie, error) {
	movie, err := au.AppRepository.GetMovie(id)
	if err != nil {
		return nil, err
	}

	getCreatedAt := movie.CreatedAt.Format("2006-01-02 15:04:05")
	getUpdatedAt := movie.UpdatedAt.Format("2006-01-02 15:04:05")

	resGetMovie := &response.GetMovie{
		ID:          movie.ID,
		Title:       movie.Title,
		Description: movie.Description,
		Rating:      movie.Rating,
		Image:       movie.Image,
		CreatedAt:   getCreatedAt,
		UpdatedAt:   getUpdatedAt,
	}

	return resGetMovie, nil
}

func (au *AppUsecase) UpdateMovie(id int64, req request.UpdateMovie) error {
	if req.Title == "" {
		return errors.New("Movie Title Cannot Be Empty")
	}
	if req.Description == "" {
		return errors.New("Movie Description Cannot Be Empty")
	}
	if req.Image == "" {
		return errors.New("Image Cannot Be Empty")
	}
	if req.Rating < 0 || req.Rating > 10 {
		return errors.New("Rating between 0 to 10")
	}
	movie, err := au.AppRepository.GetMovie(id)
	if err != nil {
		return err
	}
	movie.Title = req.Title
	movie.Description = req.Description
	movie.Image = req.Image
	movie.Rating = req.Rating
	err = au.AppRepository.UpdateMovie(id, *movie)
	if err != nil {
		return err
	}

	return nil
}

func (au *AppUsecase) DeleteMovie(id int64) error {
	_, err := au.AppRepository.GetMovie(id)
	if err != nil {
		return err
	}

	err = au.AppRepository.DeleteMovie(id)
	if err != nil {
		return err
	}

	return nil
}
