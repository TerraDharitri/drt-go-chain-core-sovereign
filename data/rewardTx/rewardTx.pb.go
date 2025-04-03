// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: rewardTx.proto

package rewardTx

import (
	bytes "bytes"
	fmt "fmt"
	github_com_TerraDharitri_drt_go_chain_core_data "github.com/TerraDharitri/drt-go-chain-core/data"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_big "math/big"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
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

// RewardTx holds the data for a reward transaction
type RewardTx struct {
	Round   uint64        `protobuf:"varint,1,opt,name=Round,proto3" json:"round"`
	Value   *math_big.Int `protobuf:"bytes,3,opt,name=Value,proto3,casttypewith=math/big.Int;github.com/TerraDharitri/drt-go-chain-core/data.BigIntCaster" json:"value"`
	RcvAddr []byte        `protobuf:"bytes,4,opt,name=RcvAddr,proto3" json:"receiver"`
	Epoch   uint32        `protobuf:"varint,2,opt,name=Epoch,proto3" json:"epoch"`
}

func (m *RewardTx) Reset()      { *m = RewardTx{} }
func (*RewardTx) ProtoMessage() {}
func (*RewardTx) Descriptor() ([]byte, []int) {
	return fileDescriptor_25dbfb608d6baddf, []int{0}
}
func (m *RewardTx) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RewardTx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *RewardTx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RewardTx.Merge(m, src)
}
func (m *RewardTx) XXX_Size() int {
	return m.Size()
}
func (m *RewardTx) XXX_DiscardUnknown() {
	xxx_messageInfo_RewardTx.DiscardUnknown(m)
}

var xxx_messageInfo_RewardTx proto.InternalMessageInfo

func (m *RewardTx) GetRound() uint64 {
	if m != nil {
		return m.Round
	}
	return 0
}

func (m *RewardTx) GetValue() *math_big.Int {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *RewardTx) GetRcvAddr() []byte {
	if m != nil {
		return m.RcvAddr
	}
	return nil
}

func (m *RewardTx) GetEpoch() uint32 {
	if m != nil {
		return m.Epoch
	}
	return 0
}

func init() {
	proto.RegisterType((*RewardTx)(nil), "proto.RewardTx")
}

func init() { proto.RegisterFile("rewardTx.proto", fileDescriptor_25dbfb608d6baddf) }

var fileDescriptor_25dbfb608d6baddf = []byte{
	// 323 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x3f, 0x4e, 0xc3, 0x30,
	0x18, 0xc5, 0x63, 0x68, 0xa0, 0x44, 0x85, 0x21, 0x53, 0xc4, 0xf0, 0xa5, 0x62, 0x40, 0x5d, 0x92,
	0x0c, 0x8c, 0x9d, 0x1a, 0x60, 0xe8, 0x6a, 0x55, 0x1d, 0xd8, 0x9c, 0xc4, 0x24, 0x96, 0xda, 0xb8,
	0xfa, 0x70, 0x0a, 0x12, 0x0b, 0x47, 0xe0, 0x18, 0x88, 0x93, 0x30, 0x76, 0xec, 0x54, 0xa8, 0xbb,
	0xa0, 0x4e, 0x3d, 0x02, 0x8a, 0x43, 0x25, 0x56, 0x26, 0xbf, 0xf7, 0xf3, 0x9f, 0xf7, 0x6c, 0x3b,
	0x67, 0xc8, 0x1f, 0x19, 0x66, 0xa3, 0xa7, 0x70, 0x86, 0x52, 0x49, 0xd7, 0x36, 0xc3, 0x79, 0x90,
	0x0b, 0x55, 0x54, 0x49, 0x98, 0xca, 0x69, 0x94, 0xcb, 0x5c, 0x46, 0x06, 0x27, 0xd5, 0xbd, 0x71,
	0xc6, 0x18, 0xd5, 0xec, 0xba, 0xd0, 0xc4, 0x69, 0xd3, 0xdf, 0x83, 0x5c, 0xdf, 0xb1, 0xa9, 0xac,
	0xca, 0xcc, 0x23, 0x5d, 0xd2, 0x6b, 0xc5, 0x27, 0xdb, 0x95, 0x6f, 0x63, 0x0d, 0x68, 0xc3, 0xdd,
	0x89, 0x63, 0x8f, 0xd9, 0xa4, 0xe2, 0xde, 0x61, 0x97, 0xf4, 0x3a, 0xf1, 0xb8, 0x5e, 0x30, 0xaf,
	0xc1, 0xfb, 0xa7, 0x3f, 0x9c, 0x32, 0x55, 0x44, 0x89, 0xc8, 0xc3, 0x61, 0xa9, 0xfa, 0x7f, 0x5a,
	0x8c, 0x38, 0x22, 0xbb, 0x29, 0x18, 0x0a, 0x85, 0x22, 0xca, 0x50, 0x05, 0xb9, 0x0c, 0xd2, 0x82,
	0x89, 0x32, 0x48, 0x25, 0xf2, 0x28, 0x63, 0x8a, 0x85, 0xb1, 0xc8, 0x87, 0xa5, 0xba, 0x66, 0x0f,
	0x8a, 0x23, 0x6d, 0x42, 0xdc, 0x4b, 0xe7, 0x98, 0xa6, 0xf3, 0x41, 0x96, 0xa1, 0xd7, 0x32, 0x79,
	0x9d, 0xed, 0xca, 0x6f, 0x23, 0x4f, 0xb9, 0x98, 0x73, 0xa4, 0xfb, 0xc9, 0xba, 0xf6, 0xed, 0x4c,
	0xa6, 0x85, 0x77, 0xd0, 0x25, 0xbd, 0xd3, 0xa6, 0x36, 0xaf, 0x01, 0x6d, 0x78, 0xfc, 0xbc, 0x58,
	0x83, 0xb5, 0x5c, 0x83, 0xb5, 0x5b, 0x03, 0x79, 0xd1, 0x40, 0xde, 0x34, 0x90, 0x0f, 0x0d, 0x64,
	0xa1, 0x81, 0x2c, 0x35, 0x90, 0x2f, 0x0d, 0xe4, 0x5b, 0x83, 0xb5, 0xd3, 0x40, 0x5e, 0x37, 0x60,
	0x2d, 0x36, 0x60, 0x2d, 0x37, 0x60, 0xdd, 0x0d, 0xfe, 0x79, 0x8f, 0x68, 0xff, 0x29, 0xfd, 0xbd,
	0x48, 0x8e, 0xcc, 0x43, 0x5f, 0xfd, 0x04, 0x00, 0x00, 0xff, 0xff, 0x11, 0x99, 0xc6, 0x7b, 0xb0,
	0x01, 0x00, 0x00,
}

func (this *RewardTx) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*RewardTx)
	if !ok {
		that2, ok := that.(RewardTx)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Round != that1.Round {
		return false
	}
	{
		__caster := &github_com_TerraDharitri_drt_go_chain_core_data.BigIntCaster{}
		if !__caster.Equal(this.Value, that1.Value) {
			return false
		}
	}
	if !bytes.Equal(this.RcvAddr, that1.RcvAddr) {
		return false
	}
	if this.Epoch != that1.Epoch {
		return false
	}
	return true
}
func (this *RewardTx) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 8)
	s = append(s, "&rewardTx.RewardTx{")
	s = append(s, "Round: "+fmt.Sprintf("%#v", this.Round)+",\n")
	s = append(s, "Value: "+fmt.Sprintf("%#v", this.Value)+",\n")
	s = append(s, "RcvAddr: "+fmt.Sprintf("%#v", this.RcvAddr)+",\n")
	s = append(s, "Epoch: "+fmt.Sprintf("%#v", this.Epoch)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringRewardTx(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *RewardTx) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RewardTx) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RewardTx) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.RcvAddr) > 0 {
		i -= len(m.RcvAddr)
		copy(dAtA[i:], m.RcvAddr)
		i = encodeVarintRewardTx(dAtA, i, uint64(len(m.RcvAddr)))
		i--
		dAtA[i] = 0x22
	}
	{
		__caster := &github_com_TerraDharitri_drt_go_chain_core_data.BigIntCaster{}
		size := __caster.Size(m.Value)
		i -= size
		if _, err := __caster.MarshalTo(m.Value, dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintRewardTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if m.Epoch != 0 {
		i = encodeVarintRewardTx(dAtA, i, uint64(m.Epoch))
		i--
		dAtA[i] = 0x10
	}
	if m.Round != 0 {
		i = encodeVarintRewardTx(dAtA, i, uint64(m.Round))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintRewardTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovRewardTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RewardTx) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Round != 0 {
		n += 1 + sovRewardTx(uint64(m.Round))
	}
	if m.Epoch != 0 {
		n += 1 + sovRewardTx(uint64(m.Epoch))
	}
	{
		__caster := &github_com_TerraDharitri_drt_go_chain_core_data.BigIntCaster{}
		l = __caster.Size(m.Value)
		n += 1 + l + sovRewardTx(uint64(l))
	}
	l = len(m.RcvAddr)
	if l > 0 {
		n += 1 + l + sovRewardTx(uint64(l))
	}
	return n
}

func sovRewardTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRewardTx(x uint64) (n int) {
	return sovRewardTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *RewardTx) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&RewardTx{`,
		`Round:` + fmt.Sprintf("%v", this.Round) + `,`,
		`Epoch:` + fmt.Sprintf("%v", this.Epoch) + `,`,
		`Value:` + fmt.Sprintf("%v", this.Value) + `,`,
		`RcvAddr:` + fmt.Sprintf("%v", this.RcvAddr) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringRewardTx(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *RewardTx) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRewardTx
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
			return fmt.Errorf("proto: RewardTx: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RewardTx: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Round", wireType)
			}
			m.Round = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRewardTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Round |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Epoch", wireType)
			}
			m.Epoch = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRewardTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Epoch |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRewardTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthRewardTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthRewardTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			{
				__caster := &github_com_TerraDharitri_drt_go_chain_core_data.BigIntCaster{}
				if tmp, err := __caster.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
					return err
				} else {
					m.Value = tmp
				}
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RcvAddr", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRewardTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthRewardTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthRewardTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RcvAddr = append(m.RcvAddr[:0], dAtA[iNdEx:postIndex]...)
			if m.RcvAddr == nil {
				m.RcvAddr = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRewardTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRewardTx
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
func skipRewardTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRewardTx
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
					return 0, ErrIntOverflowRewardTx
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
					return 0, ErrIntOverflowRewardTx
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
				return 0, ErrInvalidLengthRewardTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRewardTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRewardTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRewardTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRewardTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRewardTx = fmt.Errorf("proto: unexpected end of group")
)
