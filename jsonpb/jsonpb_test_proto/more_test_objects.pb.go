// Code generated by protoc-gen-go.
// source: more_test_objects.proto
// DO NOT EDIT!

/*
Package jsonpb is a generated protocol buffer package.

It is generated from these files:
	more_test_objects.proto
	test_objects.proto

It has these top-level messages:
	Simple3
	Mappy
	Simple
	Repeats
	Widget
	Maps
	MsgWithOneof
	Real
	Complex
*/
package jsonpb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Simple3 struct {
	Dub float64 `protobuf:"fixed64,1,opt,name=dub" json:"dub,omitempty"`
}

func (m *Simple3) Reset()         { *m = Simple3{} }
func (m *Simple3) String() string { return proto.CompactTextString(m) }
func (*Simple3) ProtoMessage()    {}

type Mappy struct {
	Nummy map[int64]int32    `protobuf:"bytes,1,rep,name=nummy" json:"nummy,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
	Strry map[string]string  `protobuf:"bytes,2,rep,name=strry" json:"strry,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Objjy map[int32]*Simple3 `protobuf:"bytes,3,rep,name=objjy" json:"objjy,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Buggy map[int64]string   `protobuf:"bytes,4,rep,name=buggy" json:"buggy,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Booly map[bool]bool      `protobuf:"bytes,5,rep,name=booly" json:"booly,omitempty" protobuf_key:"varint,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
}

func (m *Mappy) Reset()         { *m = Mappy{} }
func (m *Mappy) String() string { return proto.CompactTextString(m) }
func (*Mappy) ProtoMessage()    {}

func (m *Mappy) GetNummy() map[int64]int32 {
	if m != nil {
		return m.Nummy
	}
	return nil
}

func (m *Mappy) GetStrry() map[string]string {
	if m != nil {
		return m.Strry
	}
	return nil
}

func (m *Mappy) GetObjjy() map[int32]*Simple3 {
	if m != nil {
		return m.Objjy
	}
	return nil
}

func (m *Mappy) GetBuggy() map[int64]string {
	if m != nil {
		return m.Buggy
	}
	return nil
}

func (m *Mappy) GetBooly() map[bool]bool {
	if m != nil {
		return m.Booly
	}
	return nil
}
