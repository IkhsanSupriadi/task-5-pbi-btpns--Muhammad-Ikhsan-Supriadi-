package tests

import (
	"net/http"

	"github.com/stretchr/testify/assert"
	"github.com/yusrilsabir22/orderfaz/auth-svc/pkg/pb"
)

func (s *SSuite) Test_4LoginSuccess() {
	out, err := s.client.Login(s.ctx, &pb.LoginRequest{
		Msisdn:   "621234567890",
		Password: "test",
	})
	assert.NoError(s.T(), err)
	expected := &pb.LoginResponse{
		Status:  http.StatusAccepted,
		Message: "success",
		Token:   out.Token,
	}

	if expected.Status != out.Status ||
		expected.Message != out.Message ||
		expected.Token != out.Token {
		s.T().Errorf("Out -> \nWant: %q\nGot : %q", expected, out)
	}
}

func (s *SSuite) Test_5LoginFailedInvalidMSISDN() {
	out, err := s.client.Login(s.ctx, &pb.LoginRequest{
		Msisdn:   "test",
		Password: "test",
	})
	assert.NoError(s.T(), err)
	expected := &pb.LoginResponse{
		Status:  http.StatusBadRequest,
		Message: "invalid parameter",
	}

	if expected.Status != out.Status ||
		expected.Message != out.Message {
		s.T().Errorf("Out -> \nWant: %q\nGot : %q", expected, out)
	}
}

func (s *SSuite) Test_6LoginFailedUserNotFound() {
	out, err := s.client.Login(s.ctx, &pb.LoginRequest{
		Msisdn:   "621234567891",
		Password: "test",
	})
	assert.NoError(s.T(), err)
	expected := &pb.LoginResponse{
		Status:  http.StatusUnauthorized,
		Message: "user not found",
	}

	if expected.Status != out.Status ||
		expected.Message != out.Message {
		s.T().Errorf("Out -> \nWant: %q\nGot : %q", expected, out)
	}
}

func (s *SSuite) Test_7LoginFailedWrongPassword() {
	out, err := s.client.Login(s.ctx, &pb.LoginRequest{
		Msisdn:   "621234567890",
		Password: "wrong",
	})
	assert.NoError(s.T(), err)
	expected := &pb.LoginResponse{
		Status:  http.StatusUnauthorized,
		Message: "user not found",
	}

	if expected.Status != out.Status ||
		expected.Message != out.Message {
		s.T().Errorf("Out -> \nWant: %q\nGot : %q", expected, out)
	}
}
