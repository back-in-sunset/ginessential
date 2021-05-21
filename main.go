package main

import (
	"gin-essential/dao"
	"gin-essential/router"
	"net/http"
	"time"
)

func main() {
	dao.InitDB()
	e := router.Init()

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      e,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	err := srv.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
