package routes

import (
	"github.com/kasfulk/golang-echo-mysql/controllers"
)

func PostRoutes() {
	// post route
	e.GET("/post", controllers.PostIndex)
}
