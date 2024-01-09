package repository

import (
	"github.com/stretchr/testify/mock"
	"xsis-code-test/models/model"
)

type AppRepositoryMock struct {
	Mock mock.Mock
}

func (arm *AppRepositoryMock) CreateMovie(movie model.Movie) error {
	arguments := arm.Mock.Called(movie)
	if arguments.Get(0) == nil {
		return nil
	}
	return arguments.Get(0).(error)
}

func (arm *AppRepositoryMock) ListMovie() (*[]model.Movie, error) {
	arguments := arm.Mock.Called()

	if arguments.Get(1) == nil {
		movie := arguments.Get(0).(*[]model.Movie)
		return movie, nil
	}
	return nil, arguments.Get(1).(error)
}

func (arm *AppRepositoryMock) GetMovie(id int64) (*model.Movie, error) {
	arguments := arm.Mock.Called(id)

	if arguments.Get(1) == nil {
		movie := arguments.Get(0).(*model.Movie)
		return movie, nil
	}
	return nil, arguments.Get(1).(error)
}

func (arm *AppRepositoryMock) UpdateMovie(id int64, movie model.Movie) error {
	arguments := arm.Mock.Called(id, movie)

	if arguments.Get(0) == nil {
		return nil
	}

	return arguments.Get(0).(error)
}

func (arm *AppRepositoryMock) DeleteMovie(id int64) error {
	arguments := arm.Mock.Called(id)

	if arguments.Get(0) == nil {
		return nil
	}

	return arguments.Get(0).(error)
}
