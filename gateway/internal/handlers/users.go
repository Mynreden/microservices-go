package handlers

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/mynreden/microservices-go/common/api/proto"
	"github.com/mynreden/microservices-go/common/models"
	"github.com/mynreden/microservices-go/common/utils"
	"log"
	"net/http"
)

func (h *Handler) userHandler() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/create", h.createUser).Methods("POST")
	r.HandleFunc("/delete/{id}", h.deleteUser).Methods("DELETE")
	r.HandleFunc("/{id}", h.getUser).Methods("GET")
	return r
}

func (h *Handler) createUser(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	err := utils.ReadJSON(r, &user)
	if err != nil {
		log.Println(err.Error())
		utils.WriteErr(w, http.StatusBadRequest, "Incorrect Request Body")
		return
	}
	createUser, err := h.userService.CreateUser(context.Background(), &proto.CreateUserRequest{
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	})
	if err != nil {
		return
	}

	newUser := models.User{ID: createUser.Id,
		Email:    createUser.Email,
		Password: createUser.Password,
		Username: createUser.Username,
	}

	utils.WriteJSON(w, 200, newUser)
}

func (h *Handler) deleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	response, err := h.userService.DeleteUser(context.Background(), &proto.DeleteUserRequest{
		Id: id,
	})
	if err != nil {
		utils.WriteErr(w, http.StatusInternalServerError, "Failed to delete user")
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{"message": response.Message})
}

func (h *Handler) getUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	userResponse, err := h.userService.GetUser(context.Background(), &proto.GetUserRequest{
		Id: id,
	})
	if err != nil {
		utils.WriteErr(w, http.StatusInternalServerError, "Failed to get user")
		return
	}

	user := models.User{
		ID:       userResponse.Id,
		Username: userResponse.Username,
		Email:    userResponse.Email,
		Password: userResponse.Password,
	}

	utils.WriteJSON(w, http.StatusOK, user)
}
