package main

import (
	"github.com/figoxu/gh"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Env struct {
	fh *gh.FormHelper
	ph *gh.ParamHelper
}

func main() {
	r := gin.Default()
	r.GET("/:domain/hello", m_gh, h_domain_hello)
	r.Run(":3456")
}

//http://localhost:3456/figo/hello
func h_domain_hello(c *gin.Context) {
	env := c.MustGet("env").(*Env)
	ph := env.ph
	domain := ph.String("domain")
	c.String(http.StatusOK, "Hello %s", domain)
}

func m_gh(c *gin.Context) {
	c.Set("env", &Env{
		fh: gh.NewFormHelper(c),
		ph: gh.NewParamHelper(c),
	})
	c.Next()
}
