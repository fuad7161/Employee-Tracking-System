package util

import (
	"strings"
	"testing"
)

func TestRandomInt(t *testing.T) {
	min, max := int64(1), int64(100)
	for i := 0; i < 100; i++ {
		result := RandomInt(min, max)
		if result < min || result > max {
			t.Errorf("RandomInt(%d, %d) = %d; out of bounds", min, max, result)
		}
	}
}

func TestRandomString(t *testing.T) {
	n := 10
	result := RandomString(n)

	if len(result) != n {
		t.Errorf("RandomString(%d) generated string with incorrect length: got %d, want %d", n, len(result), n)
	}

	for _, ch := range result {
		if !strings.Contains(alphabet, string(ch)) {
			t.Errorf("RandomString(%d) contains invalid character: %c", n, ch)
		}
	}
}

func TestRandomOwner(t *testing.T) {
	result := RandomOwner()

	if len(result) != 6 {
		t.Errorf("RandomOwner() generated owner of incorrect length: got %d, want 6", len(result))
	}

	for _, ch := range result {
		if !strings.Contains(alphabet, string(ch)) {
			t.Errorf("RandomOwner() contains invalid character: %c", ch)
		}
	}
}

func TestRandomEmail(t *testing.T) {
	result := RandomEmail()

	if !strings.HasSuffix(result, "@brainstation-23.com") {
		t.Errorf("RandomEmail() does not have the correct suffix: got %s", result)
	}

	localPart := strings.Split(result, "@")[0]
	if len(localPart) != 6 {
		t.Errorf("RandomEmail() generated email with incorrect local part length: got %d, want 6", len(localPart))
	}

	for _, ch := range localPart {
		if !strings.Contains(alphabet, string(ch)) {
			t.Errorf("RandomEmail() local part contains invalid character: %c", ch)
		}
	}
}
