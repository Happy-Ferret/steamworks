// Code generated by "stringer -type EVoiceResult -trimprefix EVoiceResult_"; DO NOT EDIT.

package internal

import "strconv"

const _EVoiceResult_name = "OKNotInitializedNotRecordingNoDataBufferTooSmallDataCorruptedRestrictedUnsupportedCodecReceiverOutOfDateReceiverDidNotAnswer"

var _EVoiceResult_index = [...]uint8{0, 2, 16, 28, 34, 48, 61, 71, 87, 104, 124}

func (i EVoiceResult) String() string {
	if i < 0 || i >= EVoiceResult(len(_EVoiceResult_index)-1) {
		return "EVoiceResult(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _EVoiceResult_name[_EVoiceResult_index[i]:_EVoiceResult_index[i+1]]
}