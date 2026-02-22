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

	log.Printf("Received input: %s", input)

	switch input {

	case "":
		resp = Response{
			Action:  "popup",
			Message: "뭔가를 쓰고 \"제출\" 버튼을 눌러야 내가 답을 줄수 있지 않을까?\n🙃🙃🙃",
		}
	
	case "163":
		resp = Response{
			Action: "fill",
			Fills: []Fill{
				{Location: 1, Letter: "A"},
				{Location: 9, Letter: "H"},
			},
		}

	case "요한복음13장34절":
		resp = Response{
			Action: "fill",
			Fills: []Fill{
				{Location: 4, Letter: "E"},
				{Location: 8, Letter: "A"},
			},
		}

	case "77":
		resp = Response{
			Action: "fill",
			Fills: []Fill{
				{Location: 2, Letter: "L"},
				{Location: 5, Letter: "L"},
			},
		}

	case "1167":
		resp = Response{
			Action: "fill",
			Fills: []Fill{
				{Location: 6, Letter: "U"},
			},
		}
	
	case "174":
		resp = Response{
			Action: "fill",
			Fills: []Fill{
				{Location: 0, Letter: "H"},
				{Location: 7, Letter: "J"},
			},
		}

	case "야고보서1장5절":
		resp = Response{
			Action: "fill",
			Fills: []Fill{
				{Location: 3, Letter: "L"},
			},
		}

	default:
		resp = Response{
			Action:  "popup",
			Message: "제출된 답변 이해 안됨. 😵‍💫\n받은 답변: [" + req.Text + "]",
		}
	}

	json.NewEncoder(w).Encode(resp)
}

func main() {
	http.HandleFunc("/api", handler)

	// Serve React static files
	fs := http.FileServer(http.Dir("./dist"))
	http.Handle("/", fs)

	log.Println("Server running on :42168")
	log.Fatal(http.ListenAndServe(":42168", nil))
}
