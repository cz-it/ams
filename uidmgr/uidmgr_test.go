/**
* Author: CZ cz.theng@gmail.com
 */

package uidmgr

import (
	"fmt"
	"testing"
)

func TestNewUID(t *testing.T) {
	uid, err := NewUID(uint8(1), uint16(1))
	if err != nil {
		t.Errorf("new UID Error %s", err.Error())
	}
	fmt.Println("uid is ", uid)

	uid, err = NewUID(uint8(1), uint16(1))
	if err != nil {
		t.Errorf("new UID Error %s", err.Error())
	}
	fmt.Println("second uid is ", uid)
}
