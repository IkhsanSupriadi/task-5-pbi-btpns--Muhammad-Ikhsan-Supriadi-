package tests

import (
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/yusrilsabir22/orderfaz/auth-svc/pkg/pb"
)

func (s *SSuite) Test_8ValidateSuccess() {
	res, err := s.client.Login(s.ctx, &pb.LoginRequest{
		Msisdn:   "621234567890",
		Password: "test",
	})
	require.NoError(s.T(), err)
	out, err := s.client.Validate(s.ctx, &pb.ValidateRequest{
		Token: res.Token,
	})
	assert.NoError(s.T(), err)
	expected := &pb.ValidateResponse{
		Status:  http.StatusOK,
		Message: "success",
		Data:    out.Data,
	}

	if expected.Status != out.Status ||
		expected.Message != out.Message ||
		expected.Data != out.Data {
		s.T().Errorf("Out -> \nWant: %q\nGot : %q", expected, out)
	}
}

func (s *SSuite) Test_9ValidateFailedToken() {
	out, err := s.client.Validate(s.ctx, &pb.ValidateRequest{
		Token: "asdasd",
	})
	assert.NoError(s.T(), err)
	expected := &pb.ValidateResponse{
		Status:  http.StatusBadRequest,
		Message: "token contains an invalid number of segments",
		Data:    out.Data,
	}

	if expected.Status != out.Status ||
		expected.Message != out.Message ||
		expected.Data != out.Data {
		s.T().Errorf("Out -> \nWant: %q\nGot : %q", expected, out)
	}
}

func (s *SSuite) Test_10ValidateFailedTokenExpired() {
	type jwtClaims struct {
		jwt.StandardClaims
		Id uuid.UUID `json:"userId"`
	}
	claims := &jwtClaims{
		Id: uuid.New(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().AddDate(0, 0, -1).Unix(),
			Issuer:    s.server.Jwt.Issuer,
		},
	}

	c := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := c.SignedString([]byte(s.server.Jwt.SecretKey))
	require.NoError(s.T(), err)
	out, err := s.client.Validate(s.ctx, &pb.ValidateRequest{
		Token: t,
	})

	assert.NoError(s.T(), err)
	expected := &pb.ValidateResponse{
		Status:  http.StatusBadRequest,
		Message: "token is expired by",
		Data:    out.Data,
	}

	if expected.Status != out.Status ||
		!strings.Contains(out.Message, expected.Message) ||
		expected.Data != out.Data {
		s.T().Errorf("Out -> \nWant: %q\nGot : %q", expected, out)
	}
}
