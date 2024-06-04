package v1

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zsandibe/command-runner-service/internal/domain"
	"github.com/zsandibe/command-runner-service/pkg"
)

func (h *Handler) CreateCommand(c *gin.Context) {
	var inp domain.CreateCommandInput

	if err := c.BindJSON(&inp); err != nil {
		errorResponse(c, http.StatusBadRequest, fmt.Errorf("invalid request body: %v", err))
		return
	}

	if err := pkg.ValidateCommandRequest(inp.Script, inp.Description); err != nil {
		errorResponse(c, http.StatusBadRequest, err)
		return
	}

	command, err := h.Service.CreateCommand(c, &inp)
	if err != nil {
		errorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, command)
}
