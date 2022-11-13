package doubler

import "testing"

func TestDoubler(t *testing.T) {
	testCases := []struct {
		in  int32
		out int32
	}{
		{
			in:  10,
			out: 20,
		},
		{
			in:  50,
			out: 100,
		},
		{
			in:  100,
			out: 200,
		},
	}

	for _, testCase := range testCases {
		got := Double(testCase.in)

		if got != testCase.out {
			t.Errorf("Double(%d) = %d, want: %d\n", testCase.in, got, testCase.out)
		}
	}
}
