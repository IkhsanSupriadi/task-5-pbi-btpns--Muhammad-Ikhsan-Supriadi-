package tests

import (
	"net/http"

	"github.com/stretchr/testify/require"
	"github.com/yusrilsabir22/orderfaz/logistic-svc/pkg/pb"
)

func (s *SSuite) Test_2FindLogisticSuccess() {
	out, err := s.client.FindOne(s.ctx, &pb.FindOneRequest{
		OriginName:      "BANDUNG",
		DestinationName: "JAKARTA",
	})
	require.NoError(s.T(), err)
	expected := &pb.FindOneResponse{
		Status:  http.StatusOK,
		Message: "success",
	}

	if expected.Status != out.Status ||
		expected.Message != out.Message {
		s.T().Errorf("Out -> \nWant: %q\nGot : %q", expected, out)
	}
}

func (s *SSuite) Test_3FindLogisticFailedNotFound() {
	out, err := s.client.FindOne(s.ctx, &pb.FindOneRequest{
		OriginName:      "origin",
		DestinationName: "destination",
	})
	require.NoError(s.T(), err)
	expected := &pb.FindOneResponse{
		Status:  http.StatusNotFound,
		Message: "record not found",
	}

	if expected.Status != out.Status ||
		expected.Message != out.Message {
		s.T().Errorf("Out -> \nWant: %q\nGot : %q", expected, out)
	}
}
