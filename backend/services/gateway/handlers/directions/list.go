package directions_handle

import (
	"backend/pkg/server/respond"
	"backend/protos/gen/go/directions"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

func (h *Handler) list(c *gin.Context) {
	const op = "directions.Handler.list"
	log := h.logger.With(slog.String("op", op))

	resp, err := h.client.List(c.Request.Context(), &directions.ListDirectionsRequest{})
	if err != nil {
		log.Error("Failed to list directions", slog.Any("error", err))
		c.JSON(http.StatusInternalServerError, respond.ErrorResponse("internal server error"))
		return
	}

	log.Info("Directions listed successfully", slog.Any("response", resp))
	c.JSON(http.StatusOK, respond.SuccessResponse(resp))
}
