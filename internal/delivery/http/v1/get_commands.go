package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetCommands(c *gin.Context) {
	ids := c.QueryArray("id")

	res, err := h.Service.GetCommands(c, ids)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, res)
}
