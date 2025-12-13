package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Envelope struct {
	Success       bool        `json:"success"`
	Data          interface{} `json:"data,omitempty"`
	Error         string      `json:"error,omitempty"`
	ErrStatusCode int         `json:"err_status_code,omitempty"`
	RequestID     string      `json:"request_id,omitempty"`
}

func Success(c *gin.Context, data interface{}, requestID string) {
	c.JSON(http.StatusOK, Envelope{
		Success:   true,
		Data:      data,
		RequestID: requestID,
	})
}

func Error(c *gin.Context, statusCode int, errMsg string, requestID string) {
	c.JSON(statusCode, Envelope{
		Success:       false,
		Error:         errMsg,
		ErrStatusCode: statusCode,
		RequestID:     requestID,
	})
}
