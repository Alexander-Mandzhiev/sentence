package statuses_handle

import (
	"backend/pkg/server/respond"
	"backend/protos/gen/go/statuses"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"strconv"
)

func (h *Handler) update(c *gin.Context) {
	const op = "statuses.Handler.update"
	log := h.logger.With(slog.String("op", op))

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Error("Failed to convert ID to integer", slog.String("id", idStr), slog.Any("error", err))
		c.JSON(http.StatusBadRequest, respond.ErrorResponse("invalid ID format: ID must be an integer"))
		return
	}

	var req statuses.UpdateStatusRequest
	if err = c.ShouldBindJSON(&req); err != nil {
		log.Error("Failed to bind JSON", slog.Any("error", err))
		c.JSON(http.StatusBadRequest, respond.ErrorResponse("invalid request body"))
		return
	}

	req.Id = int32(id)
	resp, err := h.service.Update(c.Request.Context(), &req)
	if err != nil {
		log.Error("Failed to update status", slog.Any("error", err))
		c.JSON(http.StatusInternalServerError, respond.ErrorResponse("internal server error"))
		return
	}

	log.Info("Status updated successfully", slog.Any("response", resp))
	c.JSON(http.StatusOK, respond.SuccessResponse(resp))
}
