package departments_handle

import (
	"backend/pkg/server/respond"
	"backend/protos/gen/go/departments"
	"backend/services/gateway/models"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

func (h *Handler) create(c *gin.Context) {
	const op = "departments.Handler.create"
	log := h.logger.With(slog.String("op", op))

	var req departments.CreateDepartmentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Error("Failed to bind JSON", slog.Any("error", err))
		c.JSON(http.StatusBadRequest, respond.ErrorResponse("invalid request body"))
		return
	}

	resp, err := h.service.Create(c.Request.Context(), &req)
	if err != nil {
		log.Error("Failed to create department", slog.Any("error", err))
		c.JSON(http.StatusInternalServerError, respond.ErrorResponse("internal server error"))
		return
	}

	department := models.DepartmentFromProto(resp)
	log.Info("Department created successfully", slog.Any("response", department))
	c.JSON(http.StatusOK, respond.SuccessResponse(department))
}
