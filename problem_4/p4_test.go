package p4

import "testing"

func TestSumToNA(t *testing.T) {
    // Test cases for Solution A
    cases := []struct {
        input    int
        expected int
    }{
        {1, 1},
        {5, 15},
        {10, 55},
    }

    for _, tc := range cases {
        result := sum_to_n_a(tc.input)
        if result != tc.expected {
            t.Errorf("For input %d in solution A, expected %d, but got %d", tc.input, tc.expected, result)
        }
    }
}

func TestSumToNB(t *testing.T) {
    // Test cases for Solution B
    cases := []struct {
        input    int
        expected int
    }{
        {1, 1},
        {5, 15},
        {10, 55},
    }

    for _, tc := range cases {
        result := sum_to_n_b(tc.input)
        if result != tc.expected {
            t.Errorf("For input %d in solution B, expected %d, but got %d", tc.input, tc.expected, result)
        }
    }
}

func TestSumToNC(t *testing.T) {
    // Test cases for Solution C
    cases := []struct {
        input    int
        expected int
    }{
        {1, 1},
        {5, 15},
        {10, 55},
        {7, 28},
    }

    for _, tc := range cases {
        result := sum_to_n_c(tc.input)
        if result != tc.expected {
            t.Errorf("For input %d in solution C, expected %d, but got %d", tc.input, tc.expected, result)
        }
    }
}
