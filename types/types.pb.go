// Code generated by protoc-gen-go.
// source: types.proto
// DO NOT EDIT!

/*
Package types is a generated protocol buffer package.

It is generated from these files:
	types.proto

It has these top-level messages:
	Header
	Block
	Transaction
	ChainStatus
	RequestHash
	MerkleProof
	RequestBlocks
	Blocks
	Reply
	ReplyBlockHeight
	RequestTxList
	ReplyTxList
	TxHashList
*/
package types

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
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

type Header struct {
	Version    int64  `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	ParentHash []byte `protobuf:"bytes,2,opt,name=parentHash,proto3" json:"parentHash,omitempty"`
	TxHash     []byte `protobuf:"bytes,3,opt,name=txHash,proto3" json:"txHash,omitempty"`
	Height     int64  `protobuf:"varint,4,opt,name=height" json:"height,omitempty"`
	BlockTime  int64  `protobuf:"varint,5,opt,name=blockTime" json:"blockTime,omitempty"`
}

func (m *Header) Reset()                    { *m = Header{} }
func (m *Header) String() string            { return proto.CompactTextString(m) }
func (*Header) ProtoMessage()               {}
func (*Header) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Header) GetVersion() int64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Header) GetParentHash() []byte {
	if m != nil {
		return m.ParentHash
	}
	return nil
}

func (m *Header) GetTxHash() []byte {
	if m != nil {
		return m.TxHash
	}
	return nil
}

func (m *Header) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *Header) GetBlockTime() int64 {
	if m != nil {
		return m.BlockTime
	}
	return 0
}

type Block struct {
	Version    int64          `protobuf:"varint,1,opt,name=version" json:"version,omitempty"`
	ParentHash []byte         `protobuf:"bytes,2,opt,name=parentHash,proto3" json:"parentHash,omitempty"`
	TxHash     []byte         `protobuf:"bytes,3,opt,name=txHash,proto3" json:"txHash,omitempty"`
	Height     int64          `protobuf:"varint,4,opt,name=height" json:"height,omitempty"`
	BlockTime  int64          `protobuf:"varint,5,opt,name=blockTime" json:"blockTime,omitempty"`
	Txs        []*Transaction `protobuf:"bytes,6,rep,name=txs" json:"txs,omitempty"`
}

func (m *Block) Reset()                    { *m = Block{} }
func (m *Block) String() string            { return proto.CompactTextString(m) }
func (*Block) ProtoMessage()               {}
func (*Block) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Block) GetVersion() int64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *Block) GetParentHash() []byte {
	if m != nil {
		return m.ParentHash
	}
	return nil
}

func (m *Block) GetTxHash() []byte {
	if m != nil {
		return m.TxHash
	}
	return nil
}

func (m *Block) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *Block) GetBlockTime() int64 {
	if m != nil {
		return m.BlockTime
	}
	return 0
}

func (m *Block) GetTxs() []*Transaction {
	if m != nil {
		return m.Txs
	}
	return nil
}

type Transaction struct {
	Payload   []byte `protobuf:"bytes,1,opt,name=Payload,proto3" json:"Payload,omitempty"`
	Signature []byte `protobuf:"bytes,2,opt,name=Signature,proto3" json:"Signature,omitempty"`
	Account   []byte `protobuf:"bytes,3,opt,name=Account,proto3" json:"Account,omitempty"`
}

func (m *Transaction) Reset()                    { *m = Transaction{} }
func (m *Transaction) String() string            { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()               {}
func (*Transaction) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Transaction) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *Transaction) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *Transaction) GetAccount() []byte {
	if m != nil {
		return m.Account
	}
	return nil
}

type ChainStatus struct {
	CurrentHeight int64 `protobuf:"varint,1,opt,name=currentHeight" json:"currentHeight,omitempty"`
	MempoolSize   int64 `protobuf:"varint,2,opt,name=mempoolSize" json:"mempoolSize,omitempty"`
	MsgQueueSize  int64 `protobuf:"varint,3,opt,name=msgQueueSize" json:"msgQueueSize,omitempty"`
}

func (m *ChainStatus) Reset()                    { *m = ChainStatus{} }
func (m *ChainStatus) String() string            { return proto.CompactTextString(m) }
func (*ChainStatus) ProtoMessage()               {}
func (*ChainStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ChainStatus) GetCurrentHeight() int64 {
	if m != nil {
		return m.CurrentHeight
	}
	return 0
}

func (m *ChainStatus) GetMempoolSize() int64 {
	if m != nil {
		return m.MempoolSize
	}
	return 0
}

func (m *ChainStatus) GetMsgQueueSize() int64 {
	if m != nil {
		return m.MsgQueueSize
	}
	return 0
}

type RequestHash struct {
	Hash []byte `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (m *RequestHash) Reset()                    { *m = RequestHash{} }
func (m *RequestHash) String() string            { return proto.CompactTextString(m) }
func (*RequestHash) ProtoMessage()               {}
func (*RequestHash) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *RequestHash) GetHash() []byte {
	if m != nil {
		return m.Hash
	}
	return nil
}

type MerkleProof struct {
	Hashs [][]byte `protobuf:"bytes,1,rep,name=hashs,proto3" json:"hashs,omitempty"`
}

func (m *MerkleProof) Reset()                    { *m = MerkleProof{} }
func (m *MerkleProof) String() string            { return proto.CompactTextString(m) }
func (*MerkleProof) ProtoMessage()               {}
func (*MerkleProof) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *MerkleProof) GetHashs() [][]byte {
	if m != nil {
		return m.Hashs
	}
	return nil
}

// req
type RequestBlocks struct {
	Start int64 `protobuf:"varint,1,opt,name=start" json:"start,omitempty"`
	End   int64 `protobuf:"varint,2,opt,name=end" json:"end,omitempty"`
}

func (m *RequestBlocks) Reset()                    { *m = RequestBlocks{} }
func (m *RequestBlocks) String() string            { return proto.CompactTextString(m) }
func (*RequestBlocks) ProtoMessage()               {}
func (*RequestBlocks) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *RequestBlocks) GetStart() int64 {
	if m != nil {
		return m.Start
	}
	return 0
}

func (m *RequestBlocks) GetEnd() int64 {
	if m != nil {
		return m.End
	}
	return 0
}

// resp
type Blocks struct {
	Items []*Block `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
}

func (m *Blocks) Reset()                    { *m = Blocks{} }
func (m *Blocks) String() string            { return proto.CompactTextString(m) }
func (*Blocks) ProtoMessage()               {}
func (*Blocks) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *Blocks) GetItems() []*Block {
	if m != nil {
		return m.Items
	}
	return nil
}

type Reply struct {
	IsOk bool   `protobuf:"varint,1,opt,name=isOk" json:"isOk,omitempty"`
	Msg  []byte `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (m *Reply) Reset()                    { *m = Reply{} }
func (m *Reply) String() string            { return proto.CompactTextString(m) }
func (*Reply) ProtoMessage()               {}
func (*Reply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *Reply) GetIsOk() bool {
	if m != nil {
		return m.IsOk
	}
	return false
}

func (m *Reply) GetMsg() []byte {
	if m != nil {
		return m.Msg
	}
	return nil
}

type ReplyBlockHeight struct {
	Height int64 `protobuf:"varint,1,opt,name=height" json:"height,omitempty"`
}

func (m *ReplyBlockHeight) Reset()                    { *m = ReplyBlockHeight{} }
func (m *ReplyBlockHeight) String() string            { return proto.CompactTextString(m) }
func (*ReplyBlockHeight) ProtoMessage()               {}
func (*ReplyBlockHeight) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *ReplyBlockHeight) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

type RequestTxList struct {
	Count int64 `protobuf:"varint,1,opt,name=count" json:"count,omitempty"`
}

func (m *RequestTxList) Reset()                    { *m = RequestTxList{} }
func (m *RequestTxList) String() string            { return proto.CompactTextString(m) }
func (*RequestTxList) ProtoMessage()               {}
func (*RequestTxList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *RequestTxList) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type ReplyTxList struct {
	Txs []*Transaction `protobuf:"bytes,1,rep,name=txs" json:"txs,omitempty"`
}

func (m *ReplyTxList) Reset()                    { *m = ReplyTxList{} }
func (m *ReplyTxList) String() string            { return proto.CompactTextString(m) }
func (*ReplyTxList) ProtoMessage()               {}
func (*ReplyTxList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *ReplyTxList) GetTxs() []*Transaction {
	if m != nil {
		return m.Txs
	}
	return nil
}

type TxHashList struct {
	Hashes [][]byte `protobuf:"bytes,1,rep,name=hashes,proto3" json:"hashes,omitempty"`
}

func (m *TxHashList) Reset()                    { *m = TxHashList{} }
func (m *TxHashList) String() string            { return proto.CompactTextString(m) }
func (*TxHashList) ProtoMessage()               {}
func (*TxHashList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *TxHashList) GetHashes() [][]byte {
	if m != nil {
		return m.Hashes
	}
	return nil
}

func init() {
	proto.RegisterType((*Header)(nil), "types.Header")
	proto.RegisterType((*Block)(nil), "types.Block")
	proto.RegisterType((*Transaction)(nil), "types.Transaction")
	proto.RegisterType((*ChainStatus)(nil), "types.ChainStatus")
	proto.RegisterType((*RequestHash)(nil), "types.RequestHash")
	proto.RegisterType((*MerkleProof)(nil), "types.MerkleProof")
	proto.RegisterType((*RequestBlocks)(nil), "types.RequestBlocks")
	proto.RegisterType((*Blocks)(nil), "types.Blocks")
	proto.RegisterType((*Reply)(nil), "types.Reply")
	proto.RegisterType((*ReplyBlockHeight)(nil), "types.ReplyBlockHeight")
	proto.RegisterType((*RequestTxList)(nil), "types.RequestTxList")
	proto.RegisterType((*ReplyTxList)(nil), "types.ReplyTxList")
	proto.RegisterType((*TxHashList)(nil), "types.TxHashList")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Grpcservice service

type GrpcserviceClient interface {
	SendTransaction(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*Reply, error)
	QueryTransaction(ctx context.Context, in *RequestHash, opts ...grpc.CallOption) (*Reply, error)
	GetBlocks(ctx context.Context, in *RequestBlocks, opts ...grpc.CallOption) (*Reply, error)
}

type grpcserviceClient struct {
	cc *grpc.ClientConn
}

func NewGrpcserviceClient(cc *grpc.ClientConn) GrpcserviceClient {
	return &grpcserviceClient{cc}
}

func (c *grpcserviceClient) SendTransaction(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := grpc.Invoke(ctx, "/types.grpcservice/SendTransaction", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *grpcserviceClient) QueryTransaction(ctx context.Context, in *RequestHash, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := grpc.Invoke(ctx, "/types.grpcservice/QueryTransaction", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *grpcserviceClient) GetBlocks(ctx context.Context, in *RequestBlocks, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := grpc.Invoke(ctx, "/types.grpcservice/GetBlocks", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Grpcservice service

type GrpcserviceServer interface {
	SendTransaction(context.Context, *Transaction) (*Reply, error)
	QueryTransaction(context.Context, *RequestHash) (*Reply, error)
	GetBlocks(context.Context, *RequestBlocks) (*Reply, error)
}

func RegisterGrpcserviceServer(s *grpc.Server, srv GrpcserviceServer) {
	s.RegisterService(&_Grpcservice_serviceDesc, srv)
}

func _Grpcservice_SendTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GrpcserviceServer).SendTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.grpcservice/SendTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GrpcserviceServer).SendTransaction(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _Grpcservice_QueryTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestHash)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GrpcserviceServer).QueryTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.grpcservice/QueryTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GrpcserviceServer).QueryTransaction(ctx, req.(*RequestHash))
	}
	return interceptor(ctx, in, info, handler)
}

func _Grpcservice_GetBlocks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestBlocks)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GrpcserviceServer).GetBlocks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.grpcservice/GetBlocks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GrpcserviceServer).GetBlocks(ctx, req.(*RequestBlocks))
	}
	return interceptor(ctx, in, info, handler)
}

var _Grpcservice_serviceDesc = grpc.ServiceDesc{
	ServiceName: "types.grpcservice",
	HandlerType: (*GrpcserviceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SendTransaction",
			Handler:    _Grpcservice_SendTransaction_Handler,
		},
		{
			MethodName: "QueryTransaction",
			Handler:    _Grpcservice_QueryTransaction_Handler,
		},
		{
			MethodName: "GetBlocks",
			Handler:    _Grpcservice_GetBlocks_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "types.proto",
}

func init() { proto.RegisterFile("types.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 523 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xcc, 0x54, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xc5, 0xb8, 0x0e, 0x74, 0x9c, 0x8a, 0x68, 0x85, 0x90, 0x85, 0x10, 0x0a, 0x4b, 0x90, 0x22,
	0x04, 0x95, 0x48, 0x05, 0x9c, 0x81, 0x03, 0x3d, 0x80, 0x68, 0x9d, 0xdc, 0xd1, 0xd6, 0x19, 0xec,
	0x55, 0x6c, 0xaf, 0xd9, 0x5d, 0x57, 0x31, 0xbf, 0x82, 0x7f, 0xc2, 0x89, 0xff, 0x87, 0xf6, 0xc3,
	0x8d, 0x03, 0x15, 0xe7, 0xde, 0xf6, 0xbd, 0xf9, 0xf0, 0x9b, 0x37, 0x23, 0x43, 0xac, 0xbb, 0x06,
	0xd5, 0x71, 0x23, 0x85, 0x16, 0x24, 0xb2, 0x80, 0xfe, 0x0c, 0x60, 0x74, 0x8a, 0x6c, 0x8d, 0x92,
	0x24, 0x70, 0xe7, 0x12, 0xa5, 0xe2, 0xa2, 0x4e, 0x82, 0x69, 0x30, 0x0f, 0xd3, 0x1e, 0x92, 0xc7,
	0x00, 0x0d, 0x93, 0x58, 0xeb, 0x53, 0xa6, 0x8a, 0xe4, 0xf6, 0x34, 0x98, 0x8f, 0xd3, 0x01, 0x43,
	0x1e, 0xc0, 0x48, 0x6f, 0x6d, 0x2c, 0xb4, 0x31, 0x8f, 0x0c, 0x5f, 0x20, 0xcf, 0x0b, 0x9d, 0x1c,
	0xd8, 0x86, 0x1e, 0x91, 0x47, 0x70, 0x78, 0x51, 0x8a, 0x6c, 0xb3, 0xe2, 0x15, 0x26, 0x91, 0x0d,
	0xed, 0x08, 0xfa, 0x3b, 0x80, 0xe8, 0xbd, 0x41, 0x37, 0x45, 0x11, 0x99, 0x41, 0xa8, 0xb7, 0x2a,
	0x19, 0x4d, 0xc3, 0x79, 0xbc, 0x20, 0xc7, 0xce, 0xc6, 0x95, 0x64, 0xb5, 0x62, 0x99, 0xe6, 0xa2,
	0x4e, 0x4d, 0x98, 0x7e, 0x85, 0x78, 0xc0, 0x19, 0xf1, 0x67, 0xac, 0x2b, 0x05, 0x5b, 0x5b, 0xf1,
	0xe3, 0xb4, 0x87, 0xe6, 0x63, 0x4b, 0x9e, 0xd7, 0x4c, 0xb7, 0x12, 0xbd, 0xf6, 0x1d, 0x61, 0xea,
	0xde, 0x65, 0x99, 0x68, 0x6b, 0xed, 0xb5, 0xf7, 0x90, 0x76, 0x10, 0x7f, 0x28, 0x18, 0xaf, 0x97,
	0x9a, 0xe9, 0x56, 0x91, 0x19, 0x1c, 0x65, 0xad, 0xb4, 0x23, 0xbb, 0x91, 0x9c, 0x47, 0xfb, 0x24,
	0x99, 0x42, 0x5c, 0x61, 0xd5, 0x08, 0x51, 0x2e, 0xf9, 0x0f, 0xf7, 0xb9, 0x30, 0x1d, 0x52, 0x84,
	0xc2, 0xb8, 0x52, 0xf9, 0x79, 0x8b, 0x2d, 0xda, 0x94, 0xd0, 0xa6, 0xec, 0x71, 0xf4, 0x09, 0xc4,
	0x29, 0x7e, 0x6f, 0x51, 0x39, 0x7b, 0x09, 0x1c, 0x14, 0xc6, 0x5c, 0x37, 0x98, 0x7d, 0xd3, 0xa7,
	0x10, 0x7f, 0x46, 0xb9, 0x29, 0xf1, 0x4c, 0x0a, 0xf1, 0x8d, 0xdc, 0x87, 0xc8, 0xd0, 0x2a, 0x09,
	0xa6, 0xe1, 0x7c, 0x9c, 0x3a, 0x40, 0xdf, 0xc2, 0x91, 0xef, 0x63, 0x37, 0xac, 0x4c, 0x9a, 0xd2,
	0x4c, 0xf6, 0xe2, 0x1d, 0x20, 0x13, 0x08, 0xb1, 0x5e, 0x7b, 0xb1, 0xe6, 0x49, 0x5f, 0xc0, 0xc8,
	0x57, 0x50, 0x88, 0xb8, 0xc6, 0xca, 0x35, 0x8e, 0x17, 0x63, 0xbf, 0x0e, 0x1b, 0x4d, 0x5d, 0x88,
	0xbe, 0x84, 0x28, 0xc5, 0xa6, 0xec, 0x8c, 0x50, 0xae, 0xbe, 0x6c, 0x6c, 0xf7, 0xbb, 0xa9, 0x7d,
	0x9b, 0xe6, 0x95, 0xca, 0xbd, 0xf1, 0xe6, 0x49, 0x9f, 0xc3, 0xc4, 0xa6, 0xdb, 0x1e, 0xde, 0xb7,
	0xdd, 0xa5, 0x04, 0xc3, 0x4b, 0xa1, 0xcf, 0xae, 0x26, 0x58, 0x6d, 0x3f, 0x71, 0xa5, 0xcd, 0x04,
	0x6e, 0x5b, 0x7e, 0x02, 0xb7, 0xab, 0x13, 0x63, 0x58, 0x53, 0x76, 0x3e, 0xc9, 0x5f, 0x50, 0xf0,
	0xff, 0x0b, 0x9a, 0x01, 0xac, 0xec, 0x9d, 0xda, 0x1a, 0xa3, 0x80, 0xa9, 0x02, 0x7b, 0x0b, 0x3d,
	0x5a, 0xfc, 0x0a, 0x20, 0xce, 0x65, 0x93, 0x29, 0x94, 0x97, 0x3c, 0x43, 0xf2, 0x1a, 0xee, 0x2d,
	0xb1, 0x5e, 0x0f, 0x6f, 0xef, 0x9a, 0x2f, 0x3c, 0xec, 0x8d, 0xb2, 0xb2, 0xe8, 0x2d, 0xf2, 0x06,
	0x26, 0xe7, 0x2d, 0xca, 0xee, 0xba, 0xba, 0xc1, 0xae, 0xff, 0xa9, 0x7b, 0x05, 0x87, 0x1f, 0xf1,
	0x6a, 0x7d, 0xfb, 0x05, 0x8e, 0xfd, 0xbb, 0xe4, 0x62, 0x64, 0x7f, 0x39, 0x27, 0x7f, 0x02, 0x00,
	0x00, 0xff, 0xff, 0xab, 0x99, 0xcf, 0xe8, 0x81, 0x04, 0x00, 0x00,
}