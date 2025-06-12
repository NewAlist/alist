package op_test

import (
	"testing"

	_ "github.com/NewAlist/alist/drivers"
	"github.com/NewAlist/alist/internal/op"
)

func TestDriverItemsMap(t *testing.T) {
	itemsMap := op.GetDriverInfoMap()
	if len(itemsMap) != 0 {
		t.Logf("driverInfoMap: %v", itemsMap)
	} else {
		t.Errorf("expected driverInfoMap not empty, but got empty")
	}
}
