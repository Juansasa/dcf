package dto

import (
	"strconv"
	"strings"
)

type StringInt int64
type StringFloat float64

func (number *StringInt) UnmarshalJSON(bs []byte) error {
	str := string(bs)
	str = strings.Trim(str, `"`)

	if strings.ToLower(str) == "none" {
		*number = 0
		return nil
	}

	parsedInt, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return err
	}
	*number = StringInt(parsedInt)
	return nil
}

func (number *StringFloat) UnmarshalJSON(bs []byte) error {
	str := string(bs)
	str = strings.Trim(str, `"`)

	if strings.ToLower(str) == "none" {
		*number = 0
		return nil
	}

	parsedInt, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return err
	}
	*number = StringFloat(parsedInt)
	return nil
}
