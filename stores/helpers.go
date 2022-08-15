package stores

import (
	"encoding/json"
	"fmt"
)

func EncodeToJSON(obj interface{}) []byte {
	bytes, err := json.Marshal(obj)
	if err != nil {
		fmt.Println(err.Error())
		return []byte{}
	}
	return bytes
}

func DecodeFromJSON(data []byte, obj interface{}) {
	json.Unmarshal(data, obj)
}
