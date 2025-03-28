package attachment_handle

import (
	"backend/pkg/server/respond"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"strconv"
)

func (h *Handler) list(c *gin.Context) {
	const op = "attachments.Handler.list"
	log := h.logger.With(slog.String("op", op))

	limitStr := c.DefaultQuery("limit", "50")
	offsetStr := c.DefaultQuery("offset", "0")

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 || limit > 100 {
		limit = 50
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil || offset < 0 {
		offset = 0
	}

	resp, err := h.service.List(c.Request.Context(), int32(limit), int32(offset))
	if err != nil {
		log.Error("Failed to list attachments", slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, respond.ErrorResponse("failed to list attachments"))
		return
	}

	log.Info("Attachments listed successfully")
	c.JSON(http.StatusOK, respond.SuccessResponse(resp))
}
