package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) HandleInsertNote(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{
		"result": "hello world",
	})
}
