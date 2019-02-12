package luhn

import (
	"testing"
)

func TestCheckBankCardNumber(t *testing.T) {
	if b, e := CheckBankCardNumber("test"); b || e == nil {
		t.Error("function error")
	}
	if b, e := CheckBankCardNumber("123"); b || e == nil {
		t.Error("function error")
	}
	if b, e := CheckBankCardNumber("-123"); b || e == nil {
		t.Error("function error")
	}
	if b, e := CheckBankCardNumber("1233234535562345"); b || e != nil {
		t.Error("function error")
	}
	if b, e := CheckBankCardNumber("7789001182735696"); !b || e != nil {
		t.Error("function error")
	}
}

func TestLuhn(t *testing.T) {
	if b, e := CheckLuhn("test"); b || e == nil {
		t.Error("function error")
	}
	if b, e := CheckLuhn("123"); b || e != nil {
		t.Error("function error")
	}
	if b, e := CheckLuhn("125"); !b || e != nil {
		t.Error("function error")
	}
	if b, e := CheckLuhn("-123"); b || e == nil {
		t.Error("function error")
	}
	if b, e := CheckLuhn("1233234535562345"); b || e != nil {
		t.Error("function error")
	}
	if b, e := CheckLuhn("7789001182735696"); !b || e != nil {
		t.Error("function error")
	}
}
