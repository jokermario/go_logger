// Code generated by "stringer -type=LogAttribute"; DO NOT EDIT.

package go_logger

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Info-1]
	_ = x[Warn-2]
	_ = x[Error-4]
}

const (
	_LogAttribute_name_0 = "InfoWarn"
	_LogAttribute_name_1 = "Error"
)

var (
	_LogAttribute_index_0 = [...]uint8{0, 4, 8}
)

func (i LogAttribute) String() string {
	switch {
	case 1 <= i && i <= 2:
		i -= 1
		return _LogAttribute_name_0[_LogAttribute_index_0[i]:_LogAttribute_index_0[i+1]]
	case i == 4:
		return _LogAttribute_name_1
	default:
		return "LogAttribute(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}