package statuses_handle

import (
	"backend/pkg/server/respond"
	"backend/protos/gen/go/statuses"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"strconv"
)

func (h *Handler) delete(c *gin.Context) {
	const op = "statuses.Handler.delete"
	log := h.logger.With(slog.String("op", op))

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Error("Failed to convert ID to integer", slog.String("id", idStr), slog.Any("error", err))
		c.JSON(http.StatusBadRequest, respond.ErrorResponse("invalid ID format: ID must be an integer"))
		return
	}

	resp, err := h.service.Delete(c.Request.Context(), &statuses.DeleteStatusRequest{Id: int32(id)})
	if err != nil {
		log.Error("Failed to delete status", slog.Any("error", err))
		c.JSON(http.StatusInternalServerError, respond.ErrorResponse("internal server error"))
		return
	}

	log.Info("Status deleted successfully", slog.Any("response", resp))
	c.JSON(http.StatusOK, respond.SuccessResponse(resp))
}
