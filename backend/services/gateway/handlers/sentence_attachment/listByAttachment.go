package sentence_attachment_handle

import (
	"backend/pkg/server/respond"
	"backend/protos/gen/go/sentences_attachments"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"strconv"
)

func (h *Handler) listByAttachment(c *gin.Context) {
	const op = "sentence_attachment_handle.Handler.listByAttachment"
	log := h.logger.With(slog.String("op", op))

	attachmentIDStr := c.Param("attachment_id")
	attachmentID, err := strconv.ParseInt(attachmentIDStr, 10, 32)
	if err != nil {
		log.Error("Failed to parse attachment_id", slog.String("attachment_id", attachmentIDStr), slog.Any("error", err))
		c.JSON(http.StatusBadRequest, respond.ErrorResponse("invalid attachment_id format: must be an integer"))
		return
	}

	resp, err := h.service.ListByAttachment(c.Request.Context(), &sentences_attachments.ListByAttachmentRequest{AttachmentId: int32(attachmentID)})
	if err != nil {
		log.Error("Failed to list sentence attachments by attachment", slog.Any("error", err))
		c.JSON(http.StatusInternalServerError, respond.ErrorResponse("internal server error"))
		return
	}

	log.Info("Sentence attachments listed by attachment successfully", slog.Any("response", resp))
	c.JSON(http.StatusOK, respond.SuccessResponse(resp))
}
