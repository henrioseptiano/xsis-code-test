package handlers

import (
	"errors"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
	"xsis-code-test/models/request"
	"xsis-code-test/utils"
)

func (ah *AppHandler) CreateMovie(w http.ResponseWriter, r *http.Request) {
	var requestCreateMovie request.CreateMovie
	if err := utils.ReadJson(w, r, &requestCreateMovie); err != nil {
		utils.ErrorJson(w, err, http.StatusBadRequest)
		return
	}

	err := ah.AppUsecase.CreateMovie(requestCreateMovie)
	if err != nil {
		utils.ErrorJson(w, err, http.StatusNotAcceptable)
		return
	}
	jsonResponse := utils.JSONResponse{
		Error:   false,
		Message: "Success Created Movie",
	}
	utils.WriteJson(w, http.StatusCreated, jsonResponse)
	return
}

func (ah *AppHandler) ListMovie(w http.ResponseWriter, r *http.Request) {
	data, err := ah.AppUsecase.ListMovie()
	if err != nil {
		utils.ErrorJson(w, err, http.StatusNotAcceptable)
		return
	}

	jsonResponse := utils.JSONResponse{
		Error:   false,
		Message: "Success Listing Movies",
		Data:    data,
	}
	utils.WriteJson(w, http.StatusAccepted, jsonResponse)
	return
}

func (ah *AppHandler) GetMovie(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		utils.ErrorJson(w, errors.New("Id is not a numeric"), http.StatusNotAcceptable)
		return
	}
	data, err := ah.AppUsecase.GetMovie(int64(idInt))
	if err != nil {
		utils.ErrorJson(w, err, http.StatusNotAcceptable)
		return
	}

	jsonResponse := utils.JSONResponse{
		Error:   false,
		Message: "Success Getting Movie",
		Data:    data,
	}
	utils.WriteJson(w, http.StatusAccepted, jsonResponse)
	return
}

func (ah *AppHandler) UpdateMovie(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		utils.ErrorJson(w, errors.New("Id is not a numeric"), http.StatusNotAcceptable)
		return
	}

	var requestUpdateMovie request.UpdateMovie
	if err := utils.ReadJson(w, r, &requestUpdateMovie); err != nil {
		utils.ErrorJson(w, err, http.StatusBadRequest)
		return
	}

	if err := ah.AppUsecase.UpdateMovie(int64(idInt), requestUpdateMovie); err != nil {
		utils.ErrorJson(w, err, http.StatusNotAcceptable)
		return
	}

	jsonResponse := utils.JSONResponse{
		Error:   false,
		Message: "Movie Successfully Updated",
	}
	utils.WriteJson(w, http.StatusOK, jsonResponse)
	return
}

func (ah *AppHandler) DeleteMovie(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		utils.ErrorJson(w, errors.New("Id is not a numeric"), http.StatusNotAcceptable)
		return
	}

	if err := ah.AppUsecase.DeleteMovie(int64(idInt)); err != nil {
		utils.ErrorJson(w, err, http.StatusNotAcceptable)
		return
	}

	jsonResponse := utils.JSONResponse{
		Error:   false,
		Message: "Movie Sucessfully Deleted",
	}
	utils.WriteJson(w, http.StatusOK, jsonResponse)
	return
}
