package route

import (
	"github.com/awalludinfajar/note-go-api.git/app/controller"
	"github.com/awalludinfajar/note-go-api.git/app/middleware"
	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/login", controller.Login).Methods("POST")
	router.HandleFunc("/logout", controller.Logout).Methods("POST")
	router.HandleFunc("/register", controller.Register).Methods("POST")

	noteRoute := router.PathPrefix("/notes").Subrouter()
	noteRoute.Use(middleware.AuthenticateUser)

	noteRoute.HandleFunc("", controller.GetNote).Methods("GET")
	noteRoute.HandleFunc("", controller.CreateNote).Methods("POST")
	noteRoute.HandleFunc("/{id}", controller.UpdateNote).Methods("PUT")
	noteRoute.HandleFunc("/{id}", controller.DeleteNote).Methods("DELETE")

	// router.HandleFunc("/checklists", controller.GetChecklists).Methods("GET")
	// router.HandleFunc("/checklists", controller.CreateChecklist).Methods("POST")
	// router.HandleFunc("/checklists/{id}", controller.UpdateChecklist).Methods("PUT")
	// router.HandleFunc("/checklists/{id}", controller.DeleteChecklist).Methods("DELETE")

	return router
}
