package http

import (
	"context"

	"github.com/douglas-franca/hireeasy/internal/modules/sessions/domain"
	"github.com/douglas-franca/hireeasy/internal/platform/httperror"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const CookieName = "hireeasy_session"
const userIDKey = "authenticated_user_id"

type Authenticator interface {
	Execute(context.Context, string) (uuid.UUID, error)
}

func Required(auth Authenticator) gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, err := c.Cookie(CookieName)
		if err != nil {
			httperror.Abort(c, domain.ErrUnauthenticated)
			return
		}
		id, err := auth.Execute(c.Request.Context(), cookie)
		if err != nil {
			httperror.Abort(c, err)
			return
		}
		c.Set(userIDKey, id)
		c.Next()
	}
}
func UserID(c *gin.Context) uuid.UUID {
	value, _ := c.Get(userIDKey)
	id, _ := value.(uuid.UUID)
	return id
}
