// Code generated by protoc-gen-go. DO NOT EDIT.
// source: consignment.proto

/*
Package go_micro_srv_consignment is a generated protocol buffer package.

It is generated from these files:
	consignment.proto

It has these top-level messages:
	Consignment
	Container
	GetRequest
	Response
*/
package go_micro_srv_consignment

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Consignment struct {
	Id          string       `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Description string       `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	Weight      int32        `protobuf:"varint,3,opt,name=weight" json:"weight,omitempty"`
	Containers  []*Container `protobuf:"bytes,4,rep,name=containers" json:"containers,omitempty"`
	VesselId    string       `protobuf:"bytes,5,opt,name=vessel_id,json=vesselId" json:"vessel_id,omitempty"`
}

func (m *Consignment) Reset()                    { *m = Consignment{} }
func (m *Consignment) String() string            { return proto.CompactTextString(m) }
func (*Consignment) ProtoMessage()               {}
func (*Consignment) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Consignment) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Consignment) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Consignment) GetWeight() int32 {
	if m != nil {
		return m.Weight
	}
	return 0
}

func (m *Consignment) GetContainers() []*Container {
	if m != nil {
		return m.Containers
	}
	return nil
}

func (m *Consignment) GetVesselId() string {
	if m != nil {
		return m.VesselId
	}
	return ""
}

type Container struct {
	Id         string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	CustomerId string `protobuf:"bytes,2,opt,name=customer_id,json=customerId" json:"customer_id,omitempty"`
	Origin     string `protobuf:"bytes,3,opt,name=origin" json:"origin,omitempty"`
	UserId     string `protobuf:"bytes,4,opt,name=user_id,json=userId" json:"user_id,omitempty"`
}

func (m *Container) Reset()                    { *m = Container{} }
func (m *Container) String() string            { return proto.CompactTextString(m) }
func (*Container) ProtoMessage()               {}
func (*Container) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Container) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Container) GetCustomerId() string {
	if m != nil {
		return m.CustomerId
	}
	return ""
}

func (m *Container) GetOrigin() string {
	if m != nil {
		return m.Origin
	}
	return ""
}

func (m *Container) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type GetRequest struct {
}

func (m *GetRequest) Reset()                    { *m = GetRequest{} }
func (m *GetRequest) String() string            { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()               {}
func (*GetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type Response struct {
	Created     bool         `protobuf:"varint,1,opt,name=created" json:"created,omitempty"`
	Consignment *Consignment `protobuf:"bytes,2,opt,name=consignment" json:"consignment,omitempty"`
	// Added a pluralised consignment to our generic response message
	Consignments []*Consignment `protobuf:"bytes,3,rep,name=consignments" json:"consignments,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Response) GetCreated() bool {
	if m != nil {
		return m.Created
	}
	return false
}

func (m *Response) GetConsignment() *Consignment {
	if m != nil {
		return m.Consignment
	}
	return nil
}

func (m *Response) GetConsignments() []*Consignment {
	if m != nil {
		return m.Consignments
	}
	return nil
}

func init() {
	proto.RegisterType((*Consignment)(nil), "go.micro.srv.consignment.Consignment")
	proto.RegisterType((*Container)(nil), "go.micro.srv.consignment.Container")
	proto.RegisterType((*GetRequest)(nil), "go.micro.srv.consignment.GetRequest")
	proto.RegisterType((*Response)(nil), "go.micro.srv.consignment.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for ShippingService service

type ShippingServiceClient interface {
	CreateConsignment(ctx context.Context, in *Consignment, opts ...client.CallOption) (*Response, error)
	GetConsignments(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*Response, error)
}

type shippingServiceClient struct {
	c           client.Client
	serviceName string
}

func NewShippingServiceClient(serviceName string, c client.Client) ShippingServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "go.micro.srv.consignment"
	}
	return &shippingServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *shippingServiceClient) CreateConsignment(ctx context.Context, in *Consignment, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "ShippingService.CreateConsignment", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shippingServiceClient) GetConsignments(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "ShippingService.GetConsignments", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for ShippingService service

type ShippingServiceHandler interface {
	CreateConsignment(context.Context, *Consignment, *Response) error
	GetConsignments(context.Context, *GetRequest, *Response) error
}

func RegisterShippingServiceHandler(s server.Server, hdlr ShippingServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&ShippingService{hdlr}, opts...))
}

type ShippingService struct {
	ShippingServiceHandler
}

func (h *ShippingService) CreateConsignment(ctx context.Context, in *Consignment, out *Response) error {
	return h.ShippingServiceHandler.CreateConsignment(ctx, in, out)
}

func (h *ShippingService) GetConsignments(ctx context.Context, in *GetRequest, out *Response) error {
	return h.ShippingServiceHandler.GetConsignments(ctx, in, out)
}

func init() { proto.RegisterFile("consignment.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 345 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xbf, 0x4e, 0xc3, 0x30,
	0x10, 0x87, 0x49, 0xff, 0xe7, 0x52, 0x51, 0xd5, 0x03, 0x58, 0x30, 0x10, 0x05, 0x90, 0x3a, 0x65,
	0x28, 0x8f, 0x90, 0xa1, 0xca, 0xea, 0xce, 0x08, 0x4a, 0x72, 0x4a, 0x4f, 0x22, 0x76, 0xb0, 0xdd,
	0xf2, 0x6a, 0xf0, 0x1a, 0x3c, 0x11, 0x6a, 0xd2, 0x50, 0x23, 0x54, 0xd4, 0x2d, 0xbf, 0x3b, 0x7f,
	0x97, 0xcf, 0x27, 0xc3, 0x34, 0x53, 0xd2, 0x50, 0x21, 0x4b, 0x94, 0x36, 0xae, 0xb4, 0xb2, 0x8a,
	0xf1, 0x42, 0xc5, 0x25, 0x65, 0x5a, 0xc5, 0x46, 0x6f, 0x63, 0xa7, 0x1f, 0x7d, 0x7a, 0x10, 0x24,
	0x87, 0xcc, 0xce, 0xa1, 0x43, 0x39, 0xf7, 0x42, 0x6f, 0xe6, 0x8b, 0x0e, 0xe5, 0x2c, 0x84, 0x20,
	0x47, 0x93, 0x69, 0xaa, 0x2c, 0x29, 0xc9, 0x3b, 0x75, 0xc3, 0x2d, 0xb1, 0x0b, 0x18, 0xbc, 0x23,
	0x15, 0x6b, 0xcb, 0xbb, 0xa1, 0x37, 0xeb, 0x8b, 0x7d, 0x62, 0x09, 0x40, 0xa6, 0xa4, 0x5d, 0x91,
	0x44, 0x6d, 0x78, 0x2f, 0xec, 0xce, 0x82, 0xf9, 0x6d, 0x7c, 0x4c, 0x24, 0x4e, 0xda, 0xb3, 0xc2,
	0xc1, 0xd8, 0x35, 0xf8, 0x5b, 0x34, 0x06, 0x5f, 0x9f, 0x28, 0xe7, 0xfd, 0xfa, 0xe7, 0xa3, 0xa6,
	0x90, 0xe6, 0x51, 0x09, 0xfe, 0x0f, 0xf5, 0x47, 0xfc, 0x06, 0x82, 0x6c, 0x63, 0xac, 0x2a, 0x51,
	0xef, 0xd8, 0x46, 0x1c, 0xda, 0x52, 0x9a, 0xef, 0xbc, 0x95, 0xa6, 0x82, 0x64, 0xed, 0xed, 0x8b,
	0x7d, 0x62, 0x97, 0x30, 0xdc, 0x98, 0x06, 0xea, 0x35, 0x8d, 0x5d, 0x4c, 0xf3, 0x68, 0x0c, 0xb0,
	0x40, 0x2b, 0xf0, 0x6d, 0x83, 0xc6, 0x46, 0x1f, 0x1e, 0x8c, 0x04, 0x9a, 0x4a, 0x49, 0x83, 0x8c,
	0xc3, 0x30, 0xd3, 0xb8, 0xb2, 0xd8, 0x18, 0x8c, 0x44, 0x1b, 0xd9, 0x02, 0x02, 0xe7, 0x96, 0xb5,
	0x46, 0x30, 0xbf, 0xff, 0x77, 0x0d, 0xed, 0xb7, 0x70, 0x49, 0x96, 0xc2, 0xd8, 0x89, 0x86, 0x77,
	0xeb, 0x85, 0x9e, 0x38, 0xe9, 0x17, 0x3a, 0xff, 0xf2, 0x60, 0xb2, 0x5c, 0x53, 0x55, 0x91, 0x2c,
	0x96, 0xa8, 0xb7, 0x94, 0x21, 0x7b, 0x86, 0x69, 0x52, 0x2b, 0xbb, 0x8f, 0xe1, 0xb4, 0xe9, 0x57,
	0xd1, 0xf1, 0x63, 0xed, 0x86, 0xa2, 0x33, 0xf6, 0x08, 0x93, 0x05, 0x5a, 0x87, 0x33, 0xec, 0xee,
	0x38, 0x78, 0xd8, 0xf4, 0x69, 0xe3, 0x5f, 0x06, 0xf5, 0x4b, 0x7f, 0xf8, 0x0e, 0x00, 0x00, 0xff,
	0xff, 0xf9, 0x6f, 0x70, 0xe2, 0xfe, 0x02, 0x00, 0x00,
}
