package filter

import (
	"encoding/json"
	"testing"
)

func TestNotionFilterAnd(t *testing.T) {
	filter := NewFilter(And(
		TextProp__contains("taskName", "Learn Rust"),
		SelectProp__equal("status", "completed"),
	))

	b, err := json.Marshal(filter)
	if err != nil {
		t.Fatalf("failed to marshal filter: %v", err)
	}

	expected := `{"filter":{"and":[{"property":"taskName","text":{"contains":"Learn Rust"}},{"property":"status","select":{"equals":"completed"}}]}}`

	if string(b) != expected {
		t.Fatalf("expected filter to be %s, got %s", expected, b)
	}
}

func TestNotionFilterOr(t *testing.T) {
	filter := NewFilter(
		Or(
			CheckboxProp__equals("Seen", false),
			NumProp__greater_than("Yearly visitor count", 100000),
		),
	)

	b, err := json.Marshal(filter)
	if err != nil {
		t.Fatalf("failed to marshal filter: %v", err)
	}

	expected := `{"filter":{"or":[{"checkbox":{"equals":false},"property":"Seen"},{"number":{"greater_than":100000},"property":"Yearly visitor count"}]}}`

	if string(b) != expected {
		t.Fatalf("expected filter to be %s, got %s", expected, b)
	}
}
