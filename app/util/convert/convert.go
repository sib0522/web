package convert

import "strconv"

// stringをuint32に変換
func ToUint32(str string) (uint32, error) {
	parse, err := strconv.ParseUint(str, 10, 32)
	if err != nil {
		return 0, err
	}
	return uint32(parse), nil
}

// stringをuint64に変換
func ToUint64(str string) (uint64, error) {
	parse, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint64(parse), nil
}
