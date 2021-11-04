package nodejs

import (
	"encoding/json"
	"fmt"
)

func NpmDeps(packageLockJsonContents []byte) (map[string]string, error) {
	var f interface{}
	err := json.Unmarshal(packageLockJsonContents, &f)
	if err != nil {
		return nil, fmt.Errorf("unable to parse package-lock.json: %v", err)
	}

	raw := f.(map[string]interface{})
	return transform(raw["dependencies"].(map[string]interface{})), nil
}

func transform(input map[string]interface{}) map[string]string {
	transformed := make(map[string]string)
	for key, value := range input {
		dependency := value.(map[string]interface{})
		transformed[key] = fmt.Sprintf("%s", dependency["version"])
	}
	return transformed
}
