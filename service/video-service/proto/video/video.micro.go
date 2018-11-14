// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/video/video.proto

/*
Package shop_srv_video is a generated protocol buffer package.

It is generated from these files:
	proto/video/video.proto

It has these top-level messages:
	VideoListReq
	VideotResp
	VideoListResp
*/
package shop_srv_video

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "context"
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

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Video service

type VideoService interface {
	GetVideoList(ctx context.Context, in *VideoListReq, opts ...client.CallOption) (*VideoListResp, error)
}

type videoService struct {
	c    client.Client
	name string
}

func NewVideoService(name string, c client.Client) VideoService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "shop.srv.video"
	}
	return &videoService{
		c:    c,
		name: name,
	}
}

func (c *videoService) GetVideoList(ctx context.Context, in *VideoListReq, opts ...client.CallOption) (*VideoListResp, error) {
	req := c.c.NewRequest(c.name, "Video.GetVideoList", in)
	out := new(VideoListResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Video service

type VideoHandler interface {
	GetVideoList(context.Context, *VideoListReq, *VideoListResp) error
}

func RegisterVideoHandler(s server.Server, hdlr VideoHandler, opts ...server.HandlerOption) error {
	type video interface {
		GetVideoList(ctx context.Context, in *VideoListReq, out *VideoListResp) error
	}
	type Video struct {
		video
	}
	h := &videoHandler{hdlr}
	return s.Handle(s.NewHandler(&Video{h}, opts...))
}

type videoHandler struct {
	VideoHandler
}

func (h *videoHandler) GetVideoList(ctx context.Context, in *VideoListReq, out *VideoListResp) error {
	return h.VideoHandler.GetVideoList(ctx, in, out)
}
