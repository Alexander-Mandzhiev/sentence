package sentence_attachment_handle

import (
	"backend/pkg/server/respond"
	"backend/protos/gen/go/sentences_attachments"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"strconv"
)

func (h *Handler) delete(c *gin.Context) {
	const op = "sentence_attachment_handle.Handler.delete"
	log := h.logger.With(slog.String("op", op))

	sentenceIDStr := c.Query("sentence_id")
	attachmentIDStr := c.Query("attachment_id")

	sentenceID, err := strconv.ParseInt(sentenceIDStr, 10, 64)
	if err != nil {
		log.Error("Failed to parse sentence_id", slog.String("sentence_id", sentenceIDStr), slog.Any("error", err))
		c.JSON(http.StatusBadRequest, respond.ErrorResponse("invalid sentence_id format: must be an integer"))
		return
	}

	attachmentID, err := strconv.ParseInt(attachmentIDStr, 10, 32)
	if err != nil {
		log.Error("Failed to parse attachment_id", slog.String("attachment_id", attachmentIDStr), slog.Any("error", err))
		c.JSON(http.StatusBadRequest, respond.ErrorResponse("invalid attachment_id format: must be an integer"))
		return
	}

	req := &sentences_attachments.DeleteSentenceAttachmentRequest{
		SentenceId:   sentenceID,
		AttachmentId: int32(attachmentID),
	}

	resp, err := h.service.Delete(c.Request.Context(), req)
	if err != nil {
		log.Error("Failed to delete sentence attachment", slog.Any("error", err))
		c.JSON(http.StatusInternalServerError, respond.ErrorResponse("internal server error"))
		return
	}

	log.Info("Sentence attachment deleted successfully", slog.Any("response", resp))
	c.JSON(http.StatusOK, respond.SuccessResponse(resp))
}
