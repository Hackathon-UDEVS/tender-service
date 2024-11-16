// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.12.4
// source: protos/clients.proto

package tender_service

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

type CreateTenderReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title       string  `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string  `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Deadline    string  `protobuf:"bytes,3,opt,name=deadline,proto3" json:"deadline,omitempty"`
	Budget      float64 `protobuf:"fixed64,4,opt,name=budget,proto3" json:"budget,omitempty"`
	ClientId    string  `protobuf:"bytes,5,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	Attachment  string  `protobuf:"bytes,6,opt,name=attachment,proto3" json:"attachment,omitempty"` // Optional file URL
}

func (x *CreateTenderReq) Reset() {
	*x = CreateTenderReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_clients_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTenderReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTenderReq) ProtoMessage() {}

func (x *CreateTenderReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_clients_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTenderReq.ProtoReflect.Descriptor instead.
func (*CreateTenderReq) Descriptor() ([]byte, []int) {
	return file_protos_clients_proto_rawDescGZIP(), []int{0}
}

func (x *CreateTenderReq) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreateTenderReq) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *CreateTenderReq) GetDeadline() string {
	if x != nil {
		return x.Deadline
	}
	return ""
}

func (x *CreateTenderReq) GetBudget() float64 {
	if x != nil {
		return x.Budget
	}
	return 0
}

func (x *CreateTenderReq) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *CreateTenderReq) GetAttachment() string {
	if x != nil {
		return x.Attachment
	}
	return ""
}

type ResponseMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *ResponseMessage) Reset() {
	*x = ResponseMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_clients_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResponseMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResponseMessage) ProtoMessage() {}

func (x *ResponseMessage) ProtoReflect() protoreflect.Message {
	mi := &file_protos_clients_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResponseMessage.ProtoReflect.Descriptor instead.
func (*ResponseMessage) Descriptor() ([]byte, []int) {
	return file_protos_clients_proto_rawDescGZIP(), []int{1}
}

func (x *ResponseMessage) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type GetMyTendersReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId string `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	Status   string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *GetMyTendersReq) Reset() {
	*x = GetMyTendersReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_clients_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetMyTendersReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetMyTendersReq) ProtoMessage() {}

func (x *GetMyTendersReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_clients_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetMyTendersReq.ProtoReflect.Descriptor instead.
func (*GetMyTendersReq) Descriptor() ([]byte, []int) {
	return file_protos_clients_proto_rawDescGZIP(), []int{2}
}

func (x *GetMyTendersReq) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *GetMyTendersReq) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type TendersList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tenders []*Tender `protobuf:"bytes,1,rep,name=tenders,proto3" json:"tenders,omitempty"`
}

func (x *TendersList) Reset() {
	*x = TendersList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_clients_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TendersList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TendersList) ProtoMessage() {}

func (x *TendersList) ProtoReflect() protoreflect.Message {
	mi := &file_protos_clients_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TendersList.ProtoReflect.Descriptor instead.
func (*TendersList) Descriptor() ([]byte, []int) {
	return file_protos_clients_proto_rawDescGZIP(), []int{3}
}

func (x *TendersList) GetTenders() []*Tender {
	if x != nil {
		return x.Tenders
	}
	return nil
}

type GetAllTendersReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *GetAllTendersReq) Reset() {
	*x = GetAllTendersReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_clients_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllTendersReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllTendersReq) ProtoMessage() {}

func (x *GetAllTendersReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_clients_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllTendersReq.ProtoReflect.Descriptor instead.
func (*GetAllTendersReq) Descriptor() ([]byte, []int) {
	return file_protos_clients_proto_rawDescGZIP(), []int{4}
}

func (x *GetAllTendersReq) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type UpdateTenderReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TenderId    string  `protobuf:"bytes,1,opt,name=tender_id,json=tenderId,proto3" json:"tender_id,omitempty"`
	Title       string  `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`             // Optional field to update title
	Description string  `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"` // Optional field to update description
	Deadline    string  `protobuf:"bytes,4,opt,name=deadline,proto3" json:"deadline,omitempty"`       // Optional field to update deadline
	Status      string  `protobuf:"bytes,5,opt,name=status,proto3" json:"status,omitempty"`           // Optional field to update status, e.g., "open", "closed", "awarded"
	Budget      float64 `protobuf:"fixed64,6,opt,name=budget,proto3" json:"budget,omitempty"`         // Optional field to update budget
	Attachment  string  `protobuf:"bytes,7,opt,name=attachment,proto3" json:"attachment,omitempty"`   // Optional field to update attachment URL
}

func (x *UpdateTenderReq) Reset() {
	*x = UpdateTenderReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_clients_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateTenderReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateTenderReq) ProtoMessage() {}

func (x *UpdateTenderReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_clients_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateTenderReq.ProtoReflect.Descriptor instead.
func (*UpdateTenderReq) Descriptor() ([]byte, []int) {
	return file_protos_clients_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateTenderReq) GetTenderId() string {
	if x != nil {
		return x.TenderId
	}
	return ""
}

func (x *UpdateTenderReq) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UpdateTenderReq) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *UpdateTenderReq) GetDeadline() string {
	if x != nil {
		return x.Deadline
	}
	return ""
}

func (x *UpdateTenderReq) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *UpdateTenderReq) GetBudget() float64 {
	if x != nil {
		return x.Budget
	}
	return 0
}

func (x *UpdateTenderReq) GetAttachment() string {
	if x != nil {
		return x.Attachment
	}
	return ""
}

type DeleteTenderReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TenderId string `protobuf:"bytes,1,opt,name=tender_id,json=tenderId,proto3" json:"tender_id,omitempty"`
	ClientId string `protobuf:"bytes,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
}

func (x *DeleteTenderReq) Reset() {
	*x = DeleteTenderReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_clients_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteTenderReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteTenderReq) ProtoMessage() {}

func (x *DeleteTenderReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_clients_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteTenderReq.ProtoReflect.Descriptor instead.
func (*DeleteTenderReq) Descriptor() ([]byte, []int) {
	return file_protos_clients_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteTenderReq) GetTenderId() string {
	if x != nil {
		return x.TenderId
	}
	return ""
}

func (x *DeleteTenderReq) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

type Tender struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ClientId    string  `protobuf:"bytes,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	Title       string  `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Description string  `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	Deadline    string  `protobuf:"bytes,5,opt,name=deadline,proto3" json:"deadline,omitempty"`
	Budget      float64 `protobuf:"fixed64,6,opt,name=budget,proto3" json:"budget,omitempty"`
	Status      string  `protobuf:"bytes,7,opt,name=status,proto3" json:"status,omitempty"`
	Attachment  string  `protobuf:"bytes,8,opt,name=attachment,proto3" json:"attachment,omitempty"`
}

func (x *Tender) Reset() {
	*x = Tender{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_clients_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tender) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tender) ProtoMessage() {}

func (x *Tender) ProtoReflect() protoreflect.Message {
	mi := &file_protos_clients_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tender.ProtoReflect.Descriptor instead.
func (*Tender) Descriptor() ([]byte, []int) {
	return file_protos_clients_proto_rawDescGZIP(), []int{7}
}

func (x *Tender) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Tender) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *Tender) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Tender) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Tender) GetDeadline() string {
	if x != nil {
		return x.Deadline
	}
	return ""
}

func (x *Tender) GetBudget() float64 {
	if x != nil {
		return x.Budget
	}
	return 0
}

func (x *Tender) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *Tender) GetAttachment() string {
	if x != nil {
		return x.Attachment
	}
	return ""
}

type SelectWinnerReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClientId     string `protobuf:"bytes,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
	TenderId     string `protobuf:"bytes,2,opt,name=tender_id,json=tenderId,proto3" json:"tender_id,omitempty"`
	ContractorId string `protobuf:"bytes,3,opt,name=contractor_id,json=contractorId,proto3" json:"contractor_id,omitempty"`
}

func (x *SelectWinnerReq) Reset() {
	*x = SelectWinnerReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_clients_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SelectWinnerReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SelectWinnerReq) ProtoMessage() {}

func (x *SelectWinnerReq) ProtoReflect() protoreflect.Message {
	mi := &file_protos_clients_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SelectWinnerReq.ProtoReflect.Descriptor instead.
func (*SelectWinnerReq) Descriptor() ([]byte, []int) {
	return file_protos_clients_proto_rawDescGZIP(), []int{8}
}

func (x *SelectWinnerReq) GetClientId() string {
	if x != nil {
		return x.ClientId
	}
	return ""
}

func (x *SelectWinnerReq) GetTenderId() string {
	if x != nil {
		return x.TenderId
	}
	return ""
}

func (x *SelectWinnerReq) GetContractorId() string {
	if x != nil {
		return x.ContractorId
	}
	return ""
}

var File_protos_clients_proto protoreflect.FileDescriptor

var file_protos_clients_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x22, 0xba,
	0x01, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64,
	0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65,
	0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65,
	0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x12, 0x1b,
	0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x61,
	0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x2b, 0x0a, 0x0f, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x46, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4d,
	0x79, 0x54, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x12, 0x1b, 0x0a, 0x09, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x22, 0x37, 0x0a, 0x0b, 0x54, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12,
	0x28, 0x0a, 0x07, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x2e, 0x54, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x52, 0x07, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x73, 0x22, 0x2a, 0x0a, 0x10, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x54, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x12, 0x16, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0xd2, 0x01, 0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x54, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a,
	0x0a, 0x08, 0x64, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x64, 0x65, 0x61, 0x64, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x06, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x74,
	0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x4b, 0x0a, 0x0f, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x54, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x1b, 0x0a,
	0x09, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63,
	0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0xd9, 0x01, 0x0a, 0x06, 0x54, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x65, 0x61, 0x64, 0x6c,
	0x69, 0x6e, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x65, 0x61, 0x64, 0x6c,
	0x69, 0x6e, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x06, 0x62, 0x75, 0x64, 0x67, 0x65, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d, 0x65, 0x6e,
	0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x74, 0x74, 0x61, 0x63, 0x68, 0x6d,
	0x65, 0x6e, 0x74, 0x22, 0x70, 0x0a, 0x0f, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x57, 0x69, 0x6e,
	0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6c, 0x69, 0x65, 0x6e,
	0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0x6f, 0x72, 0x49, 0x64, 0x32, 0x98, 0x03, 0x0a, 0x0d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x40, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x54, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x17, 0x2e, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x1a, 0x17, 0x2e, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x3c, 0x0a, 0x0c, 0x47, 0x65, 0x74,
	0x4d, 0x79, 0x54, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x73, 0x12, 0x17, 0x2e, 0x74, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x79, 0x54, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x73, 0x52,
	0x65, 0x71, 0x1a, 0x13, 0x2e, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x2e, 0x54, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x3e, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x54, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x73, 0x12, 0x18, 0x2e, 0x74, 0x65, 0x6e, 0x64, 0x65,
	0x72, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x54, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x73, 0x52,
	0x65, 0x71, 0x1a, 0x13, 0x2e, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x2e, 0x54, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x40, 0x0a, 0x0c, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x54, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x17, 0x2e, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x1a, 0x17, 0x2e, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x40, 0x0a, 0x0c, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x54, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x17, 0x2e, 0x74, 0x65, 0x6e, 0x64,
	0x65, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x54, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x1a, 0x17, 0x2e, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x43, 0x0a, 0x0f, 0x53,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x57, 0x69, 0x6e, 0x6e, 0x65, 0x72, 0x42, 0x69, 0x64, 0x12, 0x17,
	0x2e, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x57, 0x69,
	0x6e, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x42, 0x22, 0x5a, 0x20, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x65, 0x6e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x2d, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protos_clients_proto_rawDescOnce sync.Once
	file_protos_clients_proto_rawDescData = file_protos_clients_proto_rawDesc
)

func file_protos_clients_proto_rawDescGZIP() []byte {
	file_protos_clients_proto_rawDescOnce.Do(func() {
		file_protos_clients_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_clients_proto_rawDescData)
	})
	return file_protos_clients_proto_rawDescData
}

var file_protos_clients_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_protos_clients_proto_goTypes = []any{
	(*CreateTenderReq)(nil),  // 0: tender.CreateTenderReq
	(*ResponseMessage)(nil),  // 1: tender.ResponseMessage
	(*GetMyTendersReq)(nil),  // 2: tender.GetMyTendersReq
	(*TendersList)(nil),      // 3: tender.TendersList
	(*GetAllTendersReq)(nil), // 4: tender.GetAllTendersReq
	(*UpdateTenderReq)(nil),  // 5: tender.UpdateTenderReq
	(*DeleteTenderReq)(nil),  // 6: tender.DeleteTenderReq
	(*Tender)(nil),           // 7: tender.Tender
	(*SelectWinnerReq)(nil),  // 8: tender.SelectWinnerReq
}
var file_protos_clients_proto_depIdxs = []int32{
	7, // 0: tender.TendersList.tenders:type_name -> tender.Tender
	0, // 1: tender.ClientService.CreateTender:input_type -> tender.CreateTenderReq
	2, // 2: tender.ClientService.GetMyTenders:input_type -> tender.GetMyTendersReq
	4, // 3: tender.ClientService.GetAllTenders:input_type -> tender.GetAllTendersReq
	5, // 4: tender.ClientService.UpdateTender:input_type -> tender.UpdateTenderReq
	6, // 5: tender.ClientService.DeleteTender:input_type -> tender.DeleteTenderReq
	8, // 6: tender.ClientService.SelectWinnerBid:input_type -> tender.SelectWinnerReq
	1, // 7: tender.ClientService.CreateTender:output_type -> tender.ResponseMessage
	3, // 8: tender.ClientService.GetMyTenders:output_type -> tender.TendersList
	3, // 9: tender.ClientService.GetAllTenders:output_type -> tender.TendersList
	1, // 10: tender.ClientService.UpdateTender:output_type -> tender.ResponseMessage
	1, // 11: tender.ClientService.DeleteTender:output_type -> tender.ResponseMessage
	1, // 12: tender.ClientService.SelectWinnerBid:output_type -> tender.ResponseMessage
	7, // [7:13] is the sub-list for method output_type
	1, // [1:7] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protos_clients_proto_init() }
func file_protos_clients_proto_init() {
	if File_protos_clients_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_clients_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*CreateTenderReq); i {
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
		file_protos_clients_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*ResponseMessage); i {
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
		file_protos_clients_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GetMyTendersReq); i {
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
		file_protos_clients_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*TendersList); i {
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
		file_protos_clients_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*GetAllTendersReq); i {
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
		file_protos_clients_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*UpdateTenderReq); i {
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
		file_protos_clients_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*DeleteTenderReq); i {
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
		file_protos_clients_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*Tender); i {
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
		file_protos_clients_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*SelectWinnerReq); i {
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
			RawDescriptor: file_protos_clients_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_clients_proto_goTypes,
		DependencyIndexes: file_protos_clients_proto_depIdxs,
		MessageInfos:      file_protos_clients_proto_msgTypes,
	}.Build()
	File_protos_clients_proto = out.File
	file_protos_clients_proto_rawDesc = nil
	file_protos_clients_proto_goTypes = nil
	file_protos_clients_proto_depIdxs = nil
}
