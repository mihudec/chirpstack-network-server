// Code generated by protoc-gen-go. DO NOT EDIT.
// source: as.proto

package as

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common "github.com/brocaar/loraserver/api/common"
import gw "github.com/brocaar/loraserver/api/gw"
import empty "github.com/golang/protobuf/ptypes/empty"

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

type RXWindow int32

const (
	RXWindow_RX1 RXWindow = 0
	RXWindow_RX2 RXWindow = 1
)

var RXWindow_name = map[int32]string{
	0: "RX1",
	1: "RX2",
}
var RXWindow_value = map[string]int32{
	"RX1": 0,
	"RX2": 1,
}

func (x RXWindow) String() string {
	return proto.EnumName(RXWindow_name, int32(x))
}
func (RXWindow) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_as_8dba2f46bdb31a35, []int{0}
}

type ErrorType int32

const (
	ErrorType_GENERIC                ErrorType = 0
	ErrorType_OTAA                   ErrorType = 1
	ErrorType_DATA_UP_FCNT           ErrorType = 2
	ErrorType_DATA_UP_MIC            ErrorType = 3
	ErrorType_DEVICE_QUEUE_ITEM_SIZE ErrorType = 4
	ErrorType_DEVICE_QUEUE_ITEM_FCNT ErrorType = 5
)

var ErrorType_name = map[int32]string{
	0: "GENERIC",
	1: "OTAA",
	2: "DATA_UP_FCNT",
	3: "DATA_UP_MIC",
	4: "DEVICE_QUEUE_ITEM_SIZE",
	5: "DEVICE_QUEUE_ITEM_FCNT",
}
var ErrorType_value = map[string]int32{
	"GENERIC":                0,
	"OTAA":                   1,
	"DATA_UP_FCNT":           2,
	"DATA_UP_MIC":            3,
	"DEVICE_QUEUE_ITEM_SIZE": 4,
	"DEVICE_QUEUE_ITEM_FCNT": 5,
}

func (x ErrorType) String() string {
	return proto.EnumName(ErrorType_name, int32(x))
}
func (ErrorType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_as_8dba2f46bdb31a35, []int{1}
}

type DeviceActivationContext struct {
	// Assigned Device Address.
	DevAddr []byte `protobuf:"bytes,1,opt,name=dev_addr,json=devAddr,proto3" json:"dev_addr,omitempty"`
	// Application session key (envelope).
	AppSKey              *common.KeyEnvelope `protobuf:"bytes,2,opt,name=app_s_key,json=appSKey,proto3" json:"app_s_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *DeviceActivationContext) Reset()         { *m = DeviceActivationContext{} }
func (m *DeviceActivationContext) String() string { return proto.CompactTextString(m) }
func (*DeviceActivationContext) ProtoMessage()    {}
func (*DeviceActivationContext) Descriptor() ([]byte, []int) {
	return fileDescriptor_as_8dba2f46bdb31a35, []int{0}
}
func (m *DeviceActivationContext) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeviceActivationContext.Unmarshal(m, b)
}
func (m *DeviceActivationContext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeviceActivationContext.Marshal(b, m, deterministic)
}
func (dst *DeviceActivationContext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeviceActivationContext.Merge(dst, src)
}
func (m *DeviceActivationContext) XXX_Size() int {
	return xxx_messageInfo_DeviceActivationContext.Size(m)
}
func (m *DeviceActivationContext) XXX_DiscardUnknown() {
	xxx_messageInfo_DeviceActivationContext.DiscardUnknown(m)
}

var xxx_messageInfo_DeviceActivationContext proto.InternalMessageInfo

func (m *DeviceActivationContext) GetDevAddr() []byte {
	if m != nil {
		return m.DevAddr
	}
	return nil
}

func (m *DeviceActivationContext) GetAppSKey() *common.KeyEnvelope {
	if m != nil {
		return m.AppSKey
	}
	return nil
}

type HandleUplinkDataRequest struct {
	// DevEUI EUI (8 bytes).
	DevEui []byte `protobuf:"bytes,1,opt,name=dev_eui,json=devEui,proto3" json:"dev_eui,omitempty"`
	// Join EUI used for OTAA activation (8 bytes).
	JoinEui []byte `protobuf:"bytes,2,opt,name=join_eui,json=joinEui,proto3" json:"join_eui,omitempty"`
	// Frame-counter.
	FCnt uint32 `protobuf:"varint,3,opt,name=f_cnt,json=fCnt,proto3" json:"f_cnt,omitempty"`
	// Frame port.
	FPort uint32 `protobuf:"varint,4,opt,name=f_port,json=fPort,proto3" json:"f_port,omitempty"`
	// ADR enabled.
	Adr bool `protobuf:"varint,5,opt,name=adr,proto3" json:"adr,omitempty"`
	// Data-rate.
	Dr uint32 `protobuf:"varint,6,opt,name=dr,proto3" json:"dr,omitempty"`
	// TX meta-data.
	TxInfo *gw.UplinkTXInfo `protobuf:"bytes,7,opt,name=tx_info,json=txInfo,proto3" json:"tx_info,omitempty"`
	// RX meta-data.
	RxInfo []*gw.UplinkRXInfo `protobuf:"bytes,8,rep,name=rx_info,json=rxInfo,proto3" json:"rx_info,omitempty"`
	// Received data (encrypted).
	Data []byte `protobuf:"bytes,9,opt,name=data,proto3" json:"data,omitempty"`
	// Device activation context.
	//
	// This field is only set on the first uplink frame when the security
	// context has changed (e.g. a new OTAA (re)activation).
	DeviceActivationContext *DeviceActivationContext `protobuf:"bytes,10,opt,name=device_activation_context,json=deviceActivationContext,proto3" json:"device_activation_context,omitempty"`
	XXX_NoUnkeyedLiteral    struct{}                 `json:"-"`
	XXX_unrecognized        []byte                   `json:"-"`
	XXX_sizecache           int32                    `json:"-"`
}

func (m *HandleUplinkDataRequest) Reset()         { *m = HandleUplinkDataRequest{} }
func (m *HandleUplinkDataRequest) String() string { return proto.CompactTextString(m) }
func (*HandleUplinkDataRequest) ProtoMessage()    {}
func (*HandleUplinkDataRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_as_8dba2f46bdb31a35, []int{1}
}
func (m *HandleUplinkDataRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HandleUplinkDataRequest.Unmarshal(m, b)
}
func (m *HandleUplinkDataRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HandleUplinkDataRequest.Marshal(b, m, deterministic)
}
func (dst *HandleUplinkDataRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HandleUplinkDataRequest.Merge(dst, src)
}
func (m *HandleUplinkDataRequest) XXX_Size() int {
	return xxx_messageInfo_HandleUplinkDataRequest.Size(m)
}
func (m *HandleUplinkDataRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HandleUplinkDataRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HandleUplinkDataRequest proto.InternalMessageInfo

func (m *HandleUplinkDataRequest) GetDevEui() []byte {
	if m != nil {
		return m.DevEui
	}
	return nil
}

func (m *HandleUplinkDataRequest) GetJoinEui() []byte {
	if m != nil {
		return m.JoinEui
	}
	return nil
}

func (m *HandleUplinkDataRequest) GetFCnt() uint32 {
	if m != nil {
		return m.FCnt
	}
	return 0
}

func (m *HandleUplinkDataRequest) GetFPort() uint32 {
	if m != nil {
		return m.FPort
	}
	return 0
}

func (m *HandleUplinkDataRequest) GetAdr() bool {
	if m != nil {
		return m.Adr
	}
	return false
}

func (m *HandleUplinkDataRequest) GetDr() uint32 {
	if m != nil {
		return m.Dr
	}
	return 0
}

func (m *HandleUplinkDataRequest) GetTxInfo() *gw.UplinkTXInfo {
	if m != nil {
		return m.TxInfo
	}
	return nil
}

func (m *HandleUplinkDataRequest) GetRxInfo() []*gw.UplinkRXInfo {
	if m != nil {
		return m.RxInfo
	}
	return nil
}

func (m *HandleUplinkDataRequest) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *HandleUplinkDataRequest) GetDeviceActivationContext() *DeviceActivationContext {
	if m != nil {
		return m.DeviceActivationContext
	}
	return nil
}

type HandleProprietaryUplinkRequest struct {
	// MACPayload of the proprietary LoRaWAN frame.
	MacPayload []byte `protobuf:"bytes,1,opt,name=mac_payload,json=macPayload,proto3" json:"mac_payload,omitempty"`
	// MIC of the proprietary LoRaWAN frame.
	Mic []byte `protobuf:"bytes,2,opt,name=mic,proto3" json:"mic,omitempty"`
	// TXInfo contains the TX related meta-data.
	TxInfo *gw.UplinkTXInfo `protobuf:"bytes,3,opt,name=tx_info,json=txInfo,proto3" json:"tx_info,omitempty"`
	// RXInfo contains the RX related meta-data.
	RxInfo               []*gw.UplinkRXInfo `protobuf:"bytes,4,rep,name=rx_info,json=rxInfo,proto3" json:"rx_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *HandleProprietaryUplinkRequest) Reset()         { *m = HandleProprietaryUplinkRequest{} }
func (m *HandleProprietaryUplinkRequest) String() string { return proto.CompactTextString(m) }
func (*HandleProprietaryUplinkRequest) ProtoMessage()    {}
func (*HandleProprietaryUplinkRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_as_8dba2f46bdb31a35, []int{2}
}
func (m *HandleProprietaryUplinkRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HandleProprietaryUplinkRequest.Unmarshal(m, b)
}
func (m *HandleProprietaryUplinkRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HandleProprietaryUplinkRequest.Marshal(b, m, deterministic)
}
func (dst *HandleProprietaryUplinkRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HandleProprietaryUplinkRequest.Merge(dst, src)
}
func (m *HandleProprietaryUplinkRequest) XXX_Size() int {
	return xxx_messageInfo_HandleProprietaryUplinkRequest.Size(m)
}
func (m *HandleProprietaryUplinkRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HandleProprietaryUplinkRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HandleProprietaryUplinkRequest proto.InternalMessageInfo

func (m *HandleProprietaryUplinkRequest) GetMacPayload() []byte {
	if m != nil {
		return m.MacPayload
	}
	return nil
}

func (m *HandleProprietaryUplinkRequest) GetMic() []byte {
	if m != nil {
		return m.Mic
	}
	return nil
}

func (m *HandleProprietaryUplinkRequest) GetTxInfo() *gw.UplinkTXInfo {
	if m != nil {
		return m.TxInfo
	}
	return nil
}

func (m *HandleProprietaryUplinkRequest) GetRxInfo() []*gw.UplinkRXInfo {
	if m != nil {
		return m.RxInfo
	}
	return nil
}

type HandleErrorRequest struct {
	// Device EUI (8 bytes).
	DevEui []byte `protobuf:"bytes,1,opt,name=dev_eui,json=devEui,proto3" json:"dev_eui,omitempty"`
	// Type of the error.
	Type ErrorType `protobuf:"varint,3,opt,name=type,proto3,enum=as.ErrorType" json:"type,omitempty"`
	// Error string describing the error.
	Error string `protobuf:"bytes,4,opt,name=error,proto3" json:"error,omitempty"`
	// Frame-counter (if applicable) related to the error.
	FCnt                 uint32   `protobuf:"varint,5,opt,name=f_cnt,json=fCnt,proto3" json:"f_cnt,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HandleErrorRequest) Reset()         { *m = HandleErrorRequest{} }
func (m *HandleErrorRequest) String() string { return proto.CompactTextString(m) }
func (*HandleErrorRequest) ProtoMessage()    {}
func (*HandleErrorRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_as_8dba2f46bdb31a35, []int{3}
}
func (m *HandleErrorRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HandleErrorRequest.Unmarshal(m, b)
}
func (m *HandleErrorRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HandleErrorRequest.Marshal(b, m, deterministic)
}
func (dst *HandleErrorRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HandleErrorRequest.Merge(dst, src)
}
func (m *HandleErrorRequest) XXX_Size() int {
	return xxx_messageInfo_HandleErrorRequest.Size(m)
}
func (m *HandleErrorRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HandleErrorRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HandleErrorRequest proto.InternalMessageInfo

func (m *HandleErrorRequest) GetDevEui() []byte {
	if m != nil {
		return m.DevEui
	}
	return nil
}

func (m *HandleErrorRequest) GetType() ErrorType {
	if m != nil {
		return m.Type
	}
	return ErrorType_GENERIC
}

func (m *HandleErrorRequest) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *HandleErrorRequest) GetFCnt() uint32 {
	if m != nil {
		return m.FCnt
	}
	return 0
}

type HandleDownlinkACKRequest struct {
	// Device EUI (8 bytes).
	DevEui []byte `protobuf:"bytes,1,opt,name=dev_eui,json=devEui,proto3" json:"dev_eui,omitempty"`
	// Downlink frame-counter.
	FCnt uint32 `protobuf:"varint,2,opt,name=f_cnt,json=fCnt,proto3" json:"f_cnt,omitempty"`
	// Frame was acknowledged?
	Acknowledged         bool     `protobuf:"varint,3,opt,name=acknowledged,proto3" json:"acknowledged,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HandleDownlinkACKRequest) Reset()         { *m = HandleDownlinkACKRequest{} }
func (m *HandleDownlinkACKRequest) String() string { return proto.CompactTextString(m) }
func (*HandleDownlinkACKRequest) ProtoMessage()    {}
func (*HandleDownlinkACKRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_as_8dba2f46bdb31a35, []int{4}
}
func (m *HandleDownlinkACKRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HandleDownlinkACKRequest.Unmarshal(m, b)
}
func (m *HandleDownlinkACKRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HandleDownlinkACKRequest.Marshal(b, m, deterministic)
}
func (dst *HandleDownlinkACKRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HandleDownlinkACKRequest.Merge(dst, src)
}
func (m *HandleDownlinkACKRequest) XXX_Size() int {
	return xxx_messageInfo_HandleDownlinkACKRequest.Size(m)
}
func (m *HandleDownlinkACKRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HandleDownlinkACKRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HandleDownlinkACKRequest proto.InternalMessageInfo

func (m *HandleDownlinkACKRequest) GetDevEui() []byte {
	if m != nil {
		return m.DevEui
	}
	return nil
}

func (m *HandleDownlinkACKRequest) GetFCnt() uint32 {
	if m != nil {
		return m.FCnt
	}
	return 0
}

func (m *HandleDownlinkACKRequest) GetAcknowledged() bool {
	if m != nil {
		return m.Acknowledged
	}
	return false
}

type SetDeviceStatusRequest struct {
	// Device EUI (8 bytes).
	DevEui []byte `protobuf:"bytes,1,opt,name=dev_eui,json=devEui,proto3" json:"dev_eui,omitempty"`
	// Battery level (deprecated, use battery_level).
	// 0:      The end-device is connected to an external power source
	// 1..254: The battery level, 1 being at minimum and 254 being at maximum
	// 255:    The end-device was not able to measure the battery level
	Battery uint32 `protobuf:"varint,2,opt,name=battery,proto3" json:"battery,omitempty"`
	// The device margin status
	// -32..32: The demodulation SNR ration in dB
	Margin int32 `protobuf:"varint,3,opt,name=margin,proto3" json:"margin,omitempty"`
	// Device is connected to an external power source.
	ExternalPowerSource bool `protobuf:"varint,4,opt,name=external_power_source,json=externalPowerSource,proto3" json:"external_power_source,omitempty"`
	// Device battery status is not available.
	BatteryLevelUnavailable bool `protobuf:"varint,5,opt,name=battery_level_unavailable,json=batteryLevelUnavailable,proto3" json:"battery_level_unavailable,omitempty"`
	// Battery level as a percentage.
	BatteryLevel         float32  `protobuf:"fixed32,6,opt,name=battery_level,json=batteryLevel,proto3" json:"battery_level,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetDeviceStatusRequest) Reset()         { *m = SetDeviceStatusRequest{} }
func (m *SetDeviceStatusRequest) String() string { return proto.CompactTextString(m) }
func (*SetDeviceStatusRequest) ProtoMessage()    {}
func (*SetDeviceStatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_as_8dba2f46bdb31a35, []int{5}
}
func (m *SetDeviceStatusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetDeviceStatusRequest.Unmarshal(m, b)
}
func (m *SetDeviceStatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetDeviceStatusRequest.Marshal(b, m, deterministic)
}
func (dst *SetDeviceStatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetDeviceStatusRequest.Merge(dst, src)
}
func (m *SetDeviceStatusRequest) XXX_Size() int {
	return xxx_messageInfo_SetDeviceStatusRequest.Size(m)
}
func (m *SetDeviceStatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetDeviceStatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetDeviceStatusRequest proto.InternalMessageInfo

func (m *SetDeviceStatusRequest) GetDevEui() []byte {
	if m != nil {
		return m.DevEui
	}
	return nil
}

func (m *SetDeviceStatusRequest) GetBattery() uint32 {
	if m != nil {
		return m.Battery
	}
	return 0
}

func (m *SetDeviceStatusRequest) GetMargin() int32 {
	if m != nil {
		return m.Margin
	}
	return 0
}

func (m *SetDeviceStatusRequest) GetExternalPowerSource() bool {
	if m != nil {
		return m.ExternalPowerSource
	}
	return false
}

func (m *SetDeviceStatusRequest) GetBatteryLevelUnavailable() bool {
	if m != nil {
		return m.BatteryLevelUnavailable
	}
	return false
}

func (m *SetDeviceStatusRequest) GetBatteryLevel() float32 {
	if m != nil {
		return m.BatteryLevel
	}
	return 0
}

type SetDeviceLocationRequest struct {
	// Device EUI (8 bytes).
	DevEui []byte `protobuf:"bytes,1,opt,name=dev_eui,json=devEui,proto3" json:"dev_eui,omitempty"`
	// The location of the device.
	Location             *common.Location `protobuf:"bytes,2,opt,name=location,proto3" json:"location,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *SetDeviceLocationRequest) Reset()         { *m = SetDeviceLocationRequest{} }
func (m *SetDeviceLocationRequest) String() string { return proto.CompactTextString(m) }
func (*SetDeviceLocationRequest) ProtoMessage()    {}
func (*SetDeviceLocationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_as_8dba2f46bdb31a35, []int{6}
}
func (m *SetDeviceLocationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetDeviceLocationRequest.Unmarshal(m, b)
}
func (m *SetDeviceLocationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetDeviceLocationRequest.Marshal(b, m, deterministic)
}
func (dst *SetDeviceLocationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetDeviceLocationRequest.Merge(dst, src)
}
func (m *SetDeviceLocationRequest) XXX_Size() int {
	return xxx_messageInfo_SetDeviceLocationRequest.Size(m)
}
func (m *SetDeviceLocationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetDeviceLocationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetDeviceLocationRequest proto.InternalMessageInfo

func (m *SetDeviceLocationRequest) GetDevEui() []byte {
	if m != nil {
		return m.DevEui
	}
	return nil
}

func (m *SetDeviceLocationRequest) GetLocation() *common.Location {
	if m != nil {
		return m.Location
	}
	return nil
}

func init() {
	proto.RegisterType((*DeviceActivationContext)(nil), "as.DeviceActivationContext")
	proto.RegisterType((*HandleUplinkDataRequest)(nil), "as.HandleUplinkDataRequest")
	proto.RegisterType((*HandleProprietaryUplinkRequest)(nil), "as.HandleProprietaryUplinkRequest")
	proto.RegisterType((*HandleErrorRequest)(nil), "as.HandleErrorRequest")
	proto.RegisterType((*HandleDownlinkACKRequest)(nil), "as.HandleDownlinkACKRequest")
	proto.RegisterType((*SetDeviceStatusRequest)(nil), "as.SetDeviceStatusRequest")
	proto.RegisterType((*SetDeviceLocationRequest)(nil), "as.SetDeviceLocationRequest")
	proto.RegisterEnum("as.RXWindow", RXWindow_name, RXWindow_value)
	proto.RegisterEnum("as.ErrorType", ErrorType_name, ErrorType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ApplicationServerServiceClient is the client API for ApplicationServerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ApplicationServerServiceClient interface {
	// HandleUplinkData handles uplink data received from an end-device.
	HandleUplinkData(ctx context.Context, in *HandleUplinkDataRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// HandleProprietaryUplink handles proprietary uplink payloads.
	HandleProprietaryUplink(ctx context.Context, in *HandleProprietaryUplinkRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// HandleError handles an error message.
	HandleError(ctx context.Context, in *HandleErrorRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// HandleDownlinkACK handles a downlink ACK or nACK response.
	HandleDownlinkACK(ctx context.Context, in *HandleDownlinkACKRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// SetDeviceStatus updates the device-status for a device.
	SetDeviceStatus(ctx context.Context, in *SetDeviceStatusRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	// SetDeviceLocation updates the device-location for a device.
	SetDeviceLocation(ctx context.Context, in *SetDeviceLocationRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type applicationServerServiceClient struct {
	cc *grpc.ClientConn
}

func NewApplicationServerServiceClient(cc *grpc.ClientConn) ApplicationServerServiceClient {
	return &applicationServerServiceClient{cc}
}

func (c *applicationServerServiceClient) HandleUplinkData(ctx context.Context, in *HandleUplinkDataRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/as.ApplicationServerService/HandleUplinkData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationServerServiceClient) HandleProprietaryUplink(ctx context.Context, in *HandleProprietaryUplinkRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/as.ApplicationServerService/HandleProprietaryUplink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationServerServiceClient) HandleError(ctx context.Context, in *HandleErrorRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/as.ApplicationServerService/HandleError", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationServerServiceClient) HandleDownlinkACK(ctx context.Context, in *HandleDownlinkACKRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/as.ApplicationServerService/HandleDownlinkACK", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationServerServiceClient) SetDeviceStatus(ctx context.Context, in *SetDeviceStatusRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/as.ApplicationServerService/SetDeviceStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *applicationServerServiceClient) SetDeviceLocation(ctx context.Context, in *SetDeviceLocationRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/as.ApplicationServerService/SetDeviceLocation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ApplicationServerServiceServer is the server API for ApplicationServerService service.
type ApplicationServerServiceServer interface {
	// HandleUplinkData handles uplink data received from an end-device.
	HandleUplinkData(context.Context, *HandleUplinkDataRequest) (*empty.Empty, error)
	// HandleProprietaryUplink handles proprietary uplink payloads.
	HandleProprietaryUplink(context.Context, *HandleProprietaryUplinkRequest) (*empty.Empty, error)
	// HandleError handles an error message.
	HandleError(context.Context, *HandleErrorRequest) (*empty.Empty, error)
	// HandleDownlinkACK handles a downlink ACK or nACK response.
	HandleDownlinkACK(context.Context, *HandleDownlinkACKRequest) (*empty.Empty, error)
	// SetDeviceStatus updates the device-status for a device.
	SetDeviceStatus(context.Context, *SetDeviceStatusRequest) (*empty.Empty, error)
	// SetDeviceLocation updates the device-location for a device.
	SetDeviceLocation(context.Context, *SetDeviceLocationRequest) (*empty.Empty, error)
}

func RegisterApplicationServerServiceServer(s *grpc.Server, srv ApplicationServerServiceServer) {
	s.RegisterService(&_ApplicationServerService_serviceDesc, srv)
}

func _ApplicationServerService_HandleUplinkData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HandleUplinkDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServerServiceServer).HandleUplinkData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/as.ApplicationServerService/HandleUplinkData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServerServiceServer).HandleUplinkData(ctx, req.(*HandleUplinkDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApplicationServerService_HandleProprietaryUplink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HandleProprietaryUplinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServerServiceServer).HandleProprietaryUplink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/as.ApplicationServerService/HandleProprietaryUplink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServerServiceServer).HandleProprietaryUplink(ctx, req.(*HandleProprietaryUplinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApplicationServerService_HandleError_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HandleErrorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServerServiceServer).HandleError(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/as.ApplicationServerService/HandleError",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServerServiceServer).HandleError(ctx, req.(*HandleErrorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApplicationServerService_HandleDownlinkACK_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HandleDownlinkACKRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServerServiceServer).HandleDownlinkACK(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/as.ApplicationServerService/HandleDownlinkACK",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServerServiceServer).HandleDownlinkACK(ctx, req.(*HandleDownlinkACKRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApplicationServerService_SetDeviceStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetDeviceStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServerServiceServer).SetDeviceStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/as.ApplicationServerService/SetDeviceStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServerServiceServer).SetDeviceStatus(ctx, req.(*SetDeviceStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ApplicationServerService_SetDeviceLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetDeviceLocationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ApplicationServerServiceServer).SetDeviceLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/as.ApplicationServerService/SetDeviceLocation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ApplicationServerServiceServer).SetDeviceLocation(ctx, req.(*SetDeviceLocationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ApplicationServerService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "as.ApplicationServerService",
	HandlerType: (*ApplicationServerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HandleUplinkData",
			Handler:    _ApplicationServerService_HandleUplinkData_Handler,
		},
		{
			MethodName: "HandleProprietaryUplink",
			Handler:    _ApplicationServerService_HandleProprietaryUplink_Handler,
		},
		{
			MethodName: "HandleError",
			Handler:    _ApplicationServerService_HandleError_Handler,
		},
		{
			MethodName: "HandleDownlinkACK",
			Handler:    _ApplicationServerService_HandleDownlinkACK_Handler,
		},
		{
			MethodName: "SetDeviceStatus",
			Handler:    _ApplicationServerService_SetDeviceStatus_Handler,
		},
		{
			MethodName: "SetDeviceLocation",
			Handler:    _ApplicationServerService_SetDeviceLocation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "as.proto",
}

func init() { proto.RegisterFile("as.proto", fileDescriptor_as_8dba2f46bdb31a35) }

var fileDescriptor_as_8dba2f46bdb31a35 = []byte{
	// 880 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x54, 0x4d, 0x73, 0xe2, 0x46,
	0x10, 0x35, 0xdf, 0xb8, 0xb1, 0xd7, 0xca, 0x38, 0x0b, 0x32, 0xbb, 0x95, 0x10, 0xe5, 0xe2, 0x6c,
	0xa5, 0xa0, 0x42, 0x6e, 0xb9, 0xa4, 0x28, 0x50, 0x36, 0x94, 0x77, 0x37, 0x44, 0x40, 0xec, 0xca,
	0x65, 0x6a, 0x2c, 0x35, 0x94, 0x62, 0xa1, 0x51, 0x86, 0x01, 0xac, 0x4a, 0xe5, 0x07, 0xe5, 0x4f,
	0xe5, 0x90, 0x5f, 0x91, 0x63, 0x6a, 0x46, 0xe2, 0xc3, 0x6b, 0x83, 0x73, 0x81, 0x99, 0x7e, 0xad,
	0xd7, 0x4f, 0x4f, 0xdd, 0x0d, 0x65, 0x36, 0x6f, 0x46, 0x82, 0x4b, 0x4e, 0xb2, 0x6c, 0x5e, 0x7f,
	0x35, 0xe5, 0x7c, 0x1a, 0x60, 0x4b, 0x47, 0x6e, 0x17, 0x93, 0x16, 0xce, 0x22, 0x19, 0x27, 0x09,
	0xf5, 0x1a, 0x8b, 0xfc, 0x96, 0xcb, 0x67, 0x33, 0x1e, 0xa6, 0x7f, 0x29, 0x70, 0xa6, 0x80, 0xe9,
	0xaa, 0x35, 0x5d, 0x25, 0x01, 0x0b, 0xa1, 0xd6, 0xc3, 0xa5, 0xef, 0x62, 0xc7, 0x95, 0xfe, 0x92,
	0x49, 0x9f, 0x87, 0x5d, 0x1e, 0x4a, 0xbc, 0x97, 0xe4, 0x02, 0xca, 0x1e, 0x2e, 0x29, 0xf3, 0x3c,
	0x61, 0x66, 0x1a, 0x99, 0xcb, 0x13, 0xa7, 0xe4, 0xe1, 0xb2, 0xe3, 0x79, 0x82, 0xb4, 0xe0, 0x98,
	0x45, 0x11, 0x9d, 0xd3, 0x3b, 0x8c, 0xcd, 0x6c, 0x23, 0x73, 0x59, 0x69, 0x9f, 0x37, 0xd3, 0x42,
	0x57, 0x18, 0xdb, 0xe1, 0x12, 0x03, 0x1e, 0xa1, 0x53, 0x62, 0x51, 0x34, 0xbc, 0xc2, 0xd8, 0xfa,
	0x3b, 0x0b, 0xb5, 0x1f, 0x59, 0xe8, 0x05, 0x38, 0x8e, 0x02, 0x3f, 0xbc, 0xeb, 0x31, 0xc9, 0x1c,
	0xfc, 0x7d, 0x81, 0x73, 0x49, 0x6a, 0xa0, 0x78, 0x29, 0x2e, 0xfc, 0xb4, 0x4c, 0xd1, 0xc3, 0xa5,
	0xbd, 0xf0, 0x95, 0x80, 0xdf, 0xb8, 0x1f, 0x6a, 0x24, 0x9b, 0x08, 0x50, 0x77, 0x05, 0x9d, 0x43,
	0x61, 0x42, 0xdd, 0x50, 0x9a, 0xb9, 0x46, 0xe6, 0xf2, 0xd4, 0xc9, 0x4f, 0xba, 0xa1, 0x24, 0x2f,
	0xa1, 0x38, 0xa1, 0x11, 0x17, 0xd2, 0xcc, 0xeb, 0x68, 0x61, 0x32, 0xe0, 0x42, 0x12, 0x03, 0x72,
	0xcc, 0x13, 0x66, 0xa1, 0x91, 0xb9, 0x2c, 0x3b, 0xea, 0x48, 0x5e, 0x40, 0xd6, 0x13, 0x66, 0x51,
	0x27, 0x65, 0x3d, 0x41, 0xbe, 0x82, 0x92, 0xbc, 0xa7, 0x7e, 0x38, 0xe1, 0x66, 0x49, 0xbf, 0x8c,
	0xd1, 0x9c, 0xae, 0x9a, 0x89, 0xd2, 0xd1, 0x4d, 0x3f, 0x9c, 0x70, 0xa7, 0x28, 0xef, 0xd5, 0xbf,
	0x4a, 0x15, 0x69, 0x6a, 0xb9, 0x91, 0x7b, 0x98, 0xea, 0xa4, 0xa9, 0x22, 0x49, 0x25, 0x90, 0xf7,
	0x98, 0x64, 0xe6, 0xb1, 0x96, 0xae, 0xcf, 0xe4, 0x1a, 0x2e, 0x3c, 0x6d, 0x37, 0x65, 0x1b, 0xbf,
	0xa9, 0x9b, 0x18, 0x6e, 0x82, 0xae, 0xfd, 0xaa, 0xc9, 0xe6, 0xcd, 0x3d, 0xdf, 0xc4, 0xa9, 0x79,
	0x4f, 0x03, 0xd6, 0x5f, 0x19, 0xf8, 0x2c, 0x31, 0x78, 0x20, 0x78, 0x24, 0x7c, 0x94, 0x4c, 0xc4,
	0xa9, 0xac, 0xd4, 0xe7, 0xcf, 0xa1, 0x32, 0x63, 0x2e, 0x8d, 0x58, 0x1c, 0x70, 0xe6, 0xa5, 0x5e,
	0xc3, 0x8c, 0xb9, 0x83, 0x24, 0xa2, 0x8c, 0x9a, 0xf9, 0x6e, 0x6a, 0xb5, 0x3a, 0xee, 0x1a, 0x93,
	0xfb, 0xff, 0xc6, 0xe4, 0x0f, 0x1b, 0x63, 0xfd, 0x01, 0x24, 0x91, 0x6a, 0x0b, 0xc1, 0xc5, 0xb3,
	0x6d, 0xf0, 0x05, 0xe4, 0x65, 0x1c, 0xa1, 0x56, 0xf0, 0xa2, 0x7d, 0xaa, 0xec, 0xd1, 0x0f, 0x8e,
	0xe2, 0x08, 0x1d, 0x0d, 0x91, 0x4f, 0xa1, 0x80, 0x2a, 0xa4, 0x3f, 0xfc, 0xb1, 0x93, 0x5c, 0xb6,
	0x4d, 0x52, 0xd8, 0x36, 0x89, 0x15, 0x80, 0x99, 0x14, 0xef, 0xf1, 0x55, 0xa8, 0xc4, 0x75, 0xba,
	0x57, 0xcf, 0x4a, 0xd8, 0x30, 0x65, 0x77, 0xda, 0xcd, 0x82, 0x13, 0xe6, 0xde, 0x85, 0x7c, 0x15,
	0xa0, 0x37, 0x45, 0x4f, 0xeb, 0x2b, 0x3b, 0x0f, 0x62, 0xd6, 0xbf, 0x19, 0xa8, 0x0e, 0x51, 0x26,
	0x9f, 0x73, 0x28, 0x99, 0x5c, 0xcc, 0x9f, 0x2d, 0x66, 0x42, 0xe9, 0x96, 0x49, 0x89, 0x22, 0x4e,
	0xcb, 0xad, 0xaf, 0xa4, 0x0a, 0xc5, 0x19, 0x13, 0x53, 0x3f, 0xd4, 0xb5, 0x0a, 0x4e, 0x7a, 0x23,
	0x6d, 0x78, 0x89, 0xf7, 0x12, 0x45, 0xc8, 0x02, 0x1a, 0xf1, 0x15, 0x0a, 0x3a, 0xe7, 0x0b, 0xe1,
	0xa2, 0xb6, 0xa3, 0xec, 0x9c, 0xaf, 0xc1, 0x81, 0xc2, 0x86, 0x1a, 0x22, 0xdf, 0xc1, 0x45, 0x4a,
	0x4b, 0x03, 0x5c, 0x62, 0x40, 0x17, 0x21, 0x5b, 0x32, 0x3f, 0x60, 0xb7, 0x01, 0xa6, 0xb3, 0x52,
	0x4b, 0x13, 0xde, 0x29, 0x7c, 0xbc, 0x85, 0xc9, 0x97, 0x70, 0xfa, 0xe0, 0x59, 0x3d, 0x4a, 0x59,
	0xe7, 0x64, 0x37, 0xdf, 0x62, 0x60, 0x6e, 0xde, 0xfc, 0x1d, 0x77, 0x75, 0xb7, 0x3e, 0xfb, 0xee,
	0x5f, 0x43, 0x39, 0x48, 0x73, 0xd3, 0xbd, 0x62, 0xac, 0xf7, 0xca, 0x86, 0x63, 0x93, 0xf1, 0xe6,
	0x35, 0x94, 0x9d, 0x9b, 0x6b, 0x3f, 0xf4, 0xf8, 0x8a, 0x94, 0x20, 0xe7, 0xdc, 0x7c, 0x63, 0x1c,
	0x25, 0x87, 0xb6, 0x91, 0x79, 0xf3, 0x27, 0x1c, 0x6f, 0xfa, 0x84, 0x54, 0xa0, 0xf4, 0xd6, 0xfe,
	0x60, 0x3b, 0xfd, 0xae, 0x71, 0x44, 0xca, 0x90, 0xff, 0x69, 0xd4, 0xe9, 0x18, 0x19, 0x62, 0xc0,
	0x49, 0xaf, 0x33, 0xea, 0xd0, 0xf1, 0x80, 0xfe, 0xd0, 0xfd, 0x30, 0x32, 0xb2, 0xe4, 0x0c, 0x2a,
	0xeb, 0xc8, 0xfb, 0x7e, 0xd7, 0xc8, 0x91, 0x3a, 0x54, 0x7b, 0xf6, 0x2f, 0xfd, 0xae, 0x4d, 0x7f,
	0x1e, 0xdb, 0x63, 0x9b, 0xf6, 0x47, 0xf6, 0x7b, 0x3a, 0xec, 0xff, 0x6a, 0x1b, 0xf9, 0xa7, 0x31,
	0x4d, 0x54, 0x68, 0xff, 0x93, 0x03, 0xb3, 0x13, 0x45, 0x81, 0x9f, 0x88, 0x1d, 0xa2, 0x58, 0xa2,
	0x50, 0xbf, 0xbe, 0x8b, 0xa4, 0x0f, 0xc6, 0xc7, 0xeb, 0x90, 0xe8, 0xc1, 0xdf, 0xb3, 0x24, 0xeb,
	0xd5, 0x66, 0xb2, 0xef, 0x9b, 0xeb, 0x7d, 0xdf, 0xb4, 0xd5, 0xbe, 0xb7, 0x8e, 0xc8, 0xf5, 0x7a,
	0xb3, 0x3e, 0x1a, 0x7c, 0x62, 0x6d, 0x19, 0xf7, 0x6d, 0x85, 0x03, 0xc4, 0xdf, 0x43, 0x65, 0x67,
	0x4c, 0x49, 0x75, 0x4b, 0xb6, 0x3b, 0xb7, 0x07, 0x08, 0xae, 0xe0, 0x93, 0x47, 0xa3, 0x46, 0x5e,
	0x6f, 0x69, 0x1e, 0x4f, 0xe0, 0x01, 0xb2, 0xb7, 0x70, 0xf6, 0xd1, 0x20, 0x91, 0xba, 0xa2, 0x7a,
	0x7a, 0xba, 0x0e, 0xab, 0x7a, 0xd4, 0x97, 0x89, 0xaa, 0x7d, 0xed, 0xba, 0x9f, 0xec, 0xb6, 0xa8,
	0x23, 0xdf, 0xfe, 0x17, 0x00, 0x00, 0xff, 0xff, 0x3c, 0x1e, 0x41, 0x86, 0x9c, 0x07, 0x00, 0x00,
}
