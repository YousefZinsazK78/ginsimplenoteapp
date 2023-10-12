package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yousefzinsazk78/simple_note_api/internal/models"
)

func (h *handler) HandleInsertComment(c *gin.Context) {
	//bind comment models
	var comment models.Comment
	if err := c.ShouldBind(&comment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"result": err})
		return
	}
	err := h.commentStorer.Insert(comment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"result": err})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"result": "insert successfully."})
}

func (h *handler) HandleGetComments(c *gin.Context) {
	notes, err := h.commentStorer.GetComments()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"result": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": notes})
}

func (h *handler) HandleGetCommentsByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	note, err := h.commentStorer.GetCommentByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"result": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": note})
}

func (h *handler) HandleGetCommentsByBody(c *gin.Context) {
	body := c.Param("id")
	note, err := h.commentStorer.GetCommentByBody(body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"result": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": note})
}

func (h *handler) HandleUpdateComments(c *gin.Context) {
	//bind note model
	commentID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"result": err})
		return
	}
	var updateComment models.UpdateCommentParam
	if err := c.ShouldBind(&updateComment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"result": err})
		return
	}
	err = h.commentStorer.Update(commentID, updateComment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"result": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": "update successfully."})
}

func (h *handler) HandleDeleteComment(c *gin.Context) {
	//get note id
	commentID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"result": err})
		return
	}
	err = h.commentStorer.Delete(commentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"result": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": "delete successfully."})
}
