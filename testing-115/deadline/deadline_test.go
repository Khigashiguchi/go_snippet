// ref: https://go-review.googlesource.com/c/go/+/202758/8/src/cmd/go/testdata/script/test_deadline.txt

package testing_test

import (
	"testing"
	"time"
)

func TestNoDeadline(t *testing.T) {
	d, ok := t.Deadline()
	if ok || !d.IsZero() {
		t.Fatalf("t.Deadline() = %v, %v; want 0, false", d, ok)
	}
}

func TestDeadlineWithinMinute(t *testing.T) {
	now := time.Now()
	d, ok := t.Deadline()
	if !ok || d.IsZero() {
		t.Fatalf("t.Deadline() = %v, %v; want nonzero deadline", d, ok)
	}
	if !d.After(now) {
		t.Fatalf("t.Deadline() = %v; want after start of test (%v)", d, now)
	}
	if d.Sub(now) > time.Minute {
		t.Fatalf("t.Deadline() = %v; want within one minute of start of test (%v)", d, now)
	}
}

func TestSubtestDeadlineWithinMinute(t *testing.T) {
	t.Run("sub", func(t *testing.T) {
		now := time.Now()
		d, ok := t.Deadline()
		if !ok || d.IsZero() {
			t.Fatalf("t.Deadline() = %v, %v; want nonzero deadline", d, ok)
		}
		if !d.After(now) {
			t.Fatalf("t.Deadline() = %v; want after start of test (%v)", d, now)
		}
		if d.Sub(now) > time.Minute {
			t.Fatalf("t.Deadline() = %v; want within one minute of start of test (%v)", d, now)
		}
	})
}
