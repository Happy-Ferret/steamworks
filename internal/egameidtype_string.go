// Code generated by "stringer -type EGameIDType -trimprefix EGameIDType_"; DO NOT EDIT.

package internal

import "strconv"

const _EGameIDType_name = "AppGameModShortcutP2P"

var _EGameIDType_index = [...]uint8{0, 3, 10, 18, 21}

func (i EGameIDType) String() string {
	if i < 0 || i >= EGameIDType(len(_EGameIDType_index)-1) {
		return "EGameIDType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _EGameIDType_name[_EGameIDType_index[i]:_EGameIDType_index[i+1]]
}
