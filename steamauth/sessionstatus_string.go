// Code generated by "stringer -type SessionStatus -trimprefix Status"; DO NOT EDIT.

package steamauth

import "strconv"

const _SessionStatus_name = "UnknownOKUserNotConnectedtoSteamNoLicenseOrExpiredVACBannedLoggedInElsewhereVACCheckTimedOutCanceledInvalidAlreadyUsedInvalidPublisherIssuedBan"

var _SessionStatus_index = [...]uint8{0, 7, 9, 32, 50, 59, 76, 92, 100, 118, 125, 143}

func (i SessionStatus) String() string {
	if i < 0 || i >= SessionStatus(len(_SessionStatus_index)-1) {
		return "SessionStatus(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _SessionStatus_name[_SessionStatus_index[i]:_SessionStatus_index[i+1]]
}
