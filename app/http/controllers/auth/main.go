package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/go-monsters/sulivan/resources/views"
)

func Login(c *gin.Context) {
	view := views.NewView("guest", "auth/login")
	view.Render(c.Writer, nil)
}
