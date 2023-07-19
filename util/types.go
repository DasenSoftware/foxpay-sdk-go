package util

func IntPtr(v int) *int {
	return &v
}

func Int64Ptr(v int64) *int64 {
	return &v
}

func Int32Ptr(v int32) *int32 {
	return &v
}

func UintPtr(v uint) *uint {
	return &v
}

func Uint64Ptr(v uint64) *uint64 {
	return &v
}

func Float64Ptr(v float64) *float64 {
	return &v
}

func BoolPtr(v bool) *bool {
	return &v
}

func StringPtr(v string) *string {
	return &v
}

func StringValues(ptr *string) string {
	var val string
	if ptr != nil {
		val = *ptr
	}
	return val
}

func Int32Values(ptr *int32) int32 {
	var val int32
	if ptr != nil {
		val = *ptr
	}
	return val
}

func Int64Values(ptr *int64) int64 {
	var val int64
	if ptr != nil {
		val = *ptr
	}
	return val
}
