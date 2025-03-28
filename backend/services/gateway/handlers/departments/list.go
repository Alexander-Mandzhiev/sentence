package departments_handle

import (
	"backend/pkg/server/respond"
	"backend/protos/gen/go/departments"
	"backend/services/gateway/models"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
)

func (h *Handler) list(c *gin.Context) {
	const op = "departments.Handler.list"
	log := h.logger.With(slog.String("op", op))

	resp, err := h.service.List(c.Request.Context(), &departments.ListDepartmentsRequest{})
	if err != nil {
		log.Error("Failed to list departments", slog.Any("error", err))
		c.JSON(http.StatusInternalServerError, respond.ErrorResponse("internal server error"))
		return
	}

	deps := make([]*models.Department, 0, len(resp.Data))
	for _, pbDept := range resp.Data {
		deps = append(deps, models.DepartmentFromProto(pbDept))
	}

	log.Info("Departments listed successfully", slog.Any("response", deps))
	c.JSON(http.StatusOK, respond.SuccessResponse(deps))
}
