package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type response struct {
	Message string    `json:"message"`
	Time    time.Time `json:"time"`
	Error   bool      `json:"error"`
	Data    any       `json:"data"`
}

func (app *Application) Broker(ctx *gin.Context) {
	ctx.Header("Content-Type", "application/json")

	ctx.JSON(http.StatusAccepted, response{
		Message: "Broker",
		Time:    time.Now(),
		Error:   false,
		Data:    nil,
	})
}
