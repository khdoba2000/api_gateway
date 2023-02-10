package handlers

import (
	"qurilish/api_gateway/configs"
	"qurilish/api_gateway/gateways"
	// "qurilish/api_gateway/pkg/logger"
)

type Handler struct {
	cfg configs.Config
	// log      logger.LoggerI
	gateways gateways.Gateway
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
