package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/yousefzinsazk78/simple_note_api/internal/models"
)

func (h *handler) CreateRole(c *gin.Context) {
	var Role models.Role
	c.BindJSON(&Role)
	err := h.roleStorer.InsertRole(Role)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, Role)
}

// get Roles
func (h *handler) GetRoles(c *gin.Context) {
	var Role []models.Role
	Role, err := h.roleStorer.GetRoles()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, Role)
}

// get Role by id
func (h *handler) GetRole(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var Role *models.Role
	Role, err := h.roleStorer.GetRoleByID(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, Role)
}

// update Role
func (h *handler) UpdateRole(c *gin.Context) {
	var Role models.Role
	c.BindJSON(&Role)
	err := h.roleStorer.UpdateRole(Role)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, Role)
}
