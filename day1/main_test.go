package main

import "testing"

func testCaptchaResult(t *testing.T, expected int, got *SolvedCaptcha) {
	if expected != got.Result {
		t.Errorf("Failed to solve capcha. Expected %d, got %d, captcha %s", expected, got.Result, got.Captcha)
	}
}

func TestChecksumNextDigits(t *testing.T) {
	testCaptchaResult(t, 3, SolveCaptchaNext("1122"))
	testCaptchaResult(t, 4, SolveCaptchaNext("1111"))
	testCaptchaResult(t, 0, SolveCaptchaNext("1234"))
	testCaptchaResult(t, 9, SolveCaptchaNext("91212129"))
}

func TestChecksumHalfwayDigits(t *testing.T) {
	testCaptchaResult(t, 6, SolveCaptchaHalfway("1212"))
	testCaptchaResult(t, 0, SolveCaptchaHalfway("1221"))
	testCaptchaResult(t, 4, SolveCaptchaHalfway("123425"))
	testCaptchaResult(t, 12, SolveCaptchaHalfway("123123"))
	testCaptchaResult(t, 4, SolveCaptchaHalfway("12131415"))
}
