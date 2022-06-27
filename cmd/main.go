package main

import (
	"context"
	"gin-essential/inject"
	"gin-essential/repo/dao"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	httpAddr = ":8080"
)

// HTTPConfig http config
type HTTPConfig struct {
	Addr string `toml:"addr"`
}

type options struct {
	httpConfig *HTTPConfig
}

// Option 定义配置项
type Option func(*options)

// HTTPConfigOpt 设定配置文件
func HTTPConfigOpt(c *HTTPConfig) Option {
	return func(o *options) {
		o.httpConfig = c
	}
}

// NewHTTPConfig ..
func NewHTTPConfig() *HTTPConfig {
	c := &HTTPConfig{
		Addr: httpAddr,
	}

	return c
}

func main() {
	Run(context.Background())
}

// Run 运行服务
func Run(ctx context.Context) error {
	state := 1
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	cleanFunc := Init(ctx, HTTPConfigOpt(NewHTTPConfig()))

EXIT:
	for {
		sig := <-sc
		switch sig {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			state = 0
			break EXIT
		case syscall.SIGHUP:
		default:
			break EXIT
		}
	}

	cleanFunc()
	time.Sleep(time.Second)
	os.Exit(state)
	return nil
}

// Init 初始化启动
func Init(ctx context.Context, opts ...Option) func() {
	var o options
	for _, opt := range opts {
		opt(&o)
	}

	// 初始化依赖注入器
	injector, injectorCleanFunc, err := inject.GenInjector()
	if err != nil {
		panic(err)
	}
	srv := &http.Server{
		Addr:         o.httpConfig.Addr,
		Handler:      injector.Engine,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	go srv.ListenAndServe()

	return injectorCleanFunc
}

// Config 配置文件
type Config struct {
	Addr     string
	Postgres dao.Postgres
}

// initConfig
func initConfig() {

}
