package tests

import (
	"fmt"
	"net/http"

	"github.com/stretchr/testify/require"
	"github.com/yusrilsabir22/orderfaz/logistic-svc/pkg/pb"
)

func (s *SSuite) Test_0CreateLogisticSuccess() {
	out, err := s.client.CreateLogistic(s.ctx, &pb.CreateLogisticRequest{
		LogisticName:    "jne",
		Amount:          10000,
		DestinationName: "JAKARTA",
		OriginName:      "BANDUNG",
		Duration:        "2-4",
	})
	require.NoError(s.T(), err)
	fmt.Println(out.Id)
	expected := &pb.CreateLogisticResponse{
		Status:  http.StatusCreated,
		Message: "success",
	}

	if expected.Status != out.Status ||
		expected.Message != out.Message {
		s.T().Errorf("Out -> \nWant: %q\nGot : %q", expected, out)
	}
}

func (s *SSuite) Test_1CreateLogisticFailedInvalidParams() {
	out, err := s.client.CreateLogistic(s.ctx, &pb.CreateLogisticRequest{
		LogisticName:    "",
		Amount:          0,
		DestinationName: "",
		OriginName:      "",
		Duration:        "",
	})
	require.NoError(s.T(), err)
	expected := &pb.CreateLogisticResponse{
		Status:  http.StatusBadRequest,
		Message: "invalid parameter",
	}

	if expected.Status != out.Status ||
		expected.Message != out.Message {
		s.T().Errorf("Out -> \nWant: %q\nGot : %q", expected, out)
	}
}
