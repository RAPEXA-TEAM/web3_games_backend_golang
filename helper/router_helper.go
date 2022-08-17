package helper

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SendSuccessResponse(c *gin.Context, res any, err any) {
	c.Writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		c.Writer.WriteHeader(http.StatusOK)
		c.JSON(http.StatusOK, gin.H{"error": "unknown error"})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": res})
	}
}
func SendErrorResponse(c *gin.Context, err string) {
	c.Writer.Header().Set("Content-Type", "application/json")
	c.Writer.WriteHeader(http.StatusOK)
	if ErrorIsInErrorList(err) {
		c.JSON(http.StatusOK, gin.H{"error": err})
	} else {
		c.JSON(http.StatusOK, gin.H{"error": "unknown error please try again"})
	}
}
