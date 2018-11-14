// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/video/video.proto

package shop_srv_video

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type VideoListReq struct {
	CategoryId           string   `protobuf:"bytes,1,opt,name=CategoryId,proto3" json:"CategoryId,omitempty"`
	PageNum              int32    `protobuf:"varint,2,opt,name=PageNum,proto3" json:"PageNum,omitempty"`
	PageSize             int32    `protobuf:"varint,3,opt,name=PageSize,proto3" json:"PageSize,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VideoListReq) Reset()         { *m = VideoListReq{} }
func (m *VideoListReq) String() string { return proto.CompactTextString(m) }
func (*VideoListReq) ProtoMessage()    {}
func (*VideoListReq) Descriptor() ([]byte, []int) {
	return fileDescriptor_f8a89cf20eaf7d4e, []int{0}
}

func (m *VideoListReq) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VideoListReq.Unmarshal(m, b)
}
func (m *VideoListReq) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VideoListReq.Marshal(b, m, deterministic)
}
func (m *VideoListReq) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VideoListReq.Merge(m, src)
}
func (m *VideoListReq) XXX_Size() int {
	return xxx_messageInfo_VideoListReq.Size(m)
}
func (m *VideoListReq) XXX_DiscardUnknown() {
	xxx_messageInfo_VideoListReq.DiscardUnknown(m)
}

var xxx_messageInfo_VideoListReq proto.InternalMessageInfo

func (m *VideoListReq) GetCategoryId() string {
	if m != nil {
		return m.CategoryId
	}
	return ""
}

func (m *VideoListReq) GetPageNum() int32 {
	if m != nil {
		return m.PageNum
	}
	return 0
}

func (m *VideoListReq) GetPageSize() int32 {
	if m != nil {
		return m.PageSize
	}
	return 0
}

type VideotResp struct {
	Id                   int64    `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=Title,proto3" json:"Title,omitempty"`
	ThumbUrl             string   `protobuf:"bytes,3,opt,name=ThumbUrl,proto3" json:"ThumbUrl,omitempty"`
	Avatar               string   `protobuf:"bytes,4,opt,name=Avatar,proto3" json:"Avatar,omitempty"`
	Author               string   `protobuf:"bytes,5,opt,name=Author,proto3" json:"Author,omitempty"`
	Duration             string   `protobuf:"bytes,6,opt,name=Duration,proto3" json:"Duration,omitempty"`
	ReadCount            string   `protobuf:"bytes,7,opt,name=ReadCount,proto3" json:"ReadCount,omitempty"`
	CommentCount         string   `protobuf:"bytes,8,opt,name=CommentCount,proto3" json:"CommentCount,omitempty"`
	LikeCount            int32    `protobuf:"varint,9,opt,name=LikeCount,proto3" json:"LikeCount,omitempty"`
	Category             string   `protobuf:"bytes,10,opt,name=Category,proto3" json:"Category,omitempty"`
	ViewType             int32    `protobuf:"varint,11,opt,name=ViewType,proto3" json:"ViewType,omitempty"`
	IsRecommend          int32    `protobuf:"varint,12,opt,name=IsRecommend,proto3" json:"IsRecommend,omitempty"`
	Content              string   `protobuf:"bytes,13,opt,name=Content,proto3" json:"Content,omitempty"`
	Tags                 string   `protobuf:"bytes,14,opt,name=Tags,proto3" json:"Tags,omitempty"`
	VideoDesc            string   `protobuf:"bytes,15,opt,name=VideoDesc,proto3" json:"VideoDesc,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VideotResp) Reset()         { *m = VideotResp{} }
func (m *VideotResp) String() string { return proto.CompactTextString(m) }
func (*VideotResp) ProtoMessage()    {}
func (*VideotResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_f8a89cf20eaf7d4e, []int{1}
}

func (m *VideotResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VideotResp.Unmarshal(m, b)
}
func (m *VideotResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VideotResp.Marshal(b, m, deterministic)
}
func (m *VideotResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VideotResp.Merge(m, src)
}
func (m *VideotResp) XXX_Size() int {
	return xxx_messageInfo_VideotResp.Size(m)
}
func (m *VideotResp) XXX_DiscardUnknown() {
	xxx_messageInfo_VideotResp.DiscardUnknown(m)
}

var xxx_messageInfo_VideotResp proto.InternalMessageInfo

func (m *VideotResp) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *VideotResp) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *VideotResp) GetThumbUrl() string {
	if m != nil {
		return m.ThumbUrl
	}
	return ""
}

func (m *VideotResp) GetAvatar() string {
	if m != nil {
		return m.Avatar
	}
	return ""
}

func (m *VideotResp) GetAuthor() string {
	if m != nil {
		return m.Author
	}
	return ""
}

func (m *VideotResp) GetDuration() string {
	if m != nil {
		return m.Duration
	}
	return ""
}

func (m *VideotResp) GetReadCount() string {
	if m != nil {
		return m.ReadCount
	}
	return ""
}

func (m *VideotResp) GetCommentCount() string {
	if m != nil {
		return m.CommentCount
	}
	return ""
}

func (m *VideotResp) GetLikeCount() int32 {
	if m != nil {
		return m.LikeCount
	}
	return 0
}

func (m *VideotResp) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

func (m *VideotResp) GetViewType() int32 {
	if m != nil {
		return m.ViewType
	}
	return 0
}

func (m *VideotResp) GetIsRecommend() int32 {
	if m != nil {
		return m.IsRecommend
	}
	return 0
}

func (m *VideotResp) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *VideotResp) GetTags() string {
	if m != nil {
		return m.Tags
	}
	return ""
}

func (m *VideotResp) GetVideoDesc() string {
	if m != nil {
		return m.VideoDesc
	}
	return ""
}

type VideoListResp struct {
	VideoResp            []*VideotResp `protobuf:"bytes,1,rep,name=videoResp,proto3" json:"videoResp,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *VideoListResp) Reset()         { *m = VideoListResp{} }
func (m *VideoListResp) String() string { return proto.CompactTextString(m) }
func (*VideoListResp) ProtoMessage()    {}
func (*VideoListResp) Descriptor() ([]byte, []int) {
	return fileDescriptor_f8a89cf20eaf7d4e, []int{2}
}

func (m *VideoListResp) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VideoListResp.Unmarshal(m, b)
}
func (m *VideoListResp) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VideoListResp.Marshal(b, m, deterministic)
}
func (m *VideoListResp) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VideoListResp.Merge(m, src)
}
func (m *VideoListResp) XXX_Size() int {
	return xxx_messageInfo_VideoListResp.Size(m)
}
func (m *VideoListResp) XXX_DiscardUnknown() {
	xxx_messageInfo_VideoListResp.DiscardUnknown(m)
}

var xxx_messageInfo_VideoListResp proto.InternalMessageInfo

func (m *VideoListResp) GetVideoResp() []*VideotResp {
	if m != nil {
		return m.VideoResp
	}
	return nil
}

func init() {
	proto.RegisterType((*VideoListReq)(nil), "shop.srv.video.VideoListReq")
	proto.RegisterType((*VideotResp)(nil), "shop.srv.video.VideotResp")
	proto.RegisterType((*VideoListResp)(nil), "shop.srv.video.VideoListResp")
}

func init() { proto.RegisterFile("proto/video/video.proto", fileDescriptor_f8a89cf20eaf7d4e) }

var fileDescriptor_f8a89cf20eaf7d4e = []byte{
	// 397 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x5f, 0xab, 0xd3, 0x40,
	0x10, 0xc5, 0x4d, 0x73, 0xd3, 0x7b, 0x33, 0xcd, 0xad, 0xb0, 0x88, 0x2e, 0xe5, 0x2a, 0x21, 0x4f,
	0x7d, 0x8a, 0x50, 0x5f, 0x7c, 0x95, 0x14, 0x24, 0x50, 0x45, 0xd6, 0xd8, 0xf7, 0xb4, 0x19, 0xda,
	0x60, 0x93, 0x8d, 0xd9, 0x4d, 0xa5, 0x7e, 0x1a, 0x3f, 0xaa, 0xec, 0xe4, 0x5f, 0x0b, 0x72, 0x5f,
	0xca, 0x9e, 0xdf, 0x99, 0x9e, 0xc9, 0xee, 0x0c, 0xbc, 0xa9, 0x6a, 0xa9, 0xe5, 0xfb, 0x73, 0x9e,
	0x61, 0xf7, 0x1b, 0x12, 0x61, 0x73, 0x75, 0x94, 0x55, 0xa8, 0xea, 0x73, 0x48, 0x34, 0xc8, 0xc0,
	0xdb, 0x9a, 0xc3, 0x26, 0x57, 0x5a, 0xe0, 0x2f, 0xf6, 0x0e, 0x20, 0x4a, 0x35, 0x1e, 0x64, 0x7d,
	0x89, 0x33, 0x6e, 0xf9, 0xd6, 0xd2, 0x15, 0x57, 0x84, 0x71, 0xb8, 0xff, 0x96, 0x1e, 0xf0, 0x6b,
	0x53, 0xf0, 0x89, 0x6f, 0x2d, 0x1d, 0xd1, 0x4b, 0xb6, 0x80, 0x07, 0x73, 0xfc, 0x9e, 0xff, 0x41,
	0x6e, 0x93, 0x35, 0xe8, 0xe0, 0xaf, 0x0d, 0x40, 0x6d, 0xb4, 0x40, 0x55, 0xb1, 0x39, 0x4c, 0xba,
	0x70, 0x5b, 0x4c, 0xe2, 0x8c, 0xbd, 0x02, 0x27, 0xc9, 0xf5, 0x09, 0x29, 0xd2, 0x15, 0xad, 0x30,
	0x81, 0xc9, 0xb1, 0x29, 0x76, 0x3f, 0xea, 0x13, 0x05, 0xba, 0x62, 0xd0, 0xec, 0x35, 0x4c, 0x3f,
	0x9d, 0x53, 0x9d, 0xd6, 0xfc, 0x8e, 0x9c, 0x4e, 0x11, 0x6f, 0xf4, 0x51, 0xd6, 0xdc, 0xe9, 0x38,
	0x29, 0x93, 0xb5, 0x6e, 0xea, 0x54, 0xe7, 0xb2, 0xe4, 0xd3, 0x36, 0xab, 0xd7, 0xec, 0x09, 0x5c,
	0x81, 0x69, 0x16, 0xc9, 0xa6, 0xd4, 0xfc, 0x9e, 0xcc, 0x11, 0xb0, 0x00, 0xbc, 0x48, 0x16, 0x05,
	0x96, 0xba, 0x2d, 0x78, 0xa0, 0x82, 0x1b, 0x66, 0x12, 0x36, 0xf9, 0x4f, 0x6c, 0x0b, 0x5c, 0xba,
	0xfb, 0x08, 0x4c, 0xef, 0xfe, 0x01, 0x39, 0xb4, 0xbd, 0x7b, 0x6d, 0xbc, 0x6d, 0x8e, 0xbf, 0x93,
	0x4b, 0x85, 0x7c, 0xd6, 0x3e, 0x5a, 0xaf, 0x99, 0x0f, 0xb3, 0x58, 0x09, 0xdc, 0x53, 0xa7, 0x8c,
	0x7b, 0x64, 0x5f, 0x23, 0x33, 0x8c, 0x48, 0x96, 0x1a, 0x4b, 0xcd, 0x1f, 0x29, 0xb8, 0x97, 0x8c,
	0xc1, 0x5d, 0x92, 0x1e, 0x14, 0x9f, 0x13, 0xa6, 0xb3, 0xf9, 0x4a, 0x9a, 0xc1, 0x1a, 0xd5, 0x9e,
	0xbf, 0x6c, 0xef, 0x39, 0x80, 0x20, 0x86, 0xc7, 0xab, 0x45, 0x50, 0x15, 0xfb, 0x08, 0x2e, 0xad,
	0x88, 0x11, 0xdc, 0xf2, 0xed, 0xe5, 0x6c, 0xb5, 0x08, 0x6f, 0xb7, 0x27, 0x1c, 0x67, 0x2a, 0xc6,
	0xe2, 0xd5, 0x16, 0x1c, 0x32, 0xd8, 0x17, 0xf0, 0x3e, 0xa3, 0x1e, 0x62, 0xd9, 0xd3, 0x7f, 0xff,
	0xdf, 0xad, 0xde, 0xe2, 0xed, 0x33, 0xae, 0xaa, 0x82, 0x17, 0xbb, 0x29, 0xad, 0xf0, 0x87, 0x7f,
	0x01, 0x00, 0x00, 0xff, 0xff, 0x2b, 0x33, 0x8c, 0x09, 0xdd, 0x02, 0x00, 0x00,
}
