package attachment_types

import (
	"backend/pkg/server/respond"
	"backend/protos/gen/go/attachment_types"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

func (h *Handler) list(c *gin.Context) {
	const op = "attachmentTypes.Handler.list"
	log := h.logger.With(slog.String("op", op))

	var req attachment_types.ListAttachmentTypesRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		log.Error("Failed to bind query parameters", slog.Any("error", err))
		c.JSON(http.StatusBadRequest, respond.ErrorResponse("invalid query parameters"))
		return
	}

	resp, err := h.service.List(c.Request.Context(), &req)
	if err != nil {
		log.Error("Failed to list attachment types", slog.Any("error", err))
		c.JSON(http.StatusInternalServerError, respond.ErrorResponse("internal server error"))
		return
	}

	log.Info("Attachment types listed successfully", slog.Any("response", resp))
	c.JSON(http.StatusOK, respond.SuccessResponse(resp))
}
