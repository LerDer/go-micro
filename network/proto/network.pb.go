// Code generated by protoc-gen-go. DO NOT EDIT.
// source: network.proto

package go_micro_network

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/micro/go-micro/router/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Empty request
type ListRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListRequest) Reset()         { *m = ListRequest{} }
func (m *ListRequest) String() string { return proto.CompactTextString(m) }
func (*ListRequest) ProtoMessage()    {}
func (*ListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8571034d60397816, []int{0}
}

func (m *ListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListRequest.Unmarshal(m, b)
}
func (m *ListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListRequest.Marshal(b, m, deterministic)
}
func (m *ListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRequest.Merge(m, src)
}
func (m *ListRequest) XXX_Size() int {
	return xxx_messageInfo_ListRequest.Size(m)
}
func (m *ListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ListRequest proto.InternalMessageInfo

// ListResponse is returned by ListNodes and ListNeighbours
type ListResponse struct {
	Nodes                []*Node  `protobuf:"bytes,1,rep,name=nodes,proto3" json:"nodes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ListResponse) Reset()         { *m = ListResponse{} }
func (m *ListResponse) String() string { return proto.CompactTextString(m) }
func (*ListResponse) ProtoMessage()    {}
func (*ListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8571034d60397816, []int{1}
}

func (m *ListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListResponse.Unmarshal(m, b)
}
func (m *ListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListResponse.Marshal(b, m, deterministic)
}
func (m *ListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListResponse.Merge(m, src)
}
func (m *ListResponse) XXX_Size() int {
	return xxx_messageInfo_ListResponse.Size(m)
}
func (m *ListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ListResponse proto.InternalMessageInfo

func (m *ListResponse) GetNodes() []*Node {
	if m != nil {
		return m.Nodes
	}
	return nil
}

// NeighbourhoodRequest is sent to query node neighbourhood
type NeighbourhoodRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NeighbourhoodRequest) Reset()         { *m = NeighbourhoodRequest{} }
func (m *NeighbourhoodRequest) String() string { return proto.CompactTextString(m) }
func (*NeighbourhoodRequest) ProtoMessage()    {}
func (*NeighbourhoodRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8571034d60397816, []int{2}
}

func (m *NeighbourhoodRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NeighbourhoodRequest.Unmarshal(m, b)
}
func (m *NeighbourhoodRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NeighbourhoodRequest.Marshal(b, m, deterministic)
}
func (m *NeighbourhoodRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NeighbourhoodRequest.Merge(m, src)
}
func (m *NeighbourhoodRequest) XXX_Size() int {
	return xxx_messageInfo_NeighbourhoodRequest.Size(m)
}
func (m *NeighbourhoodRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NeighbourhoodRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NeighbourhoodRequest proto.InternalMessageInfo

func (m *NeighbourhoodRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// NeighbourhoodResponse contains node neighbourhood hierarchy
type NeighbourhoodResponse struct {
	Neighbourhood        *Neighbour `protobuf:"bytes,1,opt,name=neighbourhood,proto3" json:"neighbourhood,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *NeighbourhoodResponse) Reset()         { *m = NeighbourhoodResponse{} }
func (m *NeighbourhoodResponse) String() string { return proto.CompactTextString(m) }
func (*NeighbourhoodResponse) ProtoMessage()    {}
func (*NeighbourhoodResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8571034d60397816, []int{3}
}

func (m *NeighbourhoodResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NeighbourhoodResponse.Unmarshal(m, b)
}
func (m *NeighbourhoodResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NeighbourhoodResponse.Marshal(b, m, deterministic)
}
func (m *NeighbourhoodResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NeighbourhoodResponse.Merge(m, src)
}
func (m *NeighbourhoodResponse) XXX_Size() int {
	return xxx_messageInfo_NeighbourhoodResponse.Size(m)
}
func (m *NeighbourhoodResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_NeighbourhoodResponse.DiscardUnknown(m)
}

var xxx_messageInfo_NeighbourhoodResponse proto.InternalMessageInfo

func (m *NeighbourhoodResponse) GetNeighbourhood() *Neighbour {
	if m != nil {
		return m.Neighbourhood
	}
	return nil
}

// Node is network node
type Node struct {
	// node ide
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// node address
	Address              string   `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()    {}
func (*Node) Descriptor() ([]byte, []int) {
	return fileDescriptor_8571034d60397816, []int{4}
}

func (m *Node) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Node.Unmarshal(m, b)
}
func (m *Node) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Node.Marshal(b, m, deterministic)
}
func (m *Node) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Node.Merge(m, src)
}
func (m *Node) XXX_Size() int {
	return xxx_messageInfo_Node.Size(m)
}
func (m *Node) XXX_DiscardUnknown() {
	xxx_messageInfo_Node.DiscardUnknown(m)
}

var xxx_messageInfo_Node proto.InternalMessageInfo

func (m *Node) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Node) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

// Connect is sent when the node connects to the network
type Connect struct {
	// network mode
	Node                 *Node    `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Connect) Reset()         { *m = Connect{} }
func (m *Connect) String() string { return proto.CompactTextString(m) }
func (*Connect) ProtoMessage()    {}
func (*Connect) Descriptor() ([]byte, []int) {
	return fileDescriptor_8571034d60397816, []int{5}
}

func (m *Connect) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Connect.Unmarshal(m, b)
}
func (m *Connect) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Connect.Marshal(b, m, deterministic)
}
func (m *Connect) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Connect.Merge(m, src)
}
func (m *Connect) XXX_Size() int {
	return xxx_messageInfo_Connect.Size(m)
}
func (m *Connect) XXX_DiscardUnknown() {
	xxx_messageInfo_Connect.DiscardUnknown(m)
}

var xxx_messageInfo_Connect proto.InternalMessageInfo

func (m *Connect) GetNode() *Node {
	if m != nil {
		return m.Node
	}
	return nil
}

// Close is sent when the node disconnects from the network
type Close struct {
	// network node
	Node                 *Node    `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Close) Reset()         { *m = Close{} }
func (m *Close) String() string { return proto.CompactTextString(m) }
func (*Close) ProtoMessage()    {}
func (*Close) Descriptor() ([]byte, []int) {
	return fileDescriptor_8571034d60397816, []int{6}
}

func (m *Close) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Close.Unmarshal(m, b)
}
func (m *Close) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Close.Marshal(b, m, deterministic)
}
func (m *Close) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Close.Merge(m, src)
}
func (m *Close) XXX_Size() int {
	return xxx_messageInfo_Close.Size(m)
}
func (m *Close) XXX_DiscardUnknown() {
	xxx_messageInfo_Close.DiscardUnknown(m)
}

var xxx_messageInfo_Close proto.InternalMessageInfo

func (m *Close) GetNode() *Node {
	if m != nil {
		return m.Node
	}
	return nil
}

// Neighbour is used to nnounce node neighbourhood
type Neighbour struct {
	// network node
	Node *Node `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	// neighbours
	Neighbours           []*Node  `protobuf:"bytes,3,rep,name=neighbours,proto3" json:"neighbours,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Neighbour) Reset()         { *m = Neighbour{} }
func (m *Neighbour) String() string { return proto.CompactTextString(m) }
func (*Neighbour) ProtoMessage()    {}
func (*Neighbour) Descriptor() ([]byte, []int) {
	return fileDescriptor_8571034d60397816, []int{7}
}

func (m *Neighbour) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Neighbour.Unmarshal(m, b)
}
func (m *Neighbour) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Neighbour.Marshal(b, m, deterministic)
}
func (m *Neighbour) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Neighbour.Merge(m, src)
}
func (m *Neighbour) XXX_Size() int {
	return xxx_messageInfo_Neighbour.Size(m)
}
func (m *Neighbour) XXX_DiscardUnknown() {
	xxx_messageInfo_Neighbour.DiscardUnknown(m)
}

var xxx_messageInfo_Neighbour proto.InternalMessageInfo

func (m *Neighbour) GetNode() *Node {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *Neighbour) GetNeighbours() []*Node {
	if m != nil {
		return m.Neighbours
	}
	return nil
}

func init() {
	proto.RegisterType((*ListRequest)(nil), "go.micro.network.ListRequest")
	proto.RegisterType((*ListResponse)(nil), "go.micro.network.ListResponse")
	proto.RegisterType((*NeighbourhoodRequest)(nil), "go.micro.network.NeighbourhoodRequest")
	proto.RegisterType((*NeighbourhoodResponse)(nil), "go.micro.network.NeighbourhoodResponse")
	proto.RegisterType((*Node)(nil), "go.micro.network.Node")
	proto.RegisterType((*Connect)(nil), "go.micro.network.Connect")
	proto.RegisterType((*Close)(nil), "go.micro.network.Close")
	proto.RegisterType((*Neighbour)(nil), "go.micro.network.Neighbour")
}

func init() { proto.RegisterFile("network.proto", fileDescriptor_8571034d60397816) }

var fileDescriptor_8571034d60397816 = []byte{
	// 348 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0x41, 0x4f, 0x32, 0x31,
	0x10, 0xfd, 0x58, 0xe0, 0x23, 0x0c, 0x62, 0x4c, 0xa3, 0x66, 0xb3, 0x06, 0x43, 0x7a, 0x40, 0x62,
	0x74, 0x31, 0x10, 0x3d, 0x79, 0x31, 0x1c, 0xbc, 0x10, 0x0e, 0x7b, 0xf4, 0xe6, 0xd2, 0x66, 0x69,
	0x94, 0x1d, 0x6c, 0xbb, 0xf1, 0x0f, 0xf8, 0xc3, 0x4d, 0xbb, 0x05, 0x97, 0x45, 0x30, 0xdc, 0xda,
	0x99, 0xf7, 0xe6, 0xcd, 0xe4, 0x3d, 0x68, 0xa7, 0x5c, 0x7f, 0xa2, 0x7c, 0x0b, 0x97, 0x12, 0x35,
	0x92, 0x93, 0x04, 0xc3, 0x85, 0x98, 0x49, 0x0c, 0x5d, 0x3d, 0x18, 0x25, 0x42, 0xcf, 0xb3, 0x38,
	0x9c, 0xe1, 0x62, 0x60, 0x3b, 0x83, 0x04, 0x6f, 0xf3, 0x87, 0xc4, 0x4c, 0x73, 0x39, 0xb0, 0x4c,
	0xf7, 0xc9, 0xc7, 0xd0, 0x36, 0xb4, 0x26, 0x42, 0xe9, 0x88, 0x7f, 0x64, 0x5c, 0x69, 0xfa, 0x08,
	0x47, 0xf9, 0x57, 0x2d, 0x31, 0x55, 0x9c, 0xdc, 0x40, 0x3d, 0x45, 0xc6, 0x95, 0x5f, 0xe9, 0x56,
	0xfb, 0xad, 0xe1, 0x79, 0x58, 0x56, 0x0d, 0xa7, 0xc8, 0x78, 0x94, 0x83, 0x68, 0x0f, 0x4e, 0xa7,
	0x5c, 0x24, 0xf3, 0x18, 0x33, 0x39, 0x47, 0x64, 0x6e, 0x2a, 0x39, 0x06, 0x4f, 0x30, 0xbf, 0xd2,
	0xad, 0xf4, 0x9b, 0x91, 0x27, 0x18, 0x7d, 0x81, 0xb3, 0x12, 0xce, 0xc9, 0x3d, 0x99, 0x2b, 0x0b,
	0x0d, 0xcb, 0x69, 0x0d, 0x2f, 0x7e, 0x91, 0x5d, 0xc1, 0xa2, 0x4d, 0x06, 0xbd, 0x83, 0x9a, 0x59,
	0xa9, 0xac, 0x49, 0x7c, 0x68, 0xbc, 0x32, 0x26, 0xb9, 0x52, 0xbe, 0x67, 0x8b, 0xab, 0x2f, 0xbd,
	0x87, 0xc6, 0x18, 0xd3, 0x94, 0xcf, 0x34, 0xb9, 0x86, 0x9a, 0xb9, 0xc4, 0xc9, 0xee, 0xba, 0xd6,
	0x62, 0xe8, 0x08, 0xea, 0xe3, 0x77, 0x54, 0xfc, 0x20, 0x12, 0x42, 0x73, 0xbd, 0xf9, 0x21, 0x44,
	0xf2, 0x00, 0xb0, 0xbe, 0x53, 0xf9, 0xd5, 0xbd, 0x6e, 0x14, 0x90, 0xc3, 0x2f, 0x0f, 0x1a, 0xd3,
	0xbc, 0x49, 0x9e, 0x01, 0xac, 0xb9, 0xc6, 0x7f, 0x45, 0xfc, 0x1f, 0xb6, 0x4b, 0x84, 0xb3, 0x2b,
	0xe8, 0x6c, 0x75, 0x8a, 0x99, 0xa0, 0xff, 0xc8, 0x04, 0x9a, 0xa6, 0x62, 0xc4, 0x14, 0xe9, 0x6c,
	0x6f, 0x51, 0x48, 0x54, 0x70, 0xb9, 0xab, 0xbd, 0x9e, 0x16, 0x43, 0x7b, 0x23, 0x0d, 0xa4, 0xb7,
	0xc7, 0xee, 0x42, 0xac, 0x82, 0xab, 0x3f, 0x71, 0x2b, 0x8d, 0xf8, 0xbf, 0x4d, 0xfb, 0xe8, 0x3b,
	0x00, 0x00, 0xff, 0xff, 0xfb, 0xa1, 0x6b, 0xb0, 0x45, 0x03, 0x00, 0x00,
}