// Code generated by "stringer -type ESteamControllerLEDFlag -trimprefix ESteamControllerLEDFlag_"; DO NOT EDIT.

package internal

import "strconv"

const _ESteamControllerLEDFlag_name = "SetColorRestoreUserDefault"

var _ESteamControllerLEDFlag_index = [...]uint8{0, 8, 26}

func (i ESteamControllerLEDFlag) String() string {
	if i < 0 || i >= ESteamControllerLEDFlag(len(_ESteamControllerLEDFlag_index)-1) {
		return "ESteamControllerLEDFlag(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ESteamControllerLEDFlag_name[_ESteamControllerLEDFlag_index[i]:_ESteamControllerLEDFlag_index[i+1]]
}
