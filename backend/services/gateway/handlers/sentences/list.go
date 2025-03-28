package sentences_handle

import (
	"backend/pkg/server/respond"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"strconv"
)

func (h *Handler) list(c *gin.Context) {
	const op = "sentences.Handler.list"
	log := h.logger.With(slog.String("op", op))

	limitStr := c.DefaultQuery("limit", "10")
	offsetStr := c.DefaultQuery("offset", "0")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		log.Error("Failed to convert limit to integer", slog.String("limit", limitStr), slog.Any("error", err))
		c.JSON(http.StatusBadRequest, respond.ErrorResponse("invalid limit format: must be an integer"))
		return
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		log.Error("Failed to convert offset to integer", slog.String("offset", offsetStr), slog.Any("error", err))
		c.JSON(http.StatusBadRequest, respond.ErrorResponse("invalid offset format: must be an integer"))
		return
	}

	data, total, err := h.service.List(c.Request.Context(), int32(limit), int32(offset))
	if err != nil {
		log.Error("Failed to list sentences", slog.Any("error", err))
		c.JSON(http.StatusInternalServerError, respond.ErrorResponse("internal server error"))
		return
	}

	log.Info("Sentences listed successfully", slog.Int("count", len(data)))
	c.JSON(http.StatusOK, respond.SuccessResponse(gin.H{
		"data":   data,
		"total":  total,
		"limit":  limit,
		"offset": offset,
	}))
}
