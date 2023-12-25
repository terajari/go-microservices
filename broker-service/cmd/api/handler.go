package main

import (
	"bytes"
	"encoding/json"
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

type RequestPayload struct {
	Action string `json:"action"`
	Auth   Auth   `json:"auth"`
}

type Auth struct {
	Email    string `json:"email"`
	Password string `json:"password"`
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

func (app *Application) HandleSubmission(ctx *gin.Context) {

	var payloadReq RequestPayload
	ctx.Header("Content-Type", "application/json")
	if err := ctx.ShouldBindJSON(&payloadReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	switch payloadReq.Action {
	case "auth":
		app.authenticateUser(ctx, payloadReq.Auth)
	default:
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid action"})
	}
}

func (app *Application) authenticateUser(ctx *gin.Context, a Auth) {

	jsonData, _ := json.MarshalIndent(a, "", "\t")

	req, err := http.NewRequest(http.MethodPost, "http://auth-service/auth", bytes.NewBuffer(jsonData))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusUnauthorized {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	} else if resp.StatusCode != http.StatusAccepted {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error calling auth service"})
		return
	}

	var jsonAuth response
	err = json.NewDecoder(resp.Body).Decode(&jsonAuth)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	payloadResponse := response{
		Message: "Authenticated",
		Time:    time.Now(),
		Error:   false,
		Data:    jsonAuth.Data,
	}

	ctx.JSON(http.StatusOK, payloadResponse)
}
