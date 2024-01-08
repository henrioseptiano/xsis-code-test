package handlers

import "xsis-code-test/app"

type AppHandler struct {
	AppUsecase app.IAppUsecase
}

func NewAppHandler(appUsecase app.IAppUsecase) *AppHandler {
	return &AppHandler{AppUsecase: appUsecase}
}
