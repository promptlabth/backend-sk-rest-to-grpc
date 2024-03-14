package calculator

import (
	"github.com/gin-gonic/gin"
)

type caluclatorService struct {
	calculatorClient CalculatorClient
}

func NewCalculatorService(calculatorClient CalculatorClient) *caluclatorService {
	return &caluclatorService{
		calculatorClient: calculatorClient,
	}
}

func (base *caluclatorService) Hello(c *gin.Context) {

	req := &CalculatorRequest{}
	ctx := c.Request.Context()
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(404, map[string]string{
			"err": err.Error(),
		})
	}
	resp, err := base.calculatorClient.Hello(ctx, &HelloRequest{
		Name:    req.Name,
		Surname: req.Surname,
	})
	if err != nil {
		c.JSON(404, map[string]string{
			"err": err.Error(),
		})
	}
	c.JSON(200, map[string]string{
		"message": resp.Message,
	})

}
