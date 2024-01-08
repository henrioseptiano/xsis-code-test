package app

import (
	"net/http"
	"xsis-code-test/models/model"
	"xsis-code-test/models/request"
	"xsis-code-test/models/response"
)

type IAppHandlers interface {
	CreateMovie(http.ResponseWriter, *http.Request)
	ListMovie(http.ResponseWriter, *http.Request)
	GetMovie(http.ResponseWriter, *http.Request)
	UpdateMovie(http.ResponseWriter, *http.Request)
	DeleteMovie(http.ResponseWriter, *http.Request)
}

type IAppUsecase interface {
	CreateMovie(request.CreateMovie) error
	ListMovie() (*[]response.ListMovie, error)
	GetMovie(int64) (*response.GetMovie, error)
	UpdateMovie(int64, request.UpdateMovie) error
	DeleteMovie(int64) error
}

type IAppRepository interface {
	CreateMovie(model.Movie) error
	ListMovie() (*[]model.Movie, error)
	GetMovie(int64) (*model.Movie, error)
	UpdateMovie(int64, model.Movie) error
	DeleteMovie(int64) error
}
