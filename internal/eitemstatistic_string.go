// Code generated by "stringer -type EItemStatistic -trimprefix EItemStatistic_"; DO NOT EDIT.

package internal

import "strconv"

const _EItemStatistic_name = "NumSubscriptionsNumFavoritesNumFollowersNumUniqueSubscriptionsNumUniqueFavoritesNumUniqueFollowersNumUniqueWebsiteViewsReportScoreNumSecondsPlayedNumPlaytimeSessionsNumCommentsNumSecondsPlayedDuringTimePeriodNumPlaytimeSessionsDuringTimePeriod"

var _EItemStatistic_index = [...]uint8{0, 16, 28, 40, 62, 80, 98, 119, 130, 146, 165, 176, 208, 243}

func (i EItemStatistic) String() string {
	if i < 0 || i >= EItemStatistic(len(_EItemStatistic_index)-1) {
		return "EItemStatistic(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _EItemStatistic_name[_EItemStatistic_index[i]:_EItemStatistic_index[i+1]]
}
