// Code generated by "stringer -type EFriendFlags -trimprefix EFriendFlags_"; DO NOT EDIT.

package internal

import "strconv"

const (
	_EFriendFlags_name_0 = "EFriendFlagNoneEFriendFlagBlockedEFriendFlagFriendshipRequested"
	_EFriendFlags_name_1 = "EFriendFlagImmediate"
	_EFriendFlags_name_2 = "EFriendFlagClanMember"
	_EFriendFlags_name_3 = "EFriendFlagOnGameServer"
	_EFriendFlags_name_4 = "EFriendFlagRequestingFriendship"
	_EFriendFlags_name_5 = "EFriendFlagRequestingInfo"
	_EFriendFlags_name_6 = "EFriendFlagIgnored"
	_EFriendFlags_name_7 = "EFriendFlagIgnoredFriend"
	_EFriendFlags_name_8 = "EFriendFlagChatMember"
	_EFriendFlags_name_9 = "EFriendFlagAll"
)

var (
	_EFriendFlags_index_0 = [...]uint8{0, 15, 33, 63}
)

func (i EFriendFlags) String() string {
	switch {
	case 0 <= i && i <= 2:
		return _EFriendFlags_name_0[_EFriendFlags_index_0[i]:_EFriendFlags_index_0[i+1]]
	case i == 4:
		return _EFriendFlags_name_1
	case i == 8:
		return _EFriendFlags_name_2
	case i == 16:
		return _EFriendFlags_name_3
	case i == 128:
		return _EFriendFlags_name_4
	case i == 256:
		return _EFriendFlags_name_5
	case i == 512:
		return _EFriendFlags_name_6
	case i == 1024:
		return _EFriendFlags_name_7
	case i == 4096:
		return _EFriendFlags_name_8
	case i == 65535:
		return _EFriendFlags_name_9
	default:
		return "EFriendFlags(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
