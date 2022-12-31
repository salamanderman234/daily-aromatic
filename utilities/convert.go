package utility

import "encoding/json"

func ConvertTo(source interface{}, target *interface{}) (interface{}, error) {
	var targetReturn interface{}
	temp, err := json.Marshal(source)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(temp, &targetReturn)
	return targetReturn, nil
}
