package usecase

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
	"xsis-code-test/app/repository"
	"xsis-code-test/models/model"
	"xsis-code-test/models/request"
)

var appRepo = &repository.AppRepositoryMock{Mock: mock.Mock{}}
var appUsecase = AppUsecase{AppRepository: appRepo}

func Test_CreateMovie(t *testing.T) {
	testcases := []struct {
		name                   string
		isResultNil            bool
		input                  request.CreateMovie
		isBasicValidationError bool
	}{
		{
			name:        "valid data",
			isResultNil: true,
			input: request.CreateMovie{
				Title:       "Dans 1",
				Description: "Dans 1",
				Rating:      4,
				Image:       "fafa.jpg",
			},
			isBasicValidationError: false,
		},
		{
			name:        "title empty",
			isResultNil: false,
			input: request.CreateMovie{
				Title:       "",
				Description: "Dans 1",
				Rating:      4,
				Image:       "fafa.jpg",
			},
			isBasicValidationError: true,
		},
		{
			name:        "description empty",
			isResultNil: false,
			input: request.CreateMovie{
				Title:       "Dans 1",
				Description: "",
				Rating:      4,
				Image:       "fafa.jpg",
			},
			isBasicValidationError: true,
		},
		{
			name:        "rating not valid",
			isResultNil: false,
			input: request.CreateMovie{
				Title:       "Dans 1",
				Description: "Dans 1",
				Rating:      11,
				Image:       "fafa.jpg",
			},
			isBasicValidationError: true,
		},
		{
			name:        "image empty",
			isResultNil: false,
			input: request.CreateMovie{
				Title:       "Dans 1",
				Description: "Dans 1",
				Rating:      11,
				Image:       "",
			},
			isBasicValidationError: true,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isBasicValidationError == false {
				movie := model.Movie{
					Title:       tc.input.Title,
					Description: tc.input.Description,
					Rating:      tc.input.Rating,
					Image:       tc.input.Image,
					CreatedAt:   time.Now(),
					UpdatedAt:   time.Now(),
				}
				appRepo.Mock.On("CreateMovie", movie).Return(nil)
			}
			err := appUsecase.CreateMovie(tc.input)
			if tc.isResultNil {
				assert.Nil(t, err)
			} else {
				assert.NotNil(t, err)
			}
		})
	}
}

func Test_UpdateMovie(t *testing.T) {
	createDateTime, _ := time.Parse("2006-01-02 15:04:05", "2024-01-03 00:00:00")
	testcases := []struct {
		name                   string
		isResultNil            bool
		input                  request.UpdateMovie
		existingMovieData      *model.Movie
		id                     int64
		isBasicValidationError bool
	}{
		{
			name:        "valid data",
			isResultNil: true,
			input: request.UpdateMovie{
				Title:       "Dans 4",
				Description: "Dans 4",
				Rating:      4,
				Image:       "fafa2.jpg",
			},
			existingMovieData: &model.Movie{
				ID:          1,
				Title:       "Dans 4",
				Description: "Dans 4",
				Rating:      6,
				Image:       "tori.jpg",
				CreatedAt:   createDateTime,
				UpdatedAt:   createDateTime,
			},
			id: 1,
		},
		{
			name:        "title empty",
			isResultNil: false,
			input: request.UpdateMovie{
				Title:       "",
				Description: "Dans 2",
				Rating:      4,
				Image:       "fafa2.jpg",
			},
			existingMovieData: &model.Movie{
				ID:          1,
				Title:       "",
				Description: "Dans 1",
				Rating:      6,
				Image:       "tori.jpg",
				CreatedAt:   createDateTime,
				UpdatedAt:   createDateTime,
			},
			id:                     1,
			isBasicValidationError: true,
		},
		{
			name:        "description empty",
			isResultNil: false,
			input: request.UpdateMovie{
				Title:       "Dans 2",
				Description: "",
				Rating:      4,
				Image:       "fafa2.jpg",
			},
			existingMovieData: &model.Movie{
				ID:          1,
				Title:       "Dans 1",
				Description: "Dans 1",
				Rating:      6,
				Image:       "tori.jpg",
				CreatedAt:   createDateTime,
				UpdatedAt:   createDateTime,
			},
			id:                     1,
			isBasicValidationError: true,
		},
		{
			name:        "rating not valid",
			isResultNil: false,
			input: request.UpdateMovie{
				Title:       "Dans 2",
				Description: "Dans 2",
				Rating:      11,
				Image:       "fafa2.jpg",
			},
			existingMovieData: &model.Movie{
				ID:          1,
				Title:       "Dans 1",
				Description: "Dans 1",
				Rating:      6,
				Image:       "tori.jpg",
				CreatedAt:   createDateTime,
				UpdatedAt:   createDateTime,
			},
			id:                     1,
			isBasicValidationError: true,
		},
		{
			name:        "image empty",
			isResultNil: false,
			input: request.UpdateMovie{
				Title:       "Dans 2",
				Description: "Dans 2",
				Rating:      4,
				Image:       "",
			},
			existingMovieData: &model.Movie{
				ID:          1,
				Title:       "Dans 1",
				Description: "Dans 1",
				Rating:      6,
				Image:       "tori.jpg",
				CreatedAt:   createDateTime,
				UpdatedAt:   createDateTime,
			},
			id:                     1,
			isBasicValidationError: true,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isBasicValidationError == false {
				movie := model.Movie{
					ID:          tc.existingMovieData.ID,
					Title:       tc.input.Title,
					Description: tc.input.Description,
					Rating:      tc.input.Rating,
					Image:       tc.input.Image,
					CreatedAt:   tc.existingMovieData.CreatedAt,
					UpdatedAt:   tc.existingMovieData.UpdatedAt,
				}
				appRepo.Mock.On("GetMovie", tc.id).Return(tc.existingMovieData, nil)
				appRepo.Mock.On("UpdateMovie", tc.id, movie).Return(nil)
			}
			err := appUsecase.UpdateMovie(tc.id, tc.input)
			if tc.isResultNil {
				assert.Nil(t, err)
			} else {
				assert.NotNil(t, err)
			}
		})
	}
}

func Test_DeleteMovie(t *testing.T) {
	createDateTime, _ := time.Parse("2006-01-02 15:04:05", "2024-01-03 00:00:00")
	testcases := []struct {
		name              string
		isResultNil       bool
		existingMovieData *model.Movie
		id                int64
	}{
		{
			name:        "valid data",
			isResultNil: true,
			existingMovieData: &model.Movie{
				ID:          1,
				Title:       "Dans 1",
				Description: "Dans 1",
				Rating:      6,
				Image:       "tori.jpg",
				CreatedAt:   createDateTime,
				UpdatedAt:   createDateTime,
			},
			id: 1,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			appRepo.Mock.On("GetMovie", tc.id).Return(tc.existingMovieData, nil)
			appRepo.Mock.On("DeleteMovie", tc.id).Return(nil)
			err := appUsecase.DeleteMovie(tc.id)
			if tc.isResultNil {
				assert.Nil(t, err)
			} else {
				assert.NotNil(t, err)
			}
		})
	}
}

func Test_GetMovie(t *testing.T) {
	createDateTime, _ := time.Parse("2006-01-02 15:04:05", "2024-01-03 00:00:00")
	testcases := []struct {
		name              string
		isResultNil       bool
		existingMovieData *model.Movie
		id                int64
	}{
		{
			name:        "valid data",
			isResultNil: true,
			existingMovieData: &model.Movie{
				ID:          1,
				Title:       "Dans 1",
				Description: "Dans 1",
				Rating:      6,
				Image:       "tori.jpg",
				CreatedAt:   createDateTime,
				UpdatedAt:   createDateTime,
			},
			id: 1,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			appRepo.Mock.On("GetMovie", tc.id).Return(tc.existingMovieData, nil)
			_, err := appUsecase.GetMovie(tc.id)
			if tc.isResultNil {
				assert.Nil(t, err)
			} else {
				assert.NotNil(t, err)
			}
		})
	}
}

func Test_ListMovie(t *testing.T) {
	createDateTime, _ := time.Parse("2006-01-02 15:04:05", "2024-01-03 00:00:00")
	testcases := []struct {
		name              string
		isResultNil       bool
		existingMovieData *[]model.Movie
	}{
		{
			name:        "valid data",
			isResultNil: true,
			existingMovieData: &[]model.Movie{
				{
					ID:          1,
					Title:       "Dans 1",
					Description: "Dans 1",
					Rating:      6,
					Image:       "tori.jpg",
					CreatedAt:   createDateTime,
					UpdatedAt:   createDateTime,
				},
				{
					ID:          2,
					Title:       "Dans 2",
					Description: "Dans 2",
					Rating:      7,
					Image:       "tori2.jpg",
					CreatedAt:   createDateTime,
					UpdatedAt:   createDateTime,
				},
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			appRepo.Mock.On("ListMovie").Return(tc.existingMovieData, nil)
			_, err := appUsecase.ListMovie()
			if tc.isResultNil {
				assert.Nil(t, err)
			} else {
				assert.NotNil(t, err)
			}
		})
	}
}
