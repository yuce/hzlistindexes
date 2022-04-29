package codec

import (
	"encoding/binary"

	proto "github.com/hazelcast/hazelcast-go-client"
	"github.com/hazelcast/hazelcast-go-client/types"
)

func EncodeString(message *proto.ClientMessage, value interface{}) {
	message.AddFrame(proto.NewFrame([]byte(value.(string))))
}

func DecodeString(frameIterator *proto.ForwardFrameIterator) string {
	return string(frameIterator.Next().Content)
}

func DecodeInt(buffer []byte, offset int32) int32 {
	return int32(binary.LittleEndian.Uint32(buffer[offset:]))
}

func DecodeBoolean(buffer []byte, offset int32) bool {
	return buffer[offset] == 1
}

func NextFrameIsDataStructureEndFrame(frameIterator *proto.ForwardFrameIterator) bool {
	return frameIterator.PeekNext().IsEndFrame()
}

func DecodeListMultiFrameForIndexConfig(frameIterator *proto.ForwardFrameIterator) []types.IndexConfig {
	result := make([]types.IndexConfig, 0)
	if frameIterator.HasNext() {
		frameIterator.Next()

		for !NextFrameIsDataStructureEndFrame(frameIterator) {
			result = append(result, DecodeIndexConfig(frameIterator))
		}
		frameIterator.Next()
	}
	return result
}

func NextFrameIsNullFrame(frameIterator *proto.ForwardFrameIterator) bool {
	isNullFrame := frameIterator.PeekNext().IsNullFrame()
	if isNullFrame {
		frameIterator.Next()
	}
	return isNullFrame
}

func DecodeNullableForString(frameIterator *proto.ForwardFrameIterator) string {
	if NextFrameIsNullFrame(frameIterator) {
		return ""
	}
	return DecodeString(frameIterator)
}

func DecodeListMultiFrameForString(frameIterator *proto.ForwardFrameIterator) []string {
	result := make([]string, 0)
	frameIterator.Next()
	for !NextFrameIsDataStructureEndFrame(frameIterator) {
		result = append(result, DecodeString(frameIterator))
	}
	frameIterator.Next()
	return result
}

func DecodeNullableForBitmapIndexOptions(frameIterator *proto.ForwardFrameIterator) types.BitmapIndexOptions {
	if NextFrameIsNullFrame(frameIterator) {
		return types.BitmapIndexOptions{}
	}
	return DecodeBitmapIndexOptions(frameIterator)
}

func FastForwardToEndFrame(frameIterator *proto.ForwardFrameIterator) {
	expectedEndFrames := 1
	for expectedEndFrames != 0 {
		frame := frameIterator.Next()
		if frame.IsEndFrame() {
			expectedEndFrames--
		} else if frame.IsBeginFrame() {
			expectedEndFrames++
		}
	}
}
