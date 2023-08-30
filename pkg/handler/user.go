package handler

import (
	"github.com/gin-gonic/gin"
	"github/todo-app"
	"net/http"
)

// @Summary CreateUser
// @Tags user
// @Description create a new user
// @ID create-user
// @Accept  json
// @Produce  json
// @Param input body todo.User true "user info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /user/createUser [post]
func (h *Handler) createUser(c *gin.Context) {
	var input todo.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	id, err := h.services.User.CreateUser(input) // request user
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{ //response user
		"id": id,
	})
}

// @Summary DeleteUser
// @Tags user
// @Description delete user by id
// @ID delete-user
// @Accept  json
// @Produce  json
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /user/deleteUser/:id [delete]
func (h *Handler) deleteUser(c *gin.Context) {
	id := c.Param("id")

	err := h.services.User.DeleteUser(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
