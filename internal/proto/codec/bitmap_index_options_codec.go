/*
 * Copyright (c) 2008-2021, Hazelcast, Inc. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License")
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package codec

import (
	proto "github.com/hazelcast/hazelcast-go-client"
	"github.com/hazelcast/hazelcast-go-client/types"
)

const (
	BitmapIndexOptionsCodecUniqueKeyTransformationFieldOffset      = 0
	BitmapIndexOptionsCodecUniqueKeyTransformationInitialFrameSize = BitmapIndexOptionsCodecUniqueKeyTransformationFieldOffset + proto.IntSizeInBytes
)

func DecodeBitmapIndexOptions(frameIterator *proto.ForwardFrameIterator) types.BitmapIndexOptions {
	// begin frame
	frameIterator.Next()
	initialFrame := frameIterator.Next()
	uniqueKeyTransformation := DecodeInt(initialFrame.Content, BitmapIndexOptionsCodecUniqueKeyTransformationFieldOffset)

	uniqueKey := DecodeString(frameIterator)
	FastForwardToEndFrame(frameIterator)
	return types.BitmapIndexOptions{UniqueKey: uniqueKey, UniqueKeyTransformation: types.UniqueKeyTransformation(uniqueKeyTransformation)}
}
