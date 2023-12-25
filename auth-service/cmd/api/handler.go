package main

import (
	"fmt"
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

func (app *Application) AuthenticateUser(ctx *gin.Context) {
	var payloadReq struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := ctx.ShouldBindJSON(&payloadReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := app.Models.User.GetByEmail(payloadReq.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	valid, err := user.PasswordMatches(payloadReq.Password)
	if err != nil || !valid {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	payloadResp := response{
		Message: fmt.Sprintf("Welcome %s!", user.Email),
		Time:    time.Now(),
		Error:   false,
		Data:    user,
	}

	ctx.JSON(http.StatusAccepted, payloadResp)
}
