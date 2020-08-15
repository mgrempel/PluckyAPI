package confighelper

import (
	models "PluckyAPI/Models"
	"os"

	yaml "gopkg.in/yaml.v2"
)

//GetConfigValues Accesses the api config file and returns the necessary values
func GetConfigValues() (config models.Config, err error) {
	config = models.Config{}

	//Open reader for config file
	file, err := os.Open("config.yml")
	defer file.Close()
	if err != nil {
		return models.Config{}, err
	}

	//Unmarshal file to config struct
	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return models.Config{}, err
	}

	return config, nil
}
