package handler

import (
	"github.com/gin-gonic/gin"
	"github/todo-app"
	"net/http"
)

type getAllListsResponse struct {
	Data []todo.Segment `json:"data"`
}

func (h *Handler) updateUserSegments(c *gin.Context) {
	var input todo.A

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid input body")
		return
	}

	err := h.services.UsersSegments.UpdateUserSegments(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllListsResponse{
		Data: nil,
	})
}

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
