package healthcheck

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthCheckHandler(healthService HealthCheck) func(c *gin.Context) {
	return func(c *gin.Context) {
		_, err := healthService.Health()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"status": "success"})
	}
}

//
// GetHealthCheck return HealthCheck impl
//
func GetHealthCheck(conf Config) HealthCheck {
	return &Impl{}
}
