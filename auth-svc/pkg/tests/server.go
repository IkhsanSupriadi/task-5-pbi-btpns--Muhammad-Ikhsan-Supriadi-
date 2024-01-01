package tests

import (
	"context"
	"fmt"
	"net"

	"github.com/yusrilsabir22/orderfaz/auth-svc/pkg/config"
	"github.com/yusrilsabir22/orderfaz/auth-svc/pkg/db"
	"github.com/yusrilsabir22/orderfaz/auth-svc/pkg/pb"
	"github.com/yusrilsabir22/orderfaz/auth-svc/pkg/services"
	"github.com/yusrilsabir22/orderfaz/auth-svc/pkg/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

func InitTestServer(ctx context.Context) (pb.AuthServiceClient, services.Server) {
	buffer := 101024 * 1024
	lis := bufconn.Listen(buffer)

	baseServer := grpc.NewServer()
	initClientServer, err := NewServer()
	// require.NoError(s.T(), err)
	if err != nil {
		fmt.Println(err)
	}
	pb.RegisterAuthServiceServer(baseServer, initClientServer)
	go func() {
		if err := baseServer.Serve(lis); err != nil {
			fmt.Printf("error serving server: %v", err)
		}
	}()

	conn, err := grpc.DialContext(ctx, "",
		grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
			return lis.Dial()
		}), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("error connecting to server: %v", err)
	}

	// closer := func() {
	// 	err := lis.Close()
	// 	if err != nil {
	// 		fmt.Printf("error closing listener: %v", err)
	// 	}
	// 	baseServer.Stop()
	// }

	client := pb.NewAuthServiceClient(conn)

	return client, *initClientServer
}

func NewServer() (*services.Server, error) {
	c, err := config.LoadConfig("../config/envs")
	if err != nil {
		return nil, err
	}
	// Setup DB
	h := db.Init(c.DBTestUrl)

	// Setup JWT
	jwt := utils.JwtWrapper{
		SecretKey:       c.JWTSecretKey,
		Issuer:          "auth-svc",
		ExpirationHours: 24 * 365,
	}
	t := &services.Server{
		H:   h,
		Jwt: jwt,
	}
	return t, nil
}
