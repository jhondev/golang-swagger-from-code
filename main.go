// Package classification Swagger spec from code.
// version: 1.0
// Consumes:
//     - application/json
//
// Produces:
//     - application/json
// swagger:meta
package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"

	"github.com/jhondev/golang-swagger-from-code/models"
)

func main() {
	router := chi.NewRouter()
	router.Get("/", HomeHandler)
	router.Get("/test", TestHandler)

	println("serving on 8182...")
	http.ListenAndServe(":8182", router)
}

// ApiResponse model
// swagger1:response
type ApiResponse struct {
	Code    string `json:"code"`
	Content string `json:"content"`
}

// HomeHandler swagger:route GET / home
// Responses:
// 200: ApiResponse
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, ApiResponse{Content: "description main"})
}

// TestHandler swagger:route GET /test
// Responses:
// 200: TestModel
func TestHandler(w http.ResponseWriter, r *http.Request) {
	render.JSON(w, r, models.TestModel{Desc: "description test"})
}
