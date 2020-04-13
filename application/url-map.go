package app

import (
	"fmt"

	"github.com/gyan1230/go-k8s/controller"
)

func mapurls() {
	fmt.Println("url mapping...")
	router.GET("/home", controller.Home)
	router.GET("/readiness", controller.ReadinessHandler)
	router.GET("/health", controller.HealthHandler)

}
