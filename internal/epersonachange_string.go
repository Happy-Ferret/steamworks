// Code generated by "stringer -type EPersonaChange -trimprefix EPersonaChange_"; DO NOT EDIT.

package internal

import "strconv"

const _EPersonaChange_name = "NameStatusComeOnlineGoneOfflineGamePlayedGameServerAvatarJoinedSourceLeftSourceRelationshipChangedNameFirstSetFacebookInfoNicknameSteamLevel"

var _EPersonaChange_map = map[EPersonaChange]string{
	1:    _EPersonaChange_name[0:4],
	2:    _EPersonaChange_name[4:10],
	4:    _EPersonaChange_name[10:20],
	8:    _EPersonaChange_name[20:31],
	16:   _EPersonaChange_name[31:41],
	32:   _EPersonaChange_name[41:51],
	64:   _EPersonaChange_name[51:57],
	128:  _EPersonaChange_name[57:69],
	256:  _EPersonaChange_name[69:79],
	512:  _EPersonaChange_name[79:98],
	1024: _EPersonaChange_name[98:110],
	2048: _EPersonaChange_name[110:122],
	4096: _EPersonaChange_name[122:130],
	8192: _EPersonaChange_name[130:140],
}

func (i EPersonaChange) String() string {
	if str, ok := _EPersonaChange_map[i]; ok {
		return str
	}
	return "EPersonaChange(" + strconv.FormatInt(int64(i), 10) + ")"
}
