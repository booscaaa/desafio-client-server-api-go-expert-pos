package dto

import (
	"encoding/json"
)

func Mapper[T any, J any](from T, to J) (*J, error) {
	bytes, err := json.Marshal(from)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, &to)

	if err != nil {
		return nil, err
	}

	return &to, nil
}
