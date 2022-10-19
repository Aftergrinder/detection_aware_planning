
// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.2
// source: sgtm.proto

package sgtmpb

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/alta/protopatch/patch/gopb"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Visibility int32

const (
	Visibility_UnknownVisibility Visibility = 0
	Visibility_Public            Visibility = 1
	Visibility_Draft             Visibility = 2
)

// Enum value maps for Visibility.
var (
	Visibility_name = map[int32]string{
		0: "UnknownVisibility",
		1: "Public",
		2: "Draft",
	}
	Visibility_value = map[string]int32{
		"UnknownVisibility": 0,
		"Public":            1,
		"Draft":             2,
	}
)

func (x Visibility) Enum() *Visibility {
	p := new(Visibility)
	*p = x
	return p
}

func (x Visibility) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Visibility) Descriptor() protoreflect.EnumDescriptor {
	return file_sgtm_proto_enumTypes[0].Descriptor()
}

func (Visibility) Type() protoreflect.EnumType {
	return &file_sgtm_proto_enumTypes[0]
}

func (x Visibility) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Visibility.Descriptor instead.
func (Visibility) EnumDescriptor() ([]byte, []int) {
	return file_sgtm_proto_rawDescGZIP(), []int{0}
}

type Provider int32

const (
	Provider_UnknownProvider Provider = 0
	Provider_SoundCloud      Provider = 1
	Provider_IPFS            Provider = 2
)

// Enum value maps for Provider.
var (
	Provider_name = map[int32]string{
		0: "UnknownProvider",
		1: "SoundCloud",
		2: "IPFS",
	}
	Provider_value = map[string]int32{
		"UnknownProvider": 0,
		"SoundCloud":      1,
		"IPFS":            2,
	}
)

func (x Provider) Enum() *Provider {
	p := new(Provider)
	*p = x
	return p
}

func (x Provider) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Provider) Descriptor() protoreflect.EnumDescriptor {
	return file_sgtm_proto_enumTypes[1].Descriptor()
}

func (Provider) Type() protoreflect.EnumType {
	return &file_sgtm_proto_enumTypes[1]
}

func (x Provider) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Provider.Descriptor instead.
func (Provider) EnumDescriptor() ([]byte, []int) {
	return file_sgtm_proto_rawDescGZIP(), []int{1}
}

type Post_SoundCloudKind int32

const (
	Post_UnknownSoundCloudKind Post_SoundCloudKind = 0
	Post_SoundCloudTrack       Post_SoundCloudKind = 1 //SoundCloudPlaylist = 2;
)

// Enum value maps for Post_SoundCloudKind.
var (
	Post_SoundCloudKind_name = map[int32]string{
		0: "UnknownSoundCloudKind",
		1: "SoundCloudTrack",
	}
	Post_SoundCloudKind_value = map[string]int32{
		"UnknownSoundCloudKind": 0,
		"SoundCloudTrack":       1,
	}
)

func (x Post_SoundCloudKind) Enum() *Post_SoundCloudKind {
	p := new(Post_SoundCloudKind)
	*p = x
	return p
}

func (x Post_SoundCloudKind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Post_SoundCloudKind) Descriptor() protoreflect.EnumDescriptor {
	return file_sgtm_proto_enumTypes[2].Descriptor()
}

func (Post_SoundCloudKind) Type() protoreflect.EnumType {
	return &file_sgtm_proto_enumTypes[2]
}

func (x Post_SoundCloudKind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Post_SoundCloudKind.Descriptor instead.
func (Post_SoundCloudKind) EnumDescriptor() ([]byte, []int) {
	return file_sgtm_proto_rawDescGZIP(), []int{8, 0}
}

type Post_Kind int32

const (
	Post_UnknownKind            Post_Kind = 0
	Post_TrackKind              Post_Kind = 1
	Post_LoginKind              Post_Kind = 2
	Post_RegisterKind           Post_Kind = 3
	Post_LinkDiscordAccountKind Post_Kind = 4
	Post_ViewProfileKind        Post_Kind = 5
	Post_ViewPostKind           Post_Kind = 6
	Post_ViewOpenKind           Post_Kind = 7
	Post_ViewHomeKind           Post_Kind = 8
	Post_CommentKind            Post_Kind = 9
)

// Enum value maps for Post_Kind.
var (
	Post_Kind_name = map[int32]string{
		0: "UnknownKind",
		1: "TrackKind",
		2: "LoginKind",
		3: "RegisterKind",
		4: "LinkDiscordAccountKind",
		5: "ViewProfileKind",
		6: "ViewPostKind",
		7: "ViewOpenKind",
		8: "ViewHomeKind",
		9: "CommentKind",
	}
	Post_Kind_value = map[string]int32{
		"UnknownKind":            0,
		"TrackKind":              1,
		"LoginKind":              2,
		"RegisterKind":           3,
		"LinkDiscordAccountKind": 4,
		"ViewProfileKind":        5,
		"ViewPostKind":           6,
		"ViewOpenKind":           7,
		"ViewHomeKind":           8,
		"CommentKind":            9,
	}
)

func (x Post_Kind) Enum() *Post_Kind {
	p := new(Post_Kind)
	*p = x
	return p
}

func (x Post_Kind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Post_Kind) Descriptor() protoreflect.EnumDescriptor {
	return file_sgtm_proto_enumTypes[3].Descriptor()
}

func (Post_Kind) Type() protoreflect.EnumType {
	return &file_sgtm_proto_enumTypes[3]
}

func (x Post_Kind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Post_Kind.Descriptor instead.
func (Post_Kind) EnumDescriptor() ([]byte, []int) {
	return file_sgtm_proto_rawDescGZIP(), []int{8, 1}
}

type Relationship_Kind int32

const (
	Relationship_UnknownKind           Relationship_Kind = 0
	Relationship_FeaturingUserKind     Relationship_Kind = 1
	Relationship_RemixOfTrackKind      Relationship_Kind = 2
	Relationship_NewVersionOfTrackKind Relationship_Kind = 3
	Relationship_InspiredByTrackKind   Relationship_Kind = 4
	Relationship_RemixOfUserKind       Relationship_Kind = 5
)

// Enum value maps for Relationship_Kind.
var (
	Relationship_Kind_name = map[int32]string{
		0: "UnknownKind",
		1: "FeaturingUserKind",
		2: "RemixOfTrackKind",
		3: "NewVersionOfTrackKind",
		4: "InspiredByTrackKind",
		5: "RemixOfUserKind",
	}
	Relationship_Kind_value = map[string]int32{
		"UnknownKind":           0,
		"FeaturingUserKind":     1,
		"RemixOfTrackKind":      2,
		"NewVersionOfTrackKind": 3,
		"InspiredByTrackKind":   4,
		"RemixOfUserKind":       5,
	}
)

func (x Relationship_Kind) Enum() *Relationship_Kind {
	p := new(Relationship_Kind)
	*p = x
	return p
}

func (x Relationship_Kind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Relationship_Kind) Descriptor() protoreflect.EnumDescriptor {
	return file_sgtm_proto_enumTypes[4].Descriptor()
}

func (Relationship_Kind) Type() protoreflect.EnumType {
	return &file_sgtm_proto_enumTypes[4]
}

func (x Relationship_Kind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Relationship_Kind.Descriptor instead.
func (Relationship_Kind) EnumDescriptor() ([]byte, []int) {
	return file_sgtm_proto_rawDescGZIP(), []int{9, 0}
}

type Ping struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Ping) Reset() {
	*x = Ping{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sgtm_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ping) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ping) ProtoMessage() {}

func (x *Ping) ProtoReflect() protoreflect.Message {
	mi := &file_sgtm_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ping.ProtoReflect.Descriptor instead.
func (*Ping) Descriptor() ([]byte, []int) {
	return file_sgtm_proto_rawDescGZIP(), []int{0}
}

type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sgtm_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_sgtm_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_sgtm_proto_rawDescGZIP(), []int{1}
}

type Register struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Register) Reset() {
	*x = Register{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sgtm_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Register) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Register) ProtoMessage() {}

func (x *Register) ProtoReflect() protoreflect.Message {
	mi := &file_sgtm_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Register.ProtoReflect.Descriptor instead.
func (*Register) Descriptor() ([]byte, []int) {
	return file_sgtm_proto_rawDescGZIP(), []int{2}
}

type UserList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UserList) Reset() {
	*x = UserList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sgtm_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserList) ProtoMessage() {}

func (x *UserList) ProtoReflect() protoreflect.Message {
	mi := &file_sgtm_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserList.ProtoReflect.Descriptor instead.
func (*UserList) Descriptor() ([]byte, []int) {
	return file_sgtm_proto_rawDescGZIP(), []int{3}
}

type PostList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PostList) Reset() {
	*x = PostList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sgtm_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostList) ProtoMessage() {}

func (x *PostList) ProtoReflect() protoreflect.Message {
	mi := &file_sgtm_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostList.ProtoReflect.Descriptor instead.
func (*PostList) Descriptor() ([]byte, []int) {
	return file_sgtm_proto_rawDescGZIP(), []int{4}
}

type PostSync struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PostSync) Reset() {
	*x = PostSync{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sgtm_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PostSync) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PostSync) ProtoMessage() {}

func (x *PostSync) ProtoReflect() protoreflect.Message {
	mi := &file_sgtm_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PostSync.ProtoReflect.Descriptor instead.
func (*PostSync) Descriptor() ([]byte, []int) {
	return file_sgtm_proto_rawDescGZIP(), []int{5}
}

type Me struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Me) Reset() {
	*x = Me{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sgtm_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Me) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Me) ProtoMessage() {}

func (x *Me) ProtoReflect() protoreflect.Message {
	mi := &file_sgtm_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Me.ProtoReflect.Descriptor instead.
func (*Me) Descriptor() ([]byte, []int) {
	return file_sgtm_proto_rawDescGZIP(), []int{6}
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID                    int64           `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" gorm:"primary_key"`
	CreatedAt             int64           `protobuf:"varint,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty" gorm:"autocreatetime:nano"`
	UpdatedAt             int64           `protobuf:"varint,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty" gorm:"autoupdatetime:nano"`
	DeletedAt             int64           `protobuf:"varint,4,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
	Email                 string          `protobuf:"bytes,10,opt,name=email,proto3" json:"email,omitempty" gorm:"size:255;not null;index:,unique"`
	Slug                  string          `protobuf:"bytes,11,opt,name=slug,proto3" json:"slug,omitempty" gorm:"size:32;not null;default:''"`
	Firstname             string          `protobuf:"bytes,12,opt,name=firstname,proto3" json:"firstname,omitempty" gorm:"size:255;not null;default:''"`
	Lastname              string          `protobuf:"bytes,13,opt,name=lastname,proto3" json:"lastname,omitempty" gorm:"size:255;not null;default:''"`
	Locale                string          `protobuf:"bytes,14,opt,name=locale,proto3" json:"locale,omitempty" gorm:"size:32;not null;default:''"`
	Avatar                string          `protobuf:"bytes,15,opt,name=avatar,proto3" json:"avatar,omitempty" gorm:"size:255;not null;default:''"`
	DiscordID             string          `protobuf:"bytes,16,opt,name=discord_id,json=discordId,proto3" json:"discord_id,omitempty" gorm:"size:255;not null;default:''"`
	DiscordUsername       string          `protobuf:"bytes,17,opt,name=discord_username,json=discordUsername,proto3" json:"discord_username,omitempty" gorm:"size:255;not null;default:''"`
	Headline              string          `protobuf:"bytes,18,opt,name=headline,proto3" json:"headline,omitempty"`
	Bio                   string          `protobuf:"bytes,19,opt,name=bio,proto3" json:"bio,omitempty"`
	Inspirations          string          `protobuf:"bytes,20,opt,name=inspirations,proto3" json:"inspirations,omitempty"`
	Gears                 string          `protobuf:"bytes,21,opt,name=gears,proto3" json:"gears,omitempty"`
	Genres                string          `protobuf:"bytes,22,opt,name=genres,proto3" json:"genres,omitempty"`
	Location              string          `protobuf:"bytes,23,opt,name=location,proto3" json:"location,omitempty"`
	TwitterUsername       string          `protobuf:"bytes,24,opt,name=twitter_username,json=twitterUsername,proto3" json:"twitter_username,omitempty"`
	Homepage              string          `protobuf:"bytes,25,opt,name=homepage,proto3" json:"homepage,omitempty"`
	OtherLinks            string          `protobuf:"bytes,26,opt,name=other_links,json=otherLinks,proto3" json:"other_links,omitempty"`
	Goals                 string          `protobuf:"bytes,27,opt,name=goals,proto3" json:"goals,omitempty"`
	SoundcloudUsername    string          `protobuf:"bytes,28,opt,name=soundcloud_username,json=soundcloudUsername,proto3" json:"soundcloud_username,omitempty"`
	Role                  string          `protobuf:"bytes,29,opt,name=role,proto3" json:"role,omitempty"`
	ProcessingVersion     int64           `protobuf:"varint,30,opt,name=processing_version,json=processingVersion,proto3" json:"processing_version,omitempty"`
	ProcessingError       string          `protobuf:"bytes,31,opt,name=processing_error,json=processingError,proto3" json:"processing_error,omitempty"`
	RecentPosts           []*Post         `protobuf:"bytes,50,rep,name=recent_posts,json=recentPosts,proto3" json:"recent_posts,omitempty" gorm:"foreignkey:AuthorID;PRELOAD:false"`
	RelationshipsAsSource []*Relationship `protobuf:"bytes,51,rep,name=relationships_as_source,json=relationshipsAsSource,proto3" json:"relationships_as_source,omitempty" gorm:"foreignKey:SourceUserID"`
	RelationshipsAsTarget []*Relationship `protobuf:"bytes,52,rep,name=relationships_as_target,json=relationshipsAsTarget,proto3" json:"relationships_as_target,omitempty" gorm:"foreignKey:TargetUserID"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sgtm_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_sgtm_proto_msgTypes[7]
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
	return file_sgtm_proto_rawDescGZIP(), []int{7}
}

func (x *User) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *User) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *User) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *User) GetDeletedAt() int64 {
	if x != nil {
		return x.DeletedAt
	}
	return 0
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetSlug() string {
	if x != nil {
		return x.Slug
	}
	return ""
}

func (x *User) GetFirstname() string {
	if x != nil {
		return x.Firstname
	}
	return ""
}

func (x *User) GetLastname() string {
	if x != nil {
		return x.Lastname
	}
	return ""
}

func (x *User) GetLocale() string {
	if x != nil {
		return x.Locale
	}
	return ""
}

func (x *User) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *User) GetDiscordID() string {
	if x != nil {
		return x.DiscordID
	}
	return ""
}

func (x *User) GetDiscordUsername() string {
	if x != nil {
		return x.DiscordUsername
	}
	return ""
}

func (x *User) GetHeadline() string {
	if x != nil {
		return x.Headline
	}
	return ""
}

func (x *User) GetBio() string {
	if x != nil {
		return x.Bio
	}
	return ""
}

func (x *User) GetInspirations() string {
	if x != nil {
		return x.Inspirations
	}
	return ""
}

func (x *User) GetGears() string {
	if x != nil {
		return x.Gears
	}
	return ""
}

func (x *User) GetGenres() string {
	if x != nil {
		return x.Genres
	}
	return ""
}

func (x *User) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *User) GetTwitterUsername() string {
	if x != nil {
		return x.TwitterUsername
	}
	return ""
}

func (x *User) GetHomepage() string {
	if x != nil {
		return x.Homepage
	}
	return ""
}

func (x *User) GetOtherLinks() string {
	if x != nil {
		return x.OtherLinks
	}
	return ""
}

func (x *User) GetGoals() string {
	if x != nil {
		return x.Goals
	}
	return ""
}

func (x *User) GetSoundcloudUsername() string {
	if x != nil {
		return x.SoundcloudUsername
	}
	return ""
}

func (x *User) GetRole() string {
	if x != nil {
		return x.Role
	}
	return ""
}

func (x *User) GetProcessingVersion() int64 {
	if x != nil {
		return x.ProcessingVersion
	}
	return 0
}

func (x *User) GetProcessingError() string {
	if x != nil {
		return x.ProcessingError
	}
	return ""
}

func (x *User) GetRecentPosts() []*Post {
	if x != nil {
		return x.RecentPosts
	}
	return nil
}

func (x *User) GetRelationshipsAsSource() []*Relationship {
	if x != nil {
		return x.RelationshipsAsSource
	}
	return nil
}

func (x *User) GetRelationshipsAsTarget() []*Relationship {
	if x != nil {
		return x.RelationshipsAsTarget
	}
	return nil
}

type Post struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID                int64      `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty" gorm:"primary_key"`
	CreatedAt         int64      `protobuf:"varint,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty" gorm:"autocreatetime:nano"`
	UpdatedAt         int64      `protobuf:"varint,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty" gorm:"autoupdatetime:nano"`
	DeletedAt         int64      `protobuf:"varint,4,opt,name=deleted_at,json=deletedAt,proto3" json:"deleted_at,omitempty"`
	Author            *User      `protobuf:"bytes,10,opt,name=author,proto3" json:"author,omitempty"`
	AuthorID          int64      `protobuf:"varint,11,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
	Title             string     `protobuf:"bytes,12,opt,name=title,proto3" json:"title,omitempty"`
	Slug              string     `protobuf:"bytes,13,opt,name=slug,proto3" json:"slug,omitempty"`
	Kind              Post_Kind  `protobuf:"varint,14,opt,name=kind,proto3,enum=sgtm.Post_Kind" json:"kind,omitempty"`
	Visibility        Visibility `protobuf:"varint,15,opt,name=visibility,proto3,enum=sgtm.Visibility" json:"visibility,omitempty"`
	URL               string     `protobuf:"bytes,16,opt,name=url,proto3" json:"url,omitempty"`
	Provider          Provider   `protobuf:"varint,17,opt,name=provider,proto3,enum=sgtm.Provider" json:"provider,omitempty"`
	Body              string     `protobuf:"bytes,18,opt,name=body,proto3" json:"body,omitempty"`
	SortDate          int64      `protobuf:"varint,19,opt,name=sort_date,json=sortDate,proto3" json:"sort_date,omitempty"`
	ProcessingVersion int64      `protobuf:"varint,20,opt,name=processing_version,json=processingVersion,proto3" json:"processing_version,omitempty"`
	ProcessingError   string     `protobuf:"bytes,21,opt,name=processing_error,json=processingError,proto3" json:"processing_error,omitempty"`
	// Deprecated: Do not use.
	Genre                 string              `protobuf:"bytes,40,opt,name=genre,proto3" json:"genre,omitempty"` // replaced by 'tags'
	Duration              uint64              `protobuf:"varint,41,opt,name=duration,proto3" json:"duration,omitempty"`
	ArtworkURL            string              `protobuf:"bytes,42,opt,name=artwork_url,json=artworkUrl,proto3" json:"artwork_url,omitempty"`
	BPM                   float64             `protobuf:"fixed64,43,opt,name=bpm,proto3" json:"bpm,omitempty"`
	KeySignature          string              `protobuf:"bytes,44,opt,name=key_signature,json=keySignature,proto3" json:"key_signature,omitempty"`
	ISRC                  string              `protobuf:"bytes,45,opt,name=isrc,proto3" json:"isrc,omitempty"`
	ProviderTitle         string              `protobuf:"bytes,53,opt,name=provider_title,json=providerTitle,proto3" json:"provider_title,omitempty"`
	ProviderDescription   string              `protobuf:"bytes,46,opt,name=provider_description,json=providerDescription,proto3" json:"provider_description,omitempty"`
	DownloadURL           string              `protobuf:"bytes,47,opt,name=provider_download_url,json=providerDownloadUrl,proto3" json:"provider_download_url,omitempty"`
	ProviderCreatedAt     int64               `protobuf:"varint,48,opt,name=provider_created_at,json=providerCreatedAt,proto3" json:"provider_created_at,omitempty"`
	ProviderUpdatedAt     int64               `protobuf:"varint,49,opt,name=provider_updated_at,json=providerUpdatedAt,proto3" json:"provider_updated_at,omitempty"`
	ProviderMetadata      string              `protobuf:"bytes,50,opt,name=provider_metadata,json=providerMetadata,proto3" json:"provider_metadata,omitempty"`
	Tags                  string              `protobuf:"bytes,51,opt,name=tags,proto3" json:"tags,omitempty"` // comma separated list of tags
	Lyrics                string              `protobuf:"bytes,52,opt,name=lyrics,proto3" json:"lyrics,omitempty"`
	SoundCloudSecretToken string              `protobuf:"bytes,80,opt,name=soundcloud_secret_token,json=soundcloudSecretToken,proto3" json:"soundcloud_secret_token,omitempty"`
	SoundCloudID          uint64              `protobuf:"varint,81,opt,name=soundcloud_id,json=soundcloudId,proto3" json:"soundcloud_id,omitempty"`
	SoundCloudKind        Post_SoundCloudKind `protobuf:"varint,83,opt,name=soundcloud_kind,json=soundcloudKind,proto3,enum=sgtm.Post_SoundCloudKind" json:"soundcloud_kind,omitempty"`
	IPFSCID               string              `protobuf:"bytes,90,opt,name=ipfs_cid,json=ipfsCid,proto3" json:"ipfs_cid,omitempty"`
	MIMEType              string              `protobuf:"bytes,91,opt,name=mime_type,json=mimeType,proto3" json:"mime_type,omitempty"`
	SizeBytes             int64               `protobuf:"varint,92,opt,name=size_bytes,json=sizeBytes,proto3" json:"size_bytes,omitempty"`
	FileExtension         string              `protobuf:"bytes,93,opt,name=file_extension,json=fileExtension,proto3" json:"file_extension,omitempty"`
	AttachmentFilename    string              `protobuf:"bytes,94,opt,name=attachment_filename,json=attachmentFilename,proto3" json:"attachment_filename,omitempty"`
	TargetUserID          int64               `protobuf:"varint,101,opt,name=target_user_id,json=targetUserId,proto3" json:"target_user_id,omitempty"`
	TargetUser            *User               `protobuf:"bytes,102,opt,name=target_user,json=targetUser,proto3" json:"target_user,omitempty"`
	TargetPostID          int64               `protobuf:"varint,103,opt,name=target_post_id,json=targetPostId,proto3" json:"target_post_id,omitempty"`
	TargetPost            *Post               `protobuf:"bytes,104,opt,name=target_post,json=targetPost,proto3" json:"target_post,omitempty"`
	TargetMetadata        string              `protobuf:"bytes,105,opt,name=target_metadata,json=targetMetadata,proto3" json:"target_metadata,omitempty"`
	RelationshipsAsSource []*Relationship     `protobuf:"bytes,110,rep,name=relationships_as_source,json=relationshipsAsSource,proto3" json:"relationships_as_source,omitempty" gorm:"foreignKey:SourcePostID"`
	RelationshipsAsTarget []*Relationship     `protobuf:"bytes,111,rep,name=relationships_as_target,json=relationshipsAsTarget,proto3" json:"relationships_as_target,omitempty" gorm:"foreignKey:TargetPostID"`
}

func (x *Post) Reset() {
	*x = Post{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sgtm_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Post) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Post) ProtoMessage() {}

func (x *Post) ProtoReflect() protoreflect.Message {
	mi := &file_sgtm_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Post.ProtoReflect.Descriptor instead.
func (*Post) Descriptor() ([]byte, []int) {
	return file_sgtm_proto_rawDescGZIP(), []int{8}
}

func (x *Post) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Post) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Post) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *Post) GetDeletedAt() int64 {
	if x != nil {
		return x.DeletedAt
	}
	return 0
}

func (x *Post) GetAuthor() *User {
	if x != nil {
		return x.Author
	}
	return nil
}

func (x *Post) GetAuthorID() int64 {
	if x != nil {
		return x.AuthorID
	}
	return 0