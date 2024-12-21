package util

import (
	"encoding/json"
	"fmt"

	"github.com/codingpierogi/golang-learning/streaming/types"
)

func SerializeMessage(nm types.NotifyMessage) ([]byte, error) {
	serialized, err := json.Marshal(nm)

	if err != nil {
		return nil, fmt.Errorf("failed to serialize message: %w", err)
	}

	return serialized, nil
}
