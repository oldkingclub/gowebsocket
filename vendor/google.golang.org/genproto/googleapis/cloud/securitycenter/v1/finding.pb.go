// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/securitycenter/v1/finding.proto

package securitycenter

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_struct "github.com/golang/protobuf/ptypes/struct"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// The state of the finding.
type Finding_State int32

const (
	// Unspecified state.
	Finding_STATE_UNSPECIFIED Finding_State = 0
	// The finding requires attention and has not been addressed yet.
	Finding_ACTIVE Finding_State = 1
	// The finding has been fixed, triaged as a non-issue or otherwise addressed
	// and is no longer active.
	Finding_INACTIVE Finding_State = 2
)

var Finding_State_name = map[int32]string{
	0: "STATE_UNSPECIFIED",
	1: "ACTIVE",
	2: "INACTIVE",
}

var Finding_State_value = map[string]int32{
	"STATE_UNSPECIFIED": 0,
	"ACTIVE":            1,
	"INACTIVE":          2,
}

func (x Finding_State) String() string {
	return proto.EnumName(Finding_State_name, int32(x))
}

func (Finding_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b75f23af5a63668d, []int{0, 0}
}

// Cloud Security Command Center (Cloud SCC) finding.
//
// A finding is a record of assessment data (security, risk, health or privacy)
// ingested into Cloud SCC for presentation, notification, analysis,
// policy testing, and enforcement. For example, an XSS vulnerability in an
// App Engine application is a finding.
type Finding struct {
	// The relative resource name of this finding. See:
	// https://cloud.google.com/apis/design/resource_names#relative_resource_name
	// Example:
	// "organizations/123/sources/456/findings/789"
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The relative resource name of the source the finding belongs to. See:
	// https://cloud.google.com/apis/design/resource_names#relative_resource_name
	// This field is immutable after creation time.
	// For example:
	// "organizations/123/sources/456"
	Parent string `protobuf:"bytes,2,opt,name=parent,proto3" json:"parent,omitempty"`
	// The full resource name of the Google Cloud Platform (GCP) resource this
	// finding is for. See:
	// https://cloud.google.com/apis/design/resource_names#full_resource_name
	// This field is immutable after creation time.
	ResourceName string `protobuf:"bytes,3,opt,name=resource_name,json=resourceName,proto3" json:"resource_name,omitempty"`
	// The state of the finding.
	State Finding_State `protobuf:"varint,4,opt,name=state,proto3,enum=google.cloud.securitycenter.v1.Finding_State" json:"state,omitempty"`
	// The additional taxonomy group within findings from a given source.
	// This field is immutable after creation time.
	// Example: "XSS_FLASH_INJECTION"
	Category string `protobuf:"bytes,5,opt,name=category,proto3" json:"category,omitempty"`
	// The URI that, if available, points to a web page outside of Cloud SCC
	// where additional information about the finding can be found. This field is
	// guaranteed to be either empty or a well formed URL.
	ExternalUri string `protobuf:"bytes,6,opt,name=external_uri,json=externalUri,proto3" json:"external_uri,omitempty"`
	// Source specific properties. These properties are managed by the source
	// that writes the finding. The key names in the source_properties map must be
	// between 1 and 255 characters, and must start with a letter and contain
	// alphanumeric characters or underscores only.
	SourceProperties map[string]*_struct.Value `protobuf:"bytes,7,rep,name=source_properties,json=sourceProperties,proto3" json:"source_properties,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Output only. User specified security marks. These marks are entirely
	// managed by the user and come from the SecurityMarks resource that belongs
	// to the finding.
	SecurityMarks *SecurityMarks `protobuf:"bytes,8,opt,name=security_marks,json=securityMarks,proto3" json:"security_marks,omitempty"`
	// The time at which the event took place. For example, if the finding
	// represents an open firewall it would capture the time the open firewall was
	// detected.
	EventTime *timestamp.Timestamp `protobuf:"bytes,9,opt,name=event_time,json=eventTime,proto3" json:"event_time,omitempty"`
	// The time at which the finding was created in Cloud SCC.
	CreateTime           *timestamp.Timestamp `protobuf:"bytes,10,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Finding) Reset()         { *m = Finding{} }
func (m *Finding) String() string { return proto.CompactTextString(m) }
func (*Finding) ProtoMessage()    {}
func (*Finding) Descriptor() ([]byte, []int) {
	return fileDescriptor_b75f23af5a63668d, []int{0}
}

func (m *Finding) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Finding.Unmarshal(m, b)
}
func (m *Finding) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Finding.Marshal(b, m, deterministic)
}
func (m *Finding) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Finding.Merge(m, src)
}
func (m *Finding) XXX_Size() int {
	return xxx_messageInfo_Finding.Size(m)
}
func (m *Finding) XXX_DiscardUnknown() {
	xxx_messageInfo_Finding.DiscardUnknown(m)
}

var xxx_messageInfo_Finding proto.InternalMessageInfo

func (m *Finding) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Finding) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *Finding) GetResourceName() string {
	if m != nil {
		return m.ResourceName
	}
	return ""
}

func (m *Finding) GetState() Finding_State {
	if m != nil {
		return m.State
	}
	return Finding_STATE_UNSPECIFIED
}

func (m *Finding) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *Finding) GetExternalUri() string {
	if m != nil {
		return m.ExternalUri
	}
	return ""
}

func (m *Finding) GetSourceProperties() map[string]*_struct.Value {
	if m != nil {
		return m.SourceProperties
	}
	return nil
}

func (m *Finding) GetSecurityMarks() *SecurityMarks {
	if m != nil {
		return m.SecurityMarks
	}
	return nil
}

func (m *Finding) GetEventTime() *timestamp.Timestamp {
	if m != nil {
		return m.EventTime
	}
	return nil
}

func (m *Finding) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func init() {
	proto.RegisterEnum("google.cloud.securitycenter.v1.Finding_State", Finding_State_name, Finding_State_value)
	proto.RegisterType((*Finding)(nil), "google.cloud.securitycenter.v1.Finding")
	proto.RegisterMapType((map[string]*_struct.Value)(nil), "google.cloud.securitycenter.v1.Finding.SourcePropertiesEntry")
}

func init() {
	proto.RegisterFile("google/cloud/securitycenter/v1/finding.proto", fileDescriptor_b75f23af5a63668d)
}

var fileDescriptor_b75f23af5a63668d = []byte{
	// 555 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xcf, 0x6e, 0xda, 0x40,
	0x10, 0xc6, 0x6b, 0x08, 0x04, 0x06, 0x12, 0x91, 0x95, 0x12, 0x59, 0xa8, 0x4a, 0x09, 0xbd, 0x70,
	0x48, 0x6d, 0x41, 0x2e, 0xa9, 0xa3, 0x1e, 0x52, 0x4a, 0x2a, 0xa4, 0x16, 0x21, 0x43, 0x38, 0xb4,
	0x48, 0x68, 0xe3, 0x4c, 0x2c, 0x37, 0xb0, 0x6b, 0xad, 0xd7, 0xa8, 0xbc, 0x52, 0x2f, 0x7d, 0x8f,
	0x3e, 0x42, 0x1f, 0xa1, 0x4f, 0x51, 0x79, 0xd7, 0x8e, 0xea, 0xf4, 0x0f, 0xbd, 0x79, 0xbe, 0xf9,
	0x7d, 0x33, 0xb3, 0x3b, 0x6b, 0x38, 0xf5, 0x39, 0xf7, 0x97, 0x68, 0x7b, 0x4b, 0x1e, 0xdf, 0xda,
	0x11, 0x7a, 0xb1, 0x08, 0xe4, 0xc6, 0x43, 0x26, 0x51, 0xd8, 0xeb, 0xae, 0x7d, 0x17, 0xb0, 0xdb,
	0x80, 0xf9, 0x56, 0x28, 0xb8, 0xe4, 0xe4, 0x58, 0xd3, 0x96, 0xa2, 0xad, 0x3c, 0x6d, 0xad, 0xbb,
	0xcd, 0xa7, 0x69, 0x35, 0x1a, 0x06, 0x36, 0x65, 0x8c, 0x4b, 0x2a, 0x03, 0xce, 0x22, 0xed, 0x6e,
	0x9e, 0x6d, 0xe9, 0x95, 0x29, 0x8b, 0x15, 0x15, 0xf7, 0x99, 0x29, 0x2b, 0xa9, 0xa2, 0x9b, 0xf8,
	0xce, 0x8e, 0xa4, 0x88, 0x3d, 0x99, 0x66, 0x9f, 0x3d, 0xce, 0xca, 0x60, 0x85, 0x91, 0xa4, 0xab,
	0x50, 0x03, 0xed, 0xaf, 0x25, 0xd8, 0xbd, 0xd2, 0x67, 0x20, 0x04, 0x76, 0x18, 0x5d, 0xa1, 0x69,
	0xb4, 0x8c, 0x4e, 0xd5, 0x55, 0xdf, 0xe4, 0x08, 0xca, 0x21, 0x15, 0xc8, 0xa4, 0x59, 0x50, 0x6a,
	0x1a, 0x91, 0xe7, 0xb0, 0x27, 0x30, 0xe2, 0xb1, 0xf0, 0x70, 0xa1, 0x4c, 0x45, 0x95, 0xae, 0x67,
	0xe2, 0x28, 0x31, 0xf7, 0xa1, 0x14, 0x49, 0x2a, 0xd1, 0xdc, 0x69, 0x19, 0x9d, 0xfd, 0xde, 0x0b,
	0xeb, 0xdf, 0xd7, 0x63, 0xa5, 0x83, 0x58, 0x93, 0xc4, 0xe4, 0x6a, 0x2f, 0x69, 0x42, 0xc5, 0xa3,
	0x12, 0x7d, 0x2e, 0x36, 0x66, 0x49, 0x35, 0x79, 0x88, 0xc9, 0x09, 0xd4, 0xf1, 0xb3, 0x44, 0xc1,
	0xe8, 0x72, 0x11, 0x8b, 0xc0, 0x2c, 0xab, 0x7c, 0x2d, 0xd3, 0xae, 0x45, 0x40, 0x3e, 0xc1, 0x41,
	0x3a, 0x66, 0x28, 0x78, 0x88, 0x42, 0x06, 0x18, 0x99, 0xbb, 0xad, 0x62, 0xa7, 0xd6, 0x7b, 0xf5,
	0xdf, 0xf3, 0xa8, 0x02, 0xe3, 0x07, 0xff, 0x80, 0x49, 0xb1, 0x71, 0x1b, 0xd1, 0x23, 0x99, 0x4c,
	0x61, 0x3f, 0xbf, 0x23, 0xb3, 0xd2, 0x32, 0x3a, 0xb5, 0xed, 0x07, 0x9f, 0xa4, 0xca, 0xfb, 0xc4,
	0xe4, 0xee, 0x45, 0xbf, 0x86, 0xe4, 0x25, 0x00, 0xae, 0x91, 0xc9, 0x45, 0xb2, 0x3b, 0xb3, 0xaa,
	0x2a, 0x36, 0xb3, 0x8a, 0xd9, 0x62, 0xad, 0x69, 0xb6, 0x58, 0xb7, 0xaa, 0xe8, 0x24, 0x26, 0x17,
	0x50, 0xf3, 0x04, 0x52, 0x89, 0xda, 0x0b, 0x5b, 0xbd, 0xa0, 0xf1, 0x44, 0x68, 0x7e, 0x84, 0xc3,
	0x3f, 0x1e, 0x9c, 0x34, 0xa0, 0x78, 0x8f, 0x9b, 0xf4, 0x99, 0x24, 0x9f, 0xe4, 0x14, 0x4a, 0x6b,
	0xba, 0x8c, 0x51, 0x3d, 0x92, 0x5a, 0xef, 0xe8, 0xb7, 0x0e, 0xb3, 0x24, 0xeb, 0x6a, 0xc8, 0x29,
	0x9c, 0x1b, 0xed, 0x73, 0x28, 0xa9, 0x2d, 0x93, 0x43, 0x38, 0x98, 0x4c, 0x2f, 0xa7, 0x83, 0xc5,
	0xf5, 0x68, 0x32, 0x1e, 0xf4, 0x87, 0x57, 0xc3, 0xc1, 0x9b, 0xc6, 0x13, 0x02, 0x50, 0xbe, 0xec,
	0x4f, 0x87, 0xb3, 0x41, 0xc3, 0x20, 0x75, 0xa8, 0x0c, 0x47, 0x69, 0x54, 0x78, 0xfd, 0xdd, 0x80,
	0xb6, 0xc7, 0x57, 0x5b, 0xae, 0x74, 0x6c, 0x7c, 0x78, 0x97, 0x12, 0x3e, 0x5f, 0x52, 0xe6, 0x5b,
	0x5c, 0xf8, 0xb6, 0x8f, 0x4c, 0x8d, 0x64, 0xeb, 0x14, 0x0d, 0x83, 0xe8, 0x6f, 0x7f, 0xdb, 0x45,
	0x5e, 0xf9, 0x52, 0x38, 0x7e, 0xab, 0xcb, 0xf5, 0x55, 0xc3, 0x6c, 0x63, 0x7d, 0xdd, 0x70, 0xd6,
	0xfd, 0x96, 0x01, 0x73, 0x05, 0xcc, 0xf3, 0xc0, 0x7c, 0xd6, 0xfd, 0x51, 0x38, 0xd1, 0x80, 0xe3,
	0x28, 0xc2, 0x71, 0xf2, 0x88, 0xe3, 0xcc, 0xba, 0x37, 0x65, 0x35, 0xde, 0xd9, 0xcf, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x39, 0x58, 0x82, 0xa6, 0x77, 0x04, 0x00, 0x00,
}
