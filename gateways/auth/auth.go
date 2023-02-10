package auth

import (
	"context"
	// auth_service "qurilish/api_gateway/genproto/auth_service"
)

type AuthGateway interface {
	Login(ctx context.Context) error
}
type authGateway struct {
	// sessionService auth_service.SessionServiceClient
}

func Init() AuthGateway {

	// cfg := configs.Load()
	// connAuthService, err := grpc.Dial(
	// 	cfg.AuthServiceHost+cfg.AuthGRPCPort,
	// 	grpc.WithInsecure(),
	// )
	// if err != nil {
	// 	panic(err)
	// }

	// sessionService := auth_service.NewSessionService(connAuthService)

	return authGateway{
		// sessionService: sessionService,
	}
}

func (g authGateway) Login(ctx context.Context) error {
	// r:=mappers.MapLoginReqToProto(req)
	// g.sessionService.Login(ctx, r)

	return nil
}
