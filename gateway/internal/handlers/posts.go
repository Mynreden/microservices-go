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

func (h *Handler) postHandler() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/create", h.createPost).Methods("POST")
	r.HandleFunc("/delete/{id}", h.deletePost).Methods("DELETE")
	r.HandleFunc("/{id}", h.getPost).Methods("GET")
	r.HandleFunc("/user/{id}", h.getUserPosts).Methods("GET")
	return r
}

func (h *Handler) createPost(w http.ResponseWriter, r *http.Request) {
	post := models.Post{}
	err := utils.ReadJSON(r, &post)
	if err != nil {
		log.Println(err.Error())
		utils.WriteErr(w, http.StatusBadRequest, "Incorrect Request Body")
		return
	}
	createPost, err := h.postService.CreatePost(context.Background(), &proto.CreatePostRequest{
		Title:   post.Title,
		Content: post.Content,
		UserId:  post.UserId,
	})
	if err != nil {
		return
	}

	newPost := models.Post{ID: createPost.Id,
		Title:   createPost.Title,
		Content: createPost.Content,
		UserId:  createPost.UserId,
	}

	utils.WriteJSON(w, 200, newPost)
}

func (h *Handler) deletePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	response, err := h.postService.DeletePost(context.Background(), &proto.DeletePostRequest{
		Id: id,
	})
	if err != nil {
		utils.WriteErr(w, http.StatusInternalServerError, "Failed to delete user")
		return
	}

	utils.WriteJSON(w, http.StatusOK, map[string]string{"message": response.Message})
}

func (h *Handler) getPost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	postResponse, err := h.postService.GetPost(context.Background(), &proto.GetPostRequest{
		Id: id,
	})
	if err != nil {
		utils.WriteErr(w, http.StatusInternalServerError, "Failed to get user")
		return
	}

	post := models.Post{
		ID:      postResponse.Id,
		Title:   postResponse.Title,
		Content: postResponse.Content,
		UserId:  postResponse.UserId,
	}

	utils.WriteJSON(w, http.StatusOK, post)
}

func (h *Handler) getUserPosts(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]

	postsResponse, err := h.postService.GetPostsByUserId(context.Background(), &proto.GetPostsByUserIdRequest{
		UserId: userId,
	})
	if err != nil {
		utils.WriteErr(w, http.StatusInternalServerError, "Failed to get user")
		return
	}
	posts := []models.Post{}
	for _, post := range postsResponse.Posts {
		posts = append(posts, models.Post{ID: post.Id,
			Title:   post.Title,
			Content: post.Content,
			UserId:  post.UserId,
		})
	}

	utils.WriteJSON(w, http.StatusOK, posts)
}
