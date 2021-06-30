package goconfig

import (
	"encoding/json"
	"fmt"

	ini "gopkg.in/ini.v1"
)

// Save uses the same json taggers but convert it to INI for readability
func Save(v interface{}, saveTo string) error {
	var b []byte
	// marshal to json then to map
	b, err := json.Marshal(&b)
	if err != nil {
		return err
	}
	// unmarshal to map
	kv := make(map[string]interface{})
	err = json.Unmarshal(b, &kv)
	if err != nil {
		return err
	}

	// iterate over map setting key => value in ini file
	cfg := ini.Empty()
	for k, v := range kv {
		cfg.Section("").NewKey(k, fmt.Sprint(v))
	}
	return cfg.SaveTo(saveTo)
}

// Load uses the same json taggers but reads from INI for readability
func Load(v interface{}, loadFrom string) error {
	cfg, err := ini.Load(loadFrom)
	if err != nil {
		return err
	}
	return cfg.MapTo(v)
}
