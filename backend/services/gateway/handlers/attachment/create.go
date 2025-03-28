package attachment_handle

import (
	"backend/pkg/server/respond"
	"backend/protos/gen/go/attachments"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"strconv"
)

func (h *Handler) create(c *gin.Context) {
	const op = "attachments.Handler.create"
	log := h.logger.With(slog.String("op", op))

	file, header, err := c.Request.FormFile("file")
	if err != nil {
		log.Error("Failed to get file from form", slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, respond.ErrorResponse("file is required"))
		return
	}
	defer file.Close()

	attachmentTypeID, _ := strconv.Atoi(c.PostForm("attachment_type_id"))
	metadata := &attachments.AttachmentMetadata{
		AttachmentTypeId: int32(attachmentTypeID),
		FileName:         header.Filename,
		MimeType:         header.Header.Get("Content-Type"),
	}

	attachment, err := h.service.Create(c.Request.Context(), metadata, file)
	if err != nil {
		log.Error("Failed to create attachment", slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, respond.ErrorResponse("failed to create attachment"))
		return
	}

	log.Info("Attachment created successfully", slog.Any("attachment", attachment))
	c.JSON(http.StatusCreated, respond.SuccessResponse(attachment))
}
