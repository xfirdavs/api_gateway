package services

import (
	"fmt"

	"github.com/xfirdavs/api_gateway/config"
	"github.com/xfirdavs/api_gateway/genproto/company_service"
	"github.com/xfirdavs/api_gateway/genproto/position_service"
	"google.golang.org/grpc"
)

type ServiceManager interface {
	ProfessionService() position_service.ProfessionServiceClient
	AttributeService() position_service.AttributeServiceClient
	PositionService() position_service.PositionServiceClient
	CompanyService() company_service.CompanyServiceClient
}

type grpcClients struct {
	atrributeService  position_service.AttributeServiceClient
	professionService position_service.ProfessionServiceClient
	positionService   position_service.PositionServiceClient
	companyService    company_service.CompanyServiceClient
}

// CompanyService implements ServiceManager

func NewGrpcClients(conf *config.Config) (ServiceManager, error) {
	connPositionService, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.PositionServiceHost, conf.PositionServicePort),
		grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	connCompanyService, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.CompanyServiceHost, conf.CompanyServicePort), grpc.WithInsecure())

	if err != nil {
		return nil, err
	}

	return &grpcClients{
		atrributeService:  position_service.NewAttributeServiceClient(connPositionService),
		professionService: position_service.NewProfessionServiceClient(connPositionService),
		positionService:   position_service.NewPositionServiceClient(connPositionService),
		companyService:    company_service.NewCompanyServiceClient(connCompanyService),
	}, nil
}

func (g *grpcClients) ProfessionService() position_service.ProfessionServiceClient {
	return g.professionService
}

func (g *grpcClients) CompanyService() company_service.CompanyServiceClient {
	return g.companyService
}

func (g *grpcClients) AttributeService() position_service.AttributeServiceClient {
	return g.atrributeService
}
func (g *grpcClients) PositionService() position_service.PositionServiceClient {
	return g.positionService
}
