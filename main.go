package main

import (
	"game_suit/config"
	"game_suit/core/module"
	"game_suit/handler"
	userrepository "game_suit/repository/user-repository"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
}

func main() {

	conf := config.GetConfig()
	conn := config.InitDatabaseConnection(conf)
	config.AutoMigrate(conn)

	r := mux.NewRouter()

	userRepo := userrepository.NewUserRepository(conn)
	userUsecase := module.NewUserUsecase(userRepo)
	userHandler := handler.NewUserHandler(userUsecase)

	r.HandleFunc("/users/{id_game}", userHandler.GetUser).Methods("GET")
	r.HandleFunc("/users",userHandler.GetUsers).Methods("GET")
	http.ListenAndServe(":8080", r)

}
