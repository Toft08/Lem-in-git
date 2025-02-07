package main

import (
	"reflect"
	"testing"
)

func TestFindPaths(t *testing.T) {
	tests := []struct {
		name   string
		links  map[string][]string
		start  string
		end    string
		expect [][]string
	}{
		{
			name: "Simple case",
			links: map[string][]string{
				"0": {"1", "2"},
				"1": {"0", "2", "3"},
				"2": {"0", "1", "3"},
				"3": {"1", "2"},
			},
			start: "0",
			end:   "3",
			expect: [][]string{
				{"0", "1", "2", "3"},
				{"0", "1", "3"},
				{"0", "2", "1", "3"},
				{"0", "2", "3"},
			},
		},
		{
			name: "No path",
			links: map[string][]string{
				"0": {"1"},
				"1": {"0"},
				"2": {"3"},
				"3": {"2"},
			},
			start:  "0",
			end:    "3",
			expect: [][]string{},
		},
		{
			name: "Single node",
			links: map[string][]string{
				"0": {},
			},
			start:  "0",
			end:    "0",
			expect: [][]string{{"0"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findPaths(tt.links, tt.start, tt.end)
			if len(result) == 0 && len(tt.expect) == 0 {
				// Both are empty, test passes
				return
			}
			if !reflect.DeepEqual(result, tt.expect) {
				t.Errorf("findPaths(%v, %s, %s) = %v; want %v", tt.links, tt.start, tt.end, result, tt.expect)
			}
		})
	}
}
