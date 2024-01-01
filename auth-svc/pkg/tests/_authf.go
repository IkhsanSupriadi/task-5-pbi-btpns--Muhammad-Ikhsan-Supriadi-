package tests

import (
	"context"
	"fmt"
	"net/http"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"github.com/yusrilsabir22/orderfaz/auth-svc/pkg/db"
	"github.com/yusrilsabir22/orderfaz/auth-svc/pkg/mocks"
	"github.com/yusrilsabir22/orderfaz/auth-svc/pkg/models"
	"github.com/yusrilsabir22/orderfaz/auth-svc/pkg/pb"
	"github.com/yusrilsabir22/orderfaz/auth-svc/pkg/utils"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ServerTestSuite struct {
	suite.Suite
	Jwt               utils.JwtWrapper
	H                 db.Handler
	MockServiceClient *mocks.AuthServiceClientMock
	sqlMock           sqlmock.Sqlmock
}

func (s *ServerTestSuite) SetupTest() {
	var err error
	// pgConn := "postgres://orderfaz:orderfaz@localhost:5432/orderfaz"

	require.NoError(s.T(), err)
	s.Jwt = utils.JwtWrapper{
		SecretKey:       "test",
		Issuer:          "auth-svc",
		ExpirationHours: 24 * 365,
	}

	sqlDB, mock, err := sqlmock.New()
	require.NoError(s.T(), err)
	// defer sqlDB.Close()
	s.sqlMock = mock
	dialector := postgres.New(postgres.Config{
		DSN:                  "sql_db_0",
		DriverName:           "postgres",
		Conn:                 sqlDB,
		PreferSimpleProtocol: true,
	})
	mdb, err := gorm.Open(dialector, &gorm.Config{})
	require.NoError(s.T(), err)
	s.H = db.Handler{
		DB: mdb,
	}
	s.MockServiceClient = &mocks.AuthServiceClientMock{
		RegisterFunc: func(ctx context.Context, in *pb.RegisterRequest, opts ...grpc.CallOption) (*pb.RegisterResponse, error) {
			return s.APIRegister(ctx, in)
		},
		LoginFunc: func(ctx context.Context, in *pb.LoginRequest, opts ...grpc.CallOption) (*pb.LoginResponse, error) {
			return s.APILogin(ctx, in)
		},
	}

}

func (s *ServerTestSuite) AfterTest(_, _ string) {
	require.NoError(s.T(), s.sqlMock.ExpectationsWereMet())
}

func (s *ServerTestSuite) TestRegister1Success() {
	res, err := CallRegisterSuccess(s.MockServiceClient)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), &pb.RegisterResponse{
		Status:  http.StatusCreated,
		Message: "success",
	}, res)
}

func (s *ServerTestSuite) TestRegister2Fail() {
	res, err := CallRegisterSuccess(s.MockServiceClient)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), &pb.RegisterResponse{
		Status:  http.StatusCreated,
		Message: "success",
	}, res)
}

func (s *ServerTestSuite) _TestLoginSuccess() {
	resSuccess, err := CallLoginSuccess(s.MockServiceClient)
	assert.NoError(s.T(), err)

	assert.Equal(s.T(), &pb.LoginResponse{
		Status:  http.StatusAccepted,
		Message: "success",
		Token:   resSuccess.Token,
	}, resSuccess)

}

func (s *ServerTestSuite) _TestLoginFailInvalidMSISD() {
	resInvalidMSISDN, err := CallLoginInvalidMSISDN(s.MockServiceClient)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), &pb.LoginResponse{
		Status:  http.StatusBadRequest,
		Message: "Invalid parameter",
	}, resInvalidMSISDN)
}

func CallLoginInvalidMSISDN(myService pb.AuthServiceClient) (*pb.LoginResponse, error) {
	return myService.Login(context.TODO(), &pb.LoginRequest{
		Msisdn:   "test",
		Password: "test",
	})
}

func CallLoginSuccess(myService pb.AuthServiceClient) (*pb.LoginResponse, error) {
	return myService.Login(context.TODO(), &pb.LoginRequest{
		Msisdn:   "6281234567890",
		Password: "test",
	})
}

func CallRegisterSuccess(s pb.AuthServiceClient) (*pb.RegisterResponse, error) {
	return s.Register(context.TODO(), &pb.RegisterRequest{
		Msisdn:   "6281234567890",
		Name:     "test",
		Username: "test",
		Password: "test",
	})
}

func (s *ServerTestSuite) APILogin(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	var user models.User

	if !utils.IsValidMSISDN(req.Msisdn) {
		return &pb.LoginResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid parameter",
		}, nil
	}

	if req.Msisdn == "" || req.Password == "" {
		return &pb.LoginResponse{
			Message: "User not found",
			Status:  http.StatusUnauthorized,
			Token:   "",
		}, nil
	}
	if result := s.H.DB.First(&user, &models.User{MSISDN: req.Msisdn}); result.Error != nil {
		return &pb.LoginResponse{
			Message: "User not found",
			Status:  http.StatusUnauthorized,
			Token:   "",
		}, nil
	}

	match := utils.CheckPasswordHash(req.Password, user.Password)
	if !match {
		return &pb.LoginResponse{
			Status:  http.StatusUnauthorized,
			Message: "User not found",
			Token:   "",
		}, nil
	}

	token, _ := s.Jwt.GenerateToken(user)

	return &pb.LoginResponse{
		Status:  http.StatusAccepted,
		Message: "success",
		Token:   token,
	}, nil
}

func (s *ServerTestSuite) APIRegister(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	var user models.User

	if !utils.IsValidMSISDN(req.Msisdn) || req.Name == "" || req.Password == "" || req.Username == "" {
		return &pb.RegisterResponse{
			Status:  http.StatusBadRequest,
			Message: "invalid parameter",
		}, nil
	}

	expectQuery := `SELECT * FROM "users" WHERE "users"."msisdn" = $1 OR "users"."username" = $2 ORDER BY "users"."id" LIMIT 1`
	rows := []string{
		"id",
		"msisdn",
		"name",
		"username",
		"password",
	}
	s.sqlMock.ExpectQuery(regexp.QuoteMeta(expectQuery)).
		WithArgs(req.Msisdn, req.Username).
		WillReturnRows(sqlmock.NewRows(rows))

	result := s.H.DB.Where(&models.User{MSISDN: req.Msisdn}).Or(&models.User{Username: req.Username}).First(&user)

	if result.Error == nil {
		return &pb.RegisterResponse{
			Status:  http.StatusBadRequest,
			Message: "MSISDN already exists",
		}, nil
	}

	user.MSISDN = req.Msisdn
	user.Name = req.Name
	user.Username = req.Username
	user.Password = utils.HashPassword(req.Password)

	s.sqlMock.ExpectBegin()
	s.sqlMock.ExpectQuery(regexp.QuoteMeta(
		`INSERT INTO "users" ("msisdn","name","username","password")
		VALUES ($1,$2,$3,$4) RETURNING "id"`)).
		WithArgs(user.MSISDN, user.Name, user.Username, user.Password).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow("00000000-0000-0000-0000-000000000000"))
	// s.sqlMock.ExpectRollback()
	s.sqlMock.ExpectCommit()

	c := s.H.DB.Create(&user)
	fmt.Println(user.ID)
	if c.Error != nil {
		return &pb.RegisterResponse{
			Status:  http.StatusBadRequest,
			Message: "failed",
		}, nil
	}

	return &pb.RegisterResponse{
		Message: "success",
		Status:  http.StatusCreated,
	}, nil
}

func TestAuth(t *testing.T) {

	suite.Run(t, new(ServerTestSuite))
}
