package handlers

import (
	"encoding/json"
	"net/http"
	"authgo/data"
	"authgo/models"
	"strconv"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	user := &models.User{
		ID:       data.NextUserID,
		Username: input.Username,
		Password: input.Password,
	}

	data.Users[user.ID] = user
	data.NextUserID++

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func getUserByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}
	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		http.Error(w, "ID пользователя не указан", http.StatusBadRequest)
		return
	}
	
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Некорректный ID пользователя", http.StatusBadRequest)
		return
	}

	user, exists := data.Users[id]
	if !exists {
		http.Error(w, "Пользователь не найден", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
	
