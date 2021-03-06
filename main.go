package main

import (
	"github.com/labstack/echo/v4"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	router := echo.New()
	router.GET("/",
		func(ctx echo.Context) error {
			i := time.Duration(rand.Intn(20))
			time.Sleep(i * time.Millisecond)
			return ctx.String(http.StatusOK, "hello")
		},
	)
	log.Fatal(router.Start(":80"))
}
