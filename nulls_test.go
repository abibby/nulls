package nulls_test

import (
	"testing"

	"github.com/abibby/nulls"
)

func TestStringIsNull(t *testing.T) {
	var s *nulls.String = nil

	if s.IsNull() == false {
		t.Fatalf("expected IsNull to return true")
	}
}

func TestStringOk(t *testing.T) {
	var s *nulls.String = nil

	str, ok := s.Ok()
	if ok == true {
		t.Fatalf("expected ok to return false")
	}
	if str != "" {
		t.Fatalf("expected IsNull to return empty string")
	}
}

func TestStringValue(t *testing.T) {
	s := nulls.NewString("test")

	if s.Value() != "test" {
		t.Fatalf("expected Value to return \"test\"")
	}
}
