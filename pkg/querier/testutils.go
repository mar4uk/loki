package querier

import (
	"github.com/grafana/dskit/flagext"

	"github.com/mar4uk/loki/pkg/util/validation"
)

func DefaultLimitsConfig() validation.Limits {
	limits := validation.Limits{}
	flagext.DefaultValues(&limits)
	return limits
}
