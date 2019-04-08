package fizzbuzz_test

import (
	"testing"

	"github.com/hgsgtk/go-snippets/testing-codes/fizzbuzz"
)

func TestGetMsg(t *testing.T) {
	tests := map[string]struct {
		num      int
		expected string
	}{
		"15で割り切れる場合FizzBuzz": {
			num:      45,
			expected: "FizzBuzz",
		},
		"5で割り切れる場合Buzz": {
			num:      40,
			expected: "Buzz",
		},
		"3で割り切れる場合Buzz": {
			num:      39,
			expected: "Fizz",
		},
		"15,5,3で割り切れない場合そのまま": {
			num:      37,
			expected: "37",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if actual := fizzbuzz.GetMsg(tt.num); tt.expected != actual {
				t.Errorf("Run is expected '%s', but got '%s'", tt.expected, actual)
			}
		})
	}
}

func TestGetMsg2(t *testing.T) {
	var num int
	var want string

	num = 15
	want = "FizzBuzz"
	if got := fizzbuzz.GetMsg(num); want != got {
		t.Fatalf("GetMsg(%d) = %s, want %s", num, want, got)
	}

	num = 5
	want = "Buzz"
	if got := fizzbuzz.GetMsg(num); want != got {
		t.Fatalf("GetMsg(%d) = %s, want %s", num, want, got)
	}

	num = 3
	want = "Fizz"
	if got := fizzbuzz.GetMsg(num); want != got {
		t.Fatalf("GetMsg(%d) = %s, want %s", num, want, got)
	}

	num = 1
	want = "1"
	if got := fizzbuzz.GetMsg(num); want != got {
		t.Fatalf("GetMsg(%d) = %s, want %s", num, want, got)
	}
}
