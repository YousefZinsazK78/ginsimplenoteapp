package handler

import (
	"log"
	"net/http"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

func (h *handler) HandleGetPosts(c *gin.Context) {
	posts, err := h.postStorer.GetPosts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"result": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": posts})
}

func (h *handler) HandleGetPostByTitle(c *gin.Context) {
	title := c.Param("title")
	post, err := h.postStorer.GetPostByTitle(title)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"result": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": post})
}

func (h *handler) HandleDeletePost(c *gin.Context) {
	//get note id
	strPostID := c.Param("id")
	postID, err := strconv.Atoi(strPostID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"result": err})
		return
	}
	err = h.postStorer.DeletePost(postID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"result": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": "delete successfully."})
}

func (h *handler) HandlePutPost(c *gin.Context) {
	//bind post model
	var post models.Post
	if err := c.ShouldBind(&post); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"result": err})
		return
	}
	err := h.postStorer.UpdatePost(post)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"result": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": "update successfully."})
}

func (h *handler) HandleUpload(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"result": err})
		return
	}
	log.Println(file.Filename)
	extFile := filepath.Ext(file.Filename)
	newFilename := uuid.New().String() + extFile
	if err := c.SaveUploadedFile(file, "./public"+newFilename); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to save the file",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": "image upload successfully."})

}
