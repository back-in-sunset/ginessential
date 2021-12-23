package main

import (
	"gin-essential/inject"
	"net/http"
	"time"
)

func main() {
	// 初始化依赖注入器
	injector, injectorCleanFunc, err := inject.GenInjector()
	if err != nil {
		panic(err)
	}
	srv := &http.Server{
		Addr:         ":8080",
		Handler:      injector.Engine,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	err = srv.ListenAndServe()
	if err != nil {
		panic(err)
	}
	injectorCleanFunc()
}
