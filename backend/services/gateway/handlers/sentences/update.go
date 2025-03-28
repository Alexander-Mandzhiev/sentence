package sentences_handle

import (
	"backend/pkg/server/respond"
	"backend/protos/gen/go/sentences"
	"backend/services/gateway/models"
	"github.com/gin-gonic/gin"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log/slog"
	"net/http"
	"strconv"
	"time"
)

// sentences_handle/handler.go
func (h *Handler) update(c *gin.Context) {
	const op = "sentences.Handler.update"
	log := h.logger.With(slog.String("op", op))

	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		log.Error("Failed to convert ID to integer", slog.String("id", idStr), slog.Any("error", err))
		c.JSON(http.StatusBadRequest, respond.ErrorResponse("invalid ID format: ID must be an integer"))
		return
	}

	var req struct {
		StatusID      int32  `json:"status_id"`
		Name          string `json:"name"`
		Phone         string `json:"phone"`
		DepartmentID  int32  `json:"department_id"`
		Problem       string `json:"problem"`
		Solution      string `json:"solution"`
		CreatedAt     string `json:"created_at"`
		DirectionID   int32  `json:"direction_id"`
		ImplementorID int32  `json:"implementor_id"`
		PriorityID    int32  `json:"priority_id"`
		Encouragement int32  `json:"encouragement"`
		Reason        string `json:"reason"`
		Details       string `json:"details"`
	}

	if err = c.ShouldBindJSON(&req); err != nil {
		log.Error("Failed to bind JSON", slog.Any("error", err))
		c.JSON(http.StatusBadRequest, respond.ErrorResponse("invalid request body"))
		return
	}

	var createdAt *timestamppb.Timestamp
	if req.CreatedAt != "" {
		var t time.Time
		t, err = time.Parse(time.RFC3339, req.CreatedAt)
		if err != nil {
			log.Error("Invalid created_at format", slog.Any("error", err))
			c.JSON(http.StatusBadRequest, respond.ErrorResponse("invalid created_at format, expected RFC3339"))
			return
		}
		createdAt = timestamppb.New(t)
	}

	sentence := &models.UpdateSentenceRequest{
		SentenceResponse: &sentences.SentenceResponse{
			Id:            id,
			StatusId:      req.StatusID,
			Name:          req.Name,
			Phone:         req.Phone,
			DepartmentId:  req.DepartmentID,
			Problem:       req.Problem,
			Solution:      req.Solution,
			CreatedAt:     createdAt,
			DirectionId:   req.DirectionID,
			ImplementorId: req.ImplementorID,
			PriorityId:    req.PriorityID,
			Encouragement: req.Encouragement,
		},
		Reason:  req.Reason,
		Details: req.Details,
	}

	resp, err := h.service.Update(c.Request.Context(), sentence)
	if err != nil {
		log.Error("Failed to update sentence", slog.Any("error", err))
		c.JSON(http.StatusInternalServerError, respond.ErrorResponse("internal server error"))
		return
	}

	log.Info("Sentence updated successfully", slog.Any("response", resp))
	c.JSON(http.StatusOK, respond.SuccessResponse(resp))
}
