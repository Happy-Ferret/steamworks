// Code generated by "stringer -type EUGCQuery -trimprefix EUGCQuery_"; DO NOT EDIT.

package internal

import "strconv"

const _EUGCQuery_name = "RankedByVoteRankedByPublicationDateAcceptedForGameRankedByAcceptanceDateRankedByTrendFavoritedByFriendsRankedByPublicationDateCreatedByFriendsRankedByPublicationDateRankedByNumTimesReportedCreatedByFollowedUsersRankedByPublicationDateNotYetRatedRankedByTotalVotesAscRankedByVotesUpRankedByTextSearchRankedByTotalUniqueSubscriptionsRankedByPlaytimeTrendRankedByTotalPlaytimeRankedByAveragePlaytimeTrendRankedByLifetimeAveragePlaytimeRankedByPlaytimeSessionsTrendRankedByLifetimePlaytimeSessions"

var _EUGCQuery_index = [...]uint16{0, 12, 35, 72, 85, 126, 165, 189, 234, 245, 266, 281, 299, 331, 352, 373, 401, 432, 461, 493}

func (i EUGCQuery) String() string {
	if i < 0 || i >= EUGCQuery(len(_EUGCQuery_index)-1) {
		return "EUGCQuery(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _EUGCQuery_name[_EUGCQuery_index[i]:_EUGCQuery_index[i+1]]
}
