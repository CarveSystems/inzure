// Code generated by "stringer -type RecommendationCategory,RecommendationImpact"; DO NOT EDIT.

package inzure

import "fmt"

const _RecommendationCategory_name = "RecommendationCategoryUnsetRecommendationCategoryCostRecommendationCategoryHighAvailabilityRecommendationCategoryPerformanceRecommendationCategorySecurity"

var _RecommendationCategory_index = [...]uint8{0, 27, 53, 91, 124, 154}

func (i RecommendationCategory) String() string {
	if i >= RecommendationCategory(len(_RecommendationCategory_index)-1) {
		return fmt.Sprintf("RecommendationCategory(%d)", i)
	}
	return _RecommendationCategory_name[_RecommendationCategory_index[i]:_RecommendationCategory_index[i+1]]
}

const _RecommendationImpact_name = "RecommendationImpactUnsetRecommendationImpactLowRecommendationImpactMediumRecommendationImpactHigh"

var _RecommendationImpact_index = [...]uint8{0, 25, 48, 74, 98}

func (i RecommendationImpact) String() string {
	if i >= RecommendationImpact(len(_RecommendationImpact_index)-1) {
		return fmt.Sprintf("RecommendationImpact(%d)", i)
	}
	return _RecommendationImpact_name[_RecommendationImpact_index[i]:_RecommendationImpact_index[i+1]]
}
