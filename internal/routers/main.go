package routers

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/ayushWeb07/AirBnb-Go-Api-Gateway/internal/config"
	"github.com/ayushWeb07/AirBnb-Go-Api-Gateway/internal/controllers"
	"github.com/ayushWeb07/AirBnb-Go-Api-Gateway/internal/repositories"
	"github.com/ayushWeb07/AirBnb-Go-Api-Gateway/internal/services"
	"github.com/ayushWeb07/AirBnb-Go-Api-Gateway/internal/utils"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httprate"
	"go.uber.org/zap"
)

type RouterInterface interface {
	Register(r *chi.Mux)
}

func RegisterRouters(logger *zap.Logger, db *sql.DB, serverConfig *config.ServerConfig) *chi.Mux {
	// create the router instance
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Use(httprate.Limit(
		10,
		time.Minute,
		httprate.WithKeyFuncs(httprate.KeyByIP, httprate.KeyByEndpoint),
		httprate.WithLimitHandler(func(resWriter http.ResponseWriter, req *http.Request) {
			utils.WriteJsonResponse(http.StatusTooManyRequests, resWriter, map[string]any{
				"success": false,
				"message": "Too many requests",
				"error":   "You have been rate-limited. Please, slow down and try again after sometime",
			})
		}),
	))

	// register health router
	//SetupHealthRouter(router)

	// register user router
	userRepository := repositories.NewUserRepository(logger, db, serverConfig)
	userService := services.NewUserService(userRepository, logger, serverConfig)
	userController := controllers.NewUserController(userService, logger, serverConfig)
	userRouter := NewUserRouter(userController, logger, serverConfig)

	userRouter.Register(router)

	return router
}
