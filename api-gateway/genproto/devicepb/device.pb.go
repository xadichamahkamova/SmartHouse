// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.0
// source: protos/device-service/device.proto

package devicepb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Configuration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Brightness int32  `protobuf:"varint,1,opt,name=brightness,proto3" json:"brightness,omitempty"`
	Color      string `protobuf:"bytes,2,opt,name=color,proto3" json:"color,omitempty"`
}

func (x *Configuration) Reset() {
	*x = Configuration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_device_service_device_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Configuration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Configuration) ProtoMessage() {}

func (x *Configuration) ProtoReflect() protoreflect.Message {
	mi := &file_protos_device_service_device_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Configuration.ProtoReflect.Descriptor instead.
func (*Configuration) Descriptor() ([]byte, []int) {
	return file_protos_device_service_device_proto_rawDescGZIP(), []int{0}
}

func (x *Configuration) GetBrightness() int32 {
	if x != nil {
		return x.Brightness
	}
	return 0
}

func (x *Configuration) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

type Device struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceId           string         `protobuf:"bytes,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
	UserId             string         `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	DeviceType         string         `protobuf:"bytes,3,opt,name=device_type,json=deviceType,proto3" json:"device_type,omitempty"`
	DeviceName         string         `protobuf:"bytes,4,opt,name=device_name,json=deviceName,proto3" json:"device_name,omitempty"`
	DeviceStatus       string         `protobuf:"bytes,5,opt,name=device_status,json=deviceStatus,proto3" json:"device_status,omitempty"`
	Configuration      *Configuration `protobuf:"bytes,6,opt,name=configuration,proto3" json:"configuration,omitempty"`
	Location           string         `protobuf:"bytes,7,opt,name=location,proto3" json:"location,omitempty"`
	FirmwareVersion    string         `protobuf:"bytes,8,opt,name=firmware_version,json=firmwareVersion,proto3" json:"firmware_version,omitempty"`
	ConnectivityStatus string         `protobuf:"bytes,9,opt,name=connectivity_status,json=connectivityStatus,proto3" json:"connectivity_status,omitempty"`
	CreatedAt          string         `protobuf:"bytes,10,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt          string         `protobuf:"bytes,11,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	DeletedAt          int64          `protobuf:"varint,12,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
}

func (x *Device) Reset() {
	*x = Device{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_device_service_device_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Device) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Device) ProtoMessage() {}

func (x *Device) ProtoReflect() protoreflect.Message {
	mi := &file_protos_device_service_device_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Device.ProtoReflect.Descriptor instead.
func (*Device) Descriptor() ([]byte, []int) {
	return file_protos_device_service_device_proto_rawDescGZIP(), []int{1}
}

func (x *Device) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

func (x *Device) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Device) GetDeviceType() string {
	if x != nil {
		return x.DeviceType
	}
	return ""
}

func (x *Device) GetDeviceName() string {
	if x != nil {
		return x.DeviceName
	}
	return ""
}

func (x *Device) GetDeviceStatus() string {
	if x != nil {
		return x.DeviceStatus
	}
	return ""
}

func (x *Device) GetConfiguration() *Configuration {
	if x != nil {
		return x.Configuration
	}
	return nil
}

func (x *Device) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *Device) GetFirmwareVersion() string {
	if x != nil {
		return x.FirmwareVersion
	}
	return ""
}

func (x *Device) GetConnectivityStatus() string {
	if x != nil {
		return x.ConnectivityStatus
	}
	return ""
}

func (x *Device) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Device) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *Device) GetDeletedAt() int64 {
	if x != nil {
		return x.DeletedAt
	}
	return 0
}

type SingleID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DeviceId string `protobuf:"bytes,1,opt,name=device_id,json=deviceId,proto3" json:"device_id,omitempty"`
}

func (x *SingleID) Reset() {
	*x = SingleID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_device_service_device_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SingleID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SingleID) ProtoMessage() {}

func (x *SingleID) ProtoReflect() protoreflect.Message {
	mi := &file_protos_device_service_device_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SingleID.ProtoReflect.Descriptor instead.
func (*SingleID) Descriptor() ([]byte, []int) {
	return file_protos_device_service_device_proto_rawDescGZIP(), []int{2}
}

func (x *SingleID) GetDeviceId() string {
	if x != nil {
		return x.DeviceId
	}
	return ""
}

type ListRequestOfDevice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListRequestOfDevice) Reset() {
	*x = ListRequestOfDevice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_device_service_device_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListRequestOfDevice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListRequestOfDevice) ProtoMessage() {}

func (x *ListRequestOfDevice) ProtoReflect() protoreflect.Message {
	mi := &file_protos_device_service_device_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListRequestOfDevice.ProtoReflect.Descriptor instead.
func (*ListRequestOfDevice) Descriptor() ([]byte, []int) {
	return file_protos_device_service_device_proto_rawDescGZIP(), []int{3}
}

type ListResponseOfDevice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*Device `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *ListResponseOfDevice) Reset() {
	*x = ListResponseOfDevice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_device_service_device_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListResponseOfDevice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListResponseOfDevice) ProtoMessage() {}

func (x *ListResponseOfDevice) ProtoReflect() protoreflect.Message {
	mi := &file_protos_device_service_device_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListResponseOfDevice.ProtoReflect.Descriptor instead.
func (*ListResponseOfDevice) Descriptor() ([]byte, []int) {
	return file_protos_device_service_device_proto_rawDescGZIP(), []int{4}
}

func (x *ListResponseOfDevice) GetList() []*Device {
	if x != nil {
		return x.List
	}
	return nil
}

type ResponseOfDevice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Status string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *ResponseOfDevice) Reset() {
	*x = ResponseOfDevice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_device_service_device_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseOfDevice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseOfDevice) ProtoMessage() {}

func (x *ResponseOfDevice) ProtoReflect() protoreflect.Message {
	mi := &file_protos_device_service_device_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseOfDevice.ProtoReflect.Descriptor instead.
func (*ResponseOfDevice) Descriptor() ([]byte, []int) {
	return file_protos_device_service_device_proto_rawDescGZIP(), []int{5}
}

func (x *ResponseOfDevice) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ResponseOfDevice) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_protos_device_service_device_proto protoreflect.FileDescriptor

var file_protos_device_service_device_proto_rawDesc = []byte{
	0x0a, 0x22, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2d,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x45, 0x0a, 0x0d, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e, 0x0a, 0x0a, 0x62, 0x72, 0x69, 0x67, 0x68, 0x74, 0x6e,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x62, 0x72, 0x69, 0x67, 0x68,
	0x74, 0x6e, 0x65, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x22, 0xb0, 0x03, 0x0a, 0x06,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a,
	0x0b, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23,
	0x0a, 0x0d, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x34, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x29, 0x0a, 0x10, 0x66, 0x69, 0x72, 0x6d, 0x77, 0x61, 0x72,
	0x65, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0f, 0x66, 0x69, 0x72, 0x6d, 0x77, 0x61, 0x72, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x2f, 0x0a, 0x13, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x27,
	0x0a, 0x08, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x49, 0x44, 0x12, 0x1b, 0x0a, 0x09, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x22, 0x15, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4f, 0x66, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x22, 0x33,
	0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4f, 0x66,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1b, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x04, 0x6c,
	0x69, 0x73, 0x74, 0x22, 0x3a, 0x0a, 0x10, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4f,
	0x66, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32,
	0xf3, 0x01, 0x0a, 0x0d, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x2a, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x07, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x11, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4f, 0x66, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x1f, 0x0a,
	0x09, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x09, 0x2e, 0x53, 0x69, 0x6e,
	0x67, 0x6c, 0x65, 0x49, 0x44, 0x1a, 0x07, 0x2e, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3b,
	0x0a, 0x0c, 0x4c, 0x69, 0x73, 0x74, 0x4f, 0x66, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x14,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x4f, 0x66, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x1a, 0x15, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x4f, 0x66, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2a, 0x0a, 0x0c, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x07, 0x2e, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x1a, 0x11, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4f,
	0x66, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2c, 0x0a, 0x0c, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x09, 0x2e, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65,
	0x49, 0x44, 0x1a, 0x11, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4f, 0x66, 0x44,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x42, 0x15, 0x5a, 0x13, 0x2e, 0x2f, 0x67, 0x65, 0x6e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_device_service_device_proto_rawDescOnce sync.Once
	file_protos_device_service_device_proto_rawDescData = file_protos_device_service_device_proto_rawDesc
)

func file_protos_device_service_device_proto_rawDescGZIP() []byte {
	file_protos_device_service_device_proto_rawDescOnce.Do(func() {
		file_protos_device_service_device_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_device_service_device_proto_rawDescData)
	})
	return file_protos_device_service_device_proto_rawDescData
}

var file_protos_device_service_device_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_protos_device_service_device_proto_goTypes = []any{
	(*Configuration)(nil),        // 0: Configuration
	(*Device)(nil),               // 1: Device
	(*SingleID)(nil),             // 2: SingleID
	(*ListRequestOfDevice)(nil),  // 3: ListRequestOfDevice
	(*ListResponseOfDevice)(nil), // 4: ListResponseOfDevice
	(*ResponseOfDevice)(nil),     // 5: ResponseOfDevice
}
var file_protos_device_service_device_proto_depIdxs = []int32{
	0, // 0: Device.configuration:type_name -> Configuration
	1, // 1: ListResponseOfDevice.list:type_name -> Device
	1, // 2: DeviceService.CreateDevice:input_type -> Device
	2, // 3: DeviceService.GetDevice:input_type -> SingleID
	3, // 4: DeviceService.ListOfDevice:input_type -> ListRequestOfDevice
	1, // 5: DeviceService.UpdateDevice:input_type -> Device
	2, // 6: DeviceService.DeleteDevice:input_type -> SingleID
	5, // 7: DeviceService.CreateDevice:output_type -> ResponseOfDevice
	1, // 8: DeviceService.GetDevice:output_type -> Device
	4, // 9: DeviceService.ListOfDevice:output_type -> ListResponseOfDevice
	5, // 10: DeviceService.UpdateDevice:output_type -> ResponseOfDevice
	5, // 11: DeviceService.DeleteDevice:output_type -> ResponseOfDevice
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_protos_device_service_device_proto_init() }
func file_protos_device_service_device_proto_init() {
	if File_protos_device_service_device_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_device_service_device_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Configuration); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_device_service_device_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Device); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_device_service_device_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*SingleID); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_device_service_device_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*ListRequestOfDevice); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_device_service_device_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*ListResponseOfDevice); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_protos_device_service_device_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*ResponseOfDevice); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protos_device_service_device_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_device_service_device_proto_goTypes,
		DependencyIndexes: file_protos_device_service_device_proto_depIdxs,
		MessageInfos:      file_protos_device_service_device_proto_msgTypes,
	}.Build()
	File_protos_device_service_device_proto = out.File
	file_protos_device_service_device_proto_rawDesc = nil
	file_protos_device_service_device_proto_goTypes = nil
	file_protos_device_service_device_proto_depIdxs = nil
}
