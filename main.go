package main

import (
   "fmt"
   "net/http"
   "math"
)

func main() {
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8000", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
        s := fmt.Sprintf("%f", sqrtCalc(0.0001))
	fmt.Fprintf(w, greeting("Code.education Rocks!") + s)
}

func sqrtCalc(x float64) float64 {
	for i:=0; i<=1000000; i++ {
	        x += math.Sqrt(x)
	}
	return x
}

func greeting(message string) string {
	return "<b>" + message + "</b>"
}
