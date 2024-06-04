package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAllCommands(c *gin.Context) {
	res, err := h.Service.GetAllCommands(c)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, res)
}
