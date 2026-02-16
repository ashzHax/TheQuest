package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

type Request struct {
	Text string `json:"text"`
}

type Fill struct {
	Location int    `json:"location"`
	Letter   string `json:"letter"`
}

type Response struct {
	Action  string `json:"action"`
	Message string `json:"message,omitempty"`
	Fills   []Fill `json:"fills,omitempty"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	input := strings.TrimSpace(strings.ToLower(req.Text))

	var resp Response

	switch input {

	case "":
		resp = Response{
			Action:  "popup",
			Message: "뭔가를 쓰고 제출 눌러야지? 😏",
		}

	case "grace":
		resp = Response{
			Action: "fill",
			Fills: []Fill{
				{Location: 0, Letter: "G"},
				{Location: 2, Letter: "A"},
				{Location: 4, Letter: "E"},
			},
		}

	case "faith":
		resp = Response{
			Action: "fill",
			Fills: []Fill{
				{Location: 0, Letter: "F"},
				{Location: 1, Letter: "A"},
				{Location: 2, Letter: "I"},
				{Location: 3, Letter: "T"},
				{Location: 4, Letter: "H"},
			},
		}

	default:
		resp = Response{
			Action:  "popup",
			Message: "틀렸어용~ 🤣🤣🤣🤣🤣 \n[" + req.Text + "]",
		}
	}

	json.NewEncoder(w).Encode(resp)
}

func main() {
	http.HandleFunc("/api", handler)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
