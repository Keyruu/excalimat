package handler

import (
	"time"

	"github.com/gofiber/fiber/v2/middleware/session"
)

var sessionStore = session.New(session.Config{
	Expiration: 1 * time.Hour,
})
