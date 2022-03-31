package utils

import "encoding/json"

func MergeStruct(to interface{}, from interface{}) error {
	byteTo, err := json.Marshal(to)
	if err != nil {
		return err
	}

	byteFrom, err := json.Marshal(from)
	if err != nil {
		return err
	}

	mapTo := make(map[string]interface{})
	err = json.Unmarshal(byteTo, &mapTo)

	if err != nil {
		return err
	}

	mapFrom := make(map[string]interface{})
	err = json.Unmarshal(byteFrom, &mapFrom)

	if err != nil {
		return err
	}

	for k, v := range mapFrom {
		mapTo[k] = v
	}

	byteDest, err := json.Marshal(mapTo)

	if err != nil {
		return err
	}

	err = json.Unmarshal(byteDest, to)
	
	return err
}