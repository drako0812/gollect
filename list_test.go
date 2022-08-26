package gollect

import (
	"testing"
)

func TestIsListGeneralCollector(t *testing.T) {
	list := NewListFromData[int64](1, 2, 3, 4, 5)

	if _, isGeneralCollector := interface{}(&list).(GeneralCollector[int64]); !isGeneralCollector {
		t.Fatalf("List[int64] should be a GeneralCollector[int64]")
	}
}
