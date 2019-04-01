package date

import (
	"encoding/json"
	"testing"
)

func TestJsonMarshall(t *testing.T) {
	d := New(2001, 2, 3)

	b, err := json.Marshal(d)
	if err != nil {
		t.Logf("failed to marshall date: %s", err)
		t.FailNow()
	}

	expected := `"2001-02-03"`
	if string(b) != expected {
		t.Logf("date marshalled incorrectly: expected %s, got %s", expected, string(b))
		t.Fail()
	}
}

func TestJsonUnmarshall(t *testing.T) {
	s := `"2001-02-03"`

	var d Date
	if err := json.Unmarshal([]byte(s), &d); err != nil {
		t.Logf("failed to unmarshall date: %s", err)
		t.FailNow()
	}

	if d.String() != s {
		t.Logf("date unmarshall incorrect: expected %s, got %s", s, d.String())
		t.Fail()
	}
}
