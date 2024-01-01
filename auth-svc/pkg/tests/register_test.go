package tests

import (
	"net/http"

	"github.com/stretchr/testify/require"
	"github.com/yusrilsabir22/orderfaz/auth-svc/pkg/pb"
)

func (s *SSuite) Test_0RegisterSuccess() {
	out, err := s.client.Register(s.ctx, &pb.RegisterRequest{
		Msisdn:   "621234567890",
		Name:     "test",
		Username: "test",
		Password: "test",
	})
	require.NoError(s.T(), err)
	expected := &pb.RegisterResponse{
		Status:  http.StatusCreated,
		Message: "success",
	}

	if expected.Status != out.Status ||
		expected.Message != out.Message {
		s.T().Errorf("Out -> \nWant: %q\nGot : %q", expected, out)
	}
}

func (s *SSuite) Test_1RegisterInvalidMSISDN() {
	out, err := s.client.Register(s.ctx, &pb.RegisterRequest{
		Msisdn:   "test",
		Name:     "test",
		Username: "test",
		Password: "test",
	})
	expected := &pb.RegisterResponse{
		Status:  http.StatusBadRequest,
		Message: "invalid parameter",
	}
	require.NoError(s.T(), err)
	if expected.Status != out.Status ||
		expected.Message != out.Message {
		s.T().Errorf("Out -> \nWant: %q\nGot : %q", expected, out)
	}
}

func (s *SSuite) Test_2RegisterAlreadyExists() {
	_, err := s.client.Register(s.ctx, &pb.RegisterRequest{
		Msisdn:   "621234567890",
		Name:     "test",
		Username: "test",
		Password: "test",
	})
	require.NoError(s.T(), err)
	out, err := s.client.Register(s.ctx, &pb.RegisterRequest{
		Msisdn:   "621234567890",
		Name:     "test",
		Username: "test",
		Password: "test",
	})
	expected := &pb.RegisterResponse{
		Status:  http.StatusBadRequest,
		Message: "user already registered",
	}
	require.NoError(s.T(), err)
	if expected.Status != out.Status ||
		expected.Message != out.Message {
		s.T().Errorf("Out -> \nWant: %q\nGot : %q", expected, out)
	}
}
