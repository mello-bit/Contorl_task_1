package services

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"web_calculator/utils"
)


func TestProcessRequest(t *testing.T) {
	testCases := []utils.TestRequest{
		{
			Method: http.MethodPost,
			Name: "Addition",
			Expression: strings.NewReader(`{"expression": "2+3"}`),
			StatusCode: http.StatusOK,
		},
		{
			Method: http.MethodGet,
			Name: "Subtract",
			Expression: strings.NewReader(`{"expression": "2-3"}`),
			StatusCode: http.StatusOK,
		},
		{
			Method: http.MethodPost,
			Name: "Divide by ZERO",
			Expression: strings.NewReader(`{"expression": "2/0"}`),
			StatusCode: http.StatusUnprocessableEntity,
		},
		{
			Method: http.MethodPost,
			Name: "Divide",
			Expression: strings.NewReader(`{"expression": "2/33"}`),
			StatusCode: http.StatusOK,
		},
		{
			Method: http.MethodPost,
			Name: "Multiply",
			Expression: strings.NewReader(`{"expression": "2 * 33"}`),
			StatusCode: http.StatusOK,
		},
		{
			Method: http.MethodPost,
			Name: "Long expression",
			Expression: strings.NewReader(`{"expression": "(3 + 5 * 8) - 9 * 6 + 22 - 3"}`),
			StatusCode: http.StatusOK,
		},
		{
			Method: http.MethodPost,
			Name: "Long expression with division by zero",
			Expression: strings.NewReader(`{"expression": "(3 + 5 * 8) - 9 * 6 / (22 + 3 - 4 - 21)"}`),
			StatusCode: http.StatusUnprocessableEntity,
		},
	}
	
	for _, tt := range testCases {
		t.Run(tt.Name, func(t *testing.T) {
			req := httptest.NewRequest(tt.Method, "/", tt.Expression)
			w := httptest.NewRecorder()
			calculateExpression(w, req)

			res := w.Result()
			defer res.Body.Close()

			if res.StatusCode != tt.StatusCode {
				t.Errorf("wrong status code on expression %s", tt.Name)
			}
		})
	}
}