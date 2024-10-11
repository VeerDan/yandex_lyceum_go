package main

import "testing"

func TestSum(t *testing.T) {
    cases := []struct {
		name string
        values []int
        want int
    }{
        {
			name: "positive",
            values: []int{1, 2},
            want: 3,
        },
        {
			name: "negative",
            values: []int{-1, -2},
            want: -3,
        },
        {
			name: "np",
            values: []int{-1, 2},
            want: 1,
        },
        {
			name: "pn",
            values: []int{2, -1},
            want: 1,
        },
    }

    for _, tc := range cases {
        t.Run(tc.name, func(t *testing.T) {
			a := tc.values[0]
			b := tc.values[1]
			got := Sum(a, b)
			if got != tc.want {
				t.Errorf("Sum(%v) = %v; want %v", tc.values, got, tc.want)
			}
		})
    }
}