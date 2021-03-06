package api_payloads

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// IntAsString evens out deserialization of numbers represented in JSON document sometimes as int and simetimes as string. Ugliness
type IntAsString int

func (p *IntAsString) UnmarshalJSON(data []byte) (err error) {
	// unmarshall int or string
	var obj interface{}
	if err = json.Unmarshal(data, &obj); err != nil {
		return fmt.Errorf("IntAsString.UnmarshalJSON error: %v", err)
	}

	// decide the type and convert to number
	switch val := obj.(type) {
	case float64: // json.Unmarshall recognizes numbers as float64
		*p = IntAsString(val)
	case string:
		intval, err := strconv.Atoi(val)
		if err != nil {
			return fmt.Errorf("IntAsString.UnmarshalJSON Atoi conversion error: %v", err)
		}
		*p = IntAsString(intval)
	default:
		return fmt.Errorf("IntAsString.UnmarshalJSON input should be string or int, got {%T}", obj)
	}
	return nil
}
