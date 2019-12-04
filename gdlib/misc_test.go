package gdlib

import "testing"

func TestGenerateUUID(t *testing.T) {
	uuid := GenerateUUID()

	if uuid == "" || len(uuid) == 0 {
		t.Fatal("Invalid UUID")
	}
}
