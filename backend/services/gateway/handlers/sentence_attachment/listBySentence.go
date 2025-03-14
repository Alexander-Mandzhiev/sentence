package sentence_attachment_handle

import (
	"backend/pkg/server/respond"
	"backend/protos/gen/go/sentences_attachments"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"strconv"
)

func (h *Handler) listBySentence(c *gin.Context) {
	const op = "sentence_attachment_handle.Handler.listBySentence"
	log := h.logger.With(slog.String("op", op))

	sentenceIDStr := c.Param("sentence_id")
	sentenceID, err := strconv.ParseInt(sentenceIDStr, 10, 64)
	if err != nil {
		log.Error("Failed to parse sentence_id", slog.String("sentence_id", sentenceIDStr), slog.Any("error", err))
		c.JSON(http.StatusBadRequest, respond.ErrorResponse("invalid sentence_id format: must be an integer"))
		return
	}

	resp, err := h.service.ListBySentence(c.Request.Context(), &sentences_attachments.ListBySentenceRequest{SentenceId: sentenceID})
	if err != nil {
		log.Error("Failed to list sentence attachments by sentence", slog.Any("error", err))
		c.JSON(http.StatusInternalServerError, respond.ErrorResponse("internal server error"))
		return
	}

	log.Info("Sentence attachments listed by sentence successfully", slog.Any("response", resp))
	c.JSON(http.StatusOK, respond.SuccessResponse(resp))
}
