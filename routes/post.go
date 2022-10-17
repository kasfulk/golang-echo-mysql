package routes

import (
	"github.com/kasfulk/golang-echo-mysql/handlers"
)

func PostRoutes() {
	// post route
	e.GET("/post", handlers.PostIndex)
	e.GET("/post/:id", handlers.PostDetail)
}
