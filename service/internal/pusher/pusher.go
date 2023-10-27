package pusher

import (
	"github.com/rs/zerolog"
	"github.com/samber/lo"
	"os"
)

var Logger = zerolog.New(os.Stderr).With().Timestamp().Logger()

func FilterAllExceptFirst[K any](collection []K, predicate func(v K) bool) []K {
	hit, found := lo.Find(collection, func(v K) bool {
		return predicate(v)
	})

	if found {
		r := lo.Filter(collection, func(v K, _ int) bool {
			return !predicate(v)
		})

		return append([]K{hit}, r...)
	}

	return collection
}
