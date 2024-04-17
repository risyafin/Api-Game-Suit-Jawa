package handler

import (
	"encoding/json"
	"fmt"
	"game_suit/core/module"
	"net/http"
	"strconv"

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
	id := vars["id"]
	idInt, err := strconv.Atoi(id)
	if err != nil {
		fmt.Fprintf(w, err.Error(), http.StatusInternalServerError)
	}
	user, err := e.userUsecase.GetUser(idInt)
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
	w.Write(result)

}
