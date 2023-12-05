package config

import (
	"encoding/json"
	"os"
	"strings"
)


// The Load function reads the contents of a file, 
// decodes the JSON it contains into a map, and uses the
// map to create a DefaultConfig value.
func Load(fileName string) (config Configuration, err error) {
	var data []byte
	data, err = os.ReadFile(fileName)
	if err == nil {
		decoder := json.NewDecoder(strings.NewReader(string(data)))
		m := map[string]interface{}{}
		err = decoder.Decode(&m)
		if err == nil {
			config = &DefaultConfig{configData: m}
		}
	}
	return
}
