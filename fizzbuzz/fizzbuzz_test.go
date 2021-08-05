package fizzbuzz

import "testing"

func TestFizzBuzz(t *testing.T) {
	cases := map[int]string{
		1:  "1",
		2:  "2",
		3:  "Fizz",
		6:  "Fizz",
		9:  "Fizz",
		5:  "Buzz",
		10: "Buzz",
		15: "FizzBuzz",
		30: "FizzBuzz",
	}

	for given, want := range cases {
		if get := Say(given); get != want {
			t.Errorf("Failed given[%d] %q with actual %q", given, get, want)
		}
	}
}

type stubInt struct {
	val int
}

func (s *stubInt) Intn(int) int {
	return s.val
}

func TestRandomFizzBuzz(t *testing.T) {
	want := "Fizz"

	get := RandomFizzBuzz(&stubInt{val: 2})

	if get != want {
		t.Errorf("Failed %q with actual %q", get, want)
	}
}

//	Datetime, Random ,Query Data
//	Use this style

type IntnFunc func(int) int

func (f IntnFunc) Intn(n int) int {
	return f(n)
}

func TestRandomFizzBuzz2(t *testing.T) {
	want := "Fizz"

	get := RandomFizzBuzz(IntnFunc(func(int) int { return 5 }))

	if get != want {
		t.Errorf("Failed %q with actual %q", get, want)
	}
}
