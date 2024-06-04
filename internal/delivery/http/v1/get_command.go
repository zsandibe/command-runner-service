package v1

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetCommandById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	command, err := h.service
}
