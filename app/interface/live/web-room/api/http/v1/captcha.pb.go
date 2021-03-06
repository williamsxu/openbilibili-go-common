// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api/http/v1/captcha.proto

package v1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type CreateCaptchaReq struct {
	Type                 int64    `protobuf:"varint,1,opt,name=type,proto3" json:"type,omitempty" form:"type"`
	ClientType           string   `protobuf:"bytes,2,opt,name=client_type,json=clientType,proto3" json:"client_type,omitempty" form:"client_type"`
	Height               int64    `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty" form:"height"`
	Width                int64    `protobuf:"varint,4,opt,name=width,proto3" json:"width,omitempty" form:"width"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateCaptchaReq) Reset()         { *m = CreateCaptchaReq{} }
func (m *CreateCaptchaReq) String() string { return proto.CompactTextString(m) }
func (*CreateCaptchaReq) ProtoMessage()    {}
func (*CreateCaptchaReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_captcha_835409ee8e660cbd, []int{0}
}
func (m *CreateCaptchaReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CreateCaptchaReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CreateCaptchaReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *CreateCaptchaReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateCaptchaReq.Merge(dst, src)
}
func (m *CreateCaptchaReq) XXX_Size() int {
	return m.Size()
}
func (m *CreateCaptchaReq) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateCaptchaReq.DiscardUnknown(m)
}

var xxx_messageInfo_CreateCaptchaReq proto.InternalMessageInfo

func (m *CreateCaptchaReq) GetType() int64 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *CreateCaptchaReq) GetClientType() string {
	if m != nil {
		return m.ClientType
	}
	return ""
}

func (m *CreateCaptchaReq) GetHeight() int64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *CreateCaptchaReq) GetWidth() int64 {
	if m != nil {
		return m.Width
	}
	return 0
}

type CreateCaptchaResp struct {
	Type                 int64    `protobuf:"varint,1,opt,name=type,proto3" json:"type"`
	Geetest              *GeeTest `protobuf:"bytes,2,opt,name=geetest" json:"geetest"`
	Image                *Image   `protobuf:"bytes,3,opt,name=image" json:"image"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateCaptchaResp) Reset()         { *m = CreateCaptchaResp{} }
func (m *CreateCaptchaResp) String() string { return proto.CompactTextString(m) }
func (*CreateCaptchaResp) ProtoMessage()    {}
func (*CreateCaptchaResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_captcha_835409ee8e660cbd, []int{1}
}
func (m *CreateCaptchaResp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CreateCaptchaResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CreateCaptchaResp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *CreateCaptchaResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateCaptchaResp.Merge(dst, src)
}
func (m *CreateCaptchaResp) XXX_Size() int {
	return m.Size()
}
func (m *CreateCaptchaResp) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateCaptchaResp.DiscardUnknown(m)
}

var xxx_messageInfo_CreateCaptchaResp proto.InternalMessageInfo

func (m *CreateCaptchaResp) GetType() int64 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *CreateCaptchaResp) GetGeetest() *GeeTest {
	if m != nil {
		return m.Geetest
	}
	return nil
}

func (m *CreateCaptchaResp) GetImage() *Image {
	if m != nil {
		return m.Image
	}
	return nil
}

type GeeTest struct {
	Gt                   string   `protobuf:"bytes,1,opt,name=gt,proto3" json:"gt"`
	Challenge            string   `protobuf:"bytes,2,opt,name=challenge,proto3" json:"challenge"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GeeTest) Reset()         { *m = GeeTest{} }
func (m *GeeTest) String() string { return proto.CompactTextString(m) }
func (*GeeTest) ProtoMessage()    {}
func (*GeeTest) Descriptor() ([]byte, []int) {
	return fileDescriptor_captcha_835409ee8e660cbd, []int{2}
}
func (m *GeeTest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GeeTest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GeeTest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *GeeTest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GeeTest.Merge(dst, src)
}
func (m *GeeTest) XXX_Size() int {
	return m.Size()
}
func (m *GeeTest) XXX_DiscardUnknown() {
	xxx_messageInfo_GeeTest.DiscardUnknown(m)
}

var xxx_messageInfo_GeeTest proto.InternalMessageInfo

func (m *GeeTest) GetGt() string {
	if m != nil {
		return m.Gt
	}
	return ""
}

func (m *GeeTest) GetChallenge() string {
	if m != nil {
		return m.Challenge
	}
	return ""
}

type Image struct {
	Tips                 string   `protobuf:"bytes,1,opt,name=tips,proto3" json:"tips"`
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token"`
	Content              string   `protobuf:"bytes,3,opt,name=content,proto3" json:"content"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Image) Reset()         { *m = Image{} }
func (m *Image) String() string { return proto.CompactTextString(m) }
func (*Image) ProtoMessage()    {}
func (*Image) Descriptor() ([]byte, []int) {
	return fileDescriptor_captcha_835409ee8e660cbd, []int{3}
}
func (m *Image) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Image) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Image.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *Image) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Image.Merge(dst, src)
}
func (m *Image) XXX_Size() int {
	return m.Size()
}
func (m *Image) XXX_DiscardUnknown() {
	xxx_messageInfo_Image.DiscardUnknown(m)
}

var xxx_messageInfo_Image proto.InternalMessageInfo

func (m *Image) GetTips() string {
	if m != nil {
		return m.Tips
	}
	return ""
}

func (m *Image) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *Image) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

type VerifyReq struct {
	Anti                 string   `protobuf:"bytes,1,opt,name=anti,proto3" json:"anti,omitempty" form:"anti"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VerifyReq) Reset()         { *m = VerifyReq{} }
func (m *VerifyReq) String() string { return proto.CompactTextString(m) }
func (*VerifyReq) ProtoMessage()    {}
func (*VerifyReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_captcha_835409ee8e660cbd, []int{4}
}
func (m *VerifyReq) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VerifyReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VerifyReq.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *VerifyReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifyReq.Merge(dst, src)
}
func (m *VerifyReq) XXX_Size() int {
	return m.Size()
}
func (m *VerifyReq) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifyReq.DiscardUnknown(m)
}

var xxx_messageInfo_VerifyReq proto.InternalMessageInfo

func (m *VerifyReq) GetAnti() string {
	if m != nil {
		return m.Anti
	}
	return ""
}

type VerifyResp struct {
	Type                 int64    `protobuf:"varint,1,opt,name=type,proto3" json:"type"`
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VerifyResp) Reset()         { *m = VerifyResp{} }
func (m *VerifyResp) String() string { return proto.CompactTextString(m) }
func (*VerifyResp) ProtoMessage()    {}
func (*VerifyResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_captcha_835409ee8e660cbd, []int{5}
}
func (m *VerifyResp) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VerifyResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VerifyResp.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *VerifyResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifyResp.Merge(dst, src)
}
func (m *VerifyResp) XXX_Size() int {
	return m.Size()
}
func (m *VerifyResp) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifyResp.DiscardUnknown(m)
}

var xxx_messageInfo_VerifyResp proto.InternalMessageInfo

func (m *VerifyResp) GetType() int64 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *VerifyResp) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateCaptchaReq)(nil), "live.webroom.v1.CreateCaptchaReq")
	proto.RegisterType((*CreateCaptchaResp)(nil), "live.webroom.v1.CreateCaptchaResp")
	proto.RegisterType((*GeeTest)(nil), "live.webroom.v1.GeeTest")
	proto.RegisterType((*Image)(nil), "live.webroom.v1.Image")
	proto.RegisterType((*VerifyReq)(nil), "live.webroom.v1.VerifyReq")
	proto.RegisterType((*VerifyResp)(nil), "live.webroom.v1.VerifyResp")
}
func (m *CreateCaptchaReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CreateCaptchaReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Type != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintCaptcha(dAtA, i, uint64(m.Type))
	}
	if len(m.ClientType) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintCaptcha(dAtA, i, uint64(len(m.ClientType)))
		i += copy(dAtA[i:], m.ClientType)
	}
	if m.Height != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintCaptcha(dAtA, i, uint64(m.Height))
	}
	if m.Width != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintCaptcha(dAtA, i, uint64(m.Width))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *CreateCaptchaResp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CreateCaptchaResp) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Type != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintCaptcha(dAtA, i, uint64(m.Type))
	}
	if m.Geetest != nil {
		dAtA[i] = 0x12
		i++
		i = encodeVarintCaptcha(dAtA, i, uint64(m.Geetest.Size()))
		n1, err := m.Geetest.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.Image != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintCaptcha(dAtA, i, uint64(m.Image.Size()))
		n2, err := m.Image.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *GeeTest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GeeTest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Gt) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCaptcha(dAtA, i, uint64(len(m.Gt)))
		i += copy(dAtA[i:], m.Gt)
	}
	if len(m.Challenge) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintCaptcha(dAtA, i, uint64(len(m.Challenge)))
		i += copy(dAtA[i:], m.Challenge)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *Image) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Image) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Tips) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCaptcha(dAtA, i, uint64(len(m.Tips)))
		i += copy(dAtA[i:], m.Tips)
	}
	if len(m.Token) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintCaptcha(dAtA, i, uint64(len(m.Token)))
		i += copy(dAtA[i:], m.Token)
	}
	if len(m.Content) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintCaptcha(dAtA, i, uint64(len(m.Content)))
		i += copy(dAtA[i:], m.Content)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *VerifyReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VerifyReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Anti) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCaptcha(dAtA, i, uint64(len(m.Anti)))
		i += copy(dAtA[i:], m.Anti)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *VerifyResp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VerifyResp) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Type != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintCaptcha(dAtA, i, uint64(m.Type))
	}
	if len(m.Token) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintCaptcha(dAtA, i, uint64(len(m.Token)))
		i += copy(dAtA[i:], m.Token)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintCaptcha(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *CreateCaptchaReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovCaptcha(uint64(m.Type))
	}
	l = len(m.ClientType)
	if l > 0 {
		n += 1 + l + sovCaptcha(uint64(l))
	}
	if m.Height != 0 {
		n += 1 + sovCaptcha(uint64(m.Height))
	}
	if m.Width != 0 {
		n += 1 + sovCaptcha(uint64(m.Width))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *CreateCaptchaResp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovCaptcha(uint64(m.Type))
	}
	if m.Geetest != nil {
		l = m.Geetest.Size()
		n += 1 + l + sovCaptcha(uint64(l))
	}
	if m.Image != nil {
		l = m.Image.Size()
		n += 1 + l + sovCaptcha(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *GeeTest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Gt)
	if l > 0 {
		n += 1 + l + sovCaptcha(uint64(l))
	}
	l = len(m.Challenge)
	if l > 0 {
		n += 1 + l + sovCaptcha(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Image) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Tips)
	if l > 0 {
		n += 1 + l + sovCaptcha(uint64(l))
	}
	l = len(m.Token)
	if l > 0 {
		n += 1 + l + sovCaptcha(uint64(l))
	}
	l = len(m.Content)
	if l > 0 {
		n += 1 + l + sovCaptcha(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *VerifyReq) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Anti)
	if l > 0 {
		n += 1 + l + sovCaptcha(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *VerifyResp) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Type != 0 {
		n += 1 + sovCaptcha(uint64(m.Type))
	}
	l = len(m.Token)
	if l > 0 {
		n += 1 + l + sovCaptcha(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovCaptcha(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozCaptcha(x uint64) (n int) {
	return sovCaptcha(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CreateCaptchaReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCaptcha
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CreateCaptchaReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CreateCaptchaReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCaptcha
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCaptcha
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCaptcha
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClientType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCaptcha
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Width", wireType)
			}
			m.Width = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCaptcha
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Width |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipCaptcha(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCaptcha
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CreateCaptchaResp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCaptcha
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CreateCaptchaResp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CreateCaptchaResp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCaptcha
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Geetest", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCaptcha
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCaptcha
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Geetest == nil {
				m.Geetest = &GeeTest{}
			}
			if err := m.Geetest.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Image", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCaptcha
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthCaptcha
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Image == nil {
				m.Image = &Image{}
			}
			if err := m.Image.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCaptcha(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCaptcha
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *GeeTest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCaptcha
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GeeTest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GeeTest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Gt", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCaptcha
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCaptcha
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Gt = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Challenge", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCaptcha
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCaptcha
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Challenge = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCaptcha(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCaptcha
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Image) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCaptcha
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Image: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Image: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tips", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCaptcha
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCaptcha
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tips = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCaptcha
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCaptcha
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Token = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Content", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCaptcha
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCaptcha
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Content = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCaptcha(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCaptcha
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *VerifyReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCaptcha
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: VerifyReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VerifyReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Anti", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCaptcha
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCaptcha
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Anti = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCaptcha(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCaptcha
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *VerifyResp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCaptcha
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: VerifyResp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VerifyResp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCaptcha
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCaptcha
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCaptcha
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Token = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCaptcha(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCaptcha
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipCaptcha(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCaptcha
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCaptcha
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCaptcha
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthCaptcha
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowCaptcha
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipCaptcha(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthCaptcha = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCaptcha   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("api/http/v1/captcha.proto", fileDescriptor_captcha_835409ee8e660cbd) }

var fileDescriptor_captcha_835409ee8e660cbd = []byte{
	// 512 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x95, 0xd3, 0x26, 0xc1, 0x13, 0xaa, 0xa6, 0x7b, 0x88, 0x42, 0x40, 0xd9, 0xb2, 0x08, 0x54,
	0x84, 0x70, 0x48, 0x38, 0x54, 0xe2, 0x82, 0x70, 0x0f, 0x08, 0x21, 0x81, 0xb4, 0xaa, 0x38, 0x70,
	0x41, 0x8e, 0xd9, 0xac, 0x57, 0x24, 0xde, 0x25, 0xde, 0xa6, 0xea, 0xc7, 0xf0, 0x05, 0x7c, 0x06,
	0x17, 0x8e, 0x7c, 0x81, 0x85, 0x72, 0xf4, 0x31, 0x5f, 0x80, 0x32, 0x6b, 0xd3, 0x92, 0x08, 0xe5,
	0xe2, 0x9d, 0x79, 0xf3, 0x66, 0x76, 0xde, 0xce, 0x18, 0xee, 0x44, 0x46, 0x0d, 0x12, 0x6b, 0xcd,
	0x60, 0x31, 0x1c, 0xc4, 0x91, 0xb1, 0x71, 0x12, 0x05, 0x66, 0xae, 0xad, 0x26, 0x87, 0x53, 0xb5,
	0x10, 0xc1, 0xa5, 0x18, 0xcf, 0xb5, 0x9e, 0x05, 0x8b, 0x61, 0xef, 0xa9, 0x54, 0x36, 0xb9, 0x18,
	0x07, 0xb1, 0x9e, 0x0d, 0xa4, 0x96, 0x7a, 0x80, 0xbc, 0xf1, 0xc5, 0x04, 0x3d, 0x74, 0xd0, 0x72,
	0xf9, 0xec, 0x87, 0x07, 0xed, 0xb3, 0xb9, 0x88, 0xac, 0x38, 0x73, 0x75, 0xb9, 0xf8, 0x4a, 0x1e,
	0xc0, 0xbe, 0xbd, 0x32, 0xa2, 0xeb, 0x1d, 0x7b, 0x27, 0x7b, 0xe1, 0xe1, 0x2a, 0xa7, 0xad, 0x89,
	0x9e, 0xcf, 0x5e, 0xb0, 0x35, 0xca, 0x38, 0x06, 0xc9, 0x29, 0xb4, 0xe2, 0xa9, 0x12, 0xa9, 0xfd,
	0x84, 0xdc, 0xda, 0xb1, 0x77, 0xe2, 0x87, 0x9d, 0x55, 0x4e, 0x89, 0xe3, 0xde, 0x08, 0x32, 0x0e,
	0xce, 0x3b, 0x5f, 0x27, 0x3e, 0x86, 0x46, 0x22, 0x94, 0x4c, 0x6c, 0x77, 0x0f, 0xeb, 0x1f, 0xad,
	0x72, 0x7a, 0xe0, 0x72, 0x1c, 0xce, 0x78, 0x49, 0x20, 0x8f, 0xa0, 0x7e, 0xa9, 0x3e, 0xdb, 0xa4,
	0xbb, 0x8f, 0xcc, 0xf6, 0x2a, 0xa7, 0xb7, 0x1d, 0x13, 0x61, 0xc6, 0x5d, 0x98, 0x7d, 0xf7, 0xe0,
	0x68, 0x43, 0x45, 0x66, 0xc8, 0xbd, 0x7f, 0x64, 0xdc, 0x2a, 0x72, 0x8a, 0x7e, 0xd9, 0xff, 0x4b,
	0x68, 0x4a, 0x21, 0xac, 0xc8, 0x2c, 0xf6, 0xde, 0x1a, 0x75, 0x83, 0x8d, 0xb7, 0x0c, 0x5e, 0x0b,
	0x71, 0x2e, 0x32, 0x1b, 0xb6, 0x8a, 0x9c, 0x56, 0x64, 0x5e, 0x19, 0xe4, 0x14, 0xea, 0x6a, 0x16,
	0x49, 0x81, 0x32, 0x5a, 0xa3, 0xce, 0x56, 0xfa, 0x9b, 0x75, 0x34, 0xf4, 0x8b, 0x9c, 0x3a, 0x22,
	0x77, 0x07, 0x7b, 0x07, 0xcd, 0xb2, 0x32, 0xe9, 0x40, 0x4d, 0x5a, 0x6c, 0xd0, 0x0f, 0x1b, 0x45,
	0x4e, 0x6b, 0xd2, 0xf2, 0x9a, 0xb4, 0xe4, 0x09, 0xf8, 0x71, 0x12, 0x4d, 0xa7, 0x22, 0x95, 0xd5,
	0xd3, 0x1e, 0x14, 0x39, 0xbd, 0x06, 0xf9, 0xb5, 0xc9, 0x66, 0x50, 0xc7, 0xab, 0x50, 0xb0, 0x32,
	0x59, 0x59, 0xcf, 0x09, 0x56, 0x26, 0xe3, 0xf8, 0x25, 0x14, 0xea, 0x56, 0x7f, 0x11, 0x69, 0x59,
	0x0f, 0xfb, 0x42, 0x80, 0xbb, 0x83, 0x3c, 0x84, 0x66, 0xac, 0x53, 0x2b, 0x52, 0x37, 0x19, 0xdf,
	0xe9, 0x2e, 0x21, 0x5e, 0x19, 0xec, 0x19, 0xf8, 0x1f, 0xc4, 0x5c, 0x4d, 0xae, 0xca, 0x55, 0x89,
	0x52, 0xab, 0xca, 0x2b, 0x6f, 0xac, 0xca, 0x1a, 0x65, 0x1c, 0x83, 0xec, 0x2d, 0x40, 0x95, 0xb1,
	0x73, 0x2c, 0xbb, 0xba, 0x1c, 0x7d, 0xf3, 0xa0, 0x59, 0x4e, 0x99, 0xbc, 0x87, 0x46, 0x8c, 0x63,
	0x27, 0xf7, 0xb7, 0x5e, 0x7f, 0x73, 0xab, 0x7b, 0x6c, 0x17, 0x25, 0x33, 0xe4, 0x15, 0x34, 0x16,
	0xd8, 0x29, 0xe9, 0x6d, 0xb1, 0xff, 0x8a, 0xee, 0xdd, 0xfd, 0x6f, 0x2c, 0x33, 0x61, 0xfb, 0xe7,
	0xb2, 0xef, 0xfd, 0x5a, 0xf6, 0xbd, 0xdf, 0xcb, 0xbe, 0xf7, 0xb1, 0xb6, 0x18, 0x8e, 0x1b, 0xf8,
	0xab, 0x3d, 0xff, 0x13, 0x00, 0x00, 0xff, 0xff, 0xa3, 0x54, 0x0d, 0xac, 0xc7, 0x03, 0x00, 0x00,
}
