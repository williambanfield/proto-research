// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: awesome.proto

package protos

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
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

type AwesomeMessage struct {
	AwesomeField string `protobuf:"bytes,1,opt,name=awesome_field,json=awesomeField,proto3" json:"awesome_field,omitempty"`
	AwesomeBool  bool   `protobuf:"varint,2,opt,name=awesome_bool,json=awesomeBool,proto3" json:"awesome_bool,omitempty"`
}

func (m *AwesomeMessage) Reset()         { *m = AwesomeMessage{} }
func (m *AwesomeMessage) String() string { return proto.CompactTextString(m) }
func (*AwesomeMessage) ProtoMessage()    {}
func (*AwesomeMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_b0654cb17c4c3aed, []int{0}
}
func (m *AwesomeMessage) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AwesomeMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AwesomeMessage.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AwesomeMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AwesomeMessage.Merge(m, src)
}
func (m *AwesomeMessage) XXX_Size() int {
	return m.Size()
}
func (m *AwesomeMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_AwesomeMessage.DiscardUnknown(m)
}

var xxx_messageInfo_AwesomeMessage proto.InternalMessageInfo

func (m *AwesomeMessage) GetAwesomeField() string {
	if m != nil {
		return m.AwesomeField
	}
	return ""
}

func (m *AwesomeMessage) GetAwesomeBool() bool {
	if m != nil {
		return m.AwesomeBool
	}
	return false
}

func init() {
	proto.RegisterType((*AwesomeMessage)(nil), "awesomepackage.AwesomeMessage")
}

func init() { proto.RegisterFile("awesome.proto", fileDescriptor_b0654cb17c4c3aed) }

var fileDescriptor_b0654cb17c4c3aed = []byte{
	// 146 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0x2c, 0x4f, 0x2d,
	0xce, 0xcf, 0x4d, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x83, 0x72, 0x0b, 0x12, 0x93,
	0xb3, 0x13, 0xd3, 0x53, 0x95, 0x22, 0xb8, 0xf8, 0x1c, 0x21, 0x22, 0xbe, 0xa9, 0xc5, 0xc5, 0x89,
	0xe9, 0xa9, 0x42, 0xca, 0x70, 0x2d, 0xf1, 0x69, 0x99, 0xa9, 0x39, 0x29, 0x12, 0x8c, 0x0a, 0x8c,
	0x1a, 0x9c, 0x41, 0x3c, 0x50, 0x41, 0x37, 0x90, 0x98, 0x90, 0x22, 0x17, 0x8c, 0x1f, 0x9f, 0x94,
	0x9f, 0x9f, 0x23, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0x11, 0xc4, 0x0d, 0x15, 0x73, 0xca, 0xcf, 0xcf,
	0x71, 0x52, 0x38, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x27,
	0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86, 0x28, 0x36, 0xb0, 0x53,
	0x8a, 0x93, 0x20, 0xb4, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x66, 0xfa, 0x57, 0xe0, 0xa3, 0x00,
	0x00, 0x00,
}

func (m *AwesomeMessage) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AwesomeMessage) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AwesomeMessage) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.AwesomeBool {
		i--
		if m.AwesomeBool {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if len(m.AwesomeField) > 0 {
		i -= len(m.AwesomeField)
		copy(dAtA[i:], m.AwesomeField)
		i = encodeVarintAwesome(dAtA, i, uint64(len(m.AwesomeField)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintAwesome(dAtA []byte, offset int, v uint64) int {
	offset -= sovAwesome(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AwesomeMessage) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.AwesomeField)
	if l > 0 {
		n += 1 + l + sovAwesome(uint64(l))
	}
	if m.AwesomeBool {
		n += 2
	}
	return n
}

func sovAwesome(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAwesome(x uint64) (n int) {
	return sovAwesome(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AwesomeMessage) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAwesome
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AwesomeMessage: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AwesomeMessage: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AwesomeField", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAwesome
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthAwesome
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAwesome
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AwesomeField = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AwesomeBool", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAwesome
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.AwesomeBool = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipAwesome(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAwesome
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipAwesome(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAwesome
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowAwesome
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowAwesome
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthAwesome
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAwesome
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAwesome
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAwesome        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAwesome          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAwesome = fmt.Errorf("proto: unexpected end of group")
)