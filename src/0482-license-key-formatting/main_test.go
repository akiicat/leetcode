package main
import "testing"
import . "main/pkg/testing_helper"

func TestLicenseKeyFormatting(_t *testing.T) {
  s, k, o := "5F3Z-2e-9-w", 4, "5F3Z-2E9W"
  T(_t, S(s, k), S(licenseKeyFormatting(s, k)), S(o))

  s, k, o = "2-5g-3-J", 2, "2-5G-3J"
  T(_t, S(s, k), S(licenseKeyFormatting(s, k)), S(o))

  s, k, o = "--a-a-a-a--", 2, "AA-AA"
  T(_t, S(s, k), S(licenseKeyFormatting(s, k)), S(o))
}
