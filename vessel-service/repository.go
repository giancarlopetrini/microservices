package main

import (
	pb "github.com/giancarlopetrini/microservices/vessel-service/proto/vessel"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	dbName           = "shippy"
	vesselCollection = "vessels"
)

// Repository - interface for vessels
type Repository interface {
	FindAvailable(*pb.Specification) (*pb.Vessel, error)
	Create(vessel *pb.Vessel) error
	Close()
}

// VesselRepository - mongo sessions
type VesselRepository struct {
	session *mgo.Session
}

// FindAvailable - checks a specification against a map of vessels,
// if capacity and max weight are below a vessels capacity and max weight,
// then return that vessel.
func (repo *VesselRepository) FindAvailable(spec *pb.Specification) (*pb.Vessel, error) {
	var vessel *pb.Vessel

	// Here we define a more complex query than our consignment-service's
	// GetAll function. Here we're asking for a vessel who's max weight and
	// capacity are greater than and equal to the given capacity and weight.
	// We're also using the `One` function here as that's all we want.
	err := repo.collection().Find(bson.M{
		"capacity":  bson.M{"$gte": spec.Capacity},
		"maxweight": bson.M{"gte": spec.MaxWeight},
	}).One(&vessel)
	if err != nil {
		return nil, err
	}
	return vessel, nil
}

// Create - insert vessel to collection
func (repo *VesselRepository) Create(vessel *pb.Vessel) error {
	return repo.collection().Insert(vessel)
}

// Close - close db collection
func (repo *VesselRepository) Close() {
	repo.session.Close()
}

func (repo *VesselRepository) collection() *mgo.Collection {
	return repo.session.DB(dbName).C(vesselCollection)
}
