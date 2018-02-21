package main

import (
	"errors"
	"fmt"

	pb "github.com/giancarlopetrini/microservices/vessel-service/proto/vessel"
	"github.com/micro/go-micro"
	"golang.org/x/net/context"
)

// Repository - Interface process methods
type Repository interface {
	FindAvailable(*pb.Specification) (*pb.Vessel, error)
}

// VesselRepository - Store available vessels
type VesselRepository struct {
	vessels []*pb.Vessel
}

// FindAvailable - checks a specification against a map of vessels,
// if capacity and max weight are below a vessels capacity and max weight,
// then return that vessel.
func (repo *VesselRepository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error) {
	for _, vessel := range repo.vessels {
		if spec.Capacity <= vessel.Capacity && spec.MaxWeight <= vessel.MaxWeight {
			return vessel, nil
		}
	}
	return nil, errors.New("No vessels found that match the specs")
}

// grpc service handler
type service struct {
	repo Repository
}

func (s *service) FindAvailable(ctx context.Context, req *pb.Specification, res *pb.Response) error {
	//Find available vessels
	vessel, err := s.repo.FindAvailable(req)
	if err != nil {
		return err
	}

	//Set the vessel as part of response message type
	res.Vessel = vessel
	return nil
}

func main() {
	// temp vessel data for testing...
	vessels := []*pb.Vessel{
		&pb.Vessel{Id: "vessel001", Name: "Mister Nice Vessel", MaxWeight: 10000, Capacity: 5000},
	}
	repo := &VesselRepository{vessels}

	srv := micro.NewService(
		micro.Name("go.micro.srv.vessel"),
		micro.Version("latest"),
	)

	srv.Init()

	// Register implementation handler
	pb.RegisterVesselServiceHandler(srv.Server(), &service{repo})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
