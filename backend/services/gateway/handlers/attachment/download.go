package attachment_handle

import (
	"github.com/gin-gonic/gin"
	"io"
	"log/slog"
	"net/http"
	"strconv"
)

func (h *Handler) download(c *gin.Context) {
	const op = "attachments.Handler.download"
	log := h.logger.With(slog.String("op", op))

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Error("Invalid attachment ID", slog.String("id", idStr), slog.String("error", err.Error()))
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid attachment ID"})
		return
	}

	file, metadata, err := h.service.DownloadFile(c.Request.Context(), int32(id))
	if err != nil {
		log.Error("Failed to download file", slog.String("error", err.Error()))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to download file"})
		return
	}
	defer file.Close()

	c.Header("Content-Disposition", "attachment; filename="+metadata.FileName)
	c.Header("Content-Type", metadata.MimeType)
	c.Header("Content-Length", strconv.FormatInt(metadata.FileSize, 10))

	if _, err = io.Copy(c.Writer, file); err != nil {
		log.Error("Failed to send file", slog.String("error", err.Error()))
	}
}
