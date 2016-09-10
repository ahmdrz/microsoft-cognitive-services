package face

import (
	"fmt"
)

func New(key string) (*Face, error) {
	if len(key) < 10 {
		return nil, fmt.Errorf("Invalid Key")
	}
	return &Face{
		BingKey: key,
	}, nil
}
