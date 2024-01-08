package usecase

import "xsis-code-test/app"

type AppUsecase struct {
	AppRepository app.IAppRepository
}

func NewAppUsecase(appRepo app.IAppRepository) *AppUsecase {
	return &AppUsecase{AppRepository: appRepo}
}
