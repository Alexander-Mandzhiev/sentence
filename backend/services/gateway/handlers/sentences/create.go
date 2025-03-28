package sentences_handle

import (
	"backend/pkg/server/respond"
	"backend/services/gateway/models"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

func (h *Handler) create(c *gin.Context) {
	const op = "sentences.Handler.create"
	log := h.logger.With(slog.String("op", op))

	var req models.CreateSentenceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error("Failed to bind JSON", slog.Any("error", err))
		c.JSON(http.StatusBadRequest, respond.ErrorResponse("invalid request body"))
		return
	}

	resp, err := h.service.Create(c.Request.Context(), &req)
	if err != nil {
		log.Error("Failed to create sentence", slog.Any("error", err))
		c.JSON(http.StatusInternalServerError, respond.ErrorResponse("internal server error"))
		return
	}

	log.Info("Sentence created successfully", slog.Any("response", resp))
	c.JSON(http.StatusCreated, respond.SuccessResponse(resp))
}
