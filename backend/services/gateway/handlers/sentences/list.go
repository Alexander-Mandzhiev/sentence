package sentences_handle

import (
	"backend/pkg/server/respond"
	"backend/protos/gen/go/sentences"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

func (h *Handler) list(c *gin.Context) {
	const op = "sentences.Handler.list"
	log := h.logger.With(slog.String("op", op))

	resp, err := h.service.List(c.Request.Context(), &sentences.ListSentencesRequest{})
	if err != nil {
		log.Error("Failed to list sentences", slog.Any("error", err))
		c.JSON(http.StatusInternalServerError, respond.ErrorResponse("internal server error"))
		return
	}

	log.Info("Sentences listed successfully", slog.Any("response", resp))
	c.JSON(http.StatusOK, respond.SuccessResponse(resp))
}
