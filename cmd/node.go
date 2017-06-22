package cmd

import (
	"github.com/Sirupsen/logrus"
	"github.com/pritunl/pritunl-zero/constants"
	"github.com/pritunl/pritunl-zero/handlers"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func ManagementNode() {
	if constants.Production {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	router := gin.New()

	if !constants.Production {
		router.Use(gin.Logger())
	}

	handlers.Register(router)

	server := &http.Server{
		Addr:           "0.0.0.0:8443",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 4096,
	}

	logrus.WithFields(logrus.Fields{
		"production": constants.Production,
	}).Info("cmd.app: Starting management node")

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
