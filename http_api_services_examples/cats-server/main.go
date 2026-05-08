package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"sync"
)

type Cat struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Color string `json:"color"`
	Age   int    `json:"age"`
}

// In-memory storage
var (
	cats    []Cat
	nextID  = 1
	catsMux sync.Mutex
)

func listCats(w http.ResponseWriter, r *http.Request) {
	catsMux.Lock()

	defer catsMux.Unlock()

	sendJSON(w, http.StatusOK, cats)
}

func createCat(w http.ResponseWriter, r *http.Request) {
	var c Cat

	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	catsMux.Lock()

	c.ID = nextID
	nextID++
	cats = append(cats, c)

	catsMux.Unlock()

	sendJSON(w, http.StatusCreated, c)
}

func updateCat(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(w, r)
	if !ok {
		return
	}

	var upd Cat
	if err := json.NewDecoder(r.Body).Decode(&upd); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)

		return
	}

	catsMux.Lock()
	defer catsMux.Unlock()

	for i := range cats {
		if cats[i].ID == id {
			upd.ID = id // keeping existing ID
			cats[i] = upd
			sendJSON(w, http.StatusOK, upd)

			return
		}
	}
	http.NotFound(w, r)
}

func deleteCat(w http.ResponseWriter, r *http.Request) {
	id, ok := parseID(w, r)

	if !ok {
		return
	}

	catsMux.Lock()
	defer catsMux.Unlock()

	for i, c := range cats {
		if c.ID == id {
			cats = append(cats[:i], cats[i+1:]...)
			w.WriteHeader(http.StatusNoContent)

			return
		}
	}

	http.NotFound(w, r)
}

func sendJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}

func parseID(w http.ResponseWriter, r *http.Request) (int, bool) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return 0, false
	}

	return id, true
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /cats", listCats)
	mux.HandleFunc("POST /cats", createCat)
	mux.HandleFunc("PUT /cats/{id}", updateCat)
	mux.HandleFunc("DELETE /cats/{id}", deleteCat)

	log.Println("Listening on :8080 …")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
