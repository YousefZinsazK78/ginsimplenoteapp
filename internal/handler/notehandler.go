package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yousefzinsazk78/simple_note_api/internal/models"
)

func (h *handler) HandleInsertNote(c *gin.Context) {
	//bind note models
	var note models.Note
	if err := c.ShouldBind(&note); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"result": err})
		return
	}
	err := h.noteStorer.InsertNote(note)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"result": err})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"result": "insert successfully."})
}

func (h *handler) HandleGetNotes(c *gin.Context) {
	notes, err := h.noteStorer.GetNotes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"result": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": notes})
}

func (h *handler) HandleGetNoteByTitle(c *gin.Context) {
	title := c.Param("title")
	note, err := h.noteStorer.GetNotesByTitle(title)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"result": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": note})
}

func (h *handler) HandleDeleteNote(c *gin.Context) {
	//get note id
	strNoteID := c.Param("id")
	noteID, err := strconv.Atoi(strNoteID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"result": err})
		return
	}
	err = h.noteStorer.DeleteNote(noteID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"result": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": "delete successfully."})
}

func (h *handler) HandlePutNote(c *gin.Context) {
	//bind note model
	var note models.Note
	if err := c.ShouldBind(&note); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"result": err})
		return
	}
	err := h.noteStorer.UpdateNote(note)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"result": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"result": "update successfully."})
}
