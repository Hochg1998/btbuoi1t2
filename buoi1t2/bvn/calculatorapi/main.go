package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func handleRequests() {
	http.HandleFunc("/", Handler)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	handleRequests()
}

func Handler(w http.ResponseWriter, r *http.Request) {

	//lấy tham số từ url
	calculation := r.URL.Query().Get("calculation")
	operator1, _ := strconv.ParseFloat(r.URL.Query().Get("operator1"), 64)
	operator2, _ := strconv.ParseFloat(r.URL.Query().Get("operator2"), 64)

	//hiển thị kết quả
	fmt.Fprintln(w, "Phép tính: ", operator1, calculation, operator2)
	fmt.Fprintln(w, "Kết quả: ", Calculate(calculation, operator1, operator2))

}

func Calculate(calc string, operator1 float64, operator2 float64) float64 {
	switch calc {
	case "mul":
		return operator1 * operator2
	case "sub":
		return operator1 - operator2
	case "add":
		return operator1 + operator2
	case "div":
		return operator1 / operator2
	default:
		return 0
	}
}
