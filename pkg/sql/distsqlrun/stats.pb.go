// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sql/distsqlrun/stats.proto

package distsqlrun

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import time "time"

import types "github.com/gogo/protobuf/types"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// InputStats represents the stats collected from an input.
type InputStats struct {
	// num_rows is the number of rows received from the input.
	NumRows int64 `protobuf:"varint,1,opt,name=num_rows,json=numRows,proto3" json:"num_rows,omitempty"`
	// Duration in nanoseconds of the cumulative time spent stalled.
	StallTime time.Duration `protobuf:"bytes,8,opt,name=stall_time,json=stallTime,stdduration" json:"stall_time"`
}

func (m *InputStats) Reset()                    { *m = InputStats{} }
func (m *InputStats) String() string            { return proto.CompactTextString(m) }
func (*InputStats) ProtoMessage()               {}
func (*InputStats) Descriptor() ([]byte, []int) { return fileDescriptorStats, []int{0} }

// TableReaderStats are the stats collected during a tableReader run.
type TableReaderStats struct {
	InputStats InputStats `protobuf:"bytes,1,opt,name=input_stats,json=inputStats" json:"input_stats"`
}

func (m *TableReaderStats) Reset()                    { *m = TableReaderStats{} }
func (m *TableReaderStats) String() string            { return proto.CompactTextString(m) }
func (*TableReaderStats) ProtoMessage()               {}
func (*TableReaderStats) Descriptor() ([]byte, []int) { return fileDescriptorStats, []int{1} }

// HashJoinerStats are the stats collected during a hashJoiner run.
type HashJoinerStats struct {
	LeftInputStats   InputStats `protobuf:"bytes,1,opt,name=left_input_stats,json=leftInputStats" json:"left_input_stats"`
	RightInputStats  InputStats `protobuf:"bytes,2,opt,name=right_input_stats,json=rightInputStats" json:"right_input_stats"`
	StoredSide       string     `protobuf:"bytes,3,opt,name=stored_side,json=storedSide,proto3" json:"stored_side,omitempty"`
	MaxAllocatedMem  int64      `protobuf:"varint,4,opt,name=max_allocated_mem,json=maxAllocatedMem,proto3" json:"max_allocated_mem,omitempty"`
	MaxAllocatedDisk int64      `protobuf:"varint,5,opt,name=max_allocated_disk,json=maxAllocatedDisk,proto3" json:"max_allocated_disk,omitempty"`
}

func (m *HashJoinerStats) Reset()                    { *m = HashJoinerStats{} }
func (m *HashJoinerStats) String() string            { return proto.CompactTextString(m) }
func (*HashJoinerStats) ProtoMessage()               {}
func (*HashJoinerStats) Descriptor() ([]byte, []int) { return fileDescriptorStats, []int{2} }

// AggregatorStats are the stats collected during an aggregator run.
type AggregatorStats struct {
	InputStats      InputStats `protobuf:"bytes,1,opt,name=input_stats,json=inputStats" json:"input_stats"`
	MaxAllocatedMem int64      `protobuf:"varint,2,opt,name=max_allocated_mem,json=maxAllocatedMem,proto3" json:"max_allocated_mem,omitempty"`
}

func (m *AggregatorStats) Reset()                    { *m = AggregatorStats{} }
func (m *AggregatorStats) String() string            { return proto.CompactTextString(m) }
func (*AggregatorStats) ProtoMessage()               {}
func (*AggregatorStats) Descriptor() ([]byte, []int) { return fileDescriptorStats, []int{3} }

// DistinctStats are the stats collected during a distinct run.
type DistinctStats struct {
	InputStats      InputStats `protobuf:"bytes,1,opt,name=input_stats,json=inputStats" json:"input_stats"`
	MaxAllocatedMem int64      `protobuf:"varint,2,opt,name=max_allocated_mem,json=maxAllocatedMem,proto3" json:"max_allocated_mem,omitempty"`
}

func (m *DistinctStats) Reset()                    { *m = DistinctStats{} }
func (m *DistinctStats) String() string            { return proto.CompactTextString(m) }
func (*DistinctStats) ProtoMessage()               {}
func (*DistinctStats) Descriptor() ([]byte, []int) { return fileDescriptorStats, []int{4} }

func init() {
	proto.RegisterType((*InputStats)(nil), "cockroach.sql.distsqlrun.InputStats")
	proto.RegisterType((*TableReaderStats)(nil), "cockroach.sql.distsqlrun.TableReaderStats")
	proto.RegisterType((*HashJoinerStats)(nil), "cockroach.sql.distsqlrun.HashJoinerStats")
	proto.RegisterType((*AggregatorStats)(nil), "cockroach.sql.distsqlrun.AggregatorStats")
	proto.RegisterType((*DistinctStats)(nil), "cockroach.sql.distsqlrun.DistinctStats")
}
func (m *InputStats) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InputStats) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.NumRows != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintStats(dAtA, i, uint64(m.NumRows))
	}
	dAtA[i] = 0x42
	i++
	i = encodeVarintStats(dAtA, i, uint64(types.SizeOfStdDuration(m.StallTime)))
	n1, err := types.StdDurationMarshalTo(m.StallTime, dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	return i, nil
}

func (m *TableReaderStats) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TableReaderStats) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintStats(dAtA, i, uint64(m.InputStats.Size()))
	n2, err := m.InputStats.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	return i, nil
}

func (m *HashJoinerStats) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HashJoinerStats) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintStats(dAtA, i, uint64(m.LeftInputStats.Size()))
	n3, err := m.LeftInputStats.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n3
	dAtA[i] = 0x12
	i++
	i = encodeVarintStats(dAtA, i, uint64(m.RightInputStats.Size()))
	n4, err := m.RightInputStats.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n4
	if len(m.StoredSide) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintStats(dAtA, i, uint64(len(m.StoredSide)))
		i += copy(dAtA[i:], m.StoredSide)
	}
	if m.MaxAllocatedMem != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintStats(dAtA, i, uint64(m.MaxAllocatedMem))
	}
	if m.MaxAllocatedDisk != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintStats(dAtA, i, uint64(m.MaxAllocatedDisk))
	}
	return i, nil
}

func (m *AggregatorStats) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AggregatorStats) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintStats(dAtA, i, uint64(m.InputStats.Size()))
	n5, err := m.InputStats.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n5
	if m.MaxAllocatedMem != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintStats(dAtA, i, uint64(m.MaxAllocatedMem))
	}
	return i, nil
}

func (m *DistinctStats) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DistinctStats) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintStats(dAtA, i, uint64(m.InputStats.Size()))
	n6, err := m.InputStats.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n6
	if m.MaxAllocatedMem != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintStats(dAtA, i, uint64(m.MaxAllocatedMem))
	}
	return i, nil
}

func encodeVarintStats(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *InputStats) Size() (n int) {
	var l int
	_ = l
	if m.NumRows != 0 {
		n += 1 + sovStats(uint64(m.NumRows))
	}
	l = types.SizeOfStdDuration(m.StallTime)
	n += 1 + l + sovStats(uint64(l))
	return n
}

func (m *TableReaderStats) Size() (n int) {
	var l int
	_ = l
	l = m.InputStats.Size()
	n += 1 + l + sovStats(uint64(l))
	return n
}

func (m *HashJoinerStats) Size() (n int) {
	var l int
	_ = l
	l = m.LeftInputStats.Size()
	n += 1 + l + sovStats(uint64(l))
	l = m.RightInputStats.Size()
	n += 1 + l + sovStats(uint64(l))
	l = len(m.StoredSide)
	if l > 0 {
		n += 1 + l + sovStats(uint64(l))
	}
	if m.MaxAllocatedMem != 0 {
		n += 1 + sovStats(uint64(m.MaxAllocatedMem))
	}
	if m.MaxAllocatedDisk != 0 {
		n += 1 + sovStats(uint64(m.MaxAllocatedDisk))
	}
	return n
}

func (m *AggregatorStats) Size() (n int) {
	var l int
	_ = l
	l = m.InputStats.Size()
	n += 1 + l + sovStats(uint64(l))
	if m.MaxAllocatedMem != 0 {
		n += 1 + sovStats(uint64(m.MaxAllocatedMem))
	}
	return n
}

func (m *DistinctStats) Size() (n int) {
	var l int
	_ = l
	l = m.InputStats.Size()
	n += 1 + l + sovStats(uint64(l))
	if m.MaxAllocatedMem != 0 {
		n += 1 + sovStats(uint64(m.MaxAllocatedMem))
	}
	return n
}

func sovStats(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozStats(x uint64) (n int) {
	return sovStats(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *InputStats) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStats
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
			return fmt.Errorf("proto: InputStats: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InputStats: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumRows", wireType)
			}
			m.NumRows = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NumRows |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StallTime", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStats
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
				return ErrInvalidLengthStats
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := types.StdDurationUnmarshal(&m.StallTime, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStats(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStats
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
func (m *TableReaderStats) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStats
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
			return fmt.Errorf("proto: TableReaderStats: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TableReaderStats: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InputStats", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStats
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
				return ErrInvalidLengthStats
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.InputStats.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipStats(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStats
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
func (m *HashJoinerStats) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStats
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
			return fmt.Errorf("proto: HashJoinerStats: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HashJoinerStats: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LeftInputStats", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStats
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
				return ErrInvalidLengthStats
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LeftInputStats.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RightInputStats", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStats
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
				return ErrInvalidLengthStats
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.RightInputStats.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StoredSide", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStats
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
				return ErrInvalidLengthStats
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StoredSide = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxAllocatedMem", wireType)
			}
			m.MaxAllocatedMem = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxAllocatedMem |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxAllocatedDisk", wireType)
			}
			m.MaxAllocatedDisk = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxAllocatedDisk |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipStats(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStats
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
func (m *AggregatorStats) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStats
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
			return fmt.Errorf("proto: AggregatorStats: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AggregatorStats: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InputStats", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStats
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
				return ErrInvalidLengthStats
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.InputStats.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxAllocatedMem", wireType)
			}
			m.MaxAllocatedMem = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxAllocatedMem |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipStats(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStats
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
func (m *DistinctStats) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStats
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
			return fmt.Errorf("proto: DistinctStats: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DistinctStats: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InputStats", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStats
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
				return ErrInvalidLengthStats
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.InputStats.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MaxAllocatedMem", wireType)
			}
			m.MaxAllocatedMem = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStats
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MaxAllocatedMem |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipStats(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthStats
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
func skipStats(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStats
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
					return 0, ErrIntOverflowStats
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
					return 0, ErrIntOverflowStats
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
				return 0, ErrInvalidLengthStats
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowStats
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
				next, err := skipStats(dAtA[start:])
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
	ErrInvalidLengthStats = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStats   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("sql/distsqlrun/stats.proto", fileDescriptorStats) }

var fileDescriptorStats = []byte{
	// 436 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xc4, 0x93, 0xbd, 0x8e, 0x13, 0x31,
	0x14, 0x85, 0xe3, 0xec, 0x02, 0xd9, 0x1b, 0x41, 0xb2, 0x23, 0x8a, 0xd9, 0x14, 0x93, 0x68, 0xb4,
	0x45, 0x84, 0xd0, 0x8c, 0xb4, 0x3c, 0xc1, 0x46, 0x29, 0xf8, 0x11, 0xcd, 0x6c, 0x44, 0x41, 0x63,
	0x39, 0x63, 0xaf, 0x63, 0xc5, 0x1e, 0x6f, 0x6c, 0x8f, 0x76, 0x1f, 0x01, 0x51, 0x51, 0xf2, 0x2c,
	0x3c, 0x41, 0x4a, 0x4a, 0x2a, 0x7e, 0x86, 0x17, 0x41, 0xe3, 0x21, 0x24, 0x2b, 0x85, 0x02, 0x21,
	0x44, 0x67, 0x5f, 0x1f, 0x7f, 0xe7, 0x5c, 0x5f, 0x19, 0x06, 0x76, 0x25, 0x53, 0x2a, 0xac, 0xb3,
	0x2b, 0x69, 0xca, 0x22, 0xb5, 0x8e, 0x38, 0x9b, 0x5c, 0x19, 0xed, 0x74, 0x10, 0xe6, 0x3a, 0x5f,
	0x1a, 0x4d, 0xf2, 0x45, 0x62, 0x57, 0x32, 0xd9, 0xaa, 0x06, 0x0f, 0xb9, 0xe6, 0xda, 0x8b, 0xd2,
	0x7a, 0xd5, 0xe8, 0x07, 0x11, 0xd7, 0x9a, 0x4b, 0x96, 0xfa, 0xdd, 0xbc, 0xbc, 0x4c, 0x69, 0x69,
	0x88, 0x13, 0xba, 0x68, 0xce, 0xe3, 0x25, 0xc0, 0xb3, 0xe2, 0xaa, 0x74, 0x17, 0xb5, 0x47, 0x70,
	0x02, 0x9d, 0xa2, 0x54, 0xd8, 0xe8, 0x6b, 0x1b, 0xa2, 0x11, 0x1a, 0x1f, 0x64, 0xf7, 0x8a, 0x52,
	0x65, 0xfa, 0xda, 0x06, 0x13, 0x00, 0xeb, 0x88, 0x94, 0xd8, 0x09, 0xc5, 0xc2, 0xce, 0x08, 0x8d,
	0xbb, 0x67, 0x27, 0x49, 0x43, 0x4f, 0x36, 0xf4, 0x64, 0xfa, 0x93, 0x3e, 0xe9, 0xac, 0x3f, 0x0f,
	0x5b, 0xef, 0xbf, 0x0c, 0x51, 0x76, 0xe4, 0xaf, 0xcd, 0x84, 0x62, 0x31, 0x86, 0xfe, 0x8c, 0xcc,
	0x25, 0xcb, 0x18, 0xa1, 0xcc, 0x34, 0x96, 0x2f, 0xa0, 0x2b, 0xea, 0x00, 0xd8, 0x77, 0xe9, 0x5d,
	0xbb, 0x67, 0xa7, 0xc9, 0xef, 0xda, 0x4c, 0xb6, 0x69, 0x27, 0x87, 0xb5, 0x47, 0x06, 0xe2, 0x57,
	0x25, 0xfe, 0xd0, 0x86, 0xde, 0x53, 0x62, 0x17, 0xcf, 0xb5, 0x28, 0x36, 0x06, 0x33, 0xe8, 0x4b,
	0x76, 0xe9, 0xf0, 0xdf, 0xb9, 0x3c, 0xa8, 0x19, 0x3b, 0x2f, 0xf5, 0x0a, 0x8e, 0x8d, 0xe0, 0x8b,
	0xdb, 0xd8, 0xf6, 0x1f, 0x63, 0x7b, 0x1e, 0xb2, 0xc3, 0x1d, 0x42, 0xd7, 0x3a, 0x6d, 0x18, 0xc5,
	0x56, 0x50, 0x16, 0x1e, 0x8c, 0xd0, 0xf8, 0x28, 0x83, 0xa6, 0x74, 0x21, 0x28, 0x0b, 0x1e, 0xc1,
	0xb1, 0x22, 0x37, 0x98, 0x48, 0xa9, 0x73, 0xe2, 0x18, 0xc5, 0x8a, 0xa9, 0xf0, 0xd0, 0xcf, 0xaa,
	0xa7, 0xc8, 0xcd, 0xf9, 0xa6, 0xfe, 0x92, 0xa9, 0xe0, 0x31, 0x04, 0xb7, 0xb5, 0x54, 0xd8, 0x65,
	0x78, 0xc7, 0x8b, 0xfb, 0xbb, 0xe2, 0xa9, 0xb0, 0xcb, 0xf8, 0x2d, 0x82, 0xde, 0x39, 0xe7, 0x86,
	0x71, 0xe2, 0xf4, 0x3f, 0x98, 0xce, 0xfe, 0xe8, 0xed, 0xbd, 0xd1, 0xe3, 0x37, 0x08, 0xee, 0x4f,
	0x85, 0x75, 0xa2, 0xc8, 0xdd, 0xff, 0x8d, 0x32, 0x39, 0x5d, 0x7f, 0x8b, 0x5a, 0xeb, 0x2a, 0x42,
	0x1f, 0xab, 0x08, 0x7d, 0xaa, 0x22, 0xf4, 0xb5, 0x8a, 0xd0, 0xbb, 0xef, 0x51, 0xeb, 0x35, 0x6c,
	0xed, 0xe6, 0x77, 0xfd, 0x1f, 0x78, 0xf2, 0x23, 0x00, 0x00, 0xff, 0xff, 0x55, 0xb0, 0xb1, 0x31,
	0xbd, 0x03, 0x00, 0x00,
}