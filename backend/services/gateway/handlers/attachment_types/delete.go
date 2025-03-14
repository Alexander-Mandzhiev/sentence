package attachment_types_handle

import (
	"backend/pkg/server/respond"
	"backend/protos/gen/go/attachment_types"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"strconv"
)

func (h *Handler) delete(c *gin.Context) {
	const op = "attachment_types.Handler.delete"
	log := h.logger.With(slog.String("op", op))

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Error("Invalid ID", slog.String("id", idStr), slog.Any("error", err))
		c.JSON(http.StatusBadRequest, respond.ErrorResponse("invalid ID"))
		return
	}

	resp, err := h.service.Delete(c.Request.Context(), &attachment_types.DeleteAttachmentTypeRequest{Id: int32(id)})
	if err != nil {
		log.Error("Failed to delete attachment type", slog.Any("error", err))
		c.JSON(http.StatusInternalServerError, respond.ErrorResponse("internal server error"))
		return
	}

	log.Info("Attachment type deleted successfully", slog.Any("response", resp))
	c.JSON(http.StatusOK, respond.SuccessResponse(resp))
}
