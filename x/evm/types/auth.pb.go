// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: minievm/evm/v1/auth.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	types "github.com/cosmos/cosmos-sdk/x/auth/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

// ContractAccount defines an account of contract.
type ContractAccount struct {
	*types.BaseAccount `protobuf:"bytes,1,opt,name=base_account,json=baseAccount,proto3,embedded=base_account" json:"base_account,omitempty"`
}

func (m *ContractAccount) Reset()         { *m = ContractAccount{} }
func (m *ContractAccount) String() string { return proto.CompactTextString(m) }
func (*ContractAccount) ProtoMessage()    {}
func (*ContractAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_a01464d2d75c2977, []int{0}
}
func (m *ContractAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ContractAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ContractAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ContractAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContractAccount.Merge(m, src)
}
func (m *ContractAccount) XXX_Size() int {
	return m.Size()
}
func (m *ContractAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_ContractAccount.DiscardUnknown(m)
}

var xxx_messageInfo_ContractAccount proto.InternalMessageInfo

// ShorthandAccount defines an account of shorthand address
// which is used to store the original long address (32bytes).
//
// Also it is used to check the existence of the account before
// creating a new account.
type ShorthandAccount struct {
	*types.BaseAccount `protobuf:"bytes,1,opt,name=base_account,json=baseAccount,proto3,embedded=base_account" json:"base_account,omitempty"`
	OriginalAddress    string `protobuf:"bytes,2,opt,name=original_address,json=originalAddress,proto3" json:"original_address,omitempty"`
}

func (m *ShorthandAccount) Reset()         { *m = ShorthandAccount{} }
func (m *ShorthandAccount) String() string { return proto.CompactTextString(m) }
func (*ShorthandAccount) ProtoMessage()    {}
func (*ShorthandAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_a01464d2d75c2977, []int{1}
}
func (m *ShorthandAccount) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ShorthandAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ShorthandAccount.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ShorthandAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShorthandAccount.Merge(m, src)
}
func (m *ShorthandAccount) XXX_Size() int {
	return m.Size()
}
func (m *ShorthandAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_ShorthandAccount.DiscardUnknown(m)
}

var xxx_messageInfo_ShorthandAccount proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ContractAccount)(nil), "minievm.evm.v1.ContractAccount")
	proto.RegisterType((*ShorthandAccount)(nil), "minievm.evm.v1.ShorthandAccount")
}

func init() { proto.RegisterFile("minievm/evm/v1/auth.proto", fileDescriptor_a01464d2d75c2977) }

var fileDescriptor_a01464d2d75c2977 = []byte{
	// 308 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xcc, 0xcd, 0xcc, 0xcb,
	0x4c, 0x2d, 0xcb, 0xd5, 0x07, 0xe1, 0x32, 0x43, 0xfd, 0xc4, 0xd2, 0x92, 0x0c, 0xbd, 0x82, 0xa2,
	0xfc, 0x92, 0x7c, 0x21, 0x3e, 0xa8, 0x94, 0x1e, 0x08, 0x97, 0x19, 0x4a, 0x09, 0x26, 0xe6, 0x66,
	0xe6, 0xe5, 0xeb, 0x83, 0x49, 0x88, 0x12, 0x29, 0x91, 0xf4, 0xfc, 0xf4, 0x7c, 0x30, 0x53, 0x1f,
	0xc4, 0x82, 0x8a, 0xca, 0x25, 0xe7, 0x17, 0xe7, 0xe6, 0x17, 0x83, 0xcd, 0xd2, 0x2f, 0x33, 0x4c,
	0x4a, 0x2d, 0x49, 0x44, 0x36, 0x58, 0xa9, 0x8a, 0x8b, 0xdf, 0x39, 0x3f, 0xaf, 0xa4, 0x28, 0x31,
	0xb9, 0xc4, 0x31, 0x39, 0x39, 0xbf, 0x34, 0xaf, 0x44, 0xc8, 0x93, 0x8b, 0x27, 0x29, 0xb1, 0x38,
	0x35, 0x3e, 0x11, 0xc2, 0x97, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x36, 0x52, 0xd0, 0x83, 0x98, 0xa4,
	0x07, 0xd6, 0x0c, 0x35, 0x49, 0xcf, 0x29, 0xb1, 0x38, 0x15, 0xaa, 0xcf, 0x89, 0xe5, 0xc2, 0x3d,
	0x79, 0xc6, 0x20, 0xee, 0x24, 0x84, 0x90, 0x95, 0x4c, 0xc7, 0x02, 0x79, 0x86, 0xae, 0xe7, 0x1b,
	0xb4, 0x84, 0x41, 0x5e, 0x42, 0xb3, 0x48, 0x69, 0x39, 0x23, 0x97, 0x40, 0x70, 0x46, 0x7e, 0x51,
	0x49, 0x46, 0x62, 0x5e, 0x0a, 0xf5, 0x6d, 0x17, 0xd2, 0xe4, 0x12, 0xc8, 0x2f, 0xca, 0x4c, 0xcf,
	0xcc, 0x4b, 0xcc, 0x89, 0x4f, 0x4c, 0x49, 0x29, 0x4a, 0x2d, 0x2e, 0x96, 0x60, 0x52, 0x60, 0xd4,
	0xe0, 0x0c, 0xe2, 0x87, 0x89, 0x3b, 0x42, 0x84, 0xad, 0x64, 0x61, 0x0e, 0x15, 0x01, 0x39, 0x14,
	0xdd, 0x51, 0x4e, 0x2e, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c,
	0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10, 0xa5, 0x95,
	0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x9f, 0x99, 0x97, 0x59, 0x92, 0x99,
	0xa8, 0x9b, 0x93, 0x98, 0x54, 0xac, 0x0f, 0x8b, 0xca, 0x0a, 0x70, 0x64, 0x96, 0x54, 0x16, 0xa4,
	0x16, 0x27, 0xb1, 0x81, 0x83, 0xdc, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x32, 0x2a, 0x1a, 0x97,
	0xe8, 0x01, 0x00, 0x00,
}

func (m *ContractAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ContractAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ContractAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.BaseAccount != nil {
		{
			size, err := m.BaseAccount.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAuth(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ShorthandAccount) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ShorthandAccount) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ShorthandAccount) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.OriginalAddress) > 0 {
		i -= len(m.OriginalAddress)
		copy(dAtA[i:], m.OriginalAddress)
		i = encodeVarintAuth(dAtA, i, uint64(len(m.OriginalAddress)))
		i--
		dAtA[i] = 0x12
	}
	if m.BaseAccount != nil {
		{
			size, err := m.BaseAccount.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintAuth(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintAuth(dAtA []byte, offset int, v uint64) int {
	offset -= sovAuth(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ContractAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BaseAccount != nil {
		l = m.BaseAccount.Size()
		n += 1 + l + sovAuth(uint64(l))
	}
	return n
}

func (m *ShorthandAccount) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BaseAccount != nil {
		l = m.BaseAccount.Size()
		n += 1 + l + sovAuth(uint64(l))
	}
	l = len(m.OriginalAddress)
	if l > 0 {
		n += 1 + l + sovAuth(uint64(l))
	}
	return n
}

func sovAuth(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAuth(x uint64) (n int) {
	return sovAuth(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ContractAccount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuth
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
			return fmt.Errorf("proto: ContractAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ContractAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseAccount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuth
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthAuth
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAuth
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BaseAccount == nil {
				m.BaseAccount = &types.BaseAccount{}
			}
			if err := m.BaseAccount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAuth(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAuth
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
func (m *ShorthandAccount) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAuth
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
			return fmt.Errorf("proto: ShorthandAccount: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ShorthandAccount: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseAccount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuth
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthAuth
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthAuth
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.BaseAccount == nil {
				m.BaseAccount = &types.BaseAccount{}
			}
			if err := m.BaseAccount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OriginalAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAuth
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
				return ErrInvalidLengthAuth
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAuth
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OriginalAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAuth(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAuth
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
func skipAuth(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAuth
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
					return 0, ErrIntOverflowAuth
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
					return 0, ErrIntOverflowAuth
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
				return 0, ErrInvalidLengthAuth
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAuth
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAuth
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAuth        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAuth          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAuth = fmt.Errorf("proto: unexpected end of group")
)
