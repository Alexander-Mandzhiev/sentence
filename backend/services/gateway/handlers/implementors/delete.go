package implementors_handle

import (
	"backend/pkg/server/respond"
	"backend/protos/gen/go/implementors"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"strconv"
)

func (h *Handler) delete(c *gin.Context) {
	const op = "implementors.Handler.delete"
	log := h.logger.With(slog.String("op", op))

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Error("Failed to convert ID to integer", slog.String("id", idStr), slog.Any("error", err))
		c.JSON(http.StatusBadRequest, respond.ErrorResponse("invalid ID format: ID must be an integer"))
		return
	}

	resp, err := h.service.Delete(c.Request.Context(), &implementors.DeleteImplementorRequest{Id: int32(id)})
	if err != nil {
		log.Error("Failed to delete implementor", slog.Any("error", err))
		c.JSON(http.StatusInternalServerError, respond.ErrorResponse("internal server error"))
		return
	}

	log.Info("Implementor deleted successfully", slog.Any("response", resp))
	c.JSON(http.StatusOK, resp)
}
