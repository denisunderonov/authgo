package handlers

import "net/http"

func createHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Данный метод не поддерживается", http.StatusMethodNotAllowed)
		return
	}

}
