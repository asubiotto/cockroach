// Code generated by protoc-gen-gogo.
// source: cockroach/pkg/storage/storagebase/state.proto
// DO NOT EDIT!

package storagebase

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import cockroach_storage_engine_enginepb "github.com/cockroachdb/cockroach/pkg/storage/engine/enginepb"
import cockroach_roachpb4 "github.com/cockroachdb/cockroach/pkg/roachpb"
import cockroach_roachpb "github.com/cockroachdb/cockroach/pkg/roachpb"
import cockroach_roachpb1 "github.com/cockroachdb/cockroach/pkg/roachpb"
import cockroach_util_hlc "github.com/cockroachdb/cockroach/pkg/util/hlc"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Use an enum because proto3 does not give nullable primitive values, and we
// need to be able to send EvalResults which don't specify frozen.
type ReplicaState_FrozenEnum int32

const (
	ReplicaState_FROZEN_UNSPECIFIED ReplicaState_FrozenEnum = 0
	ReplicaState_FROZEN             ReplicaState_FrozenEnum = 1
	ReplicaState_UNFROZEN           ReplicaState_FrozenEnum = 2
)

var ReplicaState_FrozenEnum_name = map[int32]string{
	0: "FROZEN_UNSPECIFIED",
	1: "FROZEN",
	2: "UNFROZEN",
}
var ReplicaState_FrozenEnum_value = map[string]int32{
	"FROZEN_UNSPECIFIED": 0,
	"FROZEN":             1,
	"UNFROZEN":           2,
}

func (x ReplicaState_FrozenEnum) String() string {
	return proto.EnumName(ReplicaState_FrozenEnum_name, int32(x))
}
func (ReplicaState_FrozenEnum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptorState, []int{0, 0}
}

// ReplicaState is the part of the Range Raft state machine which is cached in
// memory and which is manipulated exclusively through consensus.
//
// The struct is also used to transfer state to Replicas in the context of
// proposer-evaluated Raft, in which case it does not represent a complete
// state but instead an update to be applied to an existing state, with each
// field specified in the update overwriting its counterpart on the receiving
// ReplicaState.
//
// For the ReplicaState persisted on the Replica, all optional fields are
// populated (i.e. no nil pointers or enums with the default value).
type ReplicaState struct {
	// The highest (and last) index applied to the state machine.
	RaftAppliedIndex uint64 `protobuf:"varint,1,opt,name=raft_applied_index,json=raftAppliedIndex,proto3" json:"raft_applied_index,omitempty"`
	// The highest (and last) lease index applied to the state machine.
	LeaseAppliedIndex uint64 `protobuf:"varint,2,opt,name=lease_applied_index,json=leaseAppliedIndex,proto3" json:"lease_applied_index,omitempty"`
	// The Range descriptor.
	// The pointer may change, but the referenced RangeDescriptor struct itself
	// must be treated as immutable; it is leaked out of the lock.
	//
	// Changes of the descriptor should always go through one of the
	// (*Replica).setDesc* methods.
	Desc *cockroach_roachpb.RangeDescriptor `protobuf:"bytes,3,opt,name=desc" json:"desc,omitempty"`
	// The latest lease, if any.
	Lease *cockroach_roachpb1.Lease `protobuf:"bytes,4,opt,name=lease" json:"lease,omitempty"`
	// The truncation state of the Raft log.
	TruncatedState *cockroach_roachpb4.RaftTruncatedState `protobuf:"bytes,5,opt,name=truncated_state,json=truncatedState" json:"truncated_state,omitempty"`
	// gcThreshold is the GC threshold of the Range, typically updated when keys
	// are garbage collected. Reads and writes at timestamps <= this time will
	// not be served.
	//
	// TODO(tschottdorf): should be nullable to keep ReplicaState small as we are
	// sending it over the wire. Since we only ever increase gc_threshold, that's
	// the only upshot - fields which can return to the zero value must
	// special-case that value simply because otherwise there's no way of
	// distinguishing "no update" to and updating to the zero value.
	GCThreshold cockroach_util_hlc.Timestamp                `protobuf:"bytes,6,opt,name=gc_threshold,json=gcThreshold" json:"gc_threshold"`
	Stats       cockroach_storage_engine_enginepb.MVCCStats `protobuf:"bytes,7,opt,name=stats" json:"stats"`
	// txn_span_gc_threshold is the (maximum) timestamp below which transaction
	// records may have been garbage collected (as measured by txn.LastActive()).
	// Transaction at lower timestamps must not be allowed to write their initial
	// transaction entry.
	//
	// TODO(tschottdorf): should be nullable; see gc_threshold.
	TxnSpanGCThreshold cockroach_util_hlc.Timestamp `protobuf:"bytes,9,opt,name=txn_span_gc_threshold,json=txnSpanGcThreshold" json:"txn_span_gc_threshold"`
	Frozen             ReplicaState_FrozenEnum      `protobuf:"varint,10,opt,name=frozen,proto3,enum=cockroach.storage.storagebase.ReplicaState_FrozenEnum" json:"frozen,omitempty"`
}

func (m *ReplicaState) Reset()                    { *m = ReplicaState{} }
func (m *ReplicaState) String() string            { return proto.CompactTextString(m) }
func (*ReplicaState) ProtoMessage()               {}
func (*ReplicaState) Descriptor() ([]byte, []int) { return fileDescriptorState, []int{0} }

type RangeInfo struct {
	ReplicaState `protobuf:"bytes,1,opt,name=state,embedded=state" json:"state"`
	// The highest (and last) index in the Raft log.
	LastIndex  uint64 `protobuf:"varint,2,opt,name=lastIndex,proto3" json:"lastIndex,omitempty"`
	NumPending uint64 `protobuf:"varint,3,opt,name=num_pending,json=numPending,proto3" json:"num_pending,omitempty"`
	NumDropped uint64 `protobuf:"varint,5,opt,name=num_dropped,json=numDropped,proto3" json:"num_dropped,omitempty"`
	// raft_log_size may be initially inaccurate after a server restart.
	// See storage.Replica.mu.raftLogSize.
	RaftLogSize int64 `protobuf:"varint,6,opt,name=raft_log_size,json=raftLogSize,proto3" json:"raft_log_size,omitempty"`
}

func (m *RangeInfo) Reset()                    { *m = RangeInfo{} }
func (m *RangeInfo) String() string            { return proto.CompactTextString(m) }
func (*RangeInfo) ProtoMessage()               {}
func (*RangeInfo) Descriptor() ([]byte, []int) { return fileDescriptorState, []int{1} }

func init() {
	proto.RegisterType((*ReplicaState)(nil), "cockroach.storage.storagebase.ReplicaState")
	proto.RegisterType((*RangeInfo)(nil), "cockroach.storage.storagebase.RangeInfo")
	proto.RegisterEnum("cockroach.storage.storagebase.ReplicaState_FrozenEnum", ReplicaState_FrozenEnum_name, ReplicaState_FrozenEnum_value)
}
func (m *ReplicaState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReplicaState) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.RaftAppliedIndex != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintState(dAtA, i, uint64(m.RaftAppliedIndex))
	}
	if m.LeaseAppliedIndex != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintState(dAtA, i, uint64(m.LeaseAppliedIndex))
	}
	if m.Desc != nil {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintState(dAtA, i, uint64(m.Desc.Size()))
		n1, err := m.Desc.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if m.Lease != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintState(dAtA, i, uint64(m.Lease.Size()))
		n2, err := m.Lease.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	if m.TruncatedState != nil {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintState(dAtA, i, uint64(m.TruncatedState.Size()))
		n3, err := m.TruncatedState.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	dAtA[i] = 0x32
	i++
	i = encodeVarintState(dAtA, i, uint64(m.GCThreshold.Size()))
	n4, err := m.GCThreshold.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n4
	dAtA[i] = 0x3a
	i++
	i = encodeVarintState(dAtA, i, uint64(m.Stats.Size()))
	n5, err := m.Stats.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n5
	dAtA[i] = 0x4a
	i++
	i = encodeVarintState(dAtA, i, uint64(m.TxnSpanGCThreshold.Size()))
	n6, err := m.TxnSpanGCThreshold.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n6
	if m.Frozen != 0 {
		dAtA[i] = 0x50
		i++
		i = encodeVarintState(dAtA, i, uint64(m.Frozen))
	}
	return i, nil
}

func (m *RangeInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RangeInfo) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintState(dAtA, i, uint64(m.ReplicaState.Size()))
	n7, err := m.ReplicaState.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n7
	if m.LastIndex != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintState(dAtA, i, uint64(m.LastIndex))
	}
	if m.NumPending != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintState(dAtA, i, uint64(m.NumPending))
	}
	if m.NumDropped != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintState(dAtA, i, uint64(m.NumDropped))
	}
	if m.RaftLogSize != 0 {
		dAtA[i] = 0x30
		i++
		i = encodeVarintState(dAtA, i, uint64(m.RaftLogSize))
	}
	return i, nil
}

func encodeFixed64State(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32State(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintState(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ReplicaState) Size() (n int) {
	var l int
	_ = l
	if m.RaftAppliedIndex != 0 {
		n += 1 + sovState(uint64(m.RaftAppliedIndex))
	}
	if m.LeaseAppliedIndex != 0 {
		n += 1 + sovState(uint64(m.LeaseAppliedIndex))
	}
	if m.Desc != nil {
		l = m.Desc.Size()
		n += 1 + l + sovState(uint64(l))
	}
	if m.Lease != nil {
		l = m.Lease.Size()
		n += 1 + l + sovState(uint64(l))
	}
	if m.TruncatedState != nil {
		l = m.TruncatedState.Size()
		n += 1 + l + sovState(uint64(l))
	}
	l = m.GCThreshold.Size()
	n += 1 + l + sovState(uint64(l))
	l = m.Stats.Size()
	n += 1 + l + sovState(uint64(l))
	l = m.TxnSpanGCThreshold.Size()
	n += 1 + l + sovState(uint64(l))
	if m.Frozen != 0 {
		n += 1 + sovState(uint64(m.Frozen))
	}
	return n
}

func (m *RangeInfo) Size() (n int) {
	var l int
	_ = l
	l = m.ReplicaState.Size()
	n += 1 + l + sovState(uint64(l))
	if m.LastIndex != 0 {
		n += 1 + sovState(uint64(m.LastIndex))
	}
	if m.NumPending != 0 {
		n += 1 + sovState(uint64(m.NumPending))
	}
	if m.NumDropped != 0 {
		n += 1 + sovState(uint64(m.NumDropped))
	}
	if m.RaftLogSize != 0 {
		n += 1 + sovState(uint64(m.RaftLogSize))
	}
	return n
}

func sovState(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozState(x uint64) (n int) {
	return sovState(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ReplicaState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowState
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
			return fmt.Errorf("proto: ReplicaState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ReplicaState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RaftAppliedIndex", wireType)
			}
			m.RaftAppliedIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RaftAppliedIndex |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LeaseAppliedIndex", wireType)
			}
			m.LeaseAppliedIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LeaseAppliedIndex |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Desc", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Desc == nil {
				m.Desc = &cockroach_roachpb.RangeDescriptor{}
			}
			if err := m.Desc.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Lease", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Lease == nil {
				m.Lease = &cockroach_roachpb1.Lease{}
			}
			if err := m.Lease.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TruncatedState", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.TruncatedState == nil {
				m.TruncatedState = &cockroach_roachpb4.RaftTruncatedState{}
			}
			if err := m.TruncatedState.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GCThreshold", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.GCThreshold.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Stats", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Stats.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxnSpanGCThreshold", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TxnSpanGCThreshold.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Frozen", wireType)
			}
			m.Frozen = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Frozen |= (ReplicaState_FrozenEnum(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipState(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthState
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
func (m *RangeInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowState
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
			return fmt.Errorf("proto: RangeInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RangeInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReplicaState", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
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
				return ErrInvalidLengthState
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ReplicaState.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastIndex", wireType)
			}
			m.LastIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastIndex |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumPending", wireType)
			}
			m.NumPending = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NumPending |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NumDropped", wireType)
			}
			m.NumDropped = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NumDropped |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RaftLogSize", wireType)
			}
			m.RaftLogSize = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowState
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RaftLogSize |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipState(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthState
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
func skipState(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowState
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
					return 0, ErrIntOverflowState
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
					return 0, ErrIntOverflowState
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
				return 0, ErrInvalidLengthState
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowState
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
				next, err := skipState(dAtA[start:])
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
	ErrInvalidLengthState = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowState   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("cockroach/pkg/storage/storagebase/state.proto", fileDescriptorState) }

var fileDescriptorState = []byte{
	// 662 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x93, 0x51, 0x6f, 0xd3, 0x3a,
	0x18, 0x86, 0x9b, 0x2d, 0xed, 0x69, 0x9d, 0x9d, 0x9d, 0x1e, 0x0f, 0x50, 0x34, 0xb1, 0xb6, 0xaa,
	0x18, 0x1a, 0x62, 0x38, 0x68, 0x48, 0xbb, 0x44, 0xa2, 0x5d, 0x07, 0x1d, 0xa3, 0x4c, 0x69, 0xc7,
	0xc5, 0x6e, 0x22, 0x37, 0x71, 0xd3, 0x68, 0x89, 0x6d, 0x25, 0x2e, 0x9a, 0xf6, 0x2b, 0xf8, 0x59,
	0xe3, 0x6e, 0x97, 0x5c, 0x4d, 0x50, 0xfe, 0x08, 0xb2, 0x93, 0x6c, 0x29, 0x54, 0x08, 0xae, 0x62,
	0x7f, 0x7e, 0xbe, 0xb7, 0xaf, 0xbf, 0xbe, 0x06, 0xcf, 0x5c, 0xe6, 0x9e, 0xc7, 0x0c, 0xbb, 0x53,
	0x8b, 0x9f, 0xfb, 0x56, 0x22, 0x58, 0x8c, 0x7d, 0x92, 0x7f, 0xc7, 0x38, 0x91, 0x6b, 0x2c, 0x08,
	0xe2, 0x31, 0x13, 0x0c, 0x6e, 0xdd, 0xe2, 0x28, 0x43, 0x50, 0x01, 0xdd, 0x7c, 0xbe, 0x5c, 0x8d,
	0x50, 0x3f, 0xa0, 0xf9, 0x87, 0x8f, 0xad, 0xe8, 0xa3, 0xeb, 0xa6, 0x82, 0x9b, 0x4f, 0x16, 0x3b,
	0xd4, 0x8a, 0x8f, 0xad, 0x80, 0x0a, 0x12, 0x53, 0x1c, 0x3a, 0x31, 0x9e, 0x88, 0x0c, 0x7d, 0xb4,
	0x1c, 0x8d, 0x88, 0xc0, 0x1e, 0x16, 0x38, 0xa3, 0x5a, 0xcb, 0xa9, 0x02, 0xf1, 0x78, 0x91, 0x98,
	0x89, 0x20, 0xb4, 0xa6, 0xa1, 0x6b, 0x89, 0x20, 0x22, 0x89, 0xc0, 0x11, 0xcf, 0xb8, 0x7b, 0x3e,
	0xf3, 0x99, 0x5a, 0x5a, 0x72, 0x95, 0x56, 0xdb, 0x9f, 0xcb, 0x60, 0xcd, 0x26, 0x3c, 0x0c, 0x5c,
	0x3c, 0x94, 0x83, 0x81, 0xbb, 0x00, 0x4a, 0x93, 0x0e, 0xe6, 0x3c, 0x0c, 0x88, 0xe7, 0x04, 0xd4,
	0x23, 0x17, 0xa6, 0xd6, 0xd2, 0x76, 0x74, 0xbb, 0x2e, 0x4f, 0x5e, 0xa5, 0x07, 0x7d, 0x59, 0x87,
	0x08, 0x6c, 0x84, 0x04, 0x27, 0xe4, 0x27, 0x7c, 0x45, 0xe1, 0xff, 0xab, 0xa3, 0x05, 0x7e, 0x1f,
	0xe8, 0x1e, 0x49, 0x5c, 0x73, 0xb5, 0xa5, 0xed, 0x18, 0x7b, 0x6d, 0x74, 0x37, 0xff, 0xec, 0x66,
	0xc8, 0xc6, 0xd4, 0x27, 0x07, 0x24, 0x71, 0xe3, 0x80, 0x0b, 0x16, 0xdb, 0x8a, 0x87, 0x08, 0x94,
	0x95, 0x98, 0xa9, 0xab, 0x46, 0x73, 0x49, 0xe3, 0xb1, 0x3c, 0xb7, 0x53, 0x0c, 0x0e, 0xc0, 0x7f,
	0x22, 0x9e, 0x51, 0x17, 0x0b, 0xe2, 0x39, 0xea, 0x1f, 0x37, 0xcb, 0xaa, 0x73, 0x7b, 0xe9, 0x4f,
	0x4e, 0xc4, 0x28, 0xa7, 0xd5, 0x14, 0xec, 0x75, 0xb1, 0xb0, 0x87, 0xa7, 0x60, 0xcd, 0x77, 0x1d,
	0x31, 0x8d, 0x49, 0x32, 0x65, 0xa1, 0x67, 0x56, 0x94, 0xd8, 0x56, 0x41, 0x4c, 0xce, 0x1d, 0x4d,
	0x43, 0x17, 0x8d, 0xf2, 0xb9, 0x77, 0x36, 0xae, 0x6e, 0x9a, 0xa5, 0xf9, 0x4d, 0xd3, 0x78, 0xdd,
	0x1d, 0xe5, 0x9d, 0xb6, 0xe1, 0xbb, 0xb7, 0x1b, 0xf8, 0x06, 0x94, 0xa5, 0xb9, 0xc4, 0xfc, 0x47,
	0xe9, 0xed, 0xa2, 0x5f, 0xf3, 0x98, 0xa6, 0x0c, 0xe5, 0x61, 0x43, 0xef, 0x3e, 0x74, 0xbb, 0xd2,
	0x53, 0xd2, 0xd1, 0xa5, 0xbc, 0x9d, 0x0a, 0xc0, 0x10, 0xdc, 0x17, 0x17, 0xd4, 0x49, 0x38, 0xa6,
	0xce, 0x82, 0xd3, 0xda, 0x9f, 0x38, 0xdd, 0xcc, 0x9c, 0xc2, 0xd1, 0x05, 0x1d, 0x72, 0x4c, 0x8b,
	0x86, 0xa1, 0xc8, 0x6a, 0x05, 0xdf, 0x03, 0x50, 0x99, 0xc4, 0xec, 0x92, 0x50, 0x13, 0xb4, 0xb4,
	0x9d, 0xf5, 0xbd, 0x7d, 0xf4, 0xdb, 0x87, 0x84, 0x8a, 0x09, 0x43, 0x87, 0xaa, 0xb3, 0x47, 0x67,
	0x91, 0x9d, 0xa9, 0xb4, 0x5f, 0x02, 0x70, 0x57, 0x85, 0x0f, 0x00, 0x3c, 0xb4, 0xdf, 0x9f, 0xf5,
	0x06, 0xce, 0xe9, 0x60, 0x78, 0xd2, 0xeb, 0xf6, 0x0f, 0xfb, 0xbd, 0x83, 0x7a, 0x09, 0x02, 0x50,
	0x49, 0xeb, 0x75, 0x0d, 0xae, 0x81, 0xea, 0xe9, 0x20, 0xdb, 0xad, 0x1c, 0xe9, 0xd5, 0x6a, 0xbd,
	0xd6, 0x9e, 0x6b, 0xa0, 0xa6, 0xe2, 0xd3, 0xa7, 0x13, 0x06, 0xdf, 0xa6, 0xb3, 0x25, 0x2a, 0xbb,
	0xc6, 0xde, 0xd3, 0xbf, 0xb0, 0xd8, 0xa9, 0xca, 0x79, 0x5c, 0xdf, 0x34, 0xb5, 0x74, 0xbc, 0x04,
	0x3e, 0x04, 0xb5, 0x10, 0x27, 0xa2, 0x5f, 0x48, 0xf7, 0x5d, 0x01, 0x36, 0x81, 0x41, 0x67, 0x91,
	0xc3, 0x09, 0xf5, 0x02, 0xea, 0xab, 0x70, 0xeb, 0x36, 0xa0, 0xb3, 0xe8, 0x24, 0xad, 0xe4, 0x80,
	0x17, 0x33, 0xce, 0x89, 0xa7, 0xa2, 0x98, 0x02, 0x07, 0x69, 0x05, 0xb6, 0xc1, 0xbf, 0xea, 0xd5,
	0x85, 0xcc, 0x77, 0x92, 0xe0, 0x92, 0xa8, 0x80, 0xad, 0xda, 0x86, 0x2c, 0x1e, 0x33, 0x7f, 0x18,
	0x5c, 0x92, 0x23, 0xbd, 0xaa, 0xd7, 0xcb, 0x9d, 0xed, 0xab, 0x6f, 0x8d, 0xd2, 0xd5, 0xbc, 0xa1,
	0x5d, 0xcf, 0x1b, 0xda, 0x97, 0x79, 0x43, 0xfb, 0x3a, 0x6f, 0x68, 0x9f, 0xbe, 0x37, 0x4a, 0x67,
	0x46, 0xe1, 0x3a, 0xe3, 0x8a, 0x7a, 0xde, 0x2f, 0x7e, 0x04, 0x00, 0x00, 0xff, 0xff, 0x60, 0xfb,
	0x2c, 0x45, 0x11, 0x05, 0x00, 0x00,
}
