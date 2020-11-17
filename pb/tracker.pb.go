// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.11.4
// source: tracker.proto

package pb

import (
	context "context"
	duration "github.com/golang/protobuf/ptypes/duration"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Phone string `protobuf:"bytes,3,opt,name=phone,proto3" json:"phone,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tracker_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_tracker_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_tracker_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

type UpdateUser struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *UpdateUser) Reset() {
	*x = UpdateUser{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tracker_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateUser) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateUser) ProtoMessage() {}

func (x *UpdateUser) ProtoReflect() protoreflect.Message {
	mi := &file_tracker_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateUser.ProtoReflect.Descriptor instead.
func (*UpdateUser) Descriptor() ([]byte, []int) {
	return file_tracker_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateUser) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type Name struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Name) Reset() {
	*x = Name{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tracker_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Name) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Name) ProtoMessage() {}

func (x *Name) ProtoReflect() protoreflect.Message {
	mi := &file_tracker_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Name.ProtoReflect.Descriptor instead.
func (*Name) Descriptor() ([]byte, []int) {
	return file_tracker_proto_rawDescGZIP(), []int{2}
}

func (x *Name) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response string `protobuf:"bytes,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tracker_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_tracker_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_tracker_proto_rawDescGZIP(), []int{3}
}

func (x *Response) GetResponse() string {
	if x != nil {
		return x.Response
	}
	return ""
}

type Activity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type      string               `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Timestamp *timestamp.Timestamp `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Duration  *duration.Duration   `protobuf:"bytes,3,opt,name=duration,proto3" json:"duration,omitempty"`
	Label     string               `protobuf:"bytes,4,opt,name=label,proto3" json:"label,omitempty"`
	Status    string               `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`
	Valid     string               `protobuf:"bytes,6,opt,name=valid,proto3" json:"valid,omitempty"`
}

func (x *Activity) Reset() {
	*x = Activity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tracker_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Activity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Activity) ProtoMessage() {}

func (x *Activity) ProtoReflect() protoreflect.Message {
	mi := &file_tracker_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Activity.ProtoReflect.Descriptor instead.
func (*Activity) Descriptor() ([]byte, []int) {
	return file_tracker_proto_rawDescGZIP(), []int{4}
}

func (x *Activity) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Activity) GetTimestamp() *timestamp.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *Activity) GetDuration() *duration.Duration {
	if x != nil {
		return x.Duration
	}
	return nil
}

func (x *Activity) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *Activity) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Activity) GetValid() string {
	if x != nil {
		return x.Valid
	}
	return ""
}

type MapUserToActivity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Label string `protobuf:"bytes,1,opt,name=Label,proto3" json:"Label,omitempty"`
	Type  string `protobuf:"bytes,2,opt,name=Type,proto3" json:"Type,omitempty"`
}

func (x *MapUserToActivity) Reset() {
	*x = MapUserToActivity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tracker_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MapUserToActivity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MapUserToActivity) ProtoMessage() {}

func (x *MapUserToActivity) ProtoReflect() protoreflect.Message {
	mi := &file_tracker_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MapUserToActivity.ProtoReflect.Descriptor instead.
func (*MapUserToActivity) Descriptor() ([]byte, []int) {
	return file_tracker_proto_rawDescGZIP(), []int{5}
}

func (x *MapUserToActivity) GetLabel() string {
	if x != nil {
		return x.Label
	}
	return ""
}

func (x *MapUserToActivity) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type ActivityStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status bool `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *ActivityStatus) Reset() {
	*x = ActivityStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tracker_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActivityStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActivityStatus) ProtoMessage() {}

func (x *ActivityStatus) ProtoReflect() protoreflect.Message {
	mi := &file_tracker_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActivityStatus.ProtoReflect.Descriptor instead.
func (*ActivityStatus) Descriptor() ([]byte, []int) {
	return file_tracker_proto_rawDescGZIP(), []int{6}
}

func (x *ActivityStatus) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

type Valid struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Valid string `protobuf:"bytes,1,opt,name=valid,proto3" json:"valid,omitempty"`
}

func (x *Valid) Reset() {
	*x = Valid{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tracker_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Valid) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Valid) ProtoMessage() {}

func (x *Valid) ProtoReflect() protoreflect.Message {
	mi := &file_tracker_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Valid.ProtoReflect.Descriptor instead.
func (*Valid) Descriptor() ([]byte, []int) {
	return file_tracker_proto_rawDescGZIP(), []int{7}
}

func (x *Valid) GetValid() string {
	if x != nil {
		return x.Valid
	}
	return ""
}

var File_tracker_proto protoreflect.FileDescriptor

var file_tracker_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x06, 0x76, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x46, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x22, 0x2e, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x20,
	0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x76,
	0x5f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x22, 0x1a, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x26, 0x0a, 0x08,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0xd3, 0x01, 0x0a, 0x08, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12,
	0x35, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x64, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x16, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x22, 0x3d, 0x0a, 0x11, 0x4d, 0x61,
	0x70, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x22, 0x28, 0x0a, 0x0e, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x69, 0x74, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x22, 0x1d, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x32, 0xa8, 0x03, 0x0a, 0x0f, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2d, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x0c, 0x2e, 0x76, 0x5f, 0x74, 0x61, 0x73, 0x6b,
	0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x0c, 0x2e, 0x76, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55,
	0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x12, 0x2e, 0x76, 0x5f, 0x74,
	0x61, 0x73, 0x6b, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x10,
	0x2e, 0x76, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x30, 0x0a, 0x0c, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x0c, 0x2e, 0x76, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x1a, 0x10, 0x2e, 0x76, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x37, 0x0a, 0x0f, 0x41, 0x64, 0x64, 0x55, 0x73, 0x65, 0x72, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x12, 0x10, 0x2e, 0x76, 0x5f, 0x74, 0x61, 0x73, 0x6b,
	0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x1a, 0x10, 0x2e, 0x76, 0x5f, 0x74, 0x61,
	0x73, 0x6b, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x38, 0x0a,
	0x10, 0x53, 0x74, 0x6f, 0x70, 0x55, 0x73, 0x65, 0x72, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x12, 0x10, 0x2e, 0x76, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x69, 0x74, 0x79, 0x1a, 0x10, 0x2e, 0x76, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x0e, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x69, 0x74, 0x79, 0x49, 0x73, 0x44, 0x6f, 0x6e, 0x65, 0x12, 0x19, 0x2e, 0x76, 0x5f, 0x74, 0x61,
	0x73, 0x6b, 0x2e, 0x4d, 0x61, 0x70, 0x55, 0x73, 0x65, 0x72, 0x54, 0x6f, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x69, 0x74, 0x79, 0x1a, 0x16, 0x2e, 0x76, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x41, 0x63,
	0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x00, 0x12, 0x3d,
	0x0a, 0x0f, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x49, 0x73, 0x56, 0x61, 0x6c, 0x69,
	0x64, 0x12, 0x19, 0x2e, 0x76, 0x5f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x4d, 0x61, 0x70, 0x55, 0x73,
	0x65, 0x72, 0x54, 0x6f, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x1a, 0x0d, 0x2e, 0x76,
	0x5f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x22, 0x00, 0x42, 0x06, 0x5a,
	0x04, 0x2e, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tracker_proto_rawDescOnce sync.Once
	file_tracker_proto_rawDescData = file_tracker_proto_rawDesc
)

func file_tracker_proto_rawDescGZIP() []byte {
	file_tracker_proto_rawDescOnce.Do(func() {
		file_tracker_proto_rawDescData = protoimpl.X.CompressGZIP(file_tracker_proto_rawDescData)
	})
	return file_tracker_proto_rawDescData
}

var file_tracker_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_tracker_proto_goTypes = []interface{}{
	(*User)(nil),                // 0: v_task.User
	(*UpdateUser)(nil),          // 1: v_task.UpdateUser
	(*Name)(nil),                // 2: v_task.Name
	(*Response)(nil),            // 3: v_task.Response
	(*Activity)(nil),            // 4: v_task.Activity
	(*MapUserToActivity)(nil),   // 5: v_task.MapUserToActivity
	(*ActivityStatus)(nil),      // 6: v_task.ActivityStatus
	(*Valid)(nil),               // 7: v_task.Valid
	(*timestamp.Timestamp)(nil), // 8: google.protobuf.Timestamp
	(*duration.Duration)(nil),   // 9: google.protobuf.Duration
}
var file_tracker_proto_depIdxs = []int32{
	0,  // 0: v_task.UpdateUser.user:type_name -> v_task.User
	8,  // 1: v_task.Activity.timestamp:type_name -> google.protobuf.Timestamp
	9,  // 2: v_task.Activity.duration:type_name -> google.protobuf.Duration
	2,  // 3: v_task.ActivityService.GetUserByName:input_type -> v_task.Name
	1,  // 4: v_task.ActivityService.UpdateUserDetails:input_type -> v_task.UpdateUser
	0,  // 5: v_task.ActivityService.RegisterUser:input_type -> v_task.User
	4,  // 6: v_task.ActivityService.AddUserActivity:input_type -> v_task.Activity
	4,  // 7: v_task.ActivityService.StopUserActivity:input_type -> v_task.Activity
	5,  // 8: v_task.ActivityService.ActivityIsDone:input_type -> v_task.MapUserToActivity
	5,  // 9: v_task.ActivityService.ActivityIsValid:input_type -> v_task.MapUserToActivity
	0,  // 10: v_task.ActivityService.GetUserByName:output_type -> v_task.User
	3,  // 11: v_task.ActivityService.UpdateUserDetails:output_type -> v_task.Response
	3,  // 12: v_task.ActivityService.RegisterUser:output_type -> v_task.Response
	3,  // 13: v_task.ActivityService.AddUserActivity:output_type -> v_task.Response
	3,  // 14: v_task.ActivityService.StopUserActivity:output_type -> v_task.Response
	6,  // 15: v_task.ActivityService.ActivityIsDone:output_type -> v_task.ActivityStatus
	7,  // 16: v_task.ActivityService.ActivityIsValid:output_type -> v_task.Valid
	10, // [10:17] is the sub-list for method output_type
	3,  // [3:10] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_tracker_proto_init() }
func file_tracker_proto_init() {
	if File_tracker_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tracker_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_tracker_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateUser); i {
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
		file_tracker_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Name); i {
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
		file_tracker_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_tracker_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Activity); i {
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
		file_tracker_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MapUserToActivity); i {
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
		file_tracker_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActivityStatus); i {
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
		file_tracker_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Valid); i {
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
			RawDescriptor: file_tracker_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_tracker_proto_goTypes,
		DependencyIndexes: file_tracker_proto_depIdxs,
		MessageInfos:      file_tracker_proto_msgTypes,
	}.Build()
	File_tracker_proto = out.File
	file_tracker_proto_rawDesc = nil
	file_tracker_proto_goTypes = nil
	file_tracker_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ActivityServiceClient is the client API for ActivityService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ActivityServiceClient interface {
	// Get User Details by userame
	GetUserByName(ctx context.Context, in *Name, opts ...grpc.CallOption) (*User, error)
	// Update a user
	UpdateUserDetails(ctx context.Context, in *UpdateUser, opts ...grpc.CallOption) (*Response, error)
	// Add a new User
	RegisterUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*Response, error)
	//Add User's Activity
	AddUserActivity(ctx context.Context, in *Activity, opts ...grpc.CallOption) (*Response, error)
	//Stop User's Activity
	StopUserActivity(ctx context.Context, in *Activity, opts ...grpc.CallOption) (*Response, error)
	//Check if an Activity is Done
	ActivityIsDone(ctx context.Context, in *MapUserToActivity, opts ...grpc.CallOption) (*ActivityStatus, error)
	//Check if an Activity is Valid
	ActivityIsValid(ctx context.Context, in *MapUserToActivity, opts ...grpc.CallOption) (*Valid, error)
}

type activityServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewActivityServiceClient(cc grpc.ClientConnInterface) ActivityServiceClient {
	return &activityServiceClient{cc}
}

func (c *activityServiceClient) GetUserByName(ctx context.Context, in *Name, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/v_task.ActivityService/GetUserByName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activityServiceClient) UpdateUserDetails(ctx context.Context, in *UpdateUser, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/v_task.ActivityService/UpdateUserDetails", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activityServiceClient) RegisterUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/v_task.ActivityService/RegisterUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activityServiceClient) AddUserActivity(ctx context.Context, in *Activity, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/v_task.ActivityService/AddUserActivity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activityServiceClient) StopUserActivity(ctx context.Context, in *Activity, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/v_task.ActivityService/StopUserActivity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activityServiceClient) ActivityIsDone(ctx context.Context, in *MapUserToActivity, opts ...grpc.CallOption) (*ActivityStatus, error) {
	out := new(ActivityStatus)
	err := c.cc.Invoke(ctx, "/v_task.ActivityService/ActivityIsDone", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activityServiceClient) ActivityIsValid(ctx context.Context, in *MapUserToActivity, opts ...grpc.CallOption) (*Valid, error) {
	out := new(Valid)
	err := c.cc.Invoke(ctx, "/v_task.ActivityService/ActivityIsValid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ActivityServiceServer is the server API for ActivityService service.
type ActivityServiceServer interface {
	// Get User Details by userame
	GetUserByName(context.Context, *Name) (*User, error)
	// Update a user
	UpdateUserDetails(context.Context, *UpdateUser) (*Response, error)
	// Add a new User
	RegisterUser(context.Context, *User) (*Response, error)
	//Add User's Activity
	AddUserActivity(context.Context, *Activity) (*Response, error)
	//Stop User's Activity
	StopUserActivity(context.Context, *Activity) (*Response, error)
	//Check if an Activity is Done
	ActivityIsDone(context.Context, *MapUserToActivity) (*ActivityStatus, error)
	//Check if an Activity is Valid
	ActivityIsValid(context.Context, *MapUserToActivity) (*Valid, error)
}

// UnimplementedActivityServiceServer can be embedded to have forward compatible implementations.
type UnimplementedActivityServiceServer struct {
}

func (*UnimplementedActivityServiceServer) GetUserByName(context.Context, *Name) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserByName not implemented")
}
func (*UnimplementedActivityServiceServer) UpdateUserDetails(context.Context, *UpdateUser) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserDetails not implemented")
}
func (*UnimplementedActivityServiceServer) RegisterUser(context.Context, *User) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterUser not implemented")
}
func (*UnimplementedActivityServiceServer) AddUserActivity(context.Context, *Activity) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUserActivity not implemented")
}
func (*UnimplementedActivityServiceServer) StopUserActivity(context.Context, *Activity) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopUserActivity not implemented")
}
func (*UnimplementedActivityServiceServer) ActivityIsDone(context.Context, *MapUserToActivity) (*ActivityStatus, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ActivityIsDone not implemented")
}
func (*UnimplementedActivityServiceServer) ActivityIsValid(context.Context, *MapUserToActivity) (*Valid, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ActivityIsValid not implemented")
}

func RegisterActivityServiceServer(s *grpc.Server, srv ActivityServiceServer) {
	s.RegisterService(&_ActivityService_serviceDesc, srv)
}

func _ActivityService_GetUserByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Name)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityServiceServer).GetUserByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v_task.ActivityService/GetUserByName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityServiceServer).GetUserByName(ctx, req.(*Name))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActivityService_UpdateUserDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateUser)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityServiceServer).UpdateUserDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v_task.ActivityService/UpdateUserDetails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityServiceServer).UpdateUserDetails(ctx, req.(*UpdateUser))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActivityService_RegisterUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityServiceServer).RegisterUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v_task.ActivityService/RegisterUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityServiceServer).RegisterUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActivityService_AddUserActivity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Activity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityServiceServer).AddUserActivity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v_task.ActivityService/AddUserActivity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityServiceServer).AddUserActivity(ctx, req.(*Activity))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActivityService_StopUserActivity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Activity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityServiceServer).StopUserActivity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v_task.ActivityService/StopUserActivity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityServiceServer).StopUserActivity(ctx, req.(*Activity))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActivityService_ActivityIsDone_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MapUserToActivity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityServiceServer).ActivityIsDone(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v_task.ActivityService/ActivityIsDone",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityServiceServer).ActivityIsDone(ctx, req.(*MapUserToActivity))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActivityService_ActivityIsValid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MapUserToActivity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityServiceServer).ActivityIsValid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/v_task.ActivityService/ActivityIsValid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityServiceServer).ActivityIsValid(ctx, req.(*MapUserToActivity))
	}
	return interceptor(ctx, in, info, handler)
}

var _ActivityService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "v_task.ActivityService",
	HandlerType: (*ActivityServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserByName",
			Handler:    _ActivityService_GetUserByName_Handler,
		},
		{
			MethodName: "UpdateUserDetails",
			Handler:    _ActivityService_UpdateUserDetails_Handler,
		},
		{
			MethodName: "RegisterUser",
			Handler:    _ActivityService_RegisterUser_Handler,
		},
		{
			MethodName: "AddUserActivity",
			Handler:    _ActivityService_AddUserActivity_Handler,
		},
		{
			MethodName: "StopUserActivity",
			Handler:    _ActivityService_StopUserActivity_Handler,
		},
		{
			MethodName: "ActivityIsDone",
			Handler:    _ActivityService_ActivityIsDone_Handler,
		},
		{
			MethodName: "ActivityIsValid",
			Handler:    _ActivityService_ActivityIsValid_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tracker.proto",
}