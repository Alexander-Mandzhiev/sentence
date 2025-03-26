package attachment_handle

import (
	"backend/pkg/server/respond"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"strconv"
)

func (h *Handler) delete(c *gin.Context) {
	const op = "attachments.Handler.delete"
	log := h.logger.With(slog.String("op", op))

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Error("Invalid attachment ID", slog.String("id", idStr), slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, respond.ErrorResponse("invalid attachment ID"))
		return
	}

	err = h.service.DeleteAttachment(c.Request.Context(), int32(id))
	if err != nil {
		log.Error("Failed to delete attachment", slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, respond.ErrorResponse("failed to delete attachment"))
		return
	}

	log.Info("Attachment deleted successfully")
	c.JSON(http.StatusNoContent, nil)
}
