package module

import (
	"fmt"

	a "errorwrap/module/a"
)

var Aerror = fmt.Errorf("this module error, sub error :%w", a.ErrorA)
