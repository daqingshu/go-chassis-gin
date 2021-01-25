package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// RestFulHello RestFulHello
type RestFulHello struct {
}

// Root Root
func (r *RestFulHello) Hello(c *gin.Context) {
	c.Writer.WriteString(fmt.Sprintf("hello %s", c.Request.RemoteAddr))
}

//URLPatterns helps to respond for corresponding API calls
func (r *RestFulHello) URLPatterns(router *gin.RouterGroup) {
	router.GET("/hello", r.Hello)
}
