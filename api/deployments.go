package api

import (
	"fmt"
	"time"

	"golang.org/x/net/context"

	"github.com/astronomerio/commander/pkg/proto"
)

func (s *GRPCServer) Ping(ctx context.Context, in *proto.PingRequest) (*proto.PingResponse, error) {
	return &proto.PingResponse{
		Received: time.Now().UnixNano() / int64(time.Millisecond),
	}, nil
}

func (s *GRPCServer) GetDeployment(ctx context.Context, in *proto.GetDeploymentRequest) (*proto.GetDeploymentResponse, error) {
	return &proto.GetDeploymentResponse{}, nil
}

func (s *GRPCServer) CreateDeployment(ctx context.Context, in *proto.CreateDeploymentRequest) (*proto.CreateDeploymentResponse, error) {
	response := s.provisioner.InstallDeployment(in)
	fmt.Println("CreateDeployment called")
	return response, nil
}

func (s *GRPCServer) UpdateDeployment(ctx context.Context, in *proto.UpdateDeploymentRequest) (*proto.UpdateDeploymentResponse, error) {
	response := s.provisioner.UpdateDeployment(in)
	fmt.Println("UpdateDeployment called")
	return response, nil
}

func (s *GRPCServer) UpgradeDeployment(ctx context.Context, in *proto.UpgradeDeploymentRequest) (*proto.UpgradeDeploymentResponse, error) {
	fmt.Println("UpgradeDeployment called")
	return &proto.UpgradeDeploymentResponse{}, nil
}

func (s *GRPCServer) DeleteDeployment(ctx context.Context, in *proto.DeleteDeploymentRequest) (*proto.DeleteDeploymentResponse, error) {
	fmt.Println("DeleteDeployment called")
	return &proto.DeleteDeploymentResponse{}, nil
}