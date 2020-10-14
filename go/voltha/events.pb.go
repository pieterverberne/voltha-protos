// Code generated by protoc-gen-go. DO NOT EDIT.
// source: voltha_protos/events.proto

package voltha

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/opencord/voltha-protos/v4/go/common"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type ConfigEventType_Types int32

const (
	ConfigEventType_add    ConfigEventType_Types = 0
	ConfigEventType_remove ConfigEventType_Types = 1
	ConfigEventType_update ConfigEventType_Types = 2
)

var ConfigEventType_Types_name = map[int32]string{
	0: "add",
	1: "remove",
	2: "update",
}

var ConfigEventType_Types_value = map[string]int32{
	"add":    0,
	"remove": 1,
	"update": 2,
}

func (x ConfigEventType_Types) String() string {
	return proto.EnumName(ConfigEventType_Types_name, int32(x))
}

func (ConfigEventType_Types) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e63e6c07044fd2c4, []int{0, 0}
}

type KpiEventType_Types int32

const (
	KpiEventType_slice KpiEventType_Types = 0
	KpiEventType_ts    KpiEventType_Types = 1
)

var KpiEventType_Types_name = map[int32]string{
	0: "slice",
	1: "ts",
}

var KpiEventType_Types_value = map[string]int32{
	"slice": 0,
	"ts":    1,
}

func (x KpiEventType_Types) String() string {
	return proto.EnumName(KpiEventType_Types_name, int32(x))
}

func (KpiEventType_Types) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e63e6c07044fd2c4, []int{2, 0}
}

type EventCategory_Types int32

const (
	EventCategory_COMMUNICATION EventCategory_Types = 0
	EventCategory_ENVIRONMENT   EventCategory_Types = 1
	EventCategory_EQUIPMENT     EventCategory_Types = 2
	EventCategory_SERVICE       EventCategory_Types = 3
	EventCategory_PROCESSING    EventCategory_Types = 4
	EventCategory_SECURITY      EventCategory_Types = 5
)

var EventCategory_Types_name = map[int32]string{
	0: "COMMUNICATION",
	1: "ENVIRONMENT",
	2: "EQUIPMENT",
	3: "SERVICE",
	4: "PROCESSING",
	5: "SECURITY",
}

var EventCategory_Types_value = map[string]int32{
	"COMMUNICATION": 0,
	"ENVIRONMENT":   1,
	"EQUIPMENT":     2,
	"SERVICE":       3,
	"PROCESSING":    4,
	"SECURITY":      5,
}

func (x EventCategory_Types) String() string {
	return proto.EnumName(EventCategory_Types_name, int32(x))
}

func (EventCategory_Types) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e63e6c07044fd2c4, []int{9, 0}
}

type EventSubCategory_Types int32

const (
	EventSubCategory_PON EventSubCategory_Types = 0
	EventSubCategory_OLT EventSubCategory_Types = 1
	EventSubCategory_ONT EventSubCategory_Types = 2
	EventSubCategory_ONU EventSubCategory_Types = 3
	EventSubCategory_NNI EventSubCategory_Types = 4
)

var EventSubCategory_Types_name = map[int32]string{
	0: "PON",
	1: "OLT",
	2: "ONT",
	3: "ONU",
	4: "NNI",
}

var EventSubCategory_Types_value = map[string]int32{
	"PON": 0,
	"OLT": 1,
	"ONT": 2,
	"ONU": 3,
	"NNI": 4,
}

func (x EventSubCategory_Types) String() string {
	return proto.EnumName(EventSubCategory_Types_name, int32(x))
}

func (EventSubCategory_Types) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e63e6c07044fd2c4, []int{10, 0}
}

type EventType_Types int32

const (
	EventType_CONFIG_EVENT EventType_Types = 0
	EventType_KPI_EVENT    EventType_Types = 1
	EventType_KPI_EVENT2   EventType_Types = 2
	EventType_DEVICE_EVENT EventType_Types = 3
)

var EventType_Types_name = map[int32]string{
	0: "CONFIG_EVENT",
	1: "KPI_EVENT",
	2: "KPI_EVENT2",
	3: "DEVICE_EVENT",
}

var EventType_Types_value = map[string]int32{
	"CONFIG_EVENT": 0,
	"KPI_EVENT":    1,
	"KPI_EVENT2":   2,
	"DEVICE_EVENT": 3,
}

func (x EventType_Types) String() string {
	return proto.EnumName(EventType_Types_name, int32(x))
}

func (EventType_Types) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_e63e6c07044fd2c4, []int{11, 0}
}

type ConfigEventType struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ConfigEventType) Reset()         { *m = ConfigEventType{} }
func (m *ConfigEventType) String() string { return proto.CompactTextString(m) }
func (*ConfigEventType) ProtoMessage()    {}
func (*ConfigEventType) Descriptor() ([]byte, []int) {
	return fileDescriptor_e63e6c07044fd2c4, []int{0}
}

func (m *ConfigEventType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigEventType.Unmarshal(m, b)
}
func (m *ConfigEventType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigEventType.Marshal(b, m, deterministic)
}
func (m *ConfigEventType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigEventType.Merge(m, src)
}
func (m *ConfigEventType) XXX_Size() int {
	return xxx_messageInfo_ConfigEventType.Size(m)
}
func (m *ConfigEventType) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigEventType.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigEventType proto.InternalMessageInfo

type ConfigEvent struct {
	Type                 ConfigEventType_Types `protobuf:"varint,1,opt,name=type,proto3,enum=voltha.ConfigEventType_Types" json:"type,omitempty"`
	Hash                 string                `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	Data                 string                `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *ConfigEvent) Reset()         { *m = ConfigEvent{} }
func (m *ConfigEvent) String() string { return proto.CompactTextString(m) }
func (*ConfigEvent) ProtoMessage()    {}
func (*ConfigEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_e63e6c07044fd2c4, []int{1}
}

func (m *ConfigEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ConfigEvent.Unmarshal(m, b)
}
func (m *ConfigEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ConfigEvent.Marshal(b, m, deterministic)
}
func (m *ConfigEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConfigEvent.Merge(m, src)
}
func (m *ConfigEvent) XXX_Size() int {
	return xxx_messageInfo_ConfigEvent.Size(m)
}
func (m *ConfigEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_ConfigEvent.DiscardUnknown(m)
}

var xxx_messageInfo_ConfigEvent proto.InternalMessageInfo

func (m *ConfigEvent) GetType() ConfigEventType_Types {
	if m != nil {
		return m.Type
	}
	return ConfigEventType_add
}

func (m *ConfigEvent) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *ConfigEvent) GetData() string {
	if m != nil {
		return m.Data
	}
	return ""
}

type KpiEventType struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KpiEventType) Reset()         { *m = KpiEventType{} }
func (m *KpiEventType) String() string { return proto.CompactTextString(m) }
func (*KpiEventType) ProtoMessage()    {}
func (*KpiEventType) Descriptor() ([]byte, []int) {
	return fileDescriptor_e63e6c07044fd2c4, []int{2}
}

func (m *KpiEventType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KpiEventType.Unmarshal(m, b)
}
func (m *KpiEventType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KpiEventType.Marshal(b, m, deterministic)
}
func (m *KpiEventType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KpiEventType.Merge(m, src)
}
func (m *KpiEventType) XXX_Size() int {
	return xxx_messageInfo_KpiEventType.Size(m)
}
func (m *KpiEventType) XXX_DiscardUnknown() {
	xxx_messageInfo_KpiEventType.DiscardUnknown(m)
}

var xxx_messageInfo_KpiEventType proto.InternalMessageInfo

//
// Struct to convey a dictionary of metric metadata.
type MetricMetaData struct {
	Title           string  `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Ts              float64 `protobuf:"fixed64,2,opt,name=ts,proto3" json:"ts,omitempty"`
	LogicalDeviceId string  `protobuf:"bytes,3,opt,name=logical_device_id,json=logicalDeviceId,proto3" json:"logical_device_id,omitempty"`
	// (equivalent to the DPID that ONOS has
	// for the VOLTHA device without the
	//  'of:' prefix
	SerialNo             string            `protobuf:"bytes,4,opt,name=serial_no,json=serialNo,proto3" json:"serial_no,omitempty"`
	DeviceId             string            `protobuf:"bytes,5,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	Context              map[string]string `protobuf:"bytes,6,rep,name=context,proto3" json:"context,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Uuid                 string            `protobuf:"bytes,7,opt,name=uuid,proto3" json:"uuid,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *MetricMetaData) Reset()         { *m = MetricMetaData{} }
func (m *MetricMetaData) String() string { return proto.CompactTextString(m) }
func (*MetricMetaData) ProtoMessage()    {}
func (*MetricMetaData) Descriptor() ([]byte, []int) {
	return fileDescriptor_e63e6c07044fd2c4, []int{3}
}

func (m *MetricMetaData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetricMetaData.Unmarshal(m, b)
}
func (m *MetricMetaData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetricMetaData.Marshal(b, m, deterministic)
}
func (m *MetricMetaData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetricMetaData.Merge(m, src)
}
func (m *MetricMetaData) XXX_Size() int {
	return xxx_messageInfo_MetricMetaData.Size(m)
}
func (m *MetricMetaData) XXX_DiscardUnknown() {
	xxx_messageInfo_MetricMetaData.DiscardUnknown(m)
}

var xxx_messageInfo_MetricMetaData proto.InternalMessageInfo

func (m *MetricMetaData) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *MetricMetaData) GetTs() float64 {
	if m != nil {
		return m.Ts
	}
	return 0
}

func (m *MetricMetaData) GetLogicalDeviceId() string {
	if m != nil {
		return m.LogicalDeviceId
	}
	return ""
}

func (m *MetricMetaData) GetSerialNo() string {
	if m != nil {
		return m.SerialNo
	}
	return ""
}

func (m *MetricMetaData) GetDeviceId() string {
	if m != nil {
		return m.DeviceId
	}
	return ""
}

func (m *MetricMetaData) GetContext() map[string]string {
	if m != nil {
		return m.Context
	}
	return nil
}

func (m *MetricMetaData) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

//
// Struct to convey a dictionary of metric->value pairs. Typically used in
// pure shared-timestamp or shared-timestamp + shared object prefix situations.
type MetricValuePairs struct {
	// Metric / value pairs.
	Metrics              map[string]float32 `protobuf:"bytes,1,rep,name=metrics,proto3" json:"metrics,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed32,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *MetricValuePairs) Reset()         { *m = MetricValuePairs{} }
func (m *MetricValuePairs) String() string { return proto.CompactTextString(m) }
func (*MetricValuePairs) ProtoMessage()    {}
func (*MetricValuePairs) Descriptor() ([]byte, []int) {
	return fileDescriptor_e63e6c07044fd2c4, []int{4}
}

func (m *MetricValuePairs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetricValuePairs.Unmarshal(m, b)
}
func (m *MetricValuePairs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetricValuePairs.Marshal(b, m, deterministic)
}
func (m *MetricValuePairs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetricValuePairs.Merge(m, src)
}
func (m *MetricValuePairs) XXX_Size() int {
	return xxx_messageInfo_MetricValuePairs.Size(m)
}
func (m *MetricValuePairs) XXX_DiscardUnknown() {
	xxx_messageInfo_MetricValuePairs.DiscardUnknown(m)
}

var xxx_messageInfo_MetricValuePairs proto.InternalMessageInfo

func (m *MetricValuePairs) GetMetrics() map[string]float32 {
	if m != nil {
		return m.Metrics
	}
	return nil
}

//
// Struct to group metadata for a metric (or group of metrics) with the key-value
// pairs of collected metrics
type MetricInformation struct {
	Metadata             *MetricMetaData    `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Metrics              map[string]float32 `protobuf:"bytes,2,rep,name=metrics,proto3" json:"metrics,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed32,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *MetricInformation) Reset()         { *m = MetricInformation{} }
func (m *MetricInformation) String() string { return proto.CompactTextString(m) }
func (*MetricInformation) ProtoMessage()    {}
func (*MetricInformation) Descriptor() ([]byte, []int) {
	return fileDescriptor_e63e6c07044fd2c4, []int{5}
}

func (m *MetricInformation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetricInformation.Unmarshal(m, b)
}
func (m *MetricInformation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetricInformation.Marshal(b, m, deterministic)
}
func (m *MetricInformation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetricInformation.Merge(m, src)
}
func (m *MetricInformation) XXX_Size() int {
	return xxx_messageInfo_MetricInformation.Size(m)
}
func (m *MetricInformation) XXX_DiscardUnknown() {
	xxx_messageInfo_MetricInformation.DiscardUnknown(m)
}

var xxx_messageInfo_MetricInformation proto.InternalMessageInfo

func (m *MetricInformation) GetMetadata() *MetricMetaData {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *MetricInformation) GetMetrics() map[string]float32 {
	if m != nil {
		return m.Metrics
	}
	return nil
}

//
// Legacy KPI Event structured.  In mid-August, the KPI event format was updated
//                               to a more easily parsable format. See VOL-1140
//                               for more information.
type KpiEvent struct {
	Type                 KpiEventType_Types           `protobuf:"varint,1,opt,name=type,proto3,enum=voltha.KpiEventType_Types" json:"type,omitempty"`
	Ts                   float32                      `protobuf:"fixed32,2,opt,name=ts,proto3" json:"ts,omitempty"`
	Prefixes             map[string]*MetricValuePairs `protobuf:"bytes,3,rep,name=prefixes,proto3" json:"prefixes,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *KpiEvent) Reset()         { *m = KpiEvent{} }
func (m *KpiEvent) String() string { return proto.CompactTextString(m) }
func (*KpiEvent) ProtoMessage()    {}
func (*KpiEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_e63e6c07044fd2c4, []int{6}
}

func (m *KpiEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KpiEvent.Unmarshal(m, b)
}
func (m *KpiEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KpiEvent.Marshal(b, m, deterministic)
}
func (m *KpiEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KpiEvent.Merge(m, src)
}
func (m *KpiEvent) XXX_Size() int {
	return xxx_messageInfo_KpiEvent.Size(m)
}
func (m *KpiEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_KpiEvent.DiscardUnknown(m)
}

var xxx_messageInfo_KpiEvent proto.InternalMessageInfo

func (m *KpiEvent) GetType() KpiEventType_Types {
	if m != nil {
		return m.Type
	}
	return KpiEventType_slice
}

func (m *KpiEvent) GetTs() float32 {
	if m != nil {
		return m.Ts
	}
	return 0
}

func (m *KpiEvent) GetPrefixes() map[string]*MetricValuePairs {
	if m != nil {
		return m.Prefixes
	}
	return nil
}

type KpiEvent2 struct {
	// Type of KPI Event
	Type KpiEventType_Types `protobuf:"varint,1,opt,name=type,proto3,enum=voltha.KpiEventType_Types" json:"type,omitempty"`
	// Fields used when for slice:
	Ts                   float64              `protobuf:"fixed64,2,opt,name=ts,proto3" json:"ts,omitempty"`
	SliceData            []*MetricInformation `protobuf:"bytes,3,rep,name=slice_data,json=sliceData,proto3" json:"slice_data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *KpiEvent2) Reset()         { *m = KpiEvent2{} }
func (m *KpiEvent2) String() string { return proto.CompactTextString(m) }
func (*KpiEvent2) ProtoMessage()    {}
func (*KpiEvent2) Descriptor() ([]byte, []int) {
	return fileDescriptor_e63e6c07044fd2c4, []int{7}
}

func (m *KpiEvent2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KpiEvent2.Unmarshal(m, b)
}
func (m *KpiEvent2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KpiEvent2.Marshal(b, m, deterministic)
}
func (m *KpiEvent2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KpiEvent2.Merge(m, src)
}
func (m *KpiEvent2) XXX_Size() int {
	return xxx_messageInfo_KpiEvent2.Size(m)
}
func (m *KpiEvent2) XXX_DiscardUnknown() {
	xxx_messageInfo_KpiEvent2.DiscardUnknown(m)
}

var xxx_messageInfo_KpiEvent2 proto.InternalMessageInfo

func (m *KpiEvent2) GetType() KpiEventType_Types {
	if m != nil {
		return m.Type
	}
	return KpiEventType_slice
}

func (m *KpiEvent2) GetTs() float64 {
	if m != nil {
		return m.Ts
	}
	return 0
}

func (m *KpiEvent2) GetSliceData() []*MetricInformation {
	if m != nil {
		return m.SliceData
	}
	return nil
}

//
// Describes the events specific to device
type DeviceEvent struct {
	// Identifier of the originating resource of the event, for ex: device_id
	ResourceId string `protobuf:"bytes,1,opt,name=resource_id,json=resourceId,proto3" json:"resource_id,omitempty"`
	// device_event_name indicates clearly the name of the device event
	DeviceEventName string `protobuf:"bytes,2,opt,name=device_event_name,json=deviceEventName,proto3" json:"device_event_name,omitempty"`
	// Textual explanation of the device event
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// Key/Value storage for extra information that may give context to the event
	Context              map[string]string `protobuf:"bytes,4,rep,name=context,proto3" json:"context,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *DeviceEvent) Reset()         { *m = DeviceEvent{} }
func (m *DeviceEvent) String() string { return proto.CompactTextString(m) }
func (*DeviceEvent) ProtoMessage()    {}
func (*DeviceEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_e63e6c07044fd2c4, []int{8}
}

func (m *DeviceEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeviceEvent.Unmarshal(m, b)
}
func (m *DeviceEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeviceEvent.Marshal(b, m, deterministic)
}
func (m *DeviceEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeviceEvent.Merge(m, src)
}
func (m *DeviceEvent) XXX_Size() int {
	return xxx_messageInfo_DeviceEvent.Size(m)
}
func (m *DeviceEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_DeviceEvent.DiscardUnknown(m)
}

var xxx_messageInfo_DeviceEvent proto.InternalMessageInfo

func (m *DeviceEvent) GetResourceId() string {
	if m != nil {
		return m.ResourceId
	}
	return ""
}

func (m *DeviceEvent) GetDeviceEventName() string {
	if m != nil {
		return m.DeviceEventName
	}
	return ""
}

func (m *DeviceEvent) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *DeviceEvent) GetContext() map[string]string {
	if m != nil {
		return m.Context
	}
	return nil
}

//
// Identify the area of the system impacted by the event.
type EventCategory struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EventCategory) Reset()         { *m = EventCategory{} }
func (m *EventCategory) String() string { return proto.CompactTextString(m) }
func (*EventCategory) ProtoMessage()    {}
func (*EventCategory) Descriptor() ([]byte, []int) {
	return fileDescriptor_e63e6c07044fd2c4, []int{9}
}

func (m *EventCategory) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventCategory.Unmarshal(m, b)
}
func (m *EventCategory) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventCategory.Marshal(b, m, deterministic)
}
func (m *EventCategory) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventCategory.Merge(m, src)
}
func (m *EventCategory) XXX_Size() int {
	return xxx_messageInfo_EventCategory.Size(m)
}
func (m *EventCategory) XXX_DiscardUnknown() {
	xxx_messageInfo_EventCategory.DiscardUnknown(m)
}

var xxx_messageInfo_EventCategory proto.InternalMessageInfo

//
// Identify the functional category originating the event
type EventSubCategory struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EventSubCategory) Reset()         { *m = EventSubCategory{} }
func (m *EventSubCategory) String() string { return proto.CompactTextString(m) }
func (*EventSubCategory) ProtoMessage()    {}
func (*EventSubCategory) Descriptor() ([]byte, []int) {
	return fileDescriptor_e63e6c07044fd2c4, []int{10}
}

func (m *EventSubCategory) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventSubCategory.Unmarshal(m, b)
}
func (m *EventSubCategory) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventSubCategory.Marshal(b, m, deterministic)
}
func (m *EventSubCategory) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventSubCategory.Merge(m, src)
}
func (m *EventSubCategory) XXX_Size() int {
	return xxx_messageInfo_EventSubCategory.Size(m)
}
func (m *EventSubCategory) XXX_DiscardUnknown() {
	xxx_messageInfo_EventSubCategory.DiscardUnknown(m)
}

var xxx_messageInfo_EventSubCategory proto.InternalMessageInfo

//
// Identify the type of event
type EventType struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EventType) Reset()         { *m = EventType{} }
func (m *EventType) String() string { return proto.CompactTextString(m) }
func (*EventType) ProtoMessage()    {}
func (*EventType) Descriptor() ([]byte, []int) {
	return fileDescriptor_e63e6c07044fd2c4, []int{11}
}

func (m *EventType) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventType.Unmarshal(m, b)
}
func (m *EventType) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventType.Marshal(b, m, deterministic)
}
func (m *EventType) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventType.Merge(m, src)
}
func (m *EventType) XXX_Size() int {
	return xxx_messageInfo_EventType.Size(m)
}
func (m *EventType) XXX_DiscardUnknown() {
	xxx_messageInfo_EventType.DiscardUnknown(m)
}

var xxx_messageInfo_EventType proto.InternalMessageInfo

//
// Identify the functional category originating the event
type EventHeader struct {
	// Unique ID for this event.  e.g. voltha.some_olt.1234
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Refers to the functional area affect by the event
	Category EventCategory_Types `protobuf:"varint,2,opt,name=category,proto3,enum=voltha.EventCategory_Types" json:"category,omitempty"`
	// Refers to functional category of the event
	SubCategory EventSubCategory_Types `protobuf:"varint,3,opt,name=sub_category,json=subCategory,proto3,enum=voltha.EventSubCategory_Types" json:"sub_category,omitempty"`
	// Refers to the type of the event
	Type EventType_Types `protobuf:"varint,4,opt,name=type,proto3,enum=voltha.EventType_Types" json:"type,omitempty"`
	// The version identifier for this event type, thus allowing each
	// event type to evolve independently. The version should be in the
	// format “MAJOR.MINOR” format and minor changes must only be additive
	// and non-breaking.
	TypeVersion string `protobuf:"bytes,5,opt,name=type_version,json=typeVersion,proto3" json:"type_version,omitempty"`
	// Timestamp at which the event was first raised.
	// This represents the UTC time stamp since epoch (in seconds) when the
	// the event was first raised from the source entity.
	// If the source entity doesn't send the raised_ts, this shall be set
	// to timestamp when the event was received.
	RaisedTs *timestamp.Timestamp `protobuf:"bytes,6,opt,name=raised_ts,json=raisedTs,proto3" json:"raised_ts,omitempty"`
	// Timestamp at which the event was reported.
	// This represents the UTC time stamp since epoch (in seconds) when the
	// the event was reported (this time stamp is >= raised_ts).
	// If the source entity that reported this event doesn't send the
	// reported_ts, this shall be set to the same value as raised_ts.
	ReportedTs           *timestamp.Timestamp `protobuf:"bytes,7,opt,name=reported_ts,json=reportedTs,proto3" json:"reported_ts,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *EventHeader) Reset()         { *m = EventHeader{} }
func (m *EventHeader) String() string { return proto.CompactTextString(m) }
func (*EventHeader) ProtoMessage()    {}
func (*EventHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_e63e6c07044fd2c4, []int{12}
}

func (m *EventHeader) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventHeader.Unmarshal(m, b)
}
func (m *EventHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventHeader.Marshal(b, m, deterministic)
}
func (m *EventHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventHeader.Merge(m, src)
}
func (m *EventHeader) XXX_Size() int {
	return xxx_messageInfo_EventHeader.Size(m)
}
func (m *EventHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_EventHeader.DiscardUnknown(m)
}

var xxx_messageInfo_EventHeader proto.InternalMessageInfo

func (m *EventHeader) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *EventHeader) GetCategory() EventCategory_Types {
	if m != nil {
		return m.Category
	}
	return EventCategory_COMMUNICATION
}

func (m *EventHeader) GetSubCategory() EventSubCategory_Types {
	if m != nil {
		return m.SubCategory
	}
	return EventSubCategory_PON
}

func (m *EventHeader) GetType() EventType_Types {
	if m != nil {
		return m.Type
	}
	return EventType_CONFIG_EVENT
}

func (m *EventHeader) GetTypeVersion() string {
	if m != nil {
		return m.TypeVersion
	}
	return ""
}

func (m *EventHeader) GetRaisedTs() *timestamp.Timestamp {
	if m != nil {
		return m.RaisedTs
	}
	return nil
}

func (m *EventHeader) GetReportedTs() *timestamp.Timestamp {
	if m != nil {
		return m.ReportedTs
	}
	return nil
}

//
// Event Structure
type Event struct {
	// event header
	Header *EventHeader `protobuf:"bytes,1,opt,name=header,proto3" json:"header,omitempty"`
	// oneof event types referred by EventType.
	//
	// Types that are valid to be assigned to EventType:
	//	*Event_ConfigEvent
	//	*Event_KpiEvent
	//	*Event_KpiEvent2
	//	*Event_DeviceEvent
	EventType            isEvent_EventType `protobuf_oneof:"event_type"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_e63e6c07044fd2c4, []int{13}
}

func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetHeader() *EventHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type isEvent_EventType interface {
	isEvent_EventType()
}

type Event_ConfigEvent struct {
	ConfigEvent *ConfigEvent `protobuf:"bytes,2,opt,name=config_event,json=configEvent,proto3,oneof"`
}

type Event_KpiEvent struct {
	KpiEvent *KpiEvent `protobuf:"bytes,3,opt,name=kpi_event,json=kpiEvent,proto3,oneof"`
}

type Event_KpiEvent2 struct {
	KpiEvent2 *KpiEvent2 `protobuf:"bytes,4,opt,name=kpi_event2,json=kpiEvent2,proto3,oneof"`
}

type Event_DeviceEvent struct {
	DeviceEvent *DeviceEvent `protobuf:"bytes,5,opt,name=device_event,json=deviceEvent,proto3,oneof"`
}

func (*Event_ConfigEvent) isEvent_EventType() {}

func (*Event_KpiEvent) isEvent_EventType() {}

func (*Event_KpiEvent2) isEvent_EventType() {}

func (*Event_DeviceEvent) isEvent_EventType() {}

func (m *Event) GetEventType() isEvent_EventType {
	if m != nil {
		return m.EventType
	}
	return nil
}

func (m *Event) GetConfigEvent() *ConfigEvent {
	if x, ok := m.GetEventType().(*Event_ConfigEvent); ok {
		return x.ConfigEvent
	}
	return nil
}

func (m *Event) GetKpiEvent() *KpiEvent {
	if x, ok := m.GetEventType().(*Event_KpiEvent); ok {
		return x.KpiEvent
	}
	return nil
}

func (m *Event) GetKpiEvent2() *KpiEvent2 {
	if x, ok := m.GetEventType().(*Event_KpiEvent2); ok {
		return x.KpiEvent2
	}
	return nil
}

func (m *Event) GetDeviceEvent() *DeviceEvent {
	if x, ok := m.GetEventType().(*Event_DeviceEvent); ok {
		return x.DeviceEvent
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Event) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Event_ConfigEvent)(nil),
		(*Event_KpiEvent)(nil),
		(*Event_KpiEvent2)(nil),
		(*Event_DeviceEvent)(nil),
	}
}

func init() {
	proto.RegisterEnum("voltha.ConfigEventType_Types", ConfigEventType_Types_name, ConfigEventType_Types_value)
	proto.RegisterEnum("voltha.KpiEventType_Types", KpiEventType_Types_name, KpiEventType_Types_value)
	proto.RegisterEnum("voltha.EventCategory_Types", EventCategory_Types_name, EventCategory_Types_value)
	proto.RegisterEnum("voltha.EventSubCategory_Types", EventSubCategory_Types_name, EventSubCategory_Types_value)
	proto.RegisterEnum("voltha.EventType_Types", EventType_Types_name, EventType_Types_value)
	proto.RegisterType((*ConfigEventType)(nil), "voltha.ConfigEventType")
	proto.RegisterType((*ConfigEvent)(nil), "voltha.ConfigEvent")
	proto.RegisterType((*KpiEventType)(nil), "voltha.KpiEventType")
	proto.RegisterType((*MetricMetaData)(nil), "voltha.MetricMetaData")
	proto.RegisterMapType((map[string]string)(nil), "voltha.MetricMetaData.ContextEntry")
	proto.RegisterType((*MetricValuePairs)(nil), "voltha.MetricValuePairs")
	proto.RegisterMapType((map[string]float32)(nil), "voltha.MetricValuePairs.MetricsEntry")
	proto.RegisterType((*MetricInformation)(nil), "voltha.MetricInformation")
	proto.RegisterMapType((map[string]float32)(nil), "voltha.MetricInformation.MetricsEntry")
	proto.RegisterType((*KpiEvent)(nil), "voltha.KpiEvent")
	proto.RegisterMapType((map[string]*MetricValuePairs)(nil), "voltha.KpiEvent.PrefixesEntry")
	proto.RegisterType((*KpiEvent2)(nil), "voltha.KpiEvent2")
	proto.RegisterType((*DeviceEvent)(nil), "voltha.DeviceEvent")
	proto.RegisterMapType((map[string]string)(nil), "voltha.DeviceEvent.ContextEntry")
	proto.RegisterType((*EventCategory)(nil), "voltha.EventCategory")
	proto.RegisterType((*EventSubCategory)(nil), "voltha.EventSubCategory")
	proto.RegisterType((*EventType)(nil), "voltha.EventType")
	proto.RegisterType((*EventHeader)(nil), "voltha.EventHeader")
	proto.RegisterType((*Event)(nil), "voltha.Event")
}

func init() { proto.RegisterFile("voltha_protos/events.proto", fileDescriptor_e63e6c07044fd2c4) }

var fileDescriptor_e63e6c07044fd2c4 = []byte{
	// 1135 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x56, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0x8e, 0x9d, 0x34, 0x3f, 0xc7, 0x69, 0xeb, 0xce, 0x22, 0x08, 0x59, 0xd8, 0x2d, 0x46, 0xa0,
	0x6a, 0x57, 0x38, 0xc2, 0xac, 0xb4, 0x55, 0x57, 0x08, 0x76, 0x53, 0xb3, 0x35, 0x4b, 0x93, 0xe0,
	0xa6, 0x45, 0x70, 0x13, 0x4d, 0xe3, 0x69, 0x62, 0x35, 0xc9, 0x58, 0x9e, 0x49, 0xb4, 0x7d, 0x00,
	0xae, 0x79, 0x00, 0x1e, 0x81, 0xe7, 0xe0, 0x31, 0x10, 0x2f, 0xc1, 0x03, 0xa0, 0xf9, 0x71, 0x62,
	0x97, 0xae, 0xf6, 0x62, 0xc5, 0x95, 0x67, 0xce, 0x9c, 0x6f, 0xe6, 0x3b, 0x67, 0xce, 0x77, 0xc6,
	0xd0, 0x5e, 0xd1, 0x19, 0x9f, 0xe2, 0x51, 0x92, 0x52, 0x4e, 0x59, 0x87, 0xac, 0xc8, 0x82, 0x33,
	0x57, 0xce, 0x50, 0x55, 0xad, 0xb5, 0x5b, 0x45, 0x9f, 0x39, 0xe1, 0x58, 0x79, 0xb4, 0x3f, 0x9a,
	0x50, 0x3a, 0x99, 0x91, 0x0e, 0x4e, 0xe2, 0x0e, 0x5e, 0x2c, 0x28, 0xc7, 0x3c, 0xa6, 0x0b, 0x8d,
	0x6f, 0x3f, 0xd4, 0xab, 0x72, 0x76, 0xb9, 0xbc, 0xea, 0xf0, 0x78, 0x4e, 0x18, 0xc7, 0xf3, 0x44,
	0x39, 0x38, 0xcf, 0x60, 0xb7, 0x4b, 0x17, 0x57, 0xf1, 0xc4, 0x17, 0xc7, 0x0e, 0x6f, 0x12, 0xe2,
	0x1c, 0xc0, 0x96, 0xf8, 0x32, 0x54, 0x83, 0x32, 0x8e, 0x22, 0xbb, 0x84, 0x00, 0xaa, 0x29, 0x99,
	0xd3, 0x15, 0xb1, 0x0d, 0x31, 0x5e, 0x26, 0x11, 0xe6, 0xc4, 0x36, 0x9d, 0x29, 0x58, 0x39, 0x30,
	0xfa, 0x12, 0x2a, 0xfc, 0x26, 0x21, 0x2d, 0x63, 0xdf, 0x38, 0xd8, 0xf1, 0x3e, 0x76, 0x15, 0x67,
	0xf7, 0xd6, 0xfe, 0xae, 0xdc, 0x3c, 0x94, 0xae, 0x08, 0x41, 0x65, 0x8a, 0xd9, 0xb4, 0x65, 0xee,
	0x1b, 0x07, 0x8d, 0x50, 0x8e, 0x85, 0x2d, 0xc2, 0x1c, 0xb7, 0xca, 0xca, 0x26, 0xc6, 0xce, 0x23,
	0x68, 0xbe, 0x4a, 0xe2, 0x0d, 0xc7, 0x76, 0xc6, 0xb1, 0x01, 0x5b, 0x6c, 0x16, 0x8f, 0x89, 0x5d,
	0x42, 0x55, 0x30, 0x39, 0xb3, 0x0d, 0xe7, 0x0f, 0x13, 0x76, 0x4e, 0x09, 0x4f, 0xe3, 0xf1, 0x29,
	0xe1, 0xf8, 0x18, 0x73, 0x8c, 0xde, 0x83, 0x2d, 0x1e, 0xf3, 0x99, 0xa2, 0xd6, 0x08, 0xd5, 0x04,
	0xed, 0x08, 0x80, 0x3c, 0xda, 0x08, 0x4d, 0xce, 0xd0, 0x23, 0xd8, 0x9b, 0xd1, 0x49, 0x3c, 0xc6,
	0xb3, 0x51, 0x44, 0x56, 0xf1, 0x98, 0x8c, 0xe2, 0x48, 0xb3, 0xd8, 0xd5, 0x0b, 0xc7, 0xd2, 0x1e,
	0x44, 0xe8, 0x3e, 0x34, 0x18, 0x49, 0x63, 0x3c, 0x1b, 0x2d, 0x68, 0xab, 0x22, 0x7d, 0xea, 0xca,
	0xd0, 0xa3, 0x62, 0x71, 0xb3, 0xc1, 0x96, 0x5a, 0x8c, 0x32, 0xe4, 0xd7, 0x50, 0x1b, 0xd3, 0x05,
	0x27, 0xaf, 0x79, 0xab, 0xba, 0x5f, 0x3e, 0xb0, 0xbc, 0x4f, 0xb3, 0x44, 0x15, 0x49, 0x8b, 0xbc,
	0x09, 0x2f, 0x7f, 0xc1, 0xd3, 0x9b, 0x30, 0xc3, 0x88, 0xec, 0x2c, 0x97, 0x71, 0xd4, 0xaa, 0xa9,
	0xec, 0x88, 0x71, 0xfb, 0x08, 0x9a, 0x79, 0x67, 0x64, 0x43, 0xf9, 0x9a, 0xdc, 0xe8, 0x60, 0xc5,
	0x50, 0x24, 0x60, 0x85, 0x67, 0x4b, 0xa2, 0x13, 0xad, 0x26, 0x47, 0xe6, 0xa1, 0xe1, 0xfc, 0x66,
	0x80, 0xad, 0x0e, 0xbe, 0x10, 0xb6, 0x01, 0x8e, 0x53, 0x86, 0xbe, 0x81, 0xda, 0x5c, 0xda, 0x58,
	0xcb, 0x90, 0x1c, 0x3f, 0x2b, 0x72, 0xdc, 0xb8, 0x6a, 0x03, 0xd3, 0x2c, 0x35, 0x4a, 0x30, 0xca,
	0x2f, 0xbc, 0x8d, 0x91, 0x99, 0x67, 0xf4, 0xa7, 0x01, 0x7b, 0x0a, 0x1c, 0x2c, 0xae, 0x68, 0x3a,
	0x97, 0x05, 0x8d, 0x3c, 0xa8, 0x8b, 0xaa, 0x97, 0x95, 0x21, 0xb6, 0xb1, 0xbc, 0xf7, 0xef, 0xce,
	0x5b, 0xb8, 0xf6, 0x43, 0xdf, 0x6e, 0xc2, 0x30, 0x65, 0x18, 0x9f, 0x17, 0x21, 0xb9, 0xfd, 0xff,
	0x87, 0x38, 0xfe, 0x32, 0xa0, 0x9e, 0x15, 0x2d, 0x72, 0x0b, 0xda, 0x68, 0x67, 0x3c, 0xf2, 0x45,
	0x5d, 0x10, 0xc6, 0xa6, 0x36, 0x4d, 0x59, 0x9b, 0x47, 0x50, 0x4f, 0x52, 0x72, 0x15, 0xbf, 0x26,
	0xac, 0x55, 0x96, 0xb1, 0x3c, 0xb8, 0xbd, 0x87, 0x3b, 0xd0, 0x0e, 0x2a, 0x86, 0xb5, 0x7f, 0xfb,
	0x1c, 0xb6, 0x0b, 0x4b, 0x77, 0x44, 0xe1, 0xe6, 0xa3, 0xb0, 0xbc, 0xd6, 0x9b, 0xae, 0x3b, 0x1f,
	0xdf, 0xaf, 0x06, 0x34, 0xb2, 0xb3, 0xbd, 0x77, 0x08, 0x50, 0x89, 0xef, 0x10, 0x40, 0x0a, 0x79,
	0xa4, 0xb5, 0x2f, 0x42, 0xfc, 0xf0, 0x8d, 0xd7, 0x15, 0x36, 0xa4, 0xb3, 0xb8, 0x6f, 0xe7, 0x1f,
	0x03, 0x2c, 0xa5, 0x4b, 0x95, 0xea, 0x87, 0x60, 0xa5, 0x84, 0xd1, 0x65, 0xaa, 0xf4, 0xa7, 0xa2,
	0x84, 0xcc, 0x14, 0x44, 0x42, 0xe7, 0x5a, 0x9e, 0xb2, 0xd7, 0x8e, 0x16, 0x78, 0x9e, 0x09, 0x63,
	0x37, 0xda, 0x6c, 0xd4, 0xc3, 0x73, 0x82, 0xf6, 0xc1, 0x8a, 0x08, 0x1b, 0xa7, 0x71, 0x22, 0x8e,
	0xd5, 0xdd, 0x20, 0x6f, 0x42, 0x47, 0x1b, 0x3d, 0x57, 0x24, 0xeb, 0xfd, 0x8c, 0x75, 0x8e, 0xd4,
	0xdd, 0x62, 0x7e, 0x27, 0xe1, 0xae, 0x60, 0x5b, 0x6e, 0xdd, 0xc5, 0x9c, 0x4c, 0x68, 0x7a, 0xe3,
	0x90, 0xac, 0x27, 0xee, 0xc1, 0x76, 0xb7, 0x7f, 0x7a, 0x7a, 0xde, 0x0b, 0xba, 0xcf, 0x87, 0x41,
	0xbf, 0x67, 0x97, 0xd0, 0x2e, 0x58, 0x7e, 0xef, 0x22, 0x08, 0xfb, 0xbd, 0x53, 0xbf, 0x37, 0xb4,
	0x0d, 0xb4, 0x0d, 0x0d, 0xff, 0xc7, 0xf3, 0x60, 0x20, 0xa7, 0x26, 0xb2, 0xa0, 0x76, 0xe6, 0x87,
	0x17, 0x41, 0xd7, 0xb7, 0xcb, 0x68, 0x07, 0x60, 0x10, 0xf6, 0xbb, 0xfe, 0xd9, 0x59, 0xd0, 0x7b,
	0x69, 0x57, 0x50, 0x13, 0xea, 0x67, 0x7e, 0xf7, 0x3c, 0x0c, 0x86, 0x3f, 0xdb, 0x5b, 0xce, 0x09,
	0xd8, 0xf2, 0xdc, 0xb3, 0xe5, 0xe5, 0xfa, 0xe8, 0x27, 0xb9, 0x27, 0x63, 0x20, 0x0f, 0xac, 0x41,
	0xb9, 0xff, 0x83, 0x38, 0x48, 0x0c, 0xe4, 0x11, 0x72, 0x70, 0x6e, 0x97, 0xc5, 0xa0, 0xd7, 0x0b,
	0xec, 0x8a, 0xf3, 0x13, 0x34, 0x36, 0x1d, 0xfd, 0xfb, 0x6c, 0x0b, 0x1b, 0x9a, 0xdd, 0x7e, 0xef,
	0xbb, 0xe0, 0xe5, 0xc8, 0xbf, 0x10, 0xe4, 0x4a, 0x82, 0xeb, 0xab, 0x41, 0xa0, 0xa7, 0x86, 0xa0,
	0xb7, 0x9e, 0x7a, 0xb6, 0x29, 0x00, 0xc7, 0xbe, 0xa0, 0xae, 0x3d, 0xca, 0xce, 0xdf, 0x26, 0x58,
	0x72, 0xe7, 0x13, 0x82, 0x23, 0x92, 0x8a, 0x5a, 0x5b, 0x17, 0x82, 0x19, 0x47, 0xe8, 0x29, 0xd4,
	0xc7, 0x9a, 0xba, 0xcc, 0xeb, 0x8e, 0x77, 0x3f, 0xbb, 0xb3, 0x42, 0x4a, 0x75, 0xc1, 0xae, 0x9d,
	0xd1, 0x73, 0x68, 0xb2, 0xe5, 0xe5, 0x68, 0x0d, 0x2e, 0x4b, 0xf0, 0x83, 0x02, 0x38, 0x97, 0x17,
	0x8d, 0xb7, 0xd8, 0xc6, 0x84, 0x1e, 0x6b, 0x9d, 0x54, 0x24, 0xf4, 0x83, 0x02, 0xf4, 0x3f, 0x22,
	0xf9, 0x04, 0x9a, 0xe2, 0x3b, 0x5a, 0x91, 0x94, 0x89, 0xf2, 0x53, 0x6f, 0x89, 0x25, 0x6c, 0x17,
	0xca, 0x84, 0x9e, 0x42, 0x23, 0xc5, 0x31, 0x23, 0xd1, 0x88, 0xb3, 0x56, 0x55, 0xaa, 0xb7, 0xed,
	0xaa, 0x57, 0xdf, 0xcd, 0x5e, 0x7d, 0x77, 0x98, 0xbd, 0xfa, 0x61, 0x5d, 0x39, 0x0f, 0x19, 0x7a,
	0x26, 0x64, 0x92, 0xd0, 0x94, 0x2b, 0x68, 0xed, 0xad, 0x50, 0xc8, 0xdc, 0x87, 0xcc, 0xf9, 0xdd,
	0x84, 0x2d, 0xa5, 0xb6, 0xc7, 0x50, 0x9d, 0xca, 0x2c, 0xeb, 0xae, 0x7c, 0xaf, 0x10, 0x91, 0xba,
	0x80, 0x50, 0xbb, 0xa0, 0x43, 0x68, 0x8e, 0xe5, 0xdf, 0x80, 0x52, 0x9e, 0xee, 0x36, 0xf7, 0xee,
	0xf8, 0x53, 0x38, 0x29, 0x85, 0xd6, 0x38, 0xf7, 0x6f, 0xd1, 0x81, 0xc6, 0x75, 0x12, 0x6b, 0x58,
	0x59, 0xc2, 0xec, 0xdb, 0x3d, 0xe6, 0xa4, 0x14, 0xd6, 0xaf, 0xb3, 0x86, 0xeb, 0x01, 0xac, 0x01,
	0x9e, 0xcc, 0xb6, 0xe5, 0xed, 0xdd, 0x46, 0x78, 0x27, 0xa5, 0xb0, 0x71, 0xbd, 0xee, 0x61, 0x87,
	0xd0, 0xcc, 0x37, 0x06, 0x99, 0xee, 0x1c, 0xbd, 0x9c, 0x9e, 0x05, 0xbd, 0x5c, 0xab, 0x78, 0xd1,
	0x04, 0x50, 0xbd, 0x44, 0x5c, 0xcd, 0x0b, 0x1f, 0xee, 0xd1, 0x74, 0xe2, 0xd2, 0x84, 0x2c, 0xc6,
	0x34, 0x8d, 0x34, 0xfe, 0x17, 0x77, 0x12, 0xf3, 0xe9, 0xf2, 0xd2, 0x1d, 0xd3, 0x79, 0x27, 0x5b,
	0xeb, 0xa8, 0xb5, 0x2f, 0xf4, 0x8f, 0xdd, 0xea, 0x49, 0x67, 0x42, 0xb5, 0xed, 0xb2, 0x2a, 0x8d,
	0x5f, 0xfd, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x26, 0x26, 0x74, 0xd0, 0x21, 0x0a, 0x00, 0x00,
}
