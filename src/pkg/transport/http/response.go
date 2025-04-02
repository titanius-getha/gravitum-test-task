package transport

import "github.com/gin-gonic/gin"

func BadResponse(reason string) gin.H {
	return gin.H{"result": false, "error": reason}
}

func GoodResponse(data any) gin.H {
	return gin.H{"result": true, "data": data}
}
