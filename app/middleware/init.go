package middleware

import "github.com/gin-gonic/gin"

func InitMiddleware(r *gin.Engine) {
	//r.Use(NoCache)
	//r.Use(Options)
	//r.Use(Secure)
	// r.Use(Logger())
	r.Use(CORS())
}
