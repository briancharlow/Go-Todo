package api

import (
	"encoding/json"

	"net/http"

	"todo.com/go/model"
)

func CreateTodo(w http.ResponseWriter, r *http.Request){
	
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid Method", http.StatusBadRequest)
		return
	} else {
		var todo model.Todo
		
		err := json.NewDecoder(r.Body).Decode(&todo)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}
		model.AddTodo(todo)
		w.WriteHeader(http.StatusCreated)
	}
}