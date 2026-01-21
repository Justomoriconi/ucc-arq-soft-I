package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
	Domain "ucc-arq-soft-I/domain/user"
	Service "ucc-arq-soft-I/services/user"
)

func CreateUser(c *gin.Context) {
	var req Domain.User

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := Service.CreateUser(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusCreated, gin.H{"user": user})
}
