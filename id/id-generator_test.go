package id

import "testing"

func TestGenerateId(t *testing.T) {
	const idLength = 10

	id := GenerateId(idLength)

	if len(id) != idLength {
		t.Errorf("Expected id length of %d, got %d", idLength, len(id))
	}

	for _, char := range id {
		if (char < 'A' || char > 'Z') && (char < 'a' || char > 'z') {
			t.Errorf("Generated id contains invalid character: %c", char)
		}
	}
}

func TestGenerateEmptyId(t *testing.T) {
	const idLength = 0

	id := GenerateId(idLength)

	if len(id) != idLength {
		t.Errorf("Expected id length of %d, got %d", idLength, len(id))
	}

	for _, char := range id {
		if (char < 'A' || char > 'Z') && (char < 'a' || char > 'z') {
			t.Errorf("Generated id contains invalid character: %c", char)
		}
	}
}
