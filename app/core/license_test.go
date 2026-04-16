package core

import (
	"testing"
)

func TestValidateLicenseKey_유효한_키(t *testing.T) {
	validKeys := []string{
		"ABCD-1234-EFGH-5678",
		"0000-AAAA-ZZZZ-9999",
		"A1B2-C3D4-E5F6-G7H8",
	}
	for _, key := range validKeys {
		if !validateLicenseKey(key) {
			t.Errorf("유효한 키를 무효로 판단: %s", key)
		}
	}
}

func TestValidateLicenseKey_무효한_키(t *testing.T) {
	invalidKeys := []string{
		"",
		"ABCD-1234-EFGH",          // 파트 3개
		"ABCD-1234-EFGH-56789",    // 마지막 파트 5자리
		"abcd-1234-efgh-5678",     // 소문자
		"ABCD-12 4-EFGH-5678",     // 공백
		"ABCD-1234-EFGH-567!",     // 특수문자
	}
	for _, key := range invalidKeys {
		if validateLicenseKey(key) {
			t.Errorf("무효한 키를 유효로 판단: '%s'", key)
		}
	}
}
