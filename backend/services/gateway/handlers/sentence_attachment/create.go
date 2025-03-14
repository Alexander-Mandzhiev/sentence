package sentence_attachment_handle

import (
	"backend/pkg/server/respond"
	"backend/protos/gen/go/sentences_attachments"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

func (h *Handler) create(c *gin.Context) {
	const op = "sentence_attachment_handle.Handler.create"
	log := h.logger.With(slog.String("op", op))

	var req sentences_attachments.CreateSentenceAttachmentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error("Failed to bind JSON", slog.Any("error", err))
		c.JSON(http.StatusBadRequest, respond.ErrorResponse("invalid request body"))
		return
	}

	resp, err := h.service.Create(c.Request.Context(), &req)
	if err != nil {
		log.Error("Failed to create sentence attachment", slog.Any("error", err))
		c.JSON(http.StatusInternalServerError, respond.ErrorResponse("internal server error"))
		return
	}

	log.Info("Sentence attachment created successfully", slog.Any("response", resp))
	c.JSON(http.StatusOK, respond.SuccessResponse(resp))
}
