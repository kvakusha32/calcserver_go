package main_test

import (
	"testing"
)

func TestCalc(t *testing.T) {
	testCases := []struct {
		expression     string
		expectedResult string
		expectError    bool
	}{
		{
			expression:     "2 + 10 * (2 - 8)",
			expectedResult: "-58.000000",
			expectError:    false,
		},
		{
			expression:     "20 / 2",
			expectedResult: "10.000000",
			expectError:    false,
		},
		{
			expression:     "3 + g",
			expectedResult: "",
			expectError:    true,
		},
		{
			expression:     " 10 / 0",
			expectedResult: "",
			expectError:    true,
		},
		{
			expression:     "1 + 1)",
			expectedResult: "",
			expectError:    true,
		},
	}

	for _, tc := range testCases {
		result, err := Calc(tc.expression)
		if (err != nil) != tc.expectError {
			t.Fatalf("For expression %s, expected error: %v, got: %v", tc.expression, tc.expectError, err)
		}
		if !tc.expectError && result != tc.expectedResult {
			t.Fatalf("For expression %s, expected result: %s, got: %s", tc.expression, tc.expectedResult, result)
		}
	}
}
