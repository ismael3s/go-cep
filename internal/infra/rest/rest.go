package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	usecases "github.com/ismael3s/go-cep/internal/application/usecases"
	"github.com/ismael3s/go-cep/internal/infra/di"
)

type IWebServer interface {
	Run(...string) error
	ServeHTTP(w http.ResponseWriter, req *http.Request)
}

func SetupRestServer() IWebServer {
	r := gin.Default()
	r.GET("/api/v1/cep/:cep", func(c *gin.Context) {
		cep := c.Param("cep")
		useCase := di.FindAddressByCEPDI()
		output, err := useCase.Do(usecases.FindAddressByCEPInput{Value: cep})
		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(200, output.Address)
	})
	return r
}
