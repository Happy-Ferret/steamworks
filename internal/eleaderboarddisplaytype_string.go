// Code generated by "stringer -type ELeaderboardDisplayType -trimprefix ELeaderboardDisplayType_"; DO NOT EDIT.

package internal

import "strconv"

const _ELeaderboardDisplayType_name = "NoneNumericTimeSecondsTimeMilliSeconds"

var _ELeaderboardDisplayType_index = [...]uint8{0, 4, 11, 22, 38}

func (i ELeaderboardDisplayType) String() string {
	if i < 0 || i >= ELeaderboardDisplayType(len(_ELeaderboardDisplayType_index)-1) {
		return "ELeaderboardDisplayType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ELeaderboardDisplayType_name[_ELeaderboardDisplayType_index[i]:_ELeaderboardDisplayType_index[i+1]]
}
