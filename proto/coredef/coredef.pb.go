// Code generated by protoc-gen-go.
// source: coredef.proto
// DO NOT EDIT!

/*
Package coredef is a generated protocol buffer package.

It is generated from these files:
	coredef.proto

It has these top-level messages:
	SessionAccepted
	SessionConnected
	SessionClosed
	PeerInit
	PeerStart
	PeerStop
	CloseClientACK
	BroardcastACK
	TestEchoACK
*/
package coredef

import proto "github.com/golang/protobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

// 一个连接接入
type SessionAccepted struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *SessionAccepted) Reset()         { *m = SessionAccepted{} }
func (m *SessionAccepted) String() string { return proto.CompactTextString(m) }
func (*SessionAccepted) ProtoMessage()    {}

// 已连接
type SessionConnected struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *SessionConnected) Reset()         { *m = SessionConnected{} }
func (m *SessionConnected) String() string { return proto.CompactTextString(m) }
func (*SessionConnected) ProtoMessage()    {}

// 连接断开
type SessionClosed struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *SessionClosed) Reset()         { *m = SessionClosed{} }
func (m *SessionClosed) String() string { return proto.CompactTextString(m) }
func (*SessionClosed) ProtoMessage()    {}

// 端初始化
type PeerInit struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *PeerInit) Reset()         { *m = PeerInit{} }
func (m *PeerInit) String() string { return proto.CompactTextString(m) }
func (*PeerInit) ProtoMessage()    {}

// 端启动
type PeerStart struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *PeerStart) Reset()         { *m = PeerStart{} }
func (m *PeerStart) String() string { return proto.CompactTextString(m) }
func (*PeerStart) ProtoMessage()    {}

// 端停止
type PeerStop struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *PeerStop) Reset()         { *m = PeerStop{} }
func (m *PeerStop) String() string { return proto.CompactTextString(m) }
func (*PeerStop) ProtoMessage()    {}

// 关闭客户端
type CloseClientACK struct {
	ClientID         *int64 `protobuf:"varint,1,opt" json:"ClientID,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CloseClientACK) Reset()         { *m = CloseClientACK{} }
func (m *CloseClientACK) String() string { return proto.CompactTextString(m) }
func (*CloseClientACK) ProtoMessage()    {}

func (m *CloseClientACK) GetClientID() int64 {
	if m != nil && m.ClientID != nil {
		return *m.ClientID
	}
	return 0
}

// 广播
type BroardcastACK struct {
	MsgID            *uint32 `protobuf:"varint,1,opt" json:"MsgID,omitempty"`
	Data             []byte  `protobuf:"bytes,2,opt" json:"Data,omitempty"`
	ClientID         []int64 `protobuf:"varint,3,rep" json:"ClientID,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *BroardcastACK) Reset()         { *m = BroardcastACK{} }
func (m *BroardcastACK) String() string { return proto.CompactTextString(m) }
func (*BroardcastACK) ProtoMessage()    {}

func (m *BroardcastACK) GetMsgID() uint32 {
	if m != nil && m.MsgID != nil {
		return *m.MsgID
	}
	return 0
}

func (m *BroardcastACK) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *BroardcastACK) GetClientID() []int64 {
	if m != nil {
		return m.ClientID
	}
	return nil
}

// ==========================================================
// 测试用消息
// ==========================================================
type TestEchoACK struct {
	Content          *string `protobuf:"bytes,1,opt" json:"Content,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *TestEchoACK) Reset()         { *m = TestEchoACK{} }
func (m *TestEchoACK) String() string { return proto.CompactTextString(m) }
func (*TestEchoACK) ProtoMessage()    {}

func (m *TestEchoACK) GetContent() string {
	if m != nil && m.Content != nil {
		return *m.Content
	}
	return ""
}

func init() {
}
