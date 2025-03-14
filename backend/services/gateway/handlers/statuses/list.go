package statuses_handle

import (
	"backend/pkg/server/respond"
	"backend/protos/gen/go/statuses"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

func (h *Handler) list(c *gin.Context) {
	const op = "statuses.Handler.list"
	log := h.logger.With(slog.String("op", op))

	resp, err := h.service.List(c.Request.Context(), &statuses.ListStatusesRequest{})
	if err != nil {
		log.Error("Failed to list statuses", slog.Any("error", err))
		c.JSON(http.StatusInternalServerError, respond.ErrorResponse("internal server error"))
		return
	}

	log.Info("Statuses listed successfully", slog.Any("response", resp))
	c.JSON(http.StatusOK, respond.SuccessResponse(resp))
}
