package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"aviasalesTest/cmd/app"
)

func main() {
	application := app.New()

	if err := application.Init(); err != nil {
		application.Logger.Fatalf("can not init application:%v", err)
	}

	router := mux.NewRouter()
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:8080"},
	})

	router.HandleFunc("/v1/directions/DXB-BKK", application.Service.GetDirections)
	router.HandleFunc("/v1/directions/DXB-BKK/cheapest", application.Service.GetCheapest)
	router.HandleFunc("/v1/directions/DXB-BKK/mostexpensive", application.Service.GetMostExpensive)
	router.HandleFunc("/v1/directions/DXB-BKK/longest", application.Service.GetLongest)
	router.HandleFunc("/v1/directions/DXB-BKK/shortest", application.Service.GetShortest)
	handler := c.Handler(router)

	http.ListenAndServe(":2094", handler)
}