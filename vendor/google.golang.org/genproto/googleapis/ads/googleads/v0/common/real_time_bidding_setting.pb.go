// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/common/real_time_bidding_setting.proto

package common // import "google.golang.org/genproto/googleapis/ads/googleads/v0/common"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import wrappers "github.com/golang/protobuf/ptypes/wrappers"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Settings for Real-Time Bidding, a feature only available for campaigns
// targeting the Ad Exchange network.
type RealTimeBiddingSetting struct {
	// Whether the campaign is opted in to real-time bidding.
	OptIn                *wrappers.BoolValue `protobuf:"bytes,1,opt,name=opt_in,json=optIn,proto3" json:"opt_in,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *RealTimeBiddingSetting) Reset()         { *m = RealTimeBiddingSetting{} }
func (m *RealTimeBiddingSetting) String() string { return proto.CompactTextString(m) }
func (*RealTimeBiddingSetting) ProtoMessage()    {}
func (*RealTimeBiddingSetting) Descriptor() ([]byte, []int) {
	return fileDescriptor_real_time_bidding_setting_bb30570aed4e7c68, []int{0}
}
func (m *RealTimeBiddingSetting) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RealTimeBiddingSetting.Unmarshal(m, b)
}
func (m *RealTimeBiddingSetting) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RealTimeBiddingSetting.Marshal(b, m, deterministic)
}
func (dst *RealTimeBiddingSetting) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RealTimeBiddingSetting.Merge(dst, src)
}
func (m *RealTimeBiddingSetting) XXX_Size() int {
	return xxx_messageInfo_RealTimeBiddingSetting.Size(m)
}
func (m *RealTimeBiddingSetting) XXX_DiscardUnknown() {
	xxx_messageInfo_RealTimeBiddingSetting.DiscardUnknown(m)
}

var xxx_messageInfo_RealTimeBiddingSetting proto.InternalMessageInfo

func (m *RealTimeBiddingSetting) GetOptIn() *wrappers.BoolValue {
	if m != nil {
		return m.OptIn
	}
	return nil
}

func init() {
	proto.RegisterType((*RealTimeBiddingSetting)(nil), "google.ads.googleads.v0.common.RealTimeBiddingSetting")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/common/real_time_bidding_setting.proto", fileDescriptor_real_time_bidding_setting_bb30570aed4e7c68)
}

var fileDescriptor_real_time_bidding_setting_bb30570aed4e7c68 = []byte{
	// 269 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x41, 0x4b, 0xc3, 0x30,
	0x18, 0x86, 0xe9, 0xc4, 0x1d, 0xea, 0x6d, 0x07, 0x91, 0x09, 0x43, 0x76, 0xf2, 0xf4, 0xa5, 0xea,
	0x51, 0x10, 0x5a, 0x85, 0x21, 0x5e, 0xc6, 0x94, 0x1e, 0xa4, 0x50, 0xd2, 0x25, 0x86, 0x40, 0x9a,
	0x2f, 0x24, 0xd9, 0xfc, 0x3f, 0x1e, 0xfd, 0x29, 0xde, 0xfc, 0x47, 0xd2, 0x7e, 0xed, 0x4e, 0xea,
	0xa9, 0x2f, 0xf4, 0x7d, 0x9e, 0x37, 0x7c, 0xe9, 0x9d, 0x42, 0x54, 0x46, 0x32, 0x2e, 0x02, 0xa3,
	0xd8, 0xa5, 0x7d, 0xc6, 0xb6, 0xd8, 0xb6, 0x68, 0x99, 0x97, 0xdc, 0xd4, 0x51, 0xb7, 0xb2, 0x6e,
	0xb4, 0x10, 0xda, 0xaa, 0x3a, 0xc8, 0x18, 0xb5, 0x55, 0xe0, 0x3c, 0x46, 0x9c, 0x2d, 0x08, 0x02,
	0x2e, 0x02, 0x1c, 0x78, 0xd8, 0x67, 0x40, 0xfc, 0x7c, 0xf8, 0xcf, 0xfa, 0x76, 0xb3, 0x7b, 0x63,
	0xef, 0x9e, 0x3b, 0x27, 0x7d, 0x20, 0x7e, 0xf9, 0x94, 0x9e, 0x6e, 0x24, 0x37, 0x2f, 0xba, 0x95,
	0x05, 0x0d, 0x3c, 0x93, 0x7f, 0x76, 0x95, 0x4e, 0xd1, 0xc5, 0x5a, 0xdb, 0xb3, 0xe4, 0x22, 0xb9,
	0x3c, 0xb9, 0x9e, 0x0f, 0x7e, 0x18, 0x55, 0x50, 0x20, 0x9a, 0x92, 0x9b, 0x9d, 0xdc, 0x1c, 0xa3,
	0x8b, 0x8f, 0xb6, 0xf8, 0x4e, 0xd2, 0xe5, 0x16, 0x5b, 0xf8, 0xff, 0x4d, 0xc5, 0xf9, 0xef, 0x8b,
	0xeb, 0xce, 0xbb, 0x4e, 0x5e, 0x1f, 0x06, 0x5c, 0xa1, 0xe1, 0x56, 0x01, 0x7a, 0xc5, 0x94, 0xb4,
	0xfd, 0xea, 0x78, 0x22, 0xa7, 0xc3, 0x5f, 0x17, 0xbb, 0xa5, 0xcf, 0xc7, 0xe4, 0x68, 0x95, 0xe7,
	0x9f, 0x93, 0xc5, 0x8a, 0x64, 0xb9, 0x08, 0x40, 0xb1, 0x4b, 0x65, 0x06, 0xf7, 0x7d, 0xed, 0x6b,
	0x2c, 0x54, 0xb9, 0x08, 0xd5, 0xa1, 0x50, 0x95, 0x59, 0x45, 0x85, 0x66, 0xda, 0x0f, 0xdf, 0xfc,
	0x04, 0x00, 0x00, 0xff, 0xff, 0x06, 0x56, 0x26, 0xd9, 0xa9, 0x01, 0x00, 0x00,
}