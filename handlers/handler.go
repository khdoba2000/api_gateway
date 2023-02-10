package handlers

import (
	"qurilish/api_gateway/configs"
	authcontrollers "qurilish/api_gateway/controllers/auth"
	"qurilish/api_gateway/gateways"
	// "qurilish/api_gateway/pkg/logger"
)

type Handler struct {
	cfg configs.Config
	// log      logger.LoggerI
	gateways       gateways.Gateway
	authController authcontrollers.AuthController
}

func NewHandler(cfg configs.Config,
	// log logger.LoggerI,
	gateways gateways.Gateway) Handler {
	return Handler{
		cfg: cfg,
		// log:      log,
		gateways: gateways,
	}
}
