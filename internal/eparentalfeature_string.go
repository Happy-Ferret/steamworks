// Code generated by "stringer -type EParentalFeature -trimprefix EParentalFeature_"; DO NOT EDIT.

package internal

import "strconv"

const _EParentalFeature_name = "EFeatureInvalidEFeatureStoreEFeatureCommunityEFeatureProfileEFeatureFriendsEFeatureNewsEFeatureTradingEFeatureSettingsEFeatureConsoleEFeatureBrowserEFeatureParentalSetupEFeatureLibraryEFeatureTestEFeatureMax"

var _EParentalFeature_index = [...]uint8{0, 15, 28, 45, 60, 75, 87, 102, 118, 133, 148, 169, 184, 196, 207}

func (i EParentalFeature) String() string {
	if i < 0 || i >= EParentalFeature(len(_EParentalFeature_index)-1) {
		return "EParentalFeature(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _EParentalFeature_name[_EParentalFeature_index[i]:_EParentalFeature_index[i+1]]
}
