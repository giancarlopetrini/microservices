// Code generated by protoc-gen-go. DO NOT EDIT.
// source: vessel.proto

/*
Package vessel is a generated protocol buffer package.

It is generated from these files:
	vessel.proto

It has these top-level messages:
	Vessel
	Specification
	Response
*/
package vessel

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

type Vessel struct {
	Id        string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Capacity  int32  `protobuf:"varint,2,opt,name=capacity" json:"capacity,omitempty"`
	MaxWeight int32  `protobuf:"varint,3,opt,name=max_weight,json=maxWeight" json:"max_weight,omitempty"`
	Name      string `protobuf:"bytes,4,opt,name=name" json:"name,omitempty"`
	Available bool   `protobuf:"varint,5,opt,name=available" json:"available,omitempty"`
	OwnerId   string `protobuf:"bytes,6,opt,name=owner_id,json=ownerId" json:"owner_id,omitempty"`
}

func (m *Vessel) Reset()                    { *m = Vessel{} }
func (m *Vessel) String() string            { return proto.CompactTextString(m) }
func (*Vessel) ProtoMessage()               {}
func (*Vessel) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Vessel) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Vessel) GetCapacity() int32 {
	if m != nil {
		return m.Capacity
	}
	return 0
}

func (m *Vessel) GetMaxWeight() int32 {
	if m != nil {
		return m.MaxWeight
	}
	return 0
}

func (m *Vessel) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Vessel) GetAvailable() bool {
	if m != nil {
		return m.Available
	}
	return false
}

func (m *Vessel) GetOwnerId() string {
	if m != nil {
		return m.OwnerId
	}
	return ""
}

type Specification struct {
	Capacity  int32 `protobuf:"varint,1,opt,name=capacity" json:"capacity,omitempty"`
	MaxWeight int32 `protobuf:"varint,2,opt,name=max_weight,json=maxWeight" json:"max_weight,omitempty"`
}

func (m *Specification) Reset()                    { *m = Specification{} }
func (m *Specification) String() string            { return proto.CompactTextString(m) }
func (*Specification) ProtoMessage()               {}
func (*Specification) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Specification) GetCapacity() int32 {
	if m != nil {
		return m.Capacity
	}
	return 0
}

func (m *Specification) GetMaxWeight() int32 {
	if m != nil {
		return m.MaxWeight
	}
	return 0
}

type Response struct {
	Vessel  *Vessel   `protobuf:"bytes,1,opt,name=vessel" json:"vessel,omitempty"`
	Vessels []*Vessel `protobuf:"bytes,2,rep,name=vessels" json:"vessels,omitempty"`
	Created bool      `protobuf:"varint,3,opt,name=created" json:"created,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Response) GetVessel() *Vessel {
	if m != nil {
		return m.Vessel
	}
	return nil
}

func (m *Response) GetVessels() []*Vessel {
	if m != nil {
		return m.Vessels
	}
	return nil
}

func (m *Response) GetCreated() bool {
	if m != nil {
		return m.Created
	}
	return false
}

func init() {
	proto.RegisterType((*Vessel)(nil), "vessel.Vessel")
	proto.RegisterType((*Specification)(nil), "vessel.Specification")
	proto.RegisterType((*Response)(nil), "vessel.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for VesselService service

type VesselServiceClient interface {
	FindAvailable(ctx context.Context, in *Specification, opts ...client.CallOption) (*Response, error)
	Create(ctx context.Context, in *Vessel, opts ...client.CallOption) (*Response, error)
}

type vesselServiceClient struct {
	c           client.Client
	serviceName string
}

func NewVesselServiceClient(serviceName string, c client.Client) VesselServiceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "vessel"
	}
	return &vesselServiceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *vesselServiceClient) FindAvailable(ctx context.Context, in *Specification, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "VesselService.FindAvailable", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vesselServiceClient) Create(ctx context.Context, in *Vessel, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.serviceName, "VesselService.Create", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for VesselService service

type VesselServiceHandler interface {
	FindAvailable(context.Context, *Specification, *Response) error
	Create(context.Context, *Vessel, *Response) error
}

func RegisterVesselServiceHandler(s server.Server, hdlr VesselServiceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&VesselService{hdlr}, opts...))
}

type VesselService struct {
	VesselServiceHandler
}

func (h *VesselService) FindAvailable(ctx context.Context, in *Specification, out *Response) error {
	return h.VesselServiceHandler.FindAvailable(ctx, in, out)
}

func (h *VesselService) Create(ctx context.Context, in *Vessel, out *Response) error {
	return h.VesselServiceHandler.Create(ctx, in, out)
}

func init() { proto.RegisterFile("vessel.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 294 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0xcb, 0x4a, 0xc3, 0x40,
	0x14, 0x75, 0xd2, 0x36, 0x8f, 0xab, 0x29, 0x72, 0x41, 0x18, 0x8b, 0x42, 0xc8, 0x42, 0xb2, 0x90,
	0x2e, 0xea, 0xce, 0x9d, 0x08, 0x82, 0x2e, 0xa7, 0xa0, 0xcb, 0x32, 0x9d, 0xb9, 0xea, 0x40, 0x5e,
	0x24, 0x21, 0x6d, 0xff, 0xc6, 0x4f, 0x15, 0x26, 0x49, 0xa5, 0xa5, 0xb8, 0xbb, 0xe7, 0x91, 0xc3,
	0xc9, 0x19, 0xb8, 0x68, 0xa9, 0xae, 0x29, 0x9d, 0x97, 0x55, 0xd1, 0x14, 0xe8, 0x76, 0x28, 0xfe,
	0x61, 0xe0, 0xbe, 0xdb, 0x13, 0xa7, 0xe0, 0x18, 0xcd, 0x59, 0xc4, 0x92, 0x40, 0x38, 0x46, 0xe3,
	0x0c, 0x7c, 0x25, 0x4b, 0xa9, 0x4c, 0xb3, 0xe3, 0x4e, 0xc4, 0x92, 0x89, 0xd8, 0x63, 0xbc, 0x05,
	0xc8, 0xe4, 0x76, 0xb5, 0x21, 0xf3, 0xf5, 0xdd, 0xf0, 0x91, 0x55, 0x83, 0x4c, 0x6e, 0x3f, 0x2c,
	0x81, 0x08, 0xe3, 0x5c, 0x66, 0xc4, 0xc7, 0x36, 0xcc, 0xde, 0x78, 0x03, 0x81, 0x6c, 0xa5, 0x49,
	0xe5, 0x3a, 0x25, 0x3e, 0x89, 0x58, 0xe2, 0x8b, 0x3f, 0x02, 0xaf, 0xc1, 0x2f, 0x36, 0x39, 0x55,
	0x2b, 0xa3, 0xb9, 0x6b, 0xbf, 0xf2, 0x2c, 0x7e, 0xd5, 0xf1, 0x1b, 0x84, 0xcb, 0x92, 0x94, 0xf9,
	0x34, 0x4a, 0x36, 0xa6, 0xc8, 0x0f, 0x8a, 0xb1, 0x7f, 0x8b, 0x39, 0x47, 0xc5, 0xe2, 0x16, 0x7c,
	0x41, 0x75, 0x59, 0xe4, 0x35, 0xe1, 0x1d, 0xf4, 0x23, 0xd8, 0x90, 0xf3, 0xc5, 0x74, 0xde, 0x2f,
	0xd4, 0xed, 0x21, 0x7a, 0x15, 0x13, 0xf0, 0xba, 0xab, 0xe6, 0x4e, 0x34, 0x3a, 0x61, 0x1c, 0x64,
	0xe4, 0xe0, 0xa9, 0x8a, 0x64, 0x43, 0xda, 0x4e, 0xe2, 0x8b, 0x01, 0x2e, 0x76, 0x10, 0x76, 0xe6,
	0x25, 0x55, 0xad, 0x51, 0x84, 0x8f, 0x10, 0xbe, 0x98, 0x5c, 0x3f, 0xed, 0x07, 0xb8, 0x1a, 0x42,
	0x0f, 0xfe, 0x75, 0x76, 0x39, 0xd0, 0x43, 0xed, 0xf8, 0x0c, 0xef, 0xc1, 0x7d, 0xb6, 0xb9, 0x78,
	0xd4, 0xe4, 0x94, 0x7b, 0xed, 0xda, 0x07, 0x7f, 0xf8, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x03, 0xf1,
	0xd5, 0xb0, 0x00, 0x02, 0x00, 0x00,
}
