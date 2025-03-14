package attachment_types_handle

import (
	"backend/pkg/server/respond"
	"backend/protos/gen/go/attachment_types"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"strconv"
)

func (h *Handler) get(c *gin.Context) {
	const op = "attachment_types.Handler.get"
	log := h.logger.With(slog.String("op", op))

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Error("Failed to convert ID to integer", slog.String("id", idStr), slog.Any("error", err))
		c.JSON(http.StatusBadRequest, respond.ErrorResponse("invalid ID format: ID must be an integer"))
		return
	}

	resp, err := h.service.Get(c.Request.Context(), &attachment_types.GetAttachmentTypeRequest{Id: int32(id)})
	if err != nil {
		log.Error("Failed to get attachment type", slog.Any("error", err))
		c.JSON(http.StatusInternalServerError, respond.ErrorResponse("internal server error"))
		return
	}

	log.Info("Attachment type retrieved successfully", slog.Any("response", resp))
	c.JSON(http.StatusOK, respond.SuccessResponse(resp))
}
