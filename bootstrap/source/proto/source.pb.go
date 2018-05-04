// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: source.proto

/*
	Package proto is a generated protocol buffer package.

	It is generated from these files:
		source.proto

	It has these top-level messages:
		Fixture
		Source
		Position
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type Fixture struct {
	Source           *Source `protobuf:"bytes,1,req,name=source" json:"source,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Fixture) Reset()                    { *m = Fixture{} }
func (m *Fixture) String() string            { return proto1.CompactTextString(m) }
func (*Fixture) ProtoMessage()               {}
func (*Fixture) Descriptor() ([]byte, []int) { return fileDescriptorSource, []int{0} }

func (m *Fixture) GetSource() *Source {
	if m != nil {
		return m.Source
	}
	return nil
}

type Source struct {
	File             *string   `protobuf:"bytes,1,req,name=file" json:"file,omitempty"`
	Start            *Position `protobuf:"bytes,2,req,name=start" json:"start,omitempty"`
	End              *Position `protobuf:"bytes,3,req,name=end" json:"end,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *Source) Reset()                    { *m = Source{} }
func (m *Source) String() string            { return proto1.CompactTextString(m) }
func (*Source) ProtoMessage()               {}
func (*Source) Descriptor() ([]byte, []int) { return fileDescriptorSource, []int{1} }

func (m *Source) GetFile() string {
	if m != nil && m.File != nil {
		return *m.File
	}
	return ""
}

func (m *Source) GetStart() *Position {
	if m != nil {
		return m.Start
	}
	return nil
}

func (m *Source) GetEnd() *Position {
	if m != nil {
		return m.End
	}
	return nil
}

type Position struct {
	Offset           *int64 `protobuf:"varint,1,req,name=offset" json:"offset,omitempty"`
	Line             *int64 `protobuf:"varint,2,req,name=line" json:"line,omitempty"`
	Column           *int64 `protobuf:"varint,3,req,name=column" json:"column,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Position) Reset()                    { *m = Position{} }
func (m *Position) String() string            { return proto1.CompactTextString(m) }
func (*Position) ProtoMessage()               {}
func (*Position) Descriptor() ([]byte, []int) { return fileDescriptorSource, []int{2} }

func (m *Position) GetOffset() int64 {
	if m != nil && m.Offset != nil {
		return *m.Offset
	}
	return 0
}

func (m *Position) GetLine() int64 {
	if m != nil && m.Line != nil {
		return *m.Line
	}
	return 0
}

func (m *Position) GetColumn() int64 {
	if m != nil && m.Column != nil {
		return *m.Column
	}
	return 0
}

func init() {
	proto1.RegisterType((*Fixture)(nil), "ko.bootstrap.source.Fixture")
	proto1.RegisterType((*Source)(nil), "ko.bootstrap.source.Source")
	proto1.RegisterType((*Position)(nil), "ko.bootstrap.source.Position")
}
func (m *Fixture) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Fixture) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Source == nil {
		return 0, new(proto1.RequiredNotSetError)
	} else {
		dAtA[i] = 0xa
		i++
		i = encodeVarintSource(dAtA, i, uint64(m.Source.Size()))
		n1, err := m.Source.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *Source) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Source) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.File == nil {
		return 0, new(proto1.RequiredNotSetError)
	} else {
		dAtA[i] = 0xa
		i++
		i = encodeVarintSource(dAtA, i, uint64(len(*m.File)))
		i += copy(dAtA[i:], *m.File)
	}
	if m.Start == nil {
		return 0, new(proto1.RequiredNotSetError)
	} else {
		dAtA[i] = 0x12
		i++
		i = encodeVarintSource(dAtA, i, uint64(m.Start.Size()))
		n2, err := m.Start.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.End == nil {
		return 0, new(proto1.RequiredNotSetError)
	} else {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintSource(dAtA, i, uint64(m.End.Size()))
		n3, err := m.End.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *Position) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Position) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Offset == nil {
		return 0, new(proto1.RequiredNotSetError)
	} else {
		dAtA[i] = 0x8
		i++
		i = encodeVarintSource(dAtA, i, uint64(*m.Offset))
	}
	if m.Line == nil {
		return 0, new(proto1.RequiredNotSetError)
	} else {
		dAtA[i] = 0x10
		i++
		i = encodeVarintSource(dAtA, i, uint64(*m.Line))
	}
	if m.Column == nil {
		return 0, new(proto1.RequiredNotSetError)
	} else {
		dAtA[i] = 0x18
		i++
		i = encodeVarintSource(dAtA, i, uint64(*m.Column))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintSource(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Fixture) Size() (n int) {
	var l int
	_ = l
	if m.Source != nil {
		l = m.Source.Size()
		n += 1 + l + sovSource(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Source) Size() (n int) {
	var l int
	_ = l
	if m.File != nil {
		l = len(*m.File)
		n += 1 + l + sovSource(uint64(l))
	}
	if m.Start != nil {
		l = m.Start.Size()
		n += 1 + l + sovSource(uint64(l))
	}
	if m.End != nil {
		l = m.End.Size()
		n += 1 + l + sovSource(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Position) Size() (n int) {
	var l int
	_ = l
	if m.Offset != nil {
		n += 1 + sovSource(uint64(*m.Offset))
	}
	if m.Line != nil {
		n += 1 + sovSource(uint64(*m.Line))
	}
	if m.Column != nil {
		n += 1 + sovSource(uint64(*m.Column))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovSource(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozSource(x uint64) (n int) {
	return sovSource(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Fixture) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSource
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Fixture: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Fixture: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Source", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSource
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSource
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Source == nil {
				m.Source = &Source{}
			}
			if err := m.Source.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000001)
		default:
			iNdEx = preIndex
			skippy, err := skipSource(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSource
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return new(proto1.RequiredNotSetError)
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Source) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSource
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Source: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Source: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field File", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSource
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthSource
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			s := string(dAtA[iNdEx:postIndex])
			m.File = &s
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Start", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSource
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSource
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Start == nil {
				m.Start = &Position{}
			}
			if err := m.Start.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000002)
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field End", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSource
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthSource
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.End == nil {
				m.End = &Position{}
			}
			if err := m.End.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
			hasFields[0] |= uint64(0x00000004)
		default:
			iNdEx = preIndex
			skippy, err := skipSource(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSource
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return new(proto1.RequiredNotSetError)
	}
	if hasFields[0]&uint64(0x00000002) == 0 {
		return new(proto1.RequiredNotSetError)
	}
	if hasFields[0]&uint64(0x00000004) == 0 {
		return new(proto1.RequiredNotSetError)
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Position) Unmarshal(dAtA []byte) error {
	var hasFields [1]uint64
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSource
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Position: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Position: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Offset", wireType)
			}
			var v int64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSource
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Offset = &v
			hasFields[0] |= uint64(0x00000001)
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Line", wireType)
			}
			var v int64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSource
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Line = &v
			hasFields[0] |= uint64(0x00000002)
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Column", wireType)
			}
			var v int64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSource
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Column = &v
			hasFields[0] |= uint64(0x00000004)
		default:
			iNdEx = preIndex
			skippy, err := skipSource(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSource
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}
	if hasFields[0]&uint64(0x00000001) == 0 {
		return new(proto1.RequiredNotSetError)
	}
	if hasFields[0]&uint64(0x00000002) == 0 {
		return new(proto1.RequiredNotSetError)
	}
	if hasFields[0]&uint64(0x00000004) == 0 {
		return new(proto1.RequiredNotSetError)
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipSource(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSource
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
					return 0, ErrIntOverflowSource
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowSource
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthSource
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowSource
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipSource(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthSource = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSource   = fmt.Errorf("proto: integer overflow")
)

func init() { proto1.RegisterFile("source.proto", fileDescriptorSource) }

var fileDescriptorSource = []byte{
	// 222 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x29, 0xce, 0x2f, 0x2d,
	0x4a, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0xce, 0xce, 0xd7, 0x4b, 0xca, 0xcf,
	0x2f, 0x29, 0x2e, 0x29, 0x4a, 0x2c, 0xd0, 0x83, 0x48, 0x29, 0xd9, 0x71, 0xb1, 0xbb, 0x65, 0x56,
	0x94, 0x94, 0x16, 0xa5, 0x0a, 0x19, 0x73, 0xb1, 0x41, 0x04, 0x25, 0x18, 0x15, 0x98, 0x34, 0xb8,
	0x8d, 0xa4, 0xf5, 0xb0, 0x68, 0xd0, 0x0b, 0x06, 0x53, 0x41, 0x50, 0xa5, 0x4a, 0x4d, 0x8c, 0x5c,
	0x6c, 0x10, 0x21, 0x21, 0x21, 0x2e, 0x96, 0xb4, 0xcc, 0x1c, 0x88, 0x6e, 0xce, 0x20, 0x30, 0x5b,
	0xc8, 0x98, 0x8b, 0xb5, 0xb8, 0x24, 0xb1, 0xa8, 0x44, 0x82, 0x09, 0x6c, 0xa4, 0x2c, 0x56, 0x23,
	0x03, 0xf2, 0x8b, 0x33, 0x4b, 0x32, 0xf3, 0xf3, 0x82, 0x20, 0x6a, 0x85, 0xf4, 0xb9, 0x98, 0x53,
	0xf3, 0x52, 0x24, 0x98, 0x89, 0xd1, 0x02, 0x52, 0xa9, 0xe4, 0xc7, 0xc5, 0x01, 0x13, 0x10, 0x12,
	0xe3, 0x62, 0xcb, 0x4f, 0x4b, 0x2b, 0x4e, 0x2d, 0x01, 0xbb, 0x83, 0x39, 0x08, 0xca, 0x03, 0xb9,
	0x2e, 0x27, 0x33, 0x2f, 0x15, 0xec, 0x10, 0xe6, 0x20, 0x30, 0x1b, 0xa4, 0x36, 0x39, 0x3f, 0xa7,
	0x34, 0x37, 0x0f, 0x6c, 0x17, 0x73, 0x10, 0x94, 0xe7, 0x24, 0x7e, 0xe2, 0x91, 0x1c, 0xe3, 0x85,
	0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0xce, 0x78, 0x2c, 0xc7, 0x10, 0xc5, 0x0a, 0x0e, 0x42,
	0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd0, 0x73, 0xa7, 0x5f, 0x51, 0x01, 0x00, 0x00,
}
