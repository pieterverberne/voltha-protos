// Code generated by protoc-gen-go. DO NOT EDIT.
// source: voltha_protos/logical_device.proto

package voltha // import "github.com/opencord/voltha-protos/go/voltha"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/opencord/voltha-protos/go/common"
import openflow_13 "github.com/opencord/voltha-protos/go/openflow_13"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type LogicalPortId struct {
	// unique id of logical device
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// id of the port on the logical device
	PortId               string   `protobuf:"bytes,2,opt,name=port_id,json=portId,proto3" json:"port_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogicalPortId) Reset()         { *m = LogicalPortId{} }
func (m *LogicalPortId) String() string { return proto.CompactTextString(m) }
func (*LogicalPortId) ProtoMessage()    {}
func (*LogicalPortId) Descriptor() ([]byte, []int) {
	return fileDescriptor_logical_device_eab9968d916b9ed2, []int{0}
}
func (m *LogicalPortId) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogicalPortId.Unmarshal(m, b)
}
func (m *LogicalPortId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogicalPortId.Marshal(b, m, deterministic)
}
func (dst *LogicalPortId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogicalPortId.Merge(dst, src)
}
func (m *LogicalPortId) XXX_Size() int {
	return xxx_messageInfo_LogicalPortId.Size(m)
}
func (m *LogicalPortId) XXX_DiscardUnknown() {
	xxx_messageInfo_LogicalPortId.DiscardUnknown(m)
}

var xxx_messageInfo_LogicalPortId proto.InternalMessageInfo

func (m *LogicalPortId) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *LogicalPortId) GetPortId() string {
	if m != nil {
		return m.PortId
	}
	return ""
}

type LogicalPort struct {
	Id                   string                    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	OfpPort              *openflow_13.OfpPort      `protobuf:"bytes,2,opt,name=ofp_port,json=ofpPort,proto3" json:"ofp_port,omitempty"`
	DeviceId             string                    `protobuf:"bytes,3,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	DevicePortNo         uint32                    `protobuf:"varint,4,opt,name=device_port_no,json=devicePortNo,proto3" json:"device_port_no,omitempty"`
	RootPort             bool                      `protobuf:"varint,5,opt,name=root_port,json=rootPort,proto3" json:"root_port,omitempty"`
	OfpPortStats         *openflow_13.OfpPortStats `protobuf:"bytes,6,opt,name=ofp_port_stats,json=ofpPortStats,proto3" json:"ofp_port_stats,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *LogicalPort) Reset()         { *m = LogicalPort{} }
func (m *LogicalPort) String() string { return proto.CompactTextString(m) }
func (*LogicalPort) ProtoMessage()    {}
func (*LogicalPort) Descriptor() ([]byte, []int) {
	return fileDescriptor_logical_device_eab9968d916b9ed2, []int{1}
}
func (m *LogicalPort) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogicalPort.Unmarshal(m, b)
}
func (m *LogicalPort) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogicalPort.Marshal(b, m, deterministic)
}
func (dst *LogicalPort) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogicalPort.Merge(dst, src)
}
func (m *LogicalPort) XXX_Size() int {
	return xxx_messageInfo_LogicalPort.Size(m)
}
func (m *LogicalPort) XXX_DiscardUnknown() {
	xxx_messageInfo_LogicalPort.DiscardUnknown(m)
}

var xxx_messageInfo_LogicalPort proto.InternalMessageInfo

func (m *LogicalPort) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *LogicalPort) GetOfpPort() *openflow_13.OfpPort {
	if m != nil {
		return m.OfpPort
	}
	return nil
}

func (m *LogicalPort) GetDeviceId() string {
	if m != nil {
		return m.DeviceId
	}
	return ""
}

func (m *LogicalPort) GetDevicePortNo() uint32 {
	if m != nil {
		return m.DevicePortNo
	}
	return 0
}

func (m *LogicalPort) GetRootPort() bool {
	if m != nil {
		return m.RootPort
	}
	return false
}

func (m *LogicalPort) GetOfpPortStats() *openflow_13.OfpPortStats {
	if m != nil {
		return m.OfpPortStats
	}
	return nil
}

type LogicalPorts struct {
	Items                []*LogicalPort `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *LogicalPorts) Reset()         { *m = LogicalPorts{} }
func (m *LogicalPorts) String() string { return proto.CompactTextString(m) }
func (*LogicalPorts) ProtoMessage()    {}
func (*LogicalPorts) Descriptor() ([]byte, []int) {
	return fileDescriptor_logical_device_eab9968d916b9ed2, []int{2}
}
func (m *LogicalPorts) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogicalPorts.Unmarshal(m, b)
}
func (m *LogicalPorts) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogicalPorts.Marshal(b, m, deterministic)
}
func (dst *LogicalPorts) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogicalPorts.Merge(dst, src)
}
func (m *LogicalPorts) XXX_Size() int {
	return xxx_messageInfo_LogicalPorts.Size(m)
}
func (m *LogicalPorts) XXX_DiscardUnknown() {
	xxx_messageInfo_LogicalPorts.DiscardUnknown(m)
}

var xxx_messageInfo_LogicalPorts proto.InternalMessageInfo

func (m *LogicalPorts) GetItems() []*LogicalPort {
	if m != nil {
		return m.Items
	}
	return nil
}

type LogicalDevice struct {
	// unique id of logical device
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// unique datapath id for the logical device (used by the SDN controller)
	DatapathId uint64 `protobuf:"varint,2,opt,name=datapath_id,json=datapathId,proto3" json:"datapath_id,omitempty"`
	// device description
	Desc *openflow_13.OfpDesc `protobuf:"bytes,3,opt,name=desc,proto3" json:"desc,omitempty"`
	// device features
	SwitchFeatures *openflow_13.OfpSwitchFeatures `protobuf:"bytes,4,opt,name=switch_features,json=switchFeatures,proto3" json:"switch_features,omitempty"`
	// name of the root device anchoring logical device
	RootDeviceId string `protobuf:"bytes,5,opt,name=root_device_id,json=rootDeviceId,proto3" json:"root_device_id,omitempty"`
	// logical device ports
	Ports []*LogicalPort `protobuf:"bytes,128,rep,name=ports,proto3" json:"ports,omitempty"`
	// flows configured on the logical device
	Flows *openflow_13.Flows `protobuf:"bytes,129,opt,name=flows,proto3" json:"flows,omitempty"`
	// flow groups configured on the logical device
	FlowGroups           *openflow_13.FlowGroups `protobuf:"bytes,130,opt,name=flow_groups,json=flowGroups,proto3" json:"flow_groups,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *LogicalDevice) Reset()         { *m = LogicalDevice{} }
func (m *LogicalDevice) String() string { return proto.CompactTextString(m) }
func (*LogicalDevice) ProtoMessage()    {}
func (*LogicalDevice) Descriptor() ([]byte, []int) {
	return fileDescriptor_logical_device_eab9968d916b9ed2, []int{3}
}
func (m *LogicalDevice) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogicalDevice.Unmarshal(m, b)
}
func (m *LogicalDevice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogicalDevice.Marshal(b, m, deterministic)
}
func (dst *LogicalDevice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogicalDevice.Merge(dst, src)
}
func (m *LogicalDevice) XXX_Size() int {
	return xxx_messageInfo_LogicalDevice.Size(m)
}
func (m *LogicalDevice) XXX_DiscardUnknown() {
	xxx_messageInfo_LogicalDevice.DiscardUnknown(m)
}

var xxx_messageInfo_LogicalDevice proto.InternalMessageInfo

func (m *LogicalDevice) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *LogicalDevice) GetDatapathId() uint64 {
	if m != nil {
		return m.DatapathId
	}
	return 0
}

func (m *LogicalDevice) GetDesc() *openflow_13.OfpDesc {
	if m != nil {
		return m.Desc
	}
	return nil
}

func (m *LogicalDevice) GetSwitchFeatures() *openflow_13.OfpSwitchFeatures {
	if m != nil {
		return m.SwitchFeatures
	}
	return nil
}

func (m *LogicalDevice) GetRootDeviceId() string {
	if m != nil {
		return m.RootDeviceId
	}
	return ""
}

func (m *LogicalDevice) GetPorts() []*LogicalPort {
	if m != nil {
		return m.Ports
	}
	return nil
}

func (m *LogicalDevice) GetFlows() *openflow_13.Flows {
	if m != nil {
		return m.Flows
	}
	return nil
}

func (m *LogicalDevice) GetFlowGroups() *openflow_13.FlowGroups {
	if m != nil {
		return m.FlowGroups
	}
	return nil
}

type LogicalDevices struct {
	Items                []*LogicalDevice `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *LogicalDevices) Reset()         { *m = LogicalDevices{} }
func (m *LogicalDevices) String() string { return proto.CompactTextString(m) }
func (*LogicalDevices) ProtoMessage()    {}
func (*LogicalDevices) Descriptor() ([]byte, []int) {
	return fileDescriptor_logical_device_eab9968d916b9ed2, []int{4}
}
func (m *LogicalDevices) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogicalDevices.Unmarshal(m, b)
}
func (m *LogicalDevices) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogicalDevices.Marshal(b, m, deterministic)
}
func (dst *LogicalDevices) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogicalDevices.Merge(dst, src)
}
func (m *LogicalDevices) XXX_Size() int {
	return xxx_messageInfo_LogicalDevices.Size(m)
}
func (m *LogicalDevices) XXX_DiscardUnknown() {
	xxx_messageInfo_LogicalDevices.DiscardUnknown(m)
}

var xxx_messageInfo_LogicalDevices proto.InternalMessageInfo

func (m *LogicalDevices) GetItems() []*LogicalDevice {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
	proto.RegisterType((*LogicalPortId)(nil), "voltha.LogicalPortId")
	proto.RegisterType((*LogicalPort)(nil), "voltha.LogicalPort")
	proto.RegisterType((*LogicalPorts)(nil), "voltha.LogicalPorts")
	proto.RegisterType((*LogicalDevice)(nil), "voltha.LogicalDevice")
	proto.RegisterType((*LogicalDevices)(nil), "voltha.LogicalDevices")
}

func init() {
	proto.RegisterFile("voltha_protos/logical_device.proto", fileDescriptor_logical_device_eab9968d916b9ed2)
}

var fileDescriptor_logical_device_eab9968d916b9ed2 = []byte{
	// 512 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x53, 0x4f, 0x6b, 0x13, 0x41,
	0x14, 0x77, 0x9b, 0x6c, 0x9a, 0xbc, 0x4d, 0x23, 0x8c, 0x94, 0x2e, 0xad, 0xd2, 0xb0, 0x78, 0x48,
	0x29, 0x4d, 0x6a, 0x8a, 0xa0, 0x07, 0x41, 0x43, 0xa9, 0x04, 0x44, 0x64, 0xbc, 0x79, 0x59, 0xa6,
	0x3b, 0x9b, 0xcd, 0x40, 0x92, 0xb7, 0xec, 0x4c, 0xda, 0xab, 0x7a, 0xf1, 0xd3, 0xf9, 0x25, 0xfc,
	0x08, 0x1e, 0x3c, 0xcb, 0xbc, 0xd9, 0xad, 0xd9, 0xa6, 0x1e, 0xf7, 0xf7, 0xe7, 0xbd, 0xdf, 0xfc,
	0x1e, 0x0b, 0xd1, 0x0d, 0x2e, 0xcc, 0x5c, 0xc4, 0x79, 0x81, 0x06, 0xf5, 0x68, 0x81, 0x99, 0x4a,
	0xc4, 0x22, 0x96, 0xe9, 0x8d, 0x4a, 0xd2, 0x21, 0xa1, 0xac, 0xe5, 0x34, 0x87, 0x4f, 0x33, 0xc4,
	0x6c, 0x91, 0x8e, 0x44, 0xae, 0x46, 0x62, 0xb5, 0x42, 0x23, 0x8c, 0xc2, 0x95, 0x76, 0xaa, 0xc3,
	0xb0, 0x3e, 0x69, 0x99, 0x1a, 0x51, 0x32, 0xc7, 0x75, 0x06, 0xf3, 0x74, 0x35, 0x5b, 0xe0, 0x6d,
	0xfc, 0xe2, 0xc2, 0x09, 0xa2, 0x57, 0xb0, 0xf7, 0xc1, 0x2d, 0xfe, 0x84, 0x85, 0x99, 0x4a, 0xd6,
	0x83, 0x1d, 0x25, 0x43, 0xaf, 0xef, 0x0d, 0x3a, 0x7c, 0x47, 0x49, 0x76, 0x00, 0xbb, 0x39, 0x16,
	0x26, 0x56, 0x32, 0xdc, 0x21, 0xb0, 0x95, 0x93, 0x30, 0xfa, 0xed, 0x41, 0xb0, 0x61, 0xdd, 0x32,
	0x9e, 0x43, 0x1b, 0x67, 0x79, 0x6c, 0xd5, 0xe4, 0x0c, 0xc6, 0xfb, 0xc3, 0xcd, 0xfd, 0x15, 0xc9,
	0x77, 0x71, 0x96, 0xd3, 0x84, 0x23, 0xe8, 0xb8, 0xc7, 0xdb, 0x65, 0x0d, 0x1a, 0xd4, 0x76, 0xc0,
	0x54, 0xb2, 0xe7, 0xd0, 0x2b, 0x49, 0x8a, 0xb3, 0xc2, 0xb0, 0xd9, 0xf7, 0x06, 0x7b, 0xbc, 0xeb,
	0x50, 0x3b, 0xe0, 0x23, 0xda, 0x11, 0x05, 0xa2, 0x71, 0x5b, 0xfd, 0xbe, 0x37, 0x68, 0xf3, 0xb6,
	0x05, 0x68, 0xfe, 0x3b, 0xe8, 0x55, 0x4b, 0x63, 0x6d, 0x84, 0xd1, 0x61, 0x8b, 0x72, 0x1d, 0x3d,
	0x98, 0xcb, 0x49, 0x78, 0xb7, 0x4c, 0xf7, 0xd9, 0x7e, 0x45, 0xaf, 0xa1, 0xbb, 0xf1, 0x66, 0xcd,
	0x4e, 0xc0, 0x57, 0x26, 0x5d, 0xea, 0xd0, 0xeb, 0x37, 0x06, 0xc1, 0xf8, 0xc9, 0xd0, 0xf5, 0x3d,
	0xdc, 0x10, 0x71, 0xa7, 0x88, 0x7e, 0x34, 0xee, 0xaa, 0xbe, 0xa4, 0xc8, 0x5b, 0x8d, 0x1d, 0x43,
	0x20, 0x85, 0x11, 0xb9, 0x30, 0xf3, 0xaa, 0xee, 0x26, 0x87, 0x0a, 0x9a, 0x4a, 0x76, 0x02, 0x4d,
	0x99, 0xea, 0x84, 0xba, 0x79, 0xa8, 0x4e, 0x4b, 0x72, 0x92, 0xb0, 0x29, 0x3c, 0xd6, 0xb7, 0xca,
	0x24, 0xf3, 0x78, 0x96, 0x0a, 0xb3, 0x2e, 0x52, 0x4d, 0x7d, 0x05, 0xe3, 0xfe, 0x96, 0xeb, 0x9e,
	0x8e, 0xf7, 0x1c, 0x70, 0x55, 0x7e, 0xdb, 0xe6, 0xa9, 0xd3, 0x7f, 0xb7, 0xf1, 0x29, 0x72, 0xd7,
	0xa2, 0x97, 0xd5, 0x7d, 0x5e, 0x82, 0x6f, 0x5b, 0xd3, 0xe1, 0xd7, 0xff, 0x57, 0x31, 0xe9, 0xfc,
	0xfa, 0xf3, 0xf3, 0x59, 0xd3, 0x3e, 0x9b, 0x3b, 0x35, 0x3b, 0x07, 0xdf, 0x66, 0xd1, 0xe1, 0x37,
	0x8f, 0xe2, 0xb1, 0x5a, 0xbc, 0x2b, 0x4b, 0x4d, 0x7c, 0xeb, 0x7a, 0xc4, 0x9d, 0x90, 0xbd, 0x85,
	0x80, 0xe8, 0xac, 0xc0, 0x75, 0xae, 0xc3, 0xef, 0xce, 0x77, 0xb0, 0xe5, 0x7b, 0x4f, 0x7c, 0x65,
	0x86, 0xd9, 0x1d, 0x14, 0xbd, 0x81, 0x5e, 0xed, 0x10, 0x9a, 0x9d, 0xd6, 0xcf, 0xb8, 0x7f, 0x2f,
	0xbb, 0x93, 0x95, 0x87, 0x9c, 0x9c, 0x7d, 0x39, 0xcd, 0x94, 0x99, 0xaf, 0xaf, 0x87, 0x09, 0x2e,
	0xe9, 0x97, 0x4a, 0xb0, 0x90, 0x23, 0x67, 0x39, 0x2b, 0xff, 0xb4, 0x0c, 0x4b, 0xe0, 0xba, 0x45,
	0xc8, 0xc5, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x23, 0x92, 0xcd, 0x9e, 0xef, 0x03, 0x00, 0x00,
}
