package middlewares

import (
	"fmt"
	"net/http"

	"example.com/rest-api/utils"
	"github.com/gin-gonic/gin"
)

func Authorization(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized request!"})
		return
	}
	userid, err := utils.VerifyToken(token)
	if err != nil {
		fmt.Println(token)

		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "bad Json", "error": err})
		return
	}

	context.Set("userid", userid)

	context.Next()
}
