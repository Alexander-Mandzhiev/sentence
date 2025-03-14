package history_handle

import (
	"backend/pkg/server/respond"
	"backend/protos/gen/go/history"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"strconv"
)

func (h *Handler) listBySentence(c *gin.Context) {
	const op = "history.Handler.listBySentence"
	log := h.logger.With(slog.String("op", op))

	sentenceIDStr := c.Param("sentence_id")
	sentenceID, err := strconv.Atoi(sentenceIDStr)
	if err != nil {
		log.Error("Failed to convert sentence ID to integer", slog.String("sentence_id", sentenceIDStr), slog.Any("error", err))
		c.JSON(http.StatusBadRequest, respond.ErrorResponse("invalid sentence ID format: ID must be an integer"))
		return
	}

	resp, err := h.service.ListBySentence(c.Request.Context(), &history.ListBySentenceRequest{SentenceId: int64(sentenceID)})
	if err != nil {
		log.Error("Failed to list history records by sentence", slog.Any("error", err))
		c.JSON(http.StatusInternalServerError, respond.ErrorResponse("internal server error"))
		return
	}

	log.Info("History records listed by sentence successfully", slog.Any("response", resp))
	c.JSON(http.StatusOK, respond.SuccessResponse(resp))
}
