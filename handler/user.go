package handler

import (
	"encoding/json"
	"fmt"
	"game_suit/core/module"
	"net/http"

	"github.com/gorilla/mux"
)

type userHandler struct {
	userUsecase module.UserUsecase
}

func NewUserHandler(userUsecase module.UserUsecase) *userHandler {
	return &userHandler{userUsecase: userUsecase}
}

func (e *userHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id_game"]
	user, err := e.userUsecase.GetUser(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	result, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Println("hasil:", result)
	w.Write(result)

}
