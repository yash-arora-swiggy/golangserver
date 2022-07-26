package main

import (
	"fmt"
	"net/http"
	"encoding/json"
)

type mydata struct {
	Name string `json:"name"`
}

func handler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	data := mydata { Name: "Yash",}

	switch r.Method {
	case "GET":
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			fmt.Errorf("error in encoding")
		}
		fmt.Println(data)
	case "POST":
		dataposted := mydata{}
		err := json.NewDecoder(r.Body).Decode(&dataposted)
		if err != nil {
			fmt.Errorf("Error in decoding")
		}
		fmt.Println(dataposted)
		fmt.Fprint(w, dataposted)
	default:
		fmt.Fprint(w, "Sorry, only GET and POST methods are supported.")
	}
}
func main() {
	fmt.Println("Starting Server")

	http.HandleFunc("/",handler)
	http.ListenAndServe("localhost:8080", nil)
}