package interfaces

import "github.com/gin-gonic/gin"

type IHealthcheckHandler interface {
	HealthcheckHandlerHTTP(c *gin.Context)
}

type IHealthcheckServices interface {
	HealthcheckServices() (string, error)
}

type IHealthcheckRepo interface {
}
