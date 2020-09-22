// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lib/config/blockpullorder.proto

package config

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type BlockPullOrder int32

const (
	BlockPullOrderStandard BlockPullOrder = 0
	BlockPullOrderRandom   BlockPullOrder = 1
	BlockPullOrderInOrder  BlockPullOrder = 2
)

var BlockPullOrder_name = map[int32]string{
	0: "BLOCK_PULL_ORDER_STANDARD",
	1: "BLOCK_PULL_ORDER_RANDOM",
	2: "BLOCK_PULL_ORDER_IN_ORDER",
}

var BlockPullOrder_value = map[string]int32{
	"BLOCK_PULL_ORDER_STANDARD": 0,
	"BLOCK_PULL_ORDER_RANDOM":   1,
	"BLOCK_PULL_ORDER_IN_ORDER": 2,
}

func (BlockPullOrder) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_3c46a5289006da6c, []int{0}
}

func init() {
	proto.RegisterEnum("config.BlockPullOrder", BlockPullOrder_name, BlockPullOrder_value)
}

func init() { proto.RegisterFile("lib/config/blockpullorder.proto", fileDescriptor_3c46a5289006da6c) }

var fileDescriptor_3c46a5289006da6c = []byte{
	// 237 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xcf, 0xc9, 0x4c, 0xd2,
	0x4f, 0xce, 0xcf, 0x4b, 0xcb, 0x4c, 0xd7, 0x4f, 0xca, 0xc9, 0x4f, 0xce, 0x2e, 0x28, 0xcd, 0xc9,
	0xc9, 0x2f, 0x4a, 0x49, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83, 0x48, 0x6a,
	0xed, 0x61, 0xe4, 0xe2, 0x73, 0x02, 0x29, 0x08, 0x28, 0xcd, 0xc9, 0xf1, 0x07, 0x29, 0x10, 0xb2,
	0xe4, 0x92, 0x74, 0xf2, 0xf1, 0x77, 0xf6, 0x8e, 0x0f, 0x08, 0xf5, 0xf1, 0x89, 0xf7, 0x0f, 0x72,
	0x71, 0x0d, 0x8a, 0x0f, 0x0e, 0x71, 0xf4, 0x73, 0x71, 0x0c, 0x72, 0x11, 0x60, 0x90, 0x92, 0xea,
	0x9a, 0xab, 0x20, 0x86, 0xaa, 0x25, 0xb8, 0x24, 0x31, 0x2f, 0x25, 0xb1, 0x28, 0x45, 0xc8, 0x94,
	0x4b, 0x1c, 0x43, 0x6b, 0x90, 0xa3, 0x9f, 0x8b, 0xbf, 0xaf, 0x00, 0xa3, 0x94, 0x44, 0xd7, 0x5c,
	0x05, 0x11, 0x54, 0x8d, 0x41, 0x89, 0x79, 0x29, 0xf9, 0xb9, 0x42, 0x16, 0x58, 0x6c, 0xf4, 0xf4,
	0x83, 0x30, 0x04, 0x98, 0xa4, 0x24, 0xbb, 0xe6, 0x2a, 0x88, 0xa2, 0x6a, 0xf4, 0xcc, 0x03, 0x53,
	0x4e, 0xee, 0x27, 0x1e, 0xca, 0x31, 0x5c, 0x78, 0x28, 0xc7, 0xf0, 0xe2, 0x91, 0x1c, 0xc3, 0x84,
	0xc7, 0x72, 0x0c, 0x0b, 0x1e, 0xcb, 0x31, 0x5e, 0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c, 0x43,
	0x94, 0x66, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0x7e, 0x71, 0x65, 0x5e,
	0x72, 0x49, 0x46, 0x66, 0x5e, 0x3a, 0x12, 0x0b, 0x11, 0x48, 0x49, 0x6c, 0xe0, 0x60, 0x31, 0x06,
	0x04, 0x00, 0x00, 0xff, 0xff, 0x46, 0xfd, 0xeb, 0xcf, 0x39, 0x01, 0x00, 0x00,
}
