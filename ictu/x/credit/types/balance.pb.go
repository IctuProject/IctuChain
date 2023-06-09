// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ictu/credit/balance.proto

package types

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

type Balance struct {
	Uid               string `protobuf:"bytes,1,opt,name=uid,proto3" json:"uid,omitempty"`
	IdContract        string `protobuf:"bytes,2,opt,name=idContract,proto3" json:"idContract,omitempty"`
	Requester         string `protobuf:"bytes,3,opt,name=requester,proto3" json:"requester,omitempty"`
	Credited          uint64 `protobuf:"varint,4,opt,name=credited,proto3" json:"credited,omitempty"`
	Returned          uint64 `protobuf:"varint,5,opt,name=returned,proto3" json:"returned,omitempty"`
	ExpirationDate    string `protobuf:"bytes,6,opt,name=expirationDate,proto3" json:"expirationDate,omitempty"`
	ReturnPeriodicity string `protobuf:"bytes,7,opt,name=returnPeriodicity,proto3" json:"returnPeriodicity,omitempty"`
}

func (m *Balance) Reset()         { *m = Balance{} }
func (m *Balance) String() string { return proto.CompactTextString(m) }
func (*Balance) ProtoMessage()    {}
func (*Balance) Descriptor() ([]byte, []int) {
	return fileDescriptor_4272dd77ba8f67de, []int{0}
}
func (m *Balance) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Balance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Balance.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Balance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Balance.Merge(m, src)
}
func (m *Balance) XXX_Size() int {
	return m.Size()
}
func (m *Balance) XXX_DiscardUnknown() {
	xxx_messageInfo_Balance.DiscardUnknown(m)
}

var xxx_messageInfo_Balance proto.InternalMessageInfo

func (m *Balance) GetUid() string {
	if m != nil {
		return m.Uid
	}
	return ""
}

func (m *Balance) GetIdContract() string {
	if m != nil {
		return m.IdContract
	}
	return ""
}

func (m *Balance) GetRequester() string {
	if m != nil {
		return m.Requester
	}
	return ""
}

func (m *Balance) GetCredited() uint64 {
	if m != nil {
		return m.Credited
	}
	return 0
}

func (m *Balance) GetReturned() uint64 {
	if m != nil {
		return m.Returned
	}
	return 0
}

func (m *Balance) GetExpirationDate() string {
	if m != nil {
		return m.ExpirationDate
	}
	return ""
}

func (m *Balance) GetReturnPeriodicity() string {
	if m != nil {
		return m.ReturnPeriodicity
	}
	return ""
}

func init() {
	proto.RegisterType((*Balance)(nil), "ictu.credit.Balance")
}

func init() { proto.RegisterFile("ictu/credit/balance.proto", fileDescriptor_4272dd77ba8f67de) }

var fileDescriptor_4272dd77ba8f67de = []byte{
	// 244 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0xd0, 0xc1, 0x4a, 0x03, 0x31,
	0x10, 0x06, 0xe0, 0x8d, 0xad, 0xad, 0x1d, 0x41, 0x34, 0x5e, 0xa2, 0x48, 0x28, 0x1e, 0xa4, 0x07,
	0x6d, 0x0f, 0xbe, 0x41, 0xf5, 0x01, 0xa4, 0x47, 0x6f, 0x69, 0x32, 0x87, 0x01, 0xd9, 0xac, 0xd3,
	0x59, 0x68, 0xdf, 0xc2, 0xc7, 0xf2, 0xd8, 0xa3, 0x47, 0xd9, 0x3d, 0xf8, 0x1a, 0xb2, 0x09, 0x55,
	0xb1, 0xb7, 0xcc, 0xff, 0x25, 0x19, 0xf8, 0xe1, 0x82, 0xbc, 0xd4, 0x33, 0xcf, 0x18, 0x48, 0x66,
	0x4b, 0xf7, 0xe2, 0x4a, 0x8f, 0xd3, 0x8a, 0xa3, 0x44, 0x7d, 0xdc, 0xd1, 0x34, 0xd3, 0xf5, 0x97,
	0x82, 0xe1, 0x3c, 0xb3, 0x3e, 0x85, 0x5e, 0x4d, 0xc1, 0xa8, 0xb1, 0x9a, 0x8c, 0x16, 0xdd, 0x51,
	0x5b, 0x00, 0x0a, 0x0f, 0xb1, 0x14, 0x76, 0x5e, 0xcc, 0x41, 0x82, 0x3f, 0x89, 0xbe, 0x82, 0x11,
	0xe3, 0x6b, 0x8d, 0x2b, 0x41, 0x36, 0xbd, 0xc4, 0xbf, 0x81, 0xbe, 0x84, 0xa3, 0xbc, 0x05, 0x83,
	0xe9, 0x8f, 0xd5, 0xa4, 0xbf, 0xf8, 0x99, 0x3b, 0x63, 0x94, 0x9a, 0x4b, 0x0c, 0xe6, 0x30, 0xdb,
	0x6e, 0xd6, 0x37, 0x70, 0x82, 0xeb, 0x8a, 0xd8, 0x09, 0xc5, 0xf2, 0xd1, 0x09, 0x9a, 0x41, 0xfa,
	0xfa, 0x5f, 0xaa, 0x6f, 0xe1, 0x2c, 0xbf, 0x79, 0x42, 0xa6, 0x18, 0xc8, 0x93, 0x6c, 0xcc, 0x30,
	0x5d, 0xdd, 0x87, 0xf9, 0xdd, 0x7b, 0x63, 0xd5, 0xb6, 0xb1, 0xea, 0xb3, 0xb1, 0xea, 0xad, 0xb5,
	0xc5, 0xb6, 0xb5, 0xc5, 0x47, 0x6b, 0x8b, 0xe7, 0xf3, 0xd4, 0xd5, 0x7a, 0xd7, 0x96, 0x6c, 0x2a,
	0x5c, 0x2d, 0x07, 0xa9, 0xac, 0xfb, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1a, 0x16, 0x4a, 0x72,
	0x49, 0x01, 0x00, 0x00,
}

func (m *Balance) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Balance) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Balance) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ReturnPeriodicity) > 0 {
		i -= len(m.ReturnPeriodicity)
		copy(dAtA[i:], m.ReturnPeriodicity)
		i = encodeVarintBalance(dAtA, i, uint64(len(m.ReturnPeriodicity)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.ExpirationDate) > 0 {
		i -= len(m.ExpirationDate)
		copy(dAtA[i:], m.ExpirationDate)
		i = encodeVarintBalance(dAtA, i, uint64(len(m.ExpirationDate)))
		i--
		dAtA[i] = 0x32
	}
	if m.Returned != 0 {
		i = encodeVarintBalance(dAtA, i, uint64(m.Returned))
		i--
		dAtA[i] = 0x28
	}
	if m.Credited != 0 {
		i = encodeVarintBalance(dAtA, i, uint64(m.Credited))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Requester) > 0 {
		i -= len(m.Requester)
		copy(dAtA[i:], m.Requester)
		i = encodeVarintBalance(dAtA, i, uint64(len(m.Requester)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.IdContract) > 0 {
		i -= len(m.IdContract)
		copy(dAtA[i:], m.IdContract)
		i = encodeVarintBalance(dAtA, i, uint64(len(m.IdContract)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Uid) > 0 {
		i -= len(m.Uid)
		copy(dAtA[i:], m.Uid)
		i = encodeVarintBalance(dAtA, i, uint64(len(m.Uid)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintBalance(dAtA []byte, offset int, v uint64) int {
	offset -= sovBalance(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Balance) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Uid)
	if l > 0 {
		n += 1 + l + sovBalance(uint64(l))
	}
	l = len(m.IdContract)
	if l > 0 {
		n += 1 + l + sovBalance(uint64(l))
	}
	l = len(m.Requester)
	if l > 0 {
		n += 1 + l + sovBalance(uint64(l))
	}
	if m.Credited != 0 {
		n += 1 + sovBalance(uint64(m.Credited))
	}
	if m.Returned != 0 {
		n += 1 + sovBalance(uint64(m.Returned))
	}
	l = len(m.ExpirationDate)
	if l > 0 {
		n += 1 + l + sovBalance(uint64(l))
	}
	l = len(m.ReturnPeriodicity)
	if l > 0 {
		n += 1 + l + sovBalance(uint64(l))
	}
	return n
}

func sovBalance(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBalance(x uint64) (n int) {
	return sovBalance(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Balance) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBalance
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
			return fmt.Errorf("proto: Balance: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Balance: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBalance
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
				return ErrInvalidLengthBalance
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBalance
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Uid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IdContract", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBalance
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
				return ErrInvalidLengthBalance
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBalance
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IdContract = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Requester", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBalance
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
				return ErrInvalidLengthBalance
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBalance
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Requester = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Credited", wireType)
			}
			m.Credited = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBalance
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Credited |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Returned", wireType)
			}
			m.Returned = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBalance
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Returned |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpirationDate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBalance
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
				return ErrInvalidLengthBalance
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBalance
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ExpirationDate = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReturnPeriodicity", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBalance
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
				return ErrInvalidLengthBalance
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBalance
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ReturnPeriodicity = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBalance(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBalance
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
func skipBalance(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBalance
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
					return 0, ErrIntOverflowBalance
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
					return 0, ErrIntOverflowBalance
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
				return 0, ErrInvalidLengthBalance
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBalance
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBalance
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBalance        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBalance          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBalance = fmt.Errorf("proto: unexpected end of group")
)
