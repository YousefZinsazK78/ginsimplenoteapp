package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yousefzinsazk78/simple_note_api/internal/models"
)

func (h *handler) HandleInsertPost(c *gin.Context) {
	//bind post models
	var post models.Post
	if err := c.ShouldBind(&post); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"result": err})
		return
	}
	err := h.postStorer.InsertPost(post)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"result": err})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"result": "insert successfully."})
}
