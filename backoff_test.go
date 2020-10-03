package backoff

import (
	"testing"
	"time"
)

func TestBO(t *testing.T) {
	ini := uint64(1000) //<< start at 1 sec (1000ms)

	bo, err := New(ini, ini, ini, nil)

	if err != nil {
		t.Error(err)

		return
	}

	for i := 0; i < 5; i++ {
		now := time.Now()

		bo.Backoff()

		l8r := time.Since(now)

		if l8r < time.Duration(ini) {
			t.Errorf("failed at %d with %+v", i, l8r)
		}
	}
}
