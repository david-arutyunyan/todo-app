package handler

import (
	"github.com/gin-gonic/gin"
	"github/todo-app"
	"net/http"
)

// @Summary UpdateUserSegments
// @Tags user_segments
// @Description update user's segments
// @ID update-user-segments
// @Accept  json
// @Produce  json
// @Param input body todo.AlteredUserSegments true "user_id and segments to add/delete"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /users-segments [post]
func (h *Handler) updateUserSegments(c *gin.Context) {
	var input todo.AlteredUserSegments

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	err := h.services.UsersSegments.UpdateUserSegments(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

// @Summary GetUserSegments
// @Tags user_segments
// @Description get user's segments
// @ID get-user-segments
// @Accept  json
// @Produce  json
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /users-segments/:id [GET]
func (h *Handler) getUserSegments(c *gin.Context) {
	id := c.Param("id")

	lists, err := h.services.UsersSegments.GetUserSegments(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllListsResponse{
		Data: lists,
	})
}
