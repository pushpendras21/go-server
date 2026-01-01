package handlers

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pushpendras21/go-server/db"
)

type PostTaskPayload struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Status      string `json:"status"`
}

func SaveTask(ctx *gin.Context) {

	var payload PostTaskPayload
	//err := ctx.ShouldBindJSON(&payload)
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Unable to read the body"})
		return
	}
	if payload.Title == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Title is required"})
		return
	}

	var id int
	query := `INSERT INTO tasks (title, description, status) VALUES ($1, $2, $3) RETURNING id`
	err := db.DB.QueryRow(context.Background(), query, payload.Title, payload.Description, payload.Status).Scan(&id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": true, "msg": err.Error()})
		log.Fatal("Error:", err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"error": false, "msg": id})
}
