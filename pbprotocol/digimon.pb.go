// Code generated by protoc-gen-go. DO NOT EDIT.
// source: digimon.proto

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
	return fileDescriptor_9f35044cfecd686d, []int{0}
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
	return fileDescriptor_9f35044cfecd686d, []int{2, 0}
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
	return fileDescriptor_9f35044cfecd686d, []int{5, 0}
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
	return fileDescriptor_9f35044cfecd686d, []int{0}
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
	return fileDescriptor_9f35044cfecd686d, []int{1}
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
	Type                 LoginReq_LoginType `protobuf:"varint,3,opt,name=type,proto3,enum=pbprotocol.LoginReq_LoginType" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *LoginReq) Reset()         { *m = LoginReq{} }
func (m *LoginReq) String() string { return proto.CompactTextString(m) }
func (*LoginReq) ProtoMessage()    {}
func (*LoginReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f35044cfecd686d, []int{2}
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
	RoomId               uint64   `protobuf:"varint,3,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	Hero                 *Hero    `protobuf:"bytes,4,opt,name=hero,proto3" json:"hero,omitempty"`
	Seat                 int32    `protobuf:"varint,5,opt,name=seat,proto3" json:"seat,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PlayerInfo) Reset()         { *m = PlayerInfo{} }
func (m *PlayerInfo) String() string { return proto.CompactTextString(m) }
func (*PlayerInfo) ProtoMessage()    {}
func (*PlayerInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f35044cfecd686d, []int{3}
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

func (m *PlayerInfo) GetRoomId() uint64 {
	if m != nil {
		return m.RoomId
	}
	return 0
}

func (m *PlayerInfo) GetHero() *Hero {
	if m != nil {
		return m.Hero
	}
	return nil
}

func (m *PlayerInfo) GetSeat() int32 {
	if m != nil {
		return m.Seat
	}
	return 0
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
	return fileDescriptor_9f35044cfecd686d, []int{4}
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
	Type                 RoomInfo_RoomType `protobuf:"varint,2,opt,name=type,proto3,enum=pbprotocol.RoomInfo_RoomType" json:"type,omitempty"`
	IsStart              bool              `protobuf:"varint,3,opt,name=IsStart,proto3" json:"IsStart,omitempty"`
	CurrentPlayerNum     uint32            `protobuf:"varint,4,opt,name=current_player_num,json=currentPlayerNum,proto3" json:"current_player_num,omitempty"`
	PlayerInfos          []*PlayerInfo     `protobuf:"bytes,5,rep,name=player_infos,json=playerInfos,proto3" json:"player_infos,omitempty"`
	Round                int32             `protobuf:"varint,6,opt,name=round,proto3" json:"round,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *RoomInfo) Reset()         { *m = RoomInfo{} }
func (m *RoomInfo) String() string { return proto.CompactTextString(m) }
func (*RoomInfo) ProtoMessage()    {}
func (*RoomInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f35044cfecd686d, []int{5}
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

func (m *RoomInfo) GetRound() int32 {
	if m != nil {
		return m.Round
	}
	return 0
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
	return fileDescriptor_9f35044cfecd686d, []int{6}
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
	return fileDescriptor_9f35044cfecd686d, []int{7}
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

type Hero struct {
	Identity             DigimonIdentity `protobuf:"varint,1,opt,name=identity,proto3,enum=pbprotocol.DigimonIdentity" json:"identity,omitempty"`
	IdentityLevel        int32           `protobuf:"varint,2,opt,name=identity_level,json=identityLevel,proto3" json:"identity_level,omitempty"`
	SkillPoint           int32           `protobuf:"varint,3,opt,name=skill_point,json=skillPoint,proto3" json:"skill_point,omitempty"`
	SkillType            int32           `protobuf:"varint,4,opt,name=skill_type,json=skillType,proto3" json:"skill_type,omitempty"`
	SkillLevel           int32           `protobuf:"varint,5,opt,name=skill_level,json=skillLevel,proto3" json:"skill_level,omitempty"`
	SkillName            string          `protobuf:"bytes,6,opt,name=skill_name,json=skillName,proto3" json:"skill_name,omitempty"`
	SkillTargets         []uint64        `protobuf:"varint,7,rep,packed,name=skill_targets,json=skillTargets,proto3" json:"skill_targets,omitempty"`
	IsEscape             bool            `protobuf:"varint,8,opt,name=is_escape,json=isEscape,proto3" json:"is_escape,omitempty"`
	IsDead               bool            `protobuf:"varint,9,opt,name=is_dead,json=isDead,proto3" json:"is_dead,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Hero) Reset()         { *m = Hero{} }
func (m *Hero) String() string { return proto.CompactTextString(m) }
func (*Hero) ProtoMessage()    {}
func (*Hero) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f35044cfecd686d, []int{8}
}

func (m *Hero) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Hero.Unmarshal(m, b)
}
func (m *Hero) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Hero.Marshal(b, m, deterministic)
}
func (m *Hero) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Hero.Merge(m, src)
}
func (m *Hero) XXX_Size() int {
	return xxx_messageInfo_Hero.Size(m)
}
func (m *Hero) XXX_DiscardUnknown() {
	xxx_messageInfo_Hero.DiscardUnknown(m)
}

var xxx_messageInfo_Hero proto.InternalMessageInfo

func (m *Hero) GetIdentity() DigimonIdentity {
	if m != nil {
		return m.Identity
	}
	return DigimonIdentity_AGUMON
}

func (m *Hero) GetIdentityLevel() int32 {
	if m != nil {
		return m.IdentityLevel
	}
	return 0
}

func (m *Hero) GetSkillPoint() int32 {
	if m != nil {
		return m.SkillPoint
	}
	return 0
}

func (m *Hero) GetSkillType() int32 {
	if m != nil {
		return m.SkillType
	}
	return 0
}

func (m *Hero) GetSkillLevel() int32 {
	if m != nil {
		return m.SkillLevel
	}
	return 0
}

func (m *Hero) GetSkillName() string {
	if m != nil {
		return m.SkillName
	}
	return ""
}

func (m *Hero) GetSkillTargets() []uint64 {
	if m != nil {
		return m.SkillTargets
	}
	return nil
}

func (m *Hero) GetIsEscape() bool {
	if m != nil {
		return m.IsEscape
	}
	return false
}

func (m *Hero) GetIsDead() bool {
	if m != nil {
		return m.IsDead
	}
	return false
}

type StartGameAck struct {
	RoomInfo             *RoomInfo `protobuf:"bytes,1,opt,name=room_info,json=roomInfo,proto3" json:"room_info,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *StartGameAck) Reset()         { *m = StartGameAck{} }
func (m *StartGameAck) String() string { return proto.CompactTextString(m) }
func (*StartGameAck) ProtoMessage()    {}
func (*StartGameAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f35044cfecd686d, []int{9}
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

func (m *StartGameAck) GetRoomInfo() *RoomInfo {
	if m != nil {
		return m.RoomInfo
	}
	return nil
}

type ReleaseSkillReq struct {
	SkillType            int32    `protobuf:"varint,1,opt,name=skill_type,json=skillType,proto3" json:"skill_type,omitempty"`
	SkillLevel           int32    `protobuf:"varint,2,opt,name=skill_level,json=skillLevel,proto3" json:"skill_level,omitempty"`
	SkillTargets         []uint64 `protobuf:"varint,3,rep,packed,name=skill_targets,json=skillTargets,proto3" json:"skill_targets,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReleaseSkillReq) Reset()         { *m = ReleaseSkillReq{} }
func (m *ReleaseSkillReq) String() string { return proto.CompactTextString(m) }
func (*ReleaseSkillReq) ProtoMessage()    {}
func (*ReleaseSkillReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f35044cfecd686d, []int{10}
}

func (m *ReleaseSkillReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReleaseSkillReq.Unmarshal(m, b)
}
func (m *ReleaseSkillReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReleaseSkillReq.Marshal(b, m, deterministic)
}
func (m *ReleaseSkillReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReleaseSkillReq.Merge(m, src)
}
func (m *ReleaseSkillReq) XXX_Size() int {
	return xxx_messageInfo_ReleaseSkillReq.Size(m)
}
func (m *ReleaseSkillReq) XXX_DiscardUnknown() {
	xxx_messageInfo_ReleaseSkillReq.DiscardUnknown(m)
}

var xxx_messageInfo_ReleaseSkillReq proto.InternalMessageInfo

func (m *ReleaseSkillReq) GetSkillType() int32 {
	if m != nil {
		return m.SkillType
	}
	return 0
}

func (m *ReleaseSkillReq) GetSkillLevel() int32 {
	if m != nil {
		return m.SkillLevel
	}
	return 0
}

func (m *ReleaseSkillReq) GetSkillTargets() []uint64 {
	if m != nil {
		return m.SkillTargets
	}
	return nil
}

type ReleaseSkillAck struct {
	Base                 *BaseAck `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	Hero                 *Hero    `protobuf:"bytes,2,opt,name=hero,proto3" json:"hero,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ReleaseSkillAck) Reset()         { *m = ReleaseSkillAck{} }
func (m *ReleaseSkillAck) String() string { return proto.CompactTextString(m) }
func (*ReleaseSkillAck) ProtoMessage()    {}
func (*ReleaseSkillAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f35044cfecd686d, []int{11}
}

func (m *ReleaseSkillAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ReleaseSkillAck.Unmarshal(m, b)
}
func (m *ReleaseSkillAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ReleaseSkillAck.Marshal(b, m, deterministic)
}
func (m *ReleaseSkillAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReleaseSkillAck.Merge(m, src)
}
func (m *ReleaseSkillAck) XXX_Size() int {
	return xxx_messageInfo_ReleaseSkillAck.Size(m)
}
func (m *ReleaseSkillAck) XXX_DiscardUnknown() {
	xxx_messageInfo_ReleaseSkillAck.DiscardUnknown(m)
}

var xxx_messageInfo_ReleaseSkillAck proto.InternalMessageInfo

func (m *ReleaseSkillAck) GetBase() *BaseAck {
	if m != nil {
		return m.Base
	}
	return nil
}

func (m *ReleaseSkillAck) GetHero() *Hero {
	if m != nil {
		return m.Hero
	}
	return nil
}

type RulingResultAck struct {
	AttackerID           uint64   `protobuf:"varint,1,opt,name=attackerID,proto3" json:"attackerID,omitempty"`
	TargetID             uint64   `protobuf:"varint,2,opt,name=targetID,proto3" json:"targetID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RulingResultAck) Reset()         { *m = RulingResultAck{} }
func (m *RulingResultAck) String() string { return proto.CompactTextString(m) }
func (*RulingResultAck) ProtoMessage()    {}
func (*RulingResultAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f35044cfecd686d, []int{12}
}

func (m *RulingResultAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RulingResultAck.Unmarshal(m, b)
}
func (m *RulingResultAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RulingResultAck.Marshal(b, m, deterministic)
}
func (m *RulingResultAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RulingResultAck.Merge(m, src)
}
func (m *RulingResultAck) XXX_Size() int {
	return xxx_messageInfo_RulingResultAck.Size(m)
}
func (m *RulingResultAck) XXX_DiscardUnknown() {
	xxx_messageInfo_RulingResultAck.DiscardUnknown(m)
}

var xxx_messageInfo_RulingResultAck proto.InternalMessageInfo

func (m *RulingResultAck) GetAttackerID() uint64 {
	if m != nil {
		return m.AttackerID
	}
	return 0
}

func (m *RulingResultAck) GetTargetID() uint64 {
	if m != nil {
		return m.TargetID
	}
	return 0
}

type DeadNotifyAck struct {
	DeadIds              []uint64 `protobuf:"varint,1,rep,packed,name=dead_ids,json=deadIds,proto3" json:"dead_ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeadNotifyAck) Reset()         { *m = DeadNotifyAck{} }
func (m *DeadNotifyAck) String() string { return proto.CompactTextString(m) }
func (*DeadNotifyAck) ProtoMessage()    {}
func (*DeadNotifyAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f35044cfecd686d, []int{13}
}

func (m *DeadNotifyAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeadNotifyAck.Unmarshal(m, b)
}
func (m *DeadNotifyAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeadNotifyAck.Marshal(b, m, deterministic)
}
func (m *DeadNotifyAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeadNotifyAck.Merge(m, src)
}
func (m *DeadNotifyAck) XXX_Size() int {
	return xxx_messageInfo_DeadNotifyAck.Size(m)
}
func (m *DeadNotifyAck) XXX_DiscardUnknown() {
	xxx_messageInfo_DeadNotifyAck.DiscardUnknown(m)
}

var xxx_messageInfo_DeadNotifyAck proto.InternalMessageInfo

func (m *DeadNotifyAck) GetDeadIds() []uint64 {
	if m != nil {
		return m.DeadIds
	}
	return nil
}

type RPCBattleReq struct {
	Rpc                  int32    `protobuf:"varint,1,opt,name=rpc,proto3" json:"rpc,omitempty"`
	Role                 int32    `protobuf:"varint,2,opt,name=role,proto3" json:"role,omitempty"`
	OtherSideId          uint64   `protobuf:"varint,3,opt,name=other_side_id,json=otherSideId,proto3" json:"other_side_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RPCBattleReq) Reset()         { *m = RPCBattleReq{} }
func (m *RPCBattleReq) String() string { return proto.CompactTextString(m) }
func (*RPCBattleReq) ProtoMessage()    {}
func (*RPCBattleReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f35044cfecd686d, []int{14}
}

func (m *RPCBattleReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RPCBattleReq.Unmarshal(m, b)
}
func (m *RPCBattleReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RPCBattleReq.Marshal(b, m, deterministic)
}
func (m *RPCBattleReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RPCBattleReq.Merge(m, src)
}
func (m *RPCBattleReq) XXX_Size() int {
	return xxx_messageInfo_RPCBattleReq.Size(m)
}
func (m *RPCBattleReq) XXX_DiscardUnknown() {
	xxx_messageInfo_RPCBattleReq.DiscardUnknown(m)
}

var xxx_messageInfo_RPCBattleReq proto.InternalMessageInfo

func (m *RPCBattleReq) GetRpc() int32 {
	if m != nil {
		return m.Rpc
	}
	return 0
}

func (m *RPCBattleReq) GetRole() int32 {
	if m != nil {
		return m.Role
	}
	return 0
}

func (m *RPCBattleReq) GetOtherSideId() uint64 {
	if m != nil {
		return m.OtherSideId
	}
	return 0
}

type RPCBattleAck struct {
	Base                 *BaseAck `protobuf:"bytes,1,opt,name=base,proto3" json:"base,omitempty"`
	LastWinId            uint64   `protobuf:"varint,2,opt,name=last_win_id,json=lastWinId,proto3" json:"last_win_id,omitempty"`
	IsHaveNext           bool     `protobuf:"varint,3,opt,name=is_have_next,json=isHaveNext,proto3" json:"is_have_next,omitempty"`
	AttackerId           uint64   `protobuf:"varint,4,opt,name=attacker_id,json=attackerId,proto3" json:"attacker_id,omitempty"`
	TargetId             uint64   `protobuf:"varint,5,opt,name=target_id,json=targetId,proto3" json:"target_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RPCBattleAck) Reset()         { *m = RPCBattleAck{} }
func (m *RPCBattleAck) String() string { return proto.CompactTextString(m) }
func (*RPCBattleAck) ProtoMessage()    {}
func (*RPCBattleAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_9f35044cfecd686d, []int{15}
}

func (m *RPCBattleAck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RPCBattleAck.Unmarshal(m, b)
}
func (m *RPCBattleAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RPCBattleAck.Marshal(b, m, deterministic)
}
func (m *RPCBattleAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RPCBattleAck.Merge(m, src)
}
func (m *RPCBattleAck) XXX_Size() int {
	return xxx_messageInfo_RPCBattleAck.Size(m)
}
func (m *RPCBattleAck) XXX_DiscardUnknown() {
	xxx_messageInfo_RPCBattleAck.DiscardUnknown(m)
}

var xxx_messageInfo_RPCBattleAck proto.InternalMessageInfo

func (m *RPCBattleAck) GetBase() *BaseAck {
	if m != nil {
		return m.Base
	}
	return nil
}

func (m *RPCBattleAck) GetLastWinId() uint64 {
	if m != nil {
		return m.LastWinId
	}
	return 0
}

func (m *RPCBattleAck) GetIsHaveNext() bool {
	if m != nil {
		return m.IsHaveNext
	}
	return false
}

func (m *RPCBattleAck) GetAttackerId() uint64 {
	if m != nil {
		return m.AttackerId
	}
	return 0
}

func (m *RPCBattleAck) GetTargetId() uint64 {
	if m != nil {
		return m.TargetId
	}
	return 0
}

func init() {
	proto.RegisterEnum("pbprotocol.DigimonIdentity", DigimonIdentity_name, DigimonIdentity_value)
	proto.RegisterEnum("pbprotocol.LoginReq_LoginType", LoginReq_LoginType_name, LoginReq_LoginType_value)
	proto.RegisterEnum("pbprotocol.RoomInfo_RoomType", RoomInfo_RoomType_name, RoomInfo_RoomType_value)
	proto.RegisterType((*MsgPack)(nil), "pbprotocol.MsgPack")
	proto.RegisterType((*BaseAck)(nil), "pbprotocol.BaseAck")
	proto.RegisterType((*LoginReq)(nil), "pbprotocol.LoginReq")
	proto.RegisterType((*PlayerInfo)(nil), "pbprotocol.PlayerInfo")
	proto.RegisterType((*LoginAck)(nil), "pbprotocol.LoginAck")
	proto.RegisterType((*RoomInfo)(nil), "pbprotocol.RoomInfo")
	proto.RegisterType((*JoinRoomReq)(nil), "pbprotocol.JoinRoomReq")
	proto.RegisterType((*JoinRoomAck)(nil), "pbprotocol.JoinRoomAck")
	proto.RegisterType((*Hero)(nil), "pbprotocol.Hero")
	proto.RegisterType((*StartGameAck)(nil), "pbprotocol.StartGameAck")
	proto.RegisterType((*ReleaseSkillReq)(nil), "pbprotocol.ReleaseSkillReq")
	proto.RegisterType((*ReleaseSkillAck)(nil), "pbprotocol.ReleaseSkillAck")
	proto.RegisterType((*RulingResultAck)(nil), "pbprotocol.RulingResultAck")
	proto.RegisterType((*DeadNotifyAck)(nil), "pbprotocol.DeadNotifyAck")
	proto.RegisterType((*RPCBattleReq)(nil), "pbprotocol.RPCBattleReq")
	proto.RegisterType((*RPCBattleAck)(nil), "pbprotocol.RPCBattleAck")
}

func init() { proto.RegisterFile("digimon.proto", fileDescriptor_9f35044cfecd686d) }

var fileDescriptor_9f35044cfecd686d = []byte{
	// 1117 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x55, 0x4b, 0x6f, 0xdb, 0xc6,
	0x13, 0x0f, 0x25, 0xea, 0x35, 0xd4, 0x2b, 0xeb, 0xfc, 0x13, 0xe6, 0x1f, 0x24, 0x15, 0xd8, 0x97,
	0x10, 0x14, 0x01, 0xa2, 0xa0, 0x08, 0x7a, 0xa4, 0x2a, 0x41, 0x61, 0x2d, 0xc9, 0xc2, 0x5a, 0xae,
	0x9a, 0x5e, 0x58, 0x46, 0xdc, 0xc8, 0x0b, 0x53, 0xa4, 0xc2, 0xa5, 0x9c, 0xf8, 0xde, 0x5b, 0x8f,
	0xbd, 0xf7, 0x6b, 0xf4, 0xdc, 0x6f, 0x56, 0xcc, 0x90, 0x94, 0xe4, 0xb8, 0x48, 0xeb, 0xdb, 0xfc,
	0xe6, 0x3d, 0xbf, 0x59, 0x0e, 0xa1, 0xe1, 0xcb, 0x95, 0x5c, 0x47, 0xe1, 0xb3, 0x4d, 0x1c, 0x25,
	0x11, 0x83, 0xcd, 0x1b, 0x12, 0x96, 0x51, 0x60, 0x7d, 0x0b, 0x95, 0x89, 0x5a, 0xcd, 0xbc, 0xe5,
	0x05, 0xbb, 0x0f, 0xe5, 0x38, 0xda, 0x26, 0x22, 0x36, 0xb5, 0x8e, 0xd6, 0xad, 0xf1, 0x0c, 0x31,
	0x06, 0xba, 0xef, 0x25, 0x9e, 0x59, 0xe8, 0x68, 0xdd, 0x3a, 0x27, 0xd9, 0x7a, 0x01, 0x95, 0xbe,
	0xa7, 0x84, 0x9d, 0x86, 0x71, 0xa1, 0xb6, 0x41, 0x42, 0x61, 0x45, 0x9e, 0x21, 0xd6, 0x86, 0xe2,
	0x44, 0xad, 0x28, 0xaa, 0xc6, 0x51, 0xb4, 0xfe, 0xd0, 0xa0, 0x3a, 0x8e, 0x56, 0x32, 0xe4, 0xe2,
	0x1d, 0xfb, 0x3f, 0x54, 0xb7, 0x4a, 0xc4, 0xa1, 0xb7, 0x16, 0x59, 0xbd, 0x1d, 0x46, 0xdb, 0xc6,
	0x53, 0xea, 0x7d, 0x14, 0xfb, 0x59, 0xfc, 0x0e, 0xb3, 0x1e, 0xe8, 0xc9, 0xd5, 0x46, 0x98, 0xc5,
	0x8e, 0xd6, 0x6d, 0xf6, 0x9e, 0x3c, 0xdb, 0xcf, 0xf2, 0x2c, 0xcf, 0x9d, 0x0a, 0xf3, 0xab, 0x8d,
	0xe0, 0xe4, 0x6b, 0x7d, 0x05, 0xb5, 0x9d, 0x8a, 0x19, 0x50, 0xf9, 0x51, 0x2a, 0x99, 0x44, 0x71,
	0xfb, 0x0e, 0xab, 0x43, 0x75, 0xe6, 0x29, 0xb5, 0x88, 0x62, 0xbf, 0xad, 0x59, 0xbf, 0x69, 0x00,
	0xb3, 0xc0, 0xbb, 0x12, 0xb1, 0x13, 0xbe, 0x8d, 0x58, 0x13, 0x0a, 0xd2, 0xa7, 0xe6, 0x74, 0x5e,
	0x90, 0x3e, 0xb6, 0x15, 0xca, 0xe5, 0x05, 0xb5, 0x9c, 0xb5, 0x95, 0x63, 0xf6, 0x00, 0x2a, 0x71,
	0x14, 0xad, 0x5d, 0xe9, 0x53, 0x67, 0x3a, 0xb2, 0x17, 0xad, 0x1d, 0x9f, 0x7d, 0x01, 0xfa, 0xb9,
	0x88, 0x23, 0x53, 0xef, 0x68, 0x5d, 0xa3, 0xd7, 0x3e, 0xec, 0xf7, 0x95, 0x88, 0x23, 0x4e, 0x56,
	0xe4, 0x58, 0x09, 0x2f, 0x31, 0x4b, 0x1d, 0xad, 0x5b, 0xe2, 0x24, 0x5b, 0x41, 0xc6, 0x16, 0x92,
	0xfc, 0x35, 0xe8, 0x6f, 0x3c, 0x95, 0x32, 0x65, 0xf4, 0x8e, 0x0e, 0xb3, 0x64, 0x7b, 0xe0, 0xe4,
	0xc0, 0x5e, 0x82, 0xb1, 0xa1, 0x09, 0x5c, 0x19, 0xbe, 0x8d, 0xa8, 0x4d, 0xa3, 0x77, 0xff, 0xd0,
	0x7f, 0x3f, 0x20, 0x87, 0xcd, 0x4e, 0xb6, 0x7e, 0x2d, 0x40, 0x95, 0x63, 0xcb, 0x38, 0xf9, 0xc1,
	0x34, 0xda, 0xb5, 0x69, 0x9e, 0x67, 0xec, 0x17, 0x88, 0xfd, 0xc7, 0x87, 0x79, 0xf3, 0x60, 0x12,
	0xf6, 0xe4, 0x33, 0x13, 0x2a, 0x8e, 0x3a, 0x4d, 0xbc, 0x38, 0x21, 0x66, 0xaa, 0x3c, 0x87, 0xec,
	0x1b, 0x60, 0xcb, 0x6d, 0x1c, 0x8b, 0x30, 0x71, 0xb3, 0x9e, 0xc3, 0xed, 0x9a, 0x88, 0x6a, 0xf0,
	0x76, 0x66, 0x49, 0xbb, 0x9d, 0x6e, 0xd7, 0xec, 0x3b, 0xa8, 0x1f, 0x4c, 0xa6, 0xcc, 0x52, 0xa7,
	0xf8, 0x89, 0xd1, 0x8c, 0xfd, 0x68, 0x8a, 0xdd, 0x83, 0x52, 0x1c, 0x6d, 0x43, 0xdf, 0x2c, 0x13,
	0xbd, 0x29, 0xb0, 0x8e, 0xd2, 0x81, 0xe9, 0x51, 0x54, 0xa0, 0x38, 0x5f, 0x9c, 0xb4, 0xef, 0x58,
	0x0d, 0x30, 0x7e, 0x88, 0x64, 0x88, 0x06, 0x2e, 0xde, 0x59, 0x72, 0x0f, 0x6f, 0xb5, 0x86, 0xe7,
	0x50, 0x4b, 0x09, 0xdc, 0x2f, 0xe1, 0xde, 0x3f, 0x91, 0xc5, 0xab, 0x71, 0x26, 0x59, 0x7f, 0x15,
	0x40, 0xc7, 0x17, 0xc1, 0x5e, 0x42, 0x55, 0xfa, 0x22, 0x4c, 0x64, 0x72, 0x45, 0x85, 0x9a, 0xbd,
	0x47, 0x87, 0xa1, 0x83, 0xf4, 0x5b, 0x76, 0x32, 0x17, 0xbe, 0x73, 0x66, 0x5f, 0x42, 0x33, 0x97,
	0xdd, 0x40, 0x5c, 0x8a, 0x80, 0x2a, 0x97, 0x78, 0x23, 0xd7, 0x8e, 0x51, 0xc9, 0x3e, 0x03, 0x43,
	0x5d, 0xc8, 0x20, 0x70, 0x37, 0x91, 0x0c, 0xd3, 0xa5, 0x94, 0x38, 0x90, 0x6a, 0x86, 0x1a, 0xf6,
	0x18, 0x52, 0xe4, 0xd2, 0xaa, 0x75, 0xb2, 0xd7, 0x48, 0x43, 0x5c, 0xed, 0xe2, 0xd3, 0x1a, 0xa5,
	0x83, 0xf8, 0xb4, 0xc0, 0x2e, 0x9e, 0xbe, 0x94, 0x32, 0x7d, 0x29, 0x69, 0xfc, 0x14, 0x3f, 0x95,
	0xcf, 0xa1, 0x91, 0xa5, 0xf7, 0xe2, 0x95, 0x48, 0x94, 0x59, 0xe9, 0x14, 0xbb, 0x3a, 0xaf, 0xa7,
	0x15, 0x52, 0x1d, 0x7b, 0x04, 0x35, 0xa9, 0x5c, 0xa1, 0x96, 0xde, 0x46, 0x98, 0x55, 0x7a, 0x37,
	0x55, 0xa9, 0x86, 0x84, 0xf1, 0x79, 0x4a, 0xe5, 0xfa, 0xc2, 0xf3, 0xcd, 0x1a, 0x99, 0xca, 0x52,
	0x0d, 0x84, 0xe7, 0x5b, 0x36, 0xd4, 0xe9, 0x69, 0x8d, 0xbc, 0x35, 0xdd, 0xa6, 0x6b, 0x6b, 0xd0,
	0xfe, 0xd3, 0x1a, 0x2e, 0xa1, 0xc5, 0x45, 0x20, 0x3c, 0x25, 0x4e, 0xb1, 0x1f, 0x3c, 0x55, 0xd7,
	0xf9, 0xd0, 0xfe, 0x85, 0x8f, 0xc2, 0x0d, 0x3e, 0x6e, 0x0c, 0x5c, 0xbc, 0x39, 0xb0, 0xf5, 0xcb,
	0xf5, 0xba, 0xb7, 0x7a, 0x6d, 0xf9, 0x8d, 0x29, 0x7c, 0xea, 0xc6, 0x58, 0x13, 0x68, 0xf1, 0x6d,
	0x20, 0xc3, 0x55, 0x7a, 0xa0, 0xb1, 0xc2, 0x13, 0x00, 0x2f, 0x49, 0xbc, 0xe5, 0x85, 0x88, 0x9d,
	0x41, 0xf6, 0xa9, 0x1f, 0x68, 0xf0, 0xe2, 0xa5, 0x3d, 0x3b, 0x03, 0x4a, 0xae, 0xf3, 0x1d, 0xb6,
	0x9e, 0x42, 0x03, 0x39, 0x9f, 0x46, 0x89, 0x7c, 0x7b, 0x85, 0xc9, 0x1e, 0x42, 0x15, 0x57, 0xe2,
	0x4a, 0x5f, 0x99, 0x1a, 0x4d, 0x58, 0x41, 0xec, 0xf8, 0xca, 0xfa, 0x09, 0xea, 0x7c, 0xf6, 0x7d,
	0xdf, 0x4b, 0x92, 0x40, 0x20, 0xa3, 0x6d, 0x28, 0xc6, 0x9b, 0x65, 0x46, 0x25, 0x8a, 0x78, 0x00,
	0xe3, 0x28, 0x10, 0x19, 0x7b, 0x24, 0x33, 0x0b, 0x1a, 0x51, 0x72, 0x2e, 0x62, 0x57, 0x49, 0x5f,
	0xec, 0x2f, 0xab, 0x41, 0xca, 0x53, 0xe9, 0x0b, 0xc7, 0xb7, 0xfe, 0xd4, 0x0e, 0x52, 0xdf, 0x8a,
	0xb4, 0x27, 0x60, 0x04, 0x9e, 0x4a, 0xdc, 0xf7, 0x32, 0xc4, 0xdc, 0xe9, 0x78, 0x35, 0x54, 0x2d,
	0x64, 0xe8, 0xf8, 0xac, 0x03, 0x75, 0xa9, 0xdc, 0x73, 0xef, 0x52, 0xb8, 0xa1, 0xf8, 0x90, 0x1f,
	0x2f, 0x90, 0xea, 0x95, 0x77, 0x29, 0xa6, 0xe2, 0x43, 0x82, 0x8b, 0xcf, 0xb9, 0xc2, 0x0c, 0xfa,
	0x47, 0xf4, 0xf9, 0xf8, 0x88, 0x53, 0xba, 0xd0, 0x5c, 0xba, 0xc6, 0x9f, 0xff, 0xf4, 0xf7, 0x22,
	0xb4, 0x3e, 0xfa, 0x96, 0x19, 0x40, 0xd9, 0x1e, 0x9d, 0x4d, 0x4e, 0xa6, 0xed, 0x3b, 0xf8, 0x9f,
	0x1a, 0xf1, 0xe1, 0x6b, 0x04, 0x1a, 0x6b, 0x43, 0xfd, 0xf4, 0xf8, 0x6c, 0x3c, 0xce, 0x35, 0x05,
	0xd6, 0x04, 0x58, 0xd8, 0x3c, 0xc7, 0x45, 0x72, 0xb7, 0xfb, 0x14, 0xab, 0xb3, 0x06, 0xd4, 0x46,
	0x36, 0x3f, 0xe3, 0x04, 0x4b, 0xec, 0x2e, 0x34, 0x16, 0x43, 0x3e, 0xdc, 0xab, 0xca, 0x8c, 0x41,
	0x73, 0x32, 0x9c, 0xdb, 0xe3, 0xbd, 0xae, 0x82, 0x29, 0xfa, 0xce, 0xeb, 0x13, 0x04, 0x55, 0x4c,
	0xd1, 0x77, 0xf8, 0x80, 0xdb, 0x08, 0x6b, 0x79, 0xc6, 0x01, 0x41, 0xc0, 0xff, 0xe6, 0x7c, 0x38,
	0x9d, 0x93, 0xaf, 0xc1, 0x5a, 0x60, 0x1c, 0xdb, 0xfd, 0xb3, 0xf9, 0x90, 0x3b, 0xa8, 0xa8, 0xb3,
	0x23, 0x68, 0x4d, 0x86, 0x23, 0xfb, 0x50, 0xd9, 0xc0, 0xe1, 0x66, 0xf6, 0x18, 0xe5, 0x26, 0x96,
	0x9a, 0x9f, 0x8c, 0x86, 0x08, 0x5a, 0x98, 0x6c, 0xec, 0x8c, 0xc7, 0x34, 0x48, 0x9b, 0x06, 0x39,
	0x99, 0x50, 0x9d, 0xbb, 0x58, 0xd6, 0x39, 0x3e, 0xb6, 0x8f, 0xa9, 0x43, 0x86, 0xb6, 0x9f, 0xcf,
	0x06, 0x54, 0xf5, 0x08, 0xc1, 0xcc, 0x9e, 0x93, 0xe3, 0x3d, 0x04, 0xf6, 0x34, 0x4d, 0xf8, 0x3f,
	0x64, 0x6b, 0x62, 0x8f, 0xa6, 0x76, 0xae, 0xb9, 0x8f, 0xe6, 0x53, 0x7b, 0x4c, 0xbe, 0x0f, 0x52,
	0xaa, 0xd2, 0xde, 0x4d, 0xac, 0x80, 0x6e, 0x0b, 0x82, 0x0f, 0xdf, 0x94, 0xe9, 0xb5, 0xbc, 0xf8,
	0x3b, 0x00, 0x00, 0xff, 0xff, 0x90, 0xbd, 0x83, 0xa8, 0x33, 0x09, 0x00, 0x00,
}
