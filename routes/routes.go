package routes

import (
	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
	"net/http"
	"xsis-code-test/app"
	AppHandler "xsis-code-test/app/handlers"
	AppRepo "xsis-code-test/app/repository"
	AppUsecase "xsis-code-test/app/usecase"
)

func implementHandler(appHandler app.IAppHandlers) app.IAppHandlers {
	return appHandler
}

func AppRoutes(db *gorm.DB) http.Handler {
	appRepo := AppRepo.NewAppRepository(db)
	appUsecase := AppUsecase.NewAppUsecase(appRepo)
	appHandler := AppHandler.NewAppHandler(appUsecase)
	implHandler := implementHandler(appHandler)
	route := chi.NewMux()

	route.Post("/Movie", implHandler.CreateMovie)
	route.Get("/Movie", implHandler.ListMovie)
	route.Get("/Movie/{id}", implHandler.GetMovie)
	route.Put("/Movie/{id}", implHandler.UpdateMovie)
	route.Delete("/Movie/{id}", implHandler.DeleteMovie)

	return route
}
