// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protobuf/digimon.proto

package pbprotocol

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type DigimonIdentity int32

const (
	DigimonIdentity_AGUMON          DigimonIdentity = 0
	DigimonIdentity_GREYMON         DigimonIdentity = 1
	DigimonIdentity_SKULLGREYMON    DigimonIdentity = 2
	DigimonIdentity_WARGREYMON      DigimonIdentity = 3
	DigimonIdentity_GABUMON         DigimonIdentity = 4
	DigimonIdentity_GARURUMON       DigimonIdentity = 5
	DigimonIdentity_WEREGARURUMON   DigimonIdentity = 6
	DigimonIdentity_METALGARURUMON  DigimonIdentity = 7
	DigimonIdentity_BIYOMON         DigimonIdentity = 8
	DigimonIdentity_BIRDRAMON       DigimonIdentity = 9
	DigimonIdentity_GARUDAMON       DigimonIdentity = 10
	DigimonIdentity_TENTOMON        DigimonIdentity = 11
	DigimonIdentity_KABUTERIMON     DigimonIdentity = 12
	DigimonIdentity_MEGAKABUTERIMON DigimonIdentity = 13
	DigimonIdentity_PALMON          DigimonIdentity = 14
	DigimonIdentity_TOGEMON         DigimonIdentity = 15
	DigimonIdentity_LILLYMON        DigimonIdentity = 16
	DigimonIdentity_GOMAMON         DigimonIdentity = 17
	DigimonIdentity_IKKAKUMON       DigimonIdentity = 18
	DigimonIdentity_ZUDOMON         DigimonIdentity = 19
	DigimonIdentity_PATAMON         DigimonIdentity = 20
	DigimonIdentity_ANGEMON         DigimonIdentity = 21
	DigimonIdentity_MAGNAANGEMON    DigimonIdentity = 22
	DigimonIdentity_SALAMON         DigimonIdentity = 23
	DigimonIdentity_GATOMON         DigimonIdentity = 24
	DigimonIdentity_ANGEWOMON       DigimonIdentity = 25
)

var DigimonIdentity_name = map[int32]string{
	0:  "AGUMON",
	1:  "GREYMON",
	2:  "SKULLGREYMON",
	3:  "WARGREYMON",
	4:  "GABUMON",
	5:  "GARURUMON",
	6:  "WEREGARURUMON",
	7:  "METALGARURUMON",
	8:  "BIYOMON",
	9:  "BIRDRAMON",
	10: "GARUDAMON",
	11: "TENTOMON",
	12: "KABUTERIMON",
	13: "MEGAKABUTERIMON",
	14: "PALMON",
	15: "TOGEMON",
	16: "LILLYMON",
	17: "GOMAMON",
	18: "IKKAKUMON",
	19: "ZUDOMON",
	20: "PATAMON",
	21: "ANGEMON",
	22: "MAGNAANGEMON",
	23: "SALAMON",
	24: "GATOMON",
	25: "ANGEWOMON",
}

var DigimonIdentity_value = map[string]int32{
	"AGUMON":          0,
	"GREYMON":         1,
	"SKULLGREYMON":    2,
	"WARGREYMON":      3,
	"GABUMON":         4,
	"GARURUMON":       5,
	"WEREGARURUMON":   6,
	"METALGARURUMON":  7,
	"BIYOMON":         8,
	"BIRDRAMON":       9,
	"GARUDAMON":       10,
	"TENTOMON":        11,
	"KABUTERIMON":     12,
	"MEGAKABUTERIMON": 13,
	"PALMON":          14,
	"TOGEMON":         15,
	"LILLYMON":        16,
	"GOMAMON":         17,
	"IKKAKUMON":       18,
	"ZUDOMON":         19,
	"PATAMON":         20,
	"ANGEMON":         21,
	"MAGNAANGEMON":    22,
	"SALAMON":         23,
	"GATOMON":         24,
	"ANGEWOMON":       25,
}

func (x DigimonIdentity) String() string {
	return proto.EnumName(DigimonIdentity_name, int32(x))
}

func (DigimonIdentity) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5fe3dbe4306012f2, []int{0}
}

type LoginReq_LoginType int32

const (
	LoginReq_Visitor  LoginReq_LoginType = 0
	LoginReq_PassWord LoginReq_LoginType = 1
)

var LoginReq_LoginType_name = map[int32]string{
	0: "Visitor",
	1: "PassWord",
}

var LoginReq_LoginType_value = map[string]int32{
	"Visitor":  0,
	"PassWord": 1,
}

func (x LoginReq_LoginType) String() string {
	return proto.EnumName(LoginReq_LoginType_name, int32(x))
}

func (LoginReq_LoginType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5fe3dbe4306012f2, []int{2, 0}
}

type RoomInfo_RoomType int32

const (
	RoomInfo_TWO RoomInfo_RoomType = 0
)

var RoomInfo_RoomType_name = map[int32]string{
	0: "TWO",
}

var RoomInfo_RoomType_value = map[string]int32{
	"TWO": 0,
}

func (x RoomInfo_RoomType) String() string {
	return proto.EnumName(RoomInfo_RoomType_name, int32(x))
}

func (RoomInfo_RoomType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5fe3dbe4306012f2, []int{5, 0}
}

type MsgPack struct {
	Router               string   `protobuf:"bytes,1,opt,name=router,proto3" json:"router,omitempty"`
	Data                 []byte   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MsgPack) Reset()         { *m = MsgPack{} }
func (m *MsgPack) String() string { return proto.CompactTextString(m) }
func (*MsgPack) ProtoMessage()    {}
func (*MsgPack) Descriptor() ([]byte, []int) {
	return fileDescriptor_5fe3dbe4306012f2, []int{0}
}

func (m *MsgPack) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgPack.Unmarshal(m, b)
}
func (m *MsgPack) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgPack.Marshal(b, m, deterministic)
}
func (m *MsgPack) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgPack.Merge(m, src)
}
func (m *MsgPack) XXX_Size() int {
	return xxx_messageInfo_MsgPack.Size(m)
}
func (m *MsgPack) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgPack.DiscardUnknown(m)
}

var xxx_messageInfo_MsgPack proto.InternalMessageInfo

func (m *MsgPack) GetRouter() string {
	if m != nil {
		return m.Router
	}
	return ""
}

func (m *MsgPack) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type BaseAck struct {
	Result               int64    `protobuf:"varint,1,opt,name=Result,proto3" json:"Result,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=Msg,proto3" json:"Msg,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BaseAck) Reset()         { *m = BaseAck{} }
func (m *BaseAck) String() string { return proto.CompactTextString(m) }
func (*BaseAck) ProtoMessage()    {}
func (*BaseAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_5fe3dbe4306012f2, []int{1}
}

func (m *BaseAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BaseAck.Unmarshal(m, b)
}
func (m *BaseAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BaseAck.Marshal(b, m, deterministic)
}
func (m *BaseAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BaseAck.Merge(m, src)
}
func (m *BaseAck) XXX_Size() int {
	return xxx_messageInfo_BaseAck.Size(m)
}
func (m *BaseAck) XXX_DiscardUnknown() {
	xxx_messageInfo_BaseAck.DiscardUnknown(m)
}

var xxx_messageInfo_BaseAck proto.InternalMessageInfo

func (m *BaseAck) GetResult() int64 {
	if m != nil {
		return m.Result
	}
	return 0
}

func (m *BaseAck) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

type LoginReq struct {
	Username             string             `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password             string             `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	Type                 LoginReq_LoginType `protobuf:"varint,3,opt,name=type,proto3,enum=protobuf.LoginReq_LoginType" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *LoginReq) Reset()         { *m = LoginReq{} }
func (m *LoginReq) String() string { return proto.CompactTextString(m) }
func (*LoginReq) ProtoMessage()    {}
func (*LoginReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_5fe3dbe4306012f2, []int{2}
}

func (m *LoginReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginReq.Unmarshal(m, b)
}
func (m *LoginReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginReq.Marshal(b, m, deterministic)
}
func (m *LoginReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginReq.Merge(m, src)
}
func (m *LoginReq) XXX_Size() int {
	return xxx_messageInfo_LoginReq.Size(m)
}
func (m *LoginReq) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginReq.DiscardUnknown(m)
}

var xxx_messageInfo_LoginReq proto.InternalMessageInfo

func (m *LoginReq) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *LoginReq) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *LoginReq) GetType() LoginReq_LoginType {
	if m != nil {
		return m.Type
	}
	return LoginReq_Visitor
}

type PlayerInfo struct {
	Id                   uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Nickname             string   `protobuf:"bytes,2,opt,name=nickname,proto3" json:"nickname,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PlayerInfo) Reset()         { *m = PlayerInfo{} }
func (m *PlayerInfo) String() string { return proto.CompactTextString(m) }
func (*PlayerInfo) ProtoMessage()    {}
func (*PlayerInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_5fe3dbe4306012f2, []int{3}
}

func (m *PlayerInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlayerInfo.Unmarshal(m, b)
}
func (m *PlayerInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlayerInfo.Marshal(b, m, deterministic)
}
func (m *PlayerInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlayerInfo.Merge(m, src)
}
func (m *PlayerInfo) XXX_Size() int {
	return xxx_messageInfo_PlayerInfo.Size(m)
}
func (m *PlayerInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_PlayerInfo.DiscardUnknown(m)
}

var xxx_messageInfo_PlayerInfo proto.InternalMessageInfo

func (m *PlayerInfo) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *PlayerInfo) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

type LoginAck struct {
	Base                 *BaseAck    `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	PlayerInfo           *PlayerInfo `protobuf:"bytes,2,opt,name=player_info,json=playerInfo,proto3" json:"player_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *LoginAck) Reset()         { *m = LoginAck{} }
func (m *LoginAck) String() string { return proto.CompactTextString(m) }
func (*LoginAck) ProtoMessage()    {}
func (*LoginAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_5fe3dbe4306012f2, []int{4}
}

func (m *LoginAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoginAck.Unmarshal(m, b)
}
func (m *LoginAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoginAck.Marshal(b, m, deterministic)
}
func (m *LoginAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoginAck.Merge(m, src)
}
func (m *LoginAck) XXX_Size() int {
	return xxx_messageInfo_LoginAck.Size(m)
}
func (m *LoginAck) XXX_DiscardUnknown() {
	xxx_messageInfo_LoginAck.DiscardUnknown(m)
}

var xxx_messageInfo_LoginAck proto.InternalMessageInfo

func (m *LoginAck) GetBase() *BaseAck {
	if m != nil {
		return m.Base
	}
	return nil
}

func (m *LoginAck) GetPlayerInfo() *PlayerInfo {
	if m != nil {
		return m.PlayerInfo
	}
	return nil
}

type RoomInfo struct {
	RoomId               uint64            `protobuf:"varint,1,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	Type                 RoomInfo_RoomType `protobuf:"varint,2,opt,name=type,proto3,enum=protobuf.RoomInfo_RoomType" json:"type,omitempty"`
	IsStart              bool              `protobuf:"varint,3,opt,name=IsStart,proto3" json:"IsStart,omitempty"`
	CurrentPlayerNum     uint32            `protobuf:"varint,4,opt,name=current_player_num,json=currentPlayerNum,proto3" json:"current_player_num,omitempty"`
	PlayerInfos          []*PlayerInfo     `protobuf:"bytes,5,rep,name=player_infos,json=playerInfos,proto3" json:"player_infos,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *RoomInfo) Reset()         { *m = RoomInfo{} }
func (m *RoomInfo) String() string { return proto.CompactTextString(m) }
func (*RoomInfo) ProtoMessage()    {}
func (*RoomInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_5fe3dbe4306012f2, []int{5}
}

func (m *RoomInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RoomInfo.Unmarshal(m, b)
}
func (m *RoomInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RoomInfo.Marshal(b, m, deterministic)
}
func (m *RoomInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RoomInfo.Merge(m, src)
}
func (m *RoomInfo) XXX_Size() int {
	return xxx_messageInfo_RoomInfo.Size(m)
}
func (m *RoomInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RoomInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RoomInfo proto.InternalMessageInfo

func (m *RoomInfo) GetRoomId() uint64 {
	if m != nil {
		return m.RoomId
	}
	return 0
}

func (m *RoomInfo) GetType() RoomInfo_RoomType {
	if m != nil {
		return m.Type
	}
	return RoomInfo_TWO
}

func (m *RoomInfo) GetIsStart() bool {
	if m != nil {
		return m.IsStart
	}
	return false
}

func (m *RoomInfo) GetCurrentPlayerNum() uint32 {
	if m != nil {
		return m.CurrentPlayerNum
	}
	return 0
}

func (m *RoomInfo) GetPlayerInfos() []*PlayerInfo {
	if m != nil {
		return m.PlayerInfos
	}
	return nil
}

// join room
type JoinRoomReq struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JoinRoomReq) Reset()         { *m = JoinRoomReq{} }
func (m *JoinRoomReq) String() string { return proto.CompactTextString(m) }
func (*JoinRoomReq) ProtoMessage()    {}
func (*JoinRoomReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_5fe3dbe4306012f2, []int{6}
}

func (m *JoinRoomReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JoinRoomReq.Unmarshal(m, b)
}
func (m *JoinRoomReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JoinRoomReq.Marshal(b, m, deterministic)
}
func (m *JoinRoomReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JoinRoomReq.Merge(m, src)
}
func (m *JoinRoomReq) XXX_Size() int {
	return xxx_messageInfo_JoinRoomReq.Size(m)
}
func (m *JoinRoomReq) XXX_DiscardUnknown() {
	xxx_messageInfo_JoinRoomReq.DiscardUnknown(m)
}

var xxx_messageInfo_JoinRoomReq proto.InternalMessageInfo

type JoinRoomAck struct {
	Base                 *BaseAck  `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	RoomInfo             *RoomInfo `protobuf:"bytes,2,opt,name=room_info,json=roomInfo,proto3" json:"room_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *JoinRoomAck) Reset()         { *m = JoinRoomAck{} }
func (m *JoinRoomAck) String() string { return proto.CompactTextString(m) }
func (*JoinRoomAck) ProtoMessage()    {}
func (*JoinRoomAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_5fe3dbe4306012f2, []int{7}
}

func (m *JoinRoomAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JoinRoomAck.Unmarshal(m, b)
}
func (m *JoinRoomAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JoinRoomAck.Marshal(b, m, deterministic)
}
func (m *JoinRoomAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JoinRoomAck.Merge(m, src)
}
func (m *JoinRoomAck) XXX_Size() int {
	return xxx_messageInfo_JoinRoomAck.Size(m)
}
func (m *JoinRoomAck) XXX_DiscardUnknown() {
	xxx_messageInfo_JoinRoomAck.DiscardUnknown(m)
}

var xxx_messageInfo_JoinRoomAck proto.InternalMessageInfo

func (m *JoinRoomAck) GetBase() *BaseAck {
	if m != nil {
		return m.Base
	}
	return nil
}

func (m *JoinRoomAck) GetRoomInfo() *RoomInfo {
	if m != nil {
		return m.RoomInfo
	}
	return nil
}

type StartGameAck struct {
	Identity             DigimonIdentity `protobuf:"varint,1,opt,name=identity,proto3,enum=protobuf.DigimonIdentity" json:"identity,omitempty"`
	RoomInfo             *RoomInfo       `protobuf:"bytes,2,opt,name=room_info,json=roomInfo,proto3" json:"room_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *StartGameAck) Reset()         { *m = StartGameAck{} }
func (m *StartGameAck) String() string { return proto.CompactTextString(m) }
func (*StartGameAck) ProtoMessage()    {}
func (*StartGameAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_5fe3dbe4306012f2, []int{8}
}

func (m *StartGameAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StartGameAck.Unmarshal(m, b)
}
func (m *StartGameAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StartGameAck.Marshal(b, m, deterministic)
}
func (m *StartGameAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StartGameAck.Merge(m, src)
}
func (m *StartGameAck) XXX_Size() int {
	return xxx_messageInfo_StartGameAck.Size(m)
}
func (m *StartGameAck) XXX_DiscardUnknown() {
	xxx_messageInfo_StartGameAck.DiscardUnknown(m)
}

var xxx_messageInfo_StartGameAck proto.InternalMessageInfo

func (m *StartGameAck) GetIdentity() DigimonIdentity {
	if m != nil {
		return m.Identity
	}
	return DigimonIdentity_AGUMON
}

func (m *StartGameAck) GetRoomInfo() *RoomInfo {
	if m != nil {
		return m.RoomInfo
	}
	return nil
}

func init() {
	proto.RegisterEnum("protobuf.DigimonIdentity", DigimonIdentity_name, DigimonIdentity_value)
	proto.RegisterEnum("protobuf.LoginReq_LoginType", LoginReq_LoginType_name, LoginReq_LoginType_value)
	proto.RegisterEnum("protobuf.RoomInfo_RoomType", RoomInfo_RoomType_name, RoomInfo_RoomType_value)
	proto.RegisterType((*MsgPack)(nil), "protobuf.MsgPack")
	proto.RegisterType((*BaseAck)(nil), "protobuf.BaseAck")
	proto.RegisterType((*LoginReq)(nil), "protobuf.LoginReq")
	proto.RegisterType((*PlayerInfo)(nil), "protobuf.PlayerInfo")
	proto.RegisterType((*LoginAck)(nil), "protobuf.LoginAck")
	proto.RegisterType((*RoomInfo)(nil), "protobuf.RoomInfo")
	proto.RegisterType((*JoinRoomReq)(nil), "protobuf.JoinRoomReq")
	proto.RegisterType((*JoinRoomAck)(nil), "protobuf.JoinRoomAck")
	proto.RegisterType((*StartGameAck)(nil), "protobuf.StartGameAck")
}

func init() { proto.RegisterFile("protobuf/digimon.proto", fileDescriptor_5fe3dbe4306012f2) }

var fileDescriptor_5fe3dbe4306012f2 = []byte{
	// 706 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x53, 0xdb, 0x6e, 0xda, 0x4a,
	0x14, 0x8d, 0x81, 0x80, 0xd9, 0xe6, 0xe2, 0x0c, 0x39, 0x89, 0xcf, 0x39, 0x6a, 0x85, 0xfc, 0xd0,
	0xa2, 0xaa, 0x4a, 0x15, 0xa2, 0x2a, 0xed, 0xe3, 0x20, 0x2c, 0xe4, 0x62, 0x2e, 0x9a, 0x98, 0xa2,
	0xf4, 0x25, 0x72, 0xb0, 0x83, 0xac, 0x80, 0x87, 0xd8, 0x46, 0x15, 0xfd, 0x8d, 0xbe, 0xf7, 0x13,
	0xfb, 0x0d, 0xd5, 0x6c, 0xdb, 0x18, 0x55, 0x55, 0xa5, 0xf6, 0x6d, 0xaf, 0x7d, 0x5d, 0x6b, 0xd9,
	0x03, 0xda, 0xe6, 0x7e, 0x13, 0xf2, 0x98, 0x2f, 0xf8, 0xea, 0x8d, 0xeb, 0x2f, 0xfd, 0x35, 0x0f,
	0x2e, 0x30, 0x41, 0x20, 0xaf, 0xe8, 0x6f, 0xa1, 0x32, 0x8a, 0x96, 0x53, 0x67, 0xf1, 0x48, 0xce,
	0xa0, 0x1c, 0xf2, 0x6d, 0xec, 0x85, 0x9a, 0xd4, 0x96, 0x3a, 0x55, 0x96, 0x22, 0x42, 0xa0, 0xe4,
	0x3a, 0xb1, 0xa3, 0x15, 0xda, 0x52, 0xa7, 0xc6, 0x30, 0xd6, 0xaf, 0xa0, 0xd2, 0x73, 0x22, 0x8f,
	0x26, 0x63, 0xcc, 0x8b, 0xb6, 0xab, 0x18, 0xc7, 0x8a, 0x2c, 0x45, 0x44, 0x85, 0xe2, 0x28, 0x5a,
	0xe2, 0x54, 0x95, 0x89, 0x50, 0xff, 0x26, 0x81, 0x6c, 0xf1, 0xa5, 0x1f, 0x30, 0xef, 0x89, 0xfc,
	0x07, 0xf2, 0x36, 0xf2, 0xc2, 0xc0, 0x59, 0x7b, 0xe9, 0xbd, 0x3d, 0x16, 0xb5, 0x8d, 0x13, 0x45,
	0x9f, 0x79, 0xe8, 0xa6, 0xf3, 0x7b, 0x4c, 0xba, 0x50, 0x8a, 0x77, 0x1b, 0x4f, 0x2b, 0xb6, 0xa5,
	0x4e, 0xa3, 0xfb, 0xfc, 0x22, 0xd7, 0x72, 0x91, 0xed, 0x4e, 0x02, 0x7b, 0xb7, 0xf1, 0x18, 0xf6,
	0xea, 0x2f, 0xa0, 0xba, 0x4f, 0x11, 0x05, 0x2a, 0x1f, 0xfd, 0xc8, 0x8f, 0x79, 0xa8, 0x1e, 0x91,
	0x1a, 0xc8, 0x53, 0x27, 0x8a, 0xe6, 0x3c, 0x74, 0x55, 0x49, 0x7f, 0x07, 0x30, 0x5d, 0x39, 0x3b,
	0x2f, 0x34, 0x83, 0x07, 0x4e, 0x1a, 0x50, 0xf0, 0x5d, 0xe4, 0x56, 0x62, 0x05, 0xdf, 0x15, 0xac,
	0x02, 0x7f, 0xf1, 0x88, 0x8c, 0x53, 0x56, 0x19, 0xd6, 0x57, 0xa9, 0x32, 0x61, 0xc8, 0x4b, 0x28,
	0xdd, 0x3b, 0x51, 0xa2, 0x4a, 0xe9, 0xb6, 0x0e, 0x19, 0xa6, 0x9e, 0x31, 0x6c, 0x20, 0xd7, 0xa0,
	0x6c, 0xf0, 0xdc, 0x9d, 0x1f, 0x3c, 0x70, 0xdc, 0xa9, 0x74, 0xcf, 0x0e, 0xfb, 0x73, 0x36, 0x0c,
	0x36, 0xfb, 0x58, 0xff, 0x2e, 0x81, 0xcc, 0x38, 0x5f, 0x23, 0xcd, 0x73, 0xa8, 0x84, 0x9c, 0xaf,
	0xef, 0xf6, 0x5c, 0xcb, 0x02, 0x9a, 0x2e, 0xb9, 0x4c, 0x9d, 0x2a, 0xa0, 0x53, 0xcf, 0x0e, 0xf7,
	0x66, 0xc3, 0x18, 0xe4, 0x46, 0x11, 0x0d, 0x2a, 0x66, 0x74, 0x13, 0x3b, 0x61, 0x8c, 0xfe, 0xca,
	0x2c, 0x83, 0xe4, 0x35, 0x90, 0xc5, 0x36, 0x0c, 0xbd, 0x20, 0xbe, 0x4b, 0x39, 0x07, 0xdb, 0xb5,
	0x56, 0x6a, 0x4b, 0x9d, 0x3a, 0x53, 0xd3, 0x4a, 0xc2, 0x76, 0xbc, 0x5d, 0x93, 0xf7, 0x50, 0x3b,
	0x50, 0x16, 0x69, 0xc7, 0xed, 0xe2, 0x6f, 0xa4, 0x29, 0xb9, 0xb4, 0x48, 0x6f, 0x25, 0xd2, 0xf0,
	0x53, 0x55, 0xa0, 0x68, 0xcf, 0x27, 0xea, 0x91, 0x5e, 0x07, 0xe5, 0x03, 0xf7, 0x03, 0x51, 0x60,
	0xde, 0x93, 0xee, 0xe7, 0xf0, 0x8f, 0x0c, 0xbf, 0x84, 0x6a, 0x62, 0x55, 0x6e, 0xf7, 0xe9, 0xaf,
	0x6c, 0x61, 0x72, 0x98, 0x46, 0xfa, 0x17, 0xa8, 0xa1, 0x01, 0x03, 0x67, 0x8d, 0x7f, 0xfb, 0x35,
	0xc8, 0xbe, 0xeb, 0x05, 0xb1, 0x1f, 0xef, 0xf0, 0x5e, 0xa3, 0xfb, 0xff, 0xe1, 0x86, 0x7e, 0xf2,
	0xd0, 0xcc, 0xb4, 0x85, 0xed, 0x9b, 0xff, 0xe2, 0xf6, 0xab, 0xaf, 0x45, 0x68, 0xfe, 0xb4, 0x90,
	0x00, 0x94, 0xe9, 0x60, 0x36, 0x9a, 0x8c, 0xd5, 0x23, 0xf1, 0x27, 0x0f, 0x98, 0x71, 0x2b, 0x80,
	0x44, 0x54, 0xa8, 0xdd, 0x0c, 0x67, 0x96, 0x95, 0x65, 0x0a, 0xa4, 0x01, 0x30, 0xa7, 0x2c, 0xc3,
	0x45, 0x6c, 0xa7, 0x3d, 0x9c, 0x2d, 0x91, 0x3a, 0x54, 0x07, 0x94, 0xcd, 0x18, 0xc2, 0x63, 0x72,
	0x02, 0xf5, 0xb9, 0xc1, 0x8c, 0x3c, 0x55, 0x26, 0x04, 0x1a, 0x23, 0xc3, 0xa6, 0x56, 0x9e, 0xab,
	0x88, 0x15, 0x3d, 0xf3, 0x76, 0x22, 0x80, 0x2c, 0x56, 0xf4, 0x4c, 0xd6, 0x67, 0x54, 0xc0, 0x6a,
	0xb6, 0xb1, 0x8f, 0x10, 0xc4, 0xcb, 0xb2, 0x8d, 0xb1, 0x8d, 0xbd, 0x0a, 0x69, 0x82, 0x32, 0xa4,
	0xbd, 0x99, 0x6d, 0x30, 0x53, 0x24, 0x6a, 0xa4, 0x05, 0xcd, 0x91, 0x31, 0xa0, 0x87, 0xc9, 0xba,
	0x10, 0x37, 0xa5, 0x96, 0x88, 0x1b, 0xe2, 0x94, 0x3d, 0x19, 0x18, 0x02, 0x34, 0xc5, 0x32, 0xcb,
	0xb4, 0x2c, 0x14, 0xa2, 0xa2, 0x90, 0xc9, 0x08, 0xef, 0x9c, 0x88, 0xb3, 0xe6, 0x70, 0x48, 0x87,
	0xc8, 0x90, 0x88, 0xda, 0xa7, 0x59, 0x1f, 0xaf, 0xb6, 0x04, 0x98, 0x52, 0x1b, 0x1b, 0x4f, 0x05,
	0xa0, 0xe3, 0x64, 0xe1, 0x3f, 0xc2, 0xad, 0x11, 0x1d, 0x8c, 0x69, 0x96, 0x39, 0x13, 0xe5, 0x1b,
	0x6a, 0x61, 0xef, 0x79, 0x62, 0x55, 0xc2, 0x5d, 0x13, 0x17, 0x44, 0xdb, 0x1c, 0xe1, 0xbf, 0xf7,
	0x65, 0xfc, 0x64, 0x57, 0x3f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x3c, 0xd0, 0x70, 0xf5, 0x60, 0x05,
	0x00, 0x00,
}
