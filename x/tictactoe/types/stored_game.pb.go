// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tictactoe/stored_game.proto

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

type StoredGame_GameStatus int32

const (
	StoredGame_OPEN        StoredGame_GameStatus = 0
	StoredGame_IN_PROGRESS StoredGame_GameStatus = 1
	StoredGame_FINISHED    StoredGame_GameStatus = 2
)

var StoredGame_GameStatus_name = map[int32]string{
	0: "OPEN",
	1: "IN_PROGRESS",
	2: "FINISHED",
}

var StoredGame_GameStatus_value = map[string]int32{
	"OPEN":        0,
	"IN_PROGRESS": 1,
	"FINISHED":    2,
}

func (x StoredGame_GameStatus) String() string {
	return proto.EnumName(StoredGame_GameStatus_name, int32(x))
}

func (StoredGame_GameStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_890ea4c29805d5e1, []int{0, 0}
}

type StoredGame_WinnerStatus int32

const (
	StoredGame_NONE    StoredGame_WinnerStatus = 0
	StoredGame_PLAYER1 StoredGame_WinnerStatus = 1
	StoredGame_PLAYER2 StoredGame_WinnerStatus = 2
	StoredGame_DRAW    StoredGame_WinnerStatus = 3
)

var StoredGame_WinnerStatus_name = map[int32]string{
	0: "NONE",
	1: "PLAYER1",
	2: "PLAYER2",
	3: "DRAW",
}

var StoredGame_WinnerStatus_value = map[string]int32{
	"NONE":    0,
	"PLAYER1": 1,
	"PLAYER2": 2,
	"DRAW":    3,
}

func (x StoredGame_WinnerStatus) String() string {
	return proto.EnumName(StoredGame_WinnerStatus_name, int32(x))
}

func (StoredGame_WinnerStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_890ea4c29805d5e1, []int{0, 1}
}

type StoredGame struct {
	Index   string                  `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	Player1 string                  `protobuf:"bytes,2,opt,name=player1,proto3" json:"player1,omitempty"`
	Player2 string                  `protobuf:"bytes,3,opt,name=player2,proto3" json:"player2,omitempty"`
	Turn    uint64                  `protobuf:"varint,4,opt,name=turn,proto3" json:"turn,omitempty"`
	Status  StoredGame_GameStatus   `protobuf:"varint,5,opt,name=status,proto3,enum=shifty11.tictactoe.tictactoe.StoredGame_GameStatus" json:"status,omitempty"`
	Winner  StoredGame_WinnerStatus `protobuf:"varint,6,opt,name=winner,proto3,enum=shifty11.tictactoe.tictactoe.StoredGame_WinnerStatus" json:"winner,omitempty"`
	Board   []uint64                `protobuf:"varint,7,rep,packed,name=board,proto3" json:"board,omitempty"`
}

func (m *StoredGame) Reset()         { *m = StoredGame{} }
func (m *StoredGame) String() string { return proto.CompactTextString(m) }
func (*StoredGame) ProtoMessage()    {}
func (*StoredGame) Descriptor() ([]byte, []int) {
	return fileDescriptor_890ea4c29805d5e1, []int{0}
}
func (m *StoredGame) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *StoredGame) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_StoredGame.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *StoredGame) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StoredGame.Merge(m, src)
}
func (m *StoredGame) XXX_Size() int {
	return m.Size()
}
func (m *StoredGame) XXX_DiscardUnknown() {
	xxx_messageInfo_StoredGame.DiscardUnknown(m)
}

var xxx_messageInfo_StoredGame proto.InternalMessageInfo

func (m *StoredGame) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *StoredGame) GetPlayer1() string {
	if m != nil {
		return m.Player1
	}
	return ""
}

func (m *StoredGame) GetPlayer2() string {
	if m != nil {
		return m.Player2
	}
	return ""
}

func (m *StoredGame) GetTurn() uint64 {
	if m != nil {
		return m.Turn
	}
	return 0
}

func (m *StoredGame) GetStatus() StoredGame_GameStatus {
	if m != nil {
		return m.Status
	}
	return StoredGame_OPEN
}

func (m *StoredGame) GetWinner() StoredGame_WinnerStatus {
	if m != nil {
		return m.Winner
	}
	return StoredGame_NONE
}

func (m *StoredGame) GetBoard() []uint64 {
	if m != nil {
		return m.Board
	}
	return nil
}

func init() {
	proto.RegisterEnum("shifty11.tictactoe.tictactoe.StoredGame_GameStatus", StoredGame_GameStatus_name, StoredGame_GameStatus_value)
	proto.RegisterEnum("shifty11.tictactoe.tictactoe.StoredGame_WinnerStatus", StoredGame_WinnerStatus_name, StoredGame_WinnerStatus_value)
	proto.RegisterType((*StoredGame)(nil), "shifty11.tictactoe.tictactoe.StoredGame")
}

func init() { proto.RegisterFile("tictactoe/stored_game.proto", fileDescriptor_890ea4c29805d5e1) }

var fileDescriptor_890ea4c29805d5e1 = []byte{
	// 364 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x41, 0x6f, 0xaa, 0x40,
	0x10, 0xc7, 0x59, 0x41, 0xf4, 0x8d, 0xe6, 0x3d, 0xb2, 0x79, 0x87, 0x4d, 0xde, 0x0b, 0x21, 0x9e,
	0xb8, 0x08, 0x01, 0xe3, 0xad, 0x17, 0x1b, 0xa9, 0x35, 0xb6, 0x68, 0x96, 0x83, 0x69, 0x2f, 0x06,
	0x71, 0xab, 0x24, 0x15, 0x0c, 0xac, 0xa9, 0x7e, 0x8b, 0xf6, 0x5b, 0xf5, 0xe8, 0xb1, 0xc7, 0x46,
	0xbf, 0x48, 0x03, 0xd6, 0xe2, 0xa9, 0xe9, 0x85, 0xcc, 0x9f, 0x99, 0xff, 0x6f, 0x66, 0xb2, 0x03,
	0xff, 0x78, 0x18, 0x70, 0x3f, 0xe0, 0x31, 0x33, 0x53, 0x1e, 0x27, 0x6c, 0x36, 0x99, 0xfb, 0x4b,
	0x66, 0xac, 0x92, 0x98, 0xc7, 0xf8, 0x7f, 0xba, 0x08, 0x1f, 0xf8, 0xd6, 0xb2, 0x8c, 0xaf, 0xaa,
	0x22, 0x6a, 0xbc, 0x88, 0x00, 0x5e, 0xee, 0xe9, 0xf9, 0x4b, 0x86, 0xff, 0x42, 0x39, 0x8c, 0x66,
	0x6c, 0x43, 0x90, 0x86, 0xf4, 0x5f, 0xf4, 0x28, 0x30, 0x81, 0xca, 0xea, 0xd1, 0xdf, 0xb2, 0xc4,
	0x22, 0xa5, 0xfc, 0xff, 0x49, 0x16, 0x19, 0x9b, 0x88, 0xe7, 0x19, 0x1b, 0x63, 0x90, 0xf8, 0x3a,
	0x89, 0x88, 0xa4, 0x21, 0x5d, 0xa2, 0x79, 0x8c, 0x07, 0x20, 0xa7, 0xdc, 0xe7, 0xeb, 0x94, 0x94,
	0x35, 0xa4, 0xff, 0xb6, 0x5b, 0xc6, 0x77, 0xb3, 0x19, 0xc5, 0x5c, 0x46, 0xf6, 0xf1, 0x72, 0x2b,
	0xfd, 0x44, 0xe0, 0x5b, 0x90, 0x9f, 0xc2, 0x28, 0x62, 0x09, 0x91, 0x73, 0x58, 0xfb, 0xc7, 0xb0,
	0x71, 0x6e, 0x3b, 0xe1, 0x8e, 0x90, 0x6c, 0xf3, 0x69, 0xec, 0x27, 0x33, 0x52, 0xd1, 0x44, 0x5d,
	0xa2, 0x47, 0xd1, 0x68, 0x03, 0x14, 0xad, 0x71, 0x15, 0xa4, 0xe1, 0xc8, 0x71, 0x15, 0x01, 0xff,
	0x81, 0x5a, 0xdf, 0x9d, 0x8c, 0xe8, 0xb0, 0x47, 0x1d, 0xcf, 0x53, 0x10, 0xae, 0x43, 0xf5, 0xaa,
	0xef, 0xf6, 0xbd, 0x6b, 0xa7, 0xab, 0x94, 0x1a, 0x17, 0x50, 0x3f, 0x6f, 0x92, 0x19, 0xdd, 0xa1,
	0xeb, 0x28, 0x02, 0xae, 0x41, 0x65, 0x74, 0xd3, 0xb9, 0x73, 0xa8, 0xa5, 0xa0, 0x42, 0xd8, 0x4a,
	0x29, 0xab, 0xe9, 0xd2, 0xce, 0x58, 0x11, 0x2f, 0x07, 0xaf, 0x7b, 0x15, 0xed, 0xf6, 0x2a, 0x7a,
	0xdf, 0xab, 0xe8, 0xf9, 0xa0, 0x0a, 0xbb, 0x83, 0x2a, 0xbc, 0x1d, 0x54, 0xe1, 0xde, 0x9a, 0x87,
	0x7c, 0xb1, 0x9e, 0x1a, 0x41, 0xbc, 0x34, 0x4f, 0xdb, 0x9a, 0x3c, 0x0c, 0x9a, 0xdc, 0x0f, 0x9a,
	0xd9, 0xf3, 0x6f, 0xcc, 0xe2, 0x14, 0xf8, 0x76, 0xc5, 0xd2, 0xa9, 0x9c, 0x5f, 0x41, 0xeb, 0x23,
	0x00, 0x00, 0xff, 0xff, 0xb0, 0x7b, 0x78, 0xef, 0x24, 0x02, 0x00, 0x00,
}

func (m *StoredGame) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StoredGame) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *StoredGame) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Board) > 0 {
		dAtA2 := make([]byte, len(m.Board)*10)
		var j1 int
		for _, num := range m.Board {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintStoredGame(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0x3a
	}
	if m.Winner != 0 {
		i = encodeVarintStoredGame(dAtA, i, uint64(m.Winner))
		i--
		dAtA[i] = 0x30
	}
	if m.Status != 0 {
		i = encodeVarintStoredGame(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x28
	}
	if m.Turn != 0 {
		i = encodeVarintStoredGame(dAtA, i, uint64(m.Turn))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Player2) > 0 {
		i -= len(m.Player2)
		copy(dAtA[i:], m.Player2)
		i = encodeVarintStoredGame(dAtA, i, uint64(len(m.Player2)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Player1) > 0 {
		i -= len(m.Player1)
		copy(dAtA[i:], m.Player1)
		i = encodeVarintStoredGame(dAtA, i, uint64(len(m.Player1)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintStoredGame(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintStoredGame(dAtA []byte, offset int, v uint64) int {
	offset -= sovStoredGame(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *StoredGame) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovStoredGame(uint64(l))
	}
	l = len(m.Player1)
	if l > 0 {
		n += 1 + l + sovStoredGame(uint64(l))
	}
	l = len(m.Player2)
	if l > 0 {
		n += 1 + l + sovStoredGame(uint64(l))
	}
	if m.Turn != 0 {
		n += 1 + sovStoredGame(uint64(m.Turn))
	}
	if m.Status != 0 {
		n += 1 + sovStoredGame(uint64(m.Status))
	}
	if m.Winner != 0 {
		n += 1 + sovStoredGame(uint64(m.Winner))
	}
	if len(m.Board) > 0 {
		l = 0
		for _, e := range m.Board {
			l += sovStoredGame(uint64(e))
		}
		n += 1 + sovStoredGame(uint64(l)) + l
	}
	return n
}

func sovStoredGame(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozStoredGame(x uint64) (n int) {
	return sovStoredGame(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *StoredGame) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowStoredGame
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
			return fmt.Errorf("proto: StoredGame: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StoredGame: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStoredGame
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
				return ErrInvalidLengthStoredGame
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStoredGame
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Player1", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStoredGame
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
				return ErrInvalidLengthStoredGame
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStoredGame
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Player1 = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Player2", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStoredGame
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
				return ErrInvalidLengthStoredGame
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthStoredGame
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Player2 = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Turn", wireType)
			}
			m.Turn = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStoredGame
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Turn |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStoredGame
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= StoredGame_GameStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Winner", wireType)
			}
			m.Winner = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowStoredGame
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Winner |= StoredGame_WinnerStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowStoredGame
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Board = append(m.Board, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowStoredGame
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthStoredGame
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthStoredGame
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.Board) == 0 {
					m.Board = make([]uint64, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowStoredGame
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Board = append(m.Board, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Board", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipStoredGame(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthStoredGame
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
func skipStoredGame(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowStoredGame
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
					return 0, ErrIntOverflowStoredGame
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
					return 0, ErrIntOverflowStoredGame
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
				return 0, ErrInvalidLengthStoredGame
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupStoredGame
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthStoredGame
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthStoredGame        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowStoredGame          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupStoredGame = fmt.Errorf("proto: unexpected end of group")
)