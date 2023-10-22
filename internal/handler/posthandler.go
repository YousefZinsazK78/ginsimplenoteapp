package handler

import (
	"fmt"
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
	fmt.Println(post)
	//todo : add uploaded image url
	imgUrl, err := saveImage(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"result": err})
		return
	}
	post.ImgUrl = imgUrl
	//todo : added uploaded image url
	err = h.postStorer.InsertPost(post)
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
	if err := c.SaveUploadedFile(file, "./public/images/"+newFilename); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to save the file",
		})
		return
	}

	userM, ok := c.Get("userModel")
	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "unable to load user model",
		})
		return
	}
	userModel := userM.(models.User)
	image := models.Image{ImageUrl: "localhost:8000/public/images/" + newFilename, PostID: 1, UserID: userModel.ID}
	if err := h.postStorer.InsertImage(image); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to save in iamge database",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": "image upload successfully."})

}

func saveImage(c *gin.Context) (string, error) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"result": err})
		return "", err
	}
	log.Println(file.Filename)
	extFile := filepath.Ext(file.Filename)
	newFilename := uuid.New().String() + extFile

	if err := c.SaveUploadedFile(file, "./public/images/"+newFilename); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": "Unable to save the file",
		})
		return "", err
	}
	return fmt.Sprint("http://localhost:8000/public/images/" + newFilename), nil
}
