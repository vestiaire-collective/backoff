package linbo

import (
	"github.com/vestiaire-collective/backoff"
)

func New(initial uint64, max uint64, inc uint64) (*backoff.Backoff, error) {
	return backoff.New(initial, initial, max, func(current uint64) uint64 {
		return current + inc
	})
}
