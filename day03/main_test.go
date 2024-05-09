package main

import "testing"

func TestIsNeighbour(t *testing.T) {
	type test struct {
		name   string
		part   enginePart
		symbol symbol
		want   bool
	}

	tests := []test{
		{name: "above", part: enginePart{Row: 1, Start: 1, End: 5}, symbol: symbol{Row: 0, Column: 3}, want: true},
		{name: "below", part: enginePart{Row: 1, Start: 1, End: 5}, symbol: symbol{Row: 2, Column: 3}, want: true},
		{name: "topleft", part: enginePart{Row: 1, Start: 1, End: 5}, symbol: symbol{Row: 0, Column: 0}, want: true},
		{name: "topright", part: enginePart{Row: 1, Start: 1, End: 5}, symbol: symbol{Row: 0, Column: 6}, want: true},
		{name: "bottomright", part: enginePart{Row: 1, Start: 1, End: 5}, symbol: symbol{Row: 2, Column: 6}, want: true},
		{name: "bottomleft", part: enginePart{Row: 1, Start: 1, End: 5}, symbol: symbol{Row: 2, Column: 0}, want: true},
		{name: "far left", part: enginePart{Row: 2, Start: 2, End: 5}, symbol: symbol{Row: 1, Column: 0}, want: false},
		{name: "far above", part: enginePart{Row: 2, Start: 2, End: 5}, symbol: symbol{Row: 0, Column: 3}, want: false},
		{name: "far right", part: enginePart{Row: 1, Start: 1, End: 5}, symbol: symbol{Row: 8, Column: 3}, want: false},
		{name: "far bottom", part: enginePart{Row: 1, Start: 1, End: 5}, symbol: symbol{Row: 0, Column: 8}, want: false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := isNeighbour(&tc.part, &tc.symbol)
			if got != tc.want {
				t.Fatalf("expected: %v, got: %v", tc.want, got)
			}
		})
	}
}
