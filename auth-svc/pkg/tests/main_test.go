package tests

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/yusrilsabir22/orderfaz/auth-svc/pkg/pb"
	"github.com/yusrilsabir22/orderfaz/auth-svc/pkg/services"
)

type SSuite struct {
	suite.Suite
	client pb.AuthServiceClient
	ctx    context.Context
	server services.Server
}

func (s *SSuite) SetupSuite() {
	ctx := context.Background()
	s.ctx = ctx
	client, server := InitTestServer(ctx)

	s.client = client
	s.server = server
	s.server.H.DB.Exec(`DELETE FROM "users"`)
}

func (s *SSuite) TearDownAllSuite() {
	s.server.H.DB.Exec(`DELETE FROM "users"`)
	fmt.Println("db empty users table")
}

func TestInit(t *testing.T) {
	suite.Run(t, new(SSuite))
}
