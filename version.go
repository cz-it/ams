/**
* Author: CZ cz.theng@gmail.com
 */

package ams

import (
	"fmt"
)

const (
	major = 0
	minor = 1
	patch = 0
)

// Version return maglined's version
func Version() string {
	return fmt.Sprintf("maglined[%d.%d.%d]", major, minor, patch)
}
