package directions_handle

import (
	"backend/pkg/server/respond"
	"backend/protos/gen/go/directions"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

func (h *Handler) create(c *gin.Context) {
	const op = "directions.Handler.create"
	log := h.logger.With(slog.String("op", op))

	var req directions.CreateDirectionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error("Failed to bind JSON", slog.Any("error", err))
		c.JSON(http.StatusBadRequest, respond.ErrorResponse("invalid request body"))
		return
	}

	resp, err := h.client.Create(c.Request.Context(), &req)
	if err != nil {
		log.Error("Failed to create direction", slog.Any("error", err))
		c.JSON(http.StatusInternalServerError, respond.ErrorResponse("internal server error"))
		return
	}

	log.Info("Direction created successfully", slog.Any("response", resp))
	c.JSON(http.StatusOK, respond.SuccessResponse(resp))
}
