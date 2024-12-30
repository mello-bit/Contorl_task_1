package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"web_calculator/utils"
)

const url = "/api/v1/calculate"


func ProcessRequest(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(utils.Response{
				Error: "Internal server error",
			})
	
			return
		}

		next.ServeHTTP(w, r)
	})
}


func calculateExpression(w http.ResponseWriter, r *http.Request) {
	var expression utils.Expression
	json.NewDecoder(r.Body).Decode(&expression)

	answer, err := utils.Calc(expression.Expression)
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(utils.Response{
			Error: "Expression is not valid",
		})

		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(utils.Answer{
		Result: fmt.Sprintf("%.4f", answer),
	})
}

func StartServer() {
	mux := http.NewServeMux()

	mux.Handle(url, ProcessRequest(calculateExpression))

	http.ListenAndServe(":8080", mux)
}