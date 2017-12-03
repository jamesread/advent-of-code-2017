package main

import "testing"

func testCapcha(t *testing.T, expected int, captcha string) {
	got := SolveCaptcha(captcha)

	if got != expected {
		t.Errorf("Failed to solve capcha. Expected %i, got %i, captcha %s", expected, got, captcha)
	}
}

func TestExample1(t *testing.T) {
	testCapcha(t, 3, "1122")
	testCapcha(t, 4, "1111")
	testCapcha(t, 0, "1234")
	testCapcha(t, 9, "91212129")
}
