// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package names

import (
	"strings"
	"testing"
)

//cspell:disable

func TestLabelName(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"Lowercase", "myproject", "myproject"},
		{"Uppercase", "MYPROJECT", "myproject"},
		{"MixedCase", "myProject", "my-project"},
		{"MixedCaseEnd", "myProjecT", "my-project"},
		{"TitleCase", "MyProject", "my-project"},
		{"TitleCaseEnd", "MyProjecT", "my-project"},
		{"WithDot", "my.project", "my-project"},
		{"WithDotTitleCase", "My.Project", "my-project"},
		{"WithHyphen", "my-project", "my-project"},
		{"WithHyphenTitleCase", "My-Project", "my-project"},
		{"StartWithNumber", "1myproject", "1myproject"},
		{"EndWithNumber", "myproject2", "myproject2"},
		{"MixedWithNumbers", "my2Project3", "my2-project3"},
		{"SpecialCharacters", "my_project!@#", "my-project"},
		{"EmptyString", "", ""},
		{"DotOnly", ".", ""},
		{"OnlySpecialCharacters", "@#$%^&*", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := LabelName(tt.input)
			if result != tt.expected {
				t.Errorf("LabelName(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestLabelNameEdgeCases(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{"SingleCharacter", "A", "a"},
		{"TwoCharacters", "Ab", "ab"},
		{"StartEndHyphens", "-abc-", "abc"},
		{"LongString",
			"ThisIsOneVeryLongStringThatExceedsTheSixtyThreeCharacterLimitForRFC1123LabelNames",
			"this-is-one-very-long-string-that-exceeds-the-sixty-three-character-limit-for-rfc1123-label-names"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := LabelName(tt.input)
			if result != tt.expected {
				t.Errorf("LabelName(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestValidateLabelName(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{"SingleLowercase", "a", false},
		{"MaxLength", strings.Repeat("a", 63), false},
		{"WithHyphen", "a-b-c", false},
		{"EmptyString", "", true},
		{"SingleUppercase", "Z", true},
		{"TooLong", strings.Repeat("a", 64), true},
		{"StartWithUppercase", "Abcdef", true},
		{"EndsWithUppercase", "abcdefG", true},
		{"InvalidSingleHyphen", "-", true},
		{"InvalidSingleSymbol", "!", true},
		{"LabelStartingWithHyphen", "-abc", true},
		{"LabelEndingWithHyphen", "abc-", true},
		{"LabelWithInvalidCharacters", "ab#cd", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateLabelName(tt.input)
			if tt.wantErr && err == nil {
				t.Errorf("expected error for input %q, got none", tt.input)
			}
			if !tt.wantErr && err != nil {
				t.Errorf("expected no error for input %q, got %q", tt.input, err)
			}
		})
	}
}

//cspell:enable
