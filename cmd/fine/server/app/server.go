package app

import (
	"context"
	fineV1Pb "github.com/DABronskikh/ago-6/pkg/fine/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

type Server struct {
	businessSvc *Service
}

func NewServer(businessSvc *Service) *Server {
	return &Server{businessSvc}
}

func (s *Server) FindByAll(ctx context.Context, request *fineV1Pb.AutopaysRequest) (*fineV1Pb.AutopayResponse, error) {
	data, err := s.businessSvc.GetAutopayUser(ctx, request.UserId)
	if err != nil {
		log.Print(err)
		return nil, status.Errorf(codes.NotFound, err.Error())
	}
	return data, nil
}

func (s *Server) FindById(ctx context.Context, request *fineV1Pb.AutopayRequestById) (*fineV1Pb.Autopay, error) {
	data, err := s.businessSvc.GetAutopayById(ctx, request.AutopayId)
	if err != nil {
		log.Print(err)
		return nil, status.Errorf(codes.NotFound, err.Error())
	}
	return data, nil
}

func (s *Server) Create(ctx context.Context, request *fineV1Pb.Autopay) (*fineV1Pb.Autopay, error) {
	data, err := s.businessSvc.SaveAutopay(ctx, request)
	if err != nil {
		log.Print(err)
		return nil, status.Errorf(codes.NotFound, err.Error())
	}
	return data, nil
}

func (s *Server) Update(ctx context.Context, request *fineV1Pb.Autopay) (*fineV1Pb.Autopay, error) {
	data, err := s.businessSvc.SaveAutopay(ctx, request)
	if err != nil {
		log.Print(err)
		return nil, status.Errorf(codes.NotFound, err.Error())
	}
	return data, nil
}

func (s *Server) Delete(ctx context.Context, request *fineV1Pb.Autopay) (*fineV1Pb.Autopay, error) {
	data, err := s.businessSvc.DeleteAutopay(ctx, request)
	if err != nil {
		log.Print(err)
		return nil, status.Errorf(codes.NotFound, err.Error())
	}
	return data, nil
}
