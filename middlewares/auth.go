package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tyange/white-shadow-api/utils"
)

func Authenticate(context *gin.Context) {
	token, err := context.Cookie("session")

	if token == "" || err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized."})
		return
	}

	userId, err := utils.VerifyToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not authorized."})
		return
	}

	// 동일한 요청에 대해서 특정 데이터를 만들어 사용하는 것이 가능함.
	context.Set("userId", userId)
	context.Next()
}
