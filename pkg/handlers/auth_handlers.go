package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/Swatantra-66/go-bookstore/pkg/models"
	"github.com/Swatantra-66/go-bookstore/pkg/utils"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	utils.ParseBody(r, user)

	createdUser := models.CreateUser(user)

	if createdUser.ID == 0 {
		w.WriteHeader(http.StatusConflict)
		w.Write([]byte(`{"message": "User already exists!"}`))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Account created successfully"}`))
}

func Login(w http.ResponseWriter, r *http.Request) {
	user := &models.User{}
	utils.ParseBody(r, user)

	foundUser, err := models.CheckLogin(user.Email, user.Password)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized) // 401 Unauthorized
		w.Write([]byte(`{"message": "Invalid Credentials"}`))
		return
	}

	foundUser.Password = ""
	res, _ := json.Marshal(foundUser)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
