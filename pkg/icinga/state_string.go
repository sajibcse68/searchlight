// Code generated by "stringer -type=State ../../pkg/icinga/types.go"; DO NOT EDIT.

package icinga

import "fmt"

const _State_name = "OKWARNINGCRITICALUNKNOWN"

var _State_index = [...]uint8{0, 2, 9, 17, 24}

func (i State) String() string {
	if i < 0 || i >= State(len(_State_index)-1) {
		return fmt.Sprintf("State(%d)", i)
	}
	return _State_name[_State_index[i]:_State_index[i+1]]
}
