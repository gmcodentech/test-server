package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func returnJSON() ([]byte, error) {
	person := Person{Name: "ABC XYZ", Age: 33}
	jsonData, err := json.Marshal(person)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	jsonData, err := returnJSON()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
