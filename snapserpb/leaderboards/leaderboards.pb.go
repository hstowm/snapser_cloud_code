// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.3
// source: leaderboards.proto

package proto

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

type UserScore struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string  `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Rank   int64   `protobuf:"varint,2,opt,name=rank,proto3" json:"rank,omitempty"`
	Score  float64 `protobuf:"fixed64,3,opt,name=score,proto3" json:"score,omitempty"`
}

func (x *UserScore) Reset() {
	*x = UserScore{}
	if protoimpl.UnsafeEnabled {
		mi := &file_leaderboards_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserScore) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserScore) ProtoMessage() {}

func (x *UserScore) ProtoReflect() protoreflect.Message {
	mi := &file_leaderboards_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserScore.ProtoReflect.Descriptor instead.
func (*UserScore) Descriptor() ([]byte, []int) {
	return file_leaderboards_proto_rawDescGZIP(), []int{0}
}

func (x *UserScore) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UserScore) GetRank() int64 {
	if x != nil {
		return x.Rank
	}
	return 0
}

func (x *UserScore) GetScore() float64 {
	if x != nil {
		return x.Score
	}
	return 0
}

type Leaderboard struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Description string  `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Behavior    string  `protobuf:"bytes,3,opt,name=behavior,proto3" json:"behavior,omitempty"`
	Sort        string  `protobuf:"bytes,4,opt,name=sort,proto3" json:"sort,omitempty"`
	Scope       string  `protobuf:"bytes,5,opt,name=scope,proto3" json:"scope,omitempty"`
	StartTime   int64   `protobuf:"varint,6,opt,name=start_time,json=startTime,proto3" json:"start_time,omitempty"`
	Duration    uint64  `protobuf:"varint,7,opt,name=duration,proto3" json:"duration,omitempty"`
	TimeUnit    string  `protobuf:"bytes,8,opt,name=time_unit,json=timeUnit,proto3" json:"time_unit,omitempty"`
	Type        string  `protobuf:"bytes,9,opt,name=type,proto3" json:"type,omitempty"`
	Tiers       []*Tier `protobuf:"bytes,10,rep,name=tiers,proto3" json:"tiers,omitempty"`
}

func (x *Leaderboard) Reset() {
	*x = Leaderboard{}
	if protoimpl.UnsafeEnabled {
		mi := &file_leaderboards_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Leaderboard) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Leaderboard) ProtoMessage() {}

func (x *Leaderboard) ProtoReflect() protoreflect.Message {
	mi := &file_leaderboards_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Leaderboard.ProtoReflect.Descriptor instead.
func (*Leaderboard) Descriptor() ([]byte, []int) {
	return file_leaderboards_proto_rawDescGZIP(), []int{1}
}

func (x *Leaderboard) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Leaderboard) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

func (x *Leaderboard) GetBehavior() string {
	if x != nil {
		return x.Behavior
	}
	return ""
}

func (x *Leaderboard) GetSort() string {
	if x != nil {
		return x.Sort
	}
	return ""
}

func (x *Leaderboard) GetScope() string {
	if x != nil {
		return x.Scope
	}
	return ""
}

func (x *Leaderboard) GetStartTime() int64 {
	if x != nil {
		return x.StartTime
	}
	return 0
}

func (x *Leaderboard) GetDuration() uint64 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *Leaderboard) GetTimeUnit() string {
	if x != nil {
		return x.TimeUnit
	}
	return ""
}

func (x *Leaderboard) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Leaderboard) GetTiers() []*Tier {
	if x != nil {
		return x.Tiers
	}
	return nil
}

type Tier struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	BucketSize  uint64 `protobuf:"varint,2,opt,name=bucket_size,json=bucketSize,proto3" json:"bucket_size,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *Tier) Reset() {
	*x = Tier{}
	if protoimpl.UnsafeEnabled {
		mi := &file_leaderboards_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tier) ProtoMessage() {}

func (x *Tier) ProtoReflect() protoreflect.Message {
	mi := &file_leaderboards_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tier.ProtoReflect.Descriptor instead.
func (*Tier) Descriptor() ([]byte, []int) {
	return file_leaderboards_proto_rawDescGZIP(), []int{2}
}

func (x *Tier) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Tier) GetBucketSize() uint64 {
	if x != nil {
		return x.BucketSize
	}
	return 0
}

func (x *Tier) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type UserData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LeaderboardName string  `protobuf:"bytes,1,opt,name=leaderboard_name,json=leaderboardName,proto3" json:"leaderboard_name,omitempty"`
	Rank            int64   `protobuf:"varint,2,opt,name=rank,proto3" json:"rank,omitempty"`
	Score           float64 `protobuf:"fixed64,3,opt,name=score,proto3" json:"score,omitempty"`
}

func (x *UserData) Reset() {
	*x = UserData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_leaderboards_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserData) ProtoMessage() {}

func (x *UserData) ProtoReflect() protoreflect.Message {
	mi := &file_leaderboards_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserData.ProtoReflect.Descriptor instead.
func (*UserData) Descriptor() ([]byte, []int) {
	return file_leaderboards_proto_rawDescGZIP(), []int{3}
}

func (x *UserData) GetLeaderboardName() string {
	if x != nil {
		return x.LeaderboardName
	}
	return ""
}

func (x *UserData) GetRank() int64 {
	if x != nil {
		return x.Rank
	}
	return 0
}

func (x *UserData) GetScore() float64 {
	if x != nil {
		return x.Score
	}
	return 0
}

type BucketInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BucketId string `protobuf:"bytes,1,opt,name=bucket_id,json=bucketId,proto3" json:"bucket_id,omitempty"`
	Size     int32  `protobuf:"varint,2,opt,name=size,proto3" json:"size,omitempty"`
}

func (x *BucketInfo) Reset() {
	*x = BucketInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_leaderboards_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BucketInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BucketInfo) ProtoMessage() {}

func (x *BucketInfo) ProtoReflect() protoreflect.Message {
	mi := &file_leaderboards_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BucketInfo.ProtoReflect.Descriptor instead.
func (*BucketInfo) Descriptor() ([]byte, []int) {
	return file_leaderboards_proto_rawDescGZIP(), []int{4}
}

func (x *BucketInfo) GetBucketId() string {
	if x != nil {
		return x.BucketId
	}
	return ""
}

func (x *BucketInfo) GetSize() int32 {
	if x != nil {
		return x.Size
	}
	return 0
}

type TierInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TierName      string        `protobuf:"bytes,1,opt,name=tier_name,json=tierName,proto3" json:"tier_name,omitempty"`
	TotalBuckets  int32         `protobuf:"varint,2,opt,name=total_buckets,json=totalBuckets,proto3" json:"total_buckets,omitempty"`
	Buckets       []*BucketInfo `protobuf:"bytes,3,rep,name=buckets,proto3" json:"buckets,omitempty"`
	LookupBuckets []string      `protobuf:"bytes,4,rep,name=lookup_buckets,json=lookupBuckets,proto3" json:"lookup_buckets,omitempty"`
}

func (x *TierInfo) Reset() {
	*x = TierInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_leaderboards_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TierInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TierInfo) ProtoMessage() {}

func (x *TierInfo) ProtoReflect() protoreflect.Message {
	mi := &file_leaderboards_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TierInfo.ProtoReflect.Descriptor instead.
func (*TierInfo) Descriptor() ([]byte, []int) {
	return file_leaderboards_proto_rawDescGZIP(), []int{5}
}

func (x *TierInfo) GetTierName() string {
	if x != nil {
		return x.TierName
	}
	return ""
}

func (x *TierInfo) GetTotalBuckets() int32 {
	if x != nil {
		return x.TotalBuckets
	}
	return 0
}

func (x *TierInfo) GetBuckets() []*BucketInfo {
	if x != nil {
		return x.Buckets
	}
	return nil
}

func (x *TierInfo) GetLookupBuckets() []string {
	if x != nil {
		return x.LookupBuckets
	}
	return nil
}

type TieredLeaderboardInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tiers []*TierInfo `protobuf:"bytes,1,rep,name=tiers,proto3" json:"tiers,omitempty"`
}

func (x *TieredLeaderboardInfo) Reset() {
	*x = TieredLeaderboardInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_leaderboards_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TieredLeaderboardInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TieredLeaderboardInfo) ProtoMessage() {}

func (x *TieredLeaderboardInfo) ProtoReflect() protoreflect.Message {
	mi := &file_leaderboards_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TieredLeaderboardInfo.ProtoReflect.Descriptor instead.
func (*TieredLeaderboardInfo) Descriptor() ([]byte, []int) {
	return file_leaderboards_proto_rawDescGZIP(), []int{6}
}

func (x *TieredLeaderboardInfo) GetTiers() []*TierInfo {
	if x != nil {
		return x.Tiers
	}
	return nil
}

var File_leaderboards_proto protoreflect.FileDescriptor

var file_leaderboards_proto_rawDesc = []byte{
	0x0a, 0x12, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x73, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76,
	0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa9, 0x01, 0x0a, 0x09,
	0x55, 0x73, 0x65, 0x72, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x35, 0x0a, 0x07, 0x75, 0x73, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1c, 0x92, 0x41, 0x19, 0x32,
	0x13, 0x55, 0x73, 0x65, 0x72, 0x20, 0x49, 0x44, 0x20, 0x6f, 0x66, 0x20, 0x74, 0x68, 0x65, 0x20,
	0x75, 0x73, 0x65, 0x72, 0x9a, 0x02, 0x01, 0x07, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x38, 0x0a, 0x04, 0x72, 0x61, 0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x24,
	0x92, 0x41, 0x21, 0x32, 0x1b, 0x55, 0x73, 0x65, 0x72, 0x27, 0x73, 0x20, 0x72, 0x61, 0x6e, 0x6b,
	0x20, 0x73, 0x74, 0x61, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x20, 0x77, 0x69, 0x74, 0x68, 0x20, 0x31,
	0x9a, 0x02, 0x01, 0x03, 0x52, 0x04, 0x72, 0x61, 0x6e, 0x6b, 0x12, 0x2b, 0x0a, 0x05, 0x73, 0x63,
	0x6f, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x42, 0x15, 0x92, 0x41, 0x12, 0x32, 0x0c,
	0x55, 0x73, 0x65, 0x72, 0x27, 0x73, 0x20, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x9a, 0x02, 0x01, 0x05,
	0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x22, 0xe9, 0x03, 0x0a, 0x0b, 0x4c, 0x65, 0x61, 0x64,
	0x65, 0x72, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x12, 0x34, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x20, 0xfa, 0x42, 0x1d, 0x72, 0x1b, 0x32, 0x19, 0x28, 0x3f,
	0x69, 0x29, 0x5e, 0x5b, 0x61, 0x2d, 0x7a, 0x5d, 0x7b, 0x31, 0x7d, 0x5b, 0x61, 0x2d, 0x7a, 0x30,
	0x2d, 0x39, 0x2d, 0x5f, 0x5d, 0x2a, 0x24, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x4f, 0x0a, 0x08, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x33, 0xfa, 0x42, 0x30, 0x72, 0x2e, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65,
	0x72, 0x52, 0x07, 0x6d, 0x61, 0x78, 0x69, 0x6d, 0x75, 0x6d, 0x52, 0x07, 0x6d, 0x69, 0x6e, 0x69,
	0x6d, 0x75, 0x6d, 0x52, 0x11, 0x6d, 0x6f, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x63, 0x65, 0x6e, 0x74,
	0x5f, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x52, 0x08, 0x62, 0x65, 0x68, 0x61, 0x76, 0x69, 0x6f, 0x72,
	0x12, 0x30, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1c,
	0xfa, 0x42, 0x19, 0x72, 0x17, 0x52, 0x0a, 0x64, 0x65, 0x73, 0x63, 0x65, 0x6e, 0x64, 0x69, 0x6e,
	0x67, 0x52, 0x09, 0x61, 0x73, 0x63, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x04, 0x73, 0x6f,
	0x72, 0x74, 0x12, 0x2f, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x70, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x19, 0xfa, 0x42, 0x16, 0x72, 0x14, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x52, 0x08, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x52, 0x05, 0x73, 0x63,
	0x6f, 0x70, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x54, 0x69,
	0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3e,
	0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x75, 0x6e, 0x69, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x21, 0xfa, 0x42, 0x1e, 0x72, 0x1c, 0x52, 0x05, 0x68, 0x6f, 0x75, 0x72, 0x73, 0x52,
	0x04, 0x64, 0x61, 0x79, 0x73, 0x52, 0x05, 0x77, 0x65, 0x65, 0x6b, 0x73, 0x52, 0x06, 0x6d, 0x6f,
	0x6e, 0x74, 0x68, 0x73, 0x52, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x55, 0x6e, 0x69, 0x74, 0x12, 0x29,
	0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x42, 0x15, 0xfa, 0x42,
	0x12, 0x72, 0x10, 0x52, 0x06, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0x52, 0x06, 0x74, 0x69, 0x65,
	0x72, 0x65, 0x64, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x28, 0x0a, 0x05, 0x74, 0x69, 0x65,
	0x72, 0x73, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6c, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x54, 0x69, 0x65, 0x72, 0x52, 0x05, 0x74, 0x69,
	0x65, 0x72, 0x73, 0x22, 0x90, 0x01, 0x0a, 0x04, 0x54, 0x69, 0x65, 0x72, 0x12, 0x3c, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x28, 0xfa, 0x42, 0x25, 0x72,
	0x23, 0x32, 0x21, 0x28, 0x3f, 0x69, 0x29, 0x5e, 0x5b, 0x61, 0x2d, 0x7a, 0x5d, 0x28, 0x5b, 0x61,
	0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x2d, 0x5f, 0x5d, 0x2a, 0x5b, 0x61, 0x2d, 0x7a, 0x30, 0x2d, 0x39,
	0x5d, 0x29, 0x3f, 0x24, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x28, 0x0a, 0x0b, 0x62, 0x75,
	0x63, 0x6b, 0x65, 0x74, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x42,
	0x07, 0xfa, 0x42, 0x04, 0x32, 0x02, 0x20, 0x00, 0x52, 0x0a, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74,
	0x53, 0x69, 0x7a, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72,
	0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0xbe, 0x01, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x44,
	0x61, 0x74, 0x61, 0x12, 0x4b, 0x0a, 0x10, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x20, 0x92,
	0x41, 0x1d, 0x32, 0x17, 0x4e, 0x61, 0x6d, 0x65, 0x20, 0x6f, 0x66, 0x20, 0x74, 0x68, 0x65, 0x20,
	0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x9a, 0x02, 0x01, 0x07, 0x52,
	0x0f, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x38, 0x0a, 0x04, 0x72, 0x61, 0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x24,
	0x92, 0x41, 0x21, 0x32, 0x1b, 0x55, 0x73, 0x65, 0x72, 0x27, 0x73, 0x20, 0x72, 0x61, 0x6e, 0x6b,
	0x20, 0x73, 0x74, 0x61, 0x72, 0x74, 0x69, 0x6e, 0x67, 0x20, 0x77, 0x69, 0x74, 0x68, 0x20, 0x31,
	0x9a, 0x02, 0x01, 0x03, 0x52, 0x04, 0x72, 0x61, 0x6e, 0x6b, 0x12, 0x2b, 0x0a, 0x05, 0x73, 0x63,
	0x6f, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x42, 0x15, 0x92, 0x41, 0x12, 0x32, 0x0c,
	0x55, 0x73, 0x65, 0x72, 0x27, 0x73, 0x20, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x9a, 0x02, 0x01, 0x05,
	0x52, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x22, 0x67, 0x0a, 0x0a, 0x42, 0x75, 0x63, 0x6b, 0x65,
	0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2f, 0x0a, 0x09, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x12, 0x92, 0x41, 0x0f, 0x32, 0x09, 0x42,
	0x75, 0x63, 0x6b, 0x65, 0x74, 0x20, 0x49, 0x44, 0x9a, 0x02, 0x01, 0x07, 0x52, 0x08, 0x62, 0x75,
	0x63, 0x6b, 0x65, 0x74, 0x49, 0x64, 0x12, 0x28, 0x0a, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x42, 0x14, 0x92, 0x41, 0x11, 0x32, 0x0b, 0x42, 0x75, 0x63, 0x6b, 0x65,
	0x74, 0x20, 0x73, 0x69, 0x7a, 0x65, 0x9a, 0x02, 0x01, 0x03, 0x52, 0x04, 0x73, 0x69, 0x7a, 0x65,
	0x22, 0x98, 0x02, 0x0a, 0x08, 0x54, 0x69, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x2f, 0x0a,
	0x09, 0x74, 0x69, 0x65, 0x72, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x12, 0x92, 0x41, 0x0f, 0x32, 0x09, 0x54, 0x69, 0x65, 0x72, 0x20, 0x6e, 0x61, 0x6d, 0x65,
	0x9a, 0x02, 0x01, 0x07, 0x52, 0x08, 0x74, 0x69, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x45,
	0x0a, 0x0d, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x42, 0x20, 0x92, 0x41, 0x1d, 0x32, 0x17, 0x54, 0x6f, 0x74, 0x61,
	0x6c, 0x20, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x20, 0x6f, 0x66, 0x20, 0x62, 0x75, 0x63, 0x6b,
	0x65, 0x74, 0x73, 0x9a, 0x02, 0x01, 0x03, 0x52, 0x0c, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x42, 0x75,
	0x63, 0x6b, 0x65, 0x74, 0x73, 0x12, 0x4c, 0x0a, 0x07, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x73, 0x2e, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f,
	0x42, 0x18, 0x92, 0x41, 0x15, 0x32, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x20, 0x6f, 0x66, 0x20, 0x62,
	0x75, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x9a, 0x02, 0x01, 0x01, 0x52, 0x07, 0x62, 0x75, 0x63, 0x6b,
	0x65, 0x74, 0x73, 0x12, 0x46, 0x0a, 0x0e, 0x6c, 0x6f, 0x6f, 0x6b, 0x75, 0x70, 0x5f, 0x62, 0x75,
	0x63, 0x6b, 0x65, 0x74, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x42, 0x1f, 0x92, 0x41, 0x1c,
	0x32, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x20, 0x6f, 0x66, 0x20, 0x6c, 0x6f, 0x6f, 0x6b, 0x75, 0x70,
	0x20, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x9a, 0x02, 0x01, 0x01, 0x52, 0x0d, 0x6c, 0x6f,
	0x6f, 0x6b, 0x75, 0x70, 0x42, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x73, 0x22, 0x5d, 0x0a, 0x15, 0x54,
	0x69, 0x65, 0x72, 0x65, 0x64, 0x4c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x44, 0x0a, 0x05, 0x74, 0x69, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x73, 0x2e, 0x54, 0x69, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x42, 0x16, 0x92, 0x41, 0x13,
	0x32, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x20, 0x6f, 0x66, 0x20, 0x74, 0x69, 0x65, 0x72, 0x73, 0x9a,
	0x02, 0x01, 0x01, 0x52, 0x05, 0x74, 0x69, 0x65, 0x72, 0x73, 0x42, 0x37, 0x5a, 0x35, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x65, 0x72,
	0x2d, 0x65, 0x6e, 0x67, 0x69, 0x6e, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2f, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x73, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_leaderboards_proto_rawDescOnce sync.Once
	file_leaderboards_proto_rawDescData = file_leaderboards_proto_rawDesc
)

func file_leaderboards_proto_rawDescGZIP() []byte {
	file_leaderboards_proto_rawDescOnce.Do(func() {
		file_leaderboards_proto_rawDescData = protoimpl.X.CompressGZIP(file_leaderboards_proto_rawDescData)
	})
	return file_leaderboards_proto_rawDescData
}

var file_leaderboards_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_leaderboards_proto_goTypes = []interface{}{
	(*UserScore)(nil),             // 0: leaderboards.UserScore
	(*Leaderboard)(nil),           // 1: leaderboards.Leaderboard
	(*Tier)(nil),                  // 2: leaderboards.Tier
	(*UserData)(nil),              // 3: leaderboards.UserData
	(*BucketInfo)(nil),            // 4: leaderboards.BucketInfo
	(*TierInfo)(nil),              // 5: leaderboards.TierInfo
	(*TieredLeaderboardInfo)(nil), // 6: leaderboards.TieredLeaderboardInfo
}
var file_leaderboards_proto_depIdxs = []int32{
	2, // 0: leaderboards.Leaderboard.tiers:type_name -> leaderboards.Tier
	4, // 1: leaderboards.TierInfo.buckets:type_name -> leaderboards.BucketInfo
	5, // 2: leaderboards.TieredLeaderboardInfo.tiers:type_name -> leaderboards.TierInfo
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_leaderboards_proto_init() }
func file_leaderboards_proto_init() {
	if File_leaderboards_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_leaderboards_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserScore); i {
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
		file_leaderboards_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Leaderboard); i {
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
		file_leaderboards_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tier); i {
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
		file_leaderboards_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserData); i {
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
		file_leaderboards_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BucketInfo); i {
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
		file_leaderboards_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TierInfo); i {
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
		file_leaderboards_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TieredLeaderboardInfo); i {
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
			RawDescriptor: file_leaderboards_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_leaderboards_proto_goTypes,
		DependencyIndexes: file_leaderboards_proto_depIdxs,
		MessageInfos:      file_leaderboards_proto_msgTypes,
	}.Build()
	File_leaderboards_proto = out.File
	file_leaderboards_proto_rawDesc = nil
	file_leaderboards_proto_goTypes = nil
	file_leaderboards_proto_depIdxs = nil
}