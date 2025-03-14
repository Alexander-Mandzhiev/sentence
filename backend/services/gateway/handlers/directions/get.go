package directions_handle

import (
	"backend/pkg/server/respond"
	"backend/protos/gen/go/directions"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"strconv"
)

func (h *Handler) get(c *gin.Context) {
	const op = "directions.Handler.get"
	log := h.logger.With(slog.String("op", op))

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Error("Failed to convert ID to integer", slog.String("id", idStr), slog.Any("error", err))
		c.JSON(http.StatusBadRequest, respond.ErrorResponse("invalid ID format: ID must be an integer"))
		return
	}

	resp, err := h.client.Get(c.Request.Context(), &directions.GetDirectionRequest{Id: int32(id)})
	if err != nil {
		log.Error("Failed to get direction", slog.Any("error", err))
		c.JSON(http.StatusInternalServerError, respond.ErrorResponse("internal server error"))
		return
	}

	log.Info("Direction retrieved successfully", slog.Any("response", resp))
	c.JSON(http.StatusOK, respond.SuccessResponse(resp))
}
