package api

import (
	"net/http"
	"user-service-client/internal/interfaces/input/api/handler/user"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

func SetUpRoutes(userHand *user.UserHandlerImpl) *chi.Mux {

	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type", "Set-Cookie", "Cookie"},
		AllowCredentials: true,
		ExposedHeaders:   []string{},
	}))

	r.Get("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Server is running"))
	})

	r.Route("/api/user", func(r chi.Router) {

		r.Get("/{user_id}", userHand.FindUserById())

	})

	return r
}
