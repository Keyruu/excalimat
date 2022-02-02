package sessions

import (
	"time"

	"github.com/gofiber/fiber/v2/middleware/session"
)

var Store = session.New(session.Config{
	Expiration: 1 * time.Hour,
})
