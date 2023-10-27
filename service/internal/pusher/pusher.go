package pusher

import (
	"github.com/rs/zerolog"
	"github.com/samber/lo"
	"os"
)

var Logger = zerolog.New(os.Stderr).With().Timestamp().Logger()

// FilterAllExceptFirst filters a collection by a predicate, but ensures that the first hit is always included at the top
// of the result. This is useful for messages that always contain the full state about something. By moving it to the top
// the web client receives the full state first, instead of last.
// Think about a progress indicator 0...100 (in reverse), we keep 100, but filter out 0...99, because they do not contain
// any relevant state and the client receives the progress latest state first.
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
