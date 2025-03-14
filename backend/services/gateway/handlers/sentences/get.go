package sentences_handle

import (
	"backend/pkg/server/respond"
	"backend/protos/gen/go/sentences"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"strconv"
)

func (h *Handler) get(c *gin.Context) {
	const op = "sentences.Handler.get"
	log := h.logger.With(slog.String("op", op))

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Error("Failed to convert ID to integer", slog.String("id", idStr), slog.Any("error", err))
		c.JSON(http.StatusBadRequest, respond.ErrorResponse("invalid ID format: ID must be an integer"))
		return
	}

	resp, err := h.service.Get(c.Request.Context(), &sentences.GetSentenceRequest{Id: int64(id)})
	if err != nil {
		log.Error("Failed to get sentence", slog.Any("error", err))
		c.JSON(http.StatusInternalServerError, respond.ErrorResponse("internal server error"))
		return
	}

	log.Info("Sentence retrieved successfully", slog.Any("response", resp))
	c.JSON(http.StatusOK, respond.SuccessResponse(resp))
}
