package configs

import (
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

type Configurations struct {
	Monitor Monitor `yaml:"monitor"`
}

type Monitor struct {
	Port string `yaml:"port"`
}

var fileCnf *Configurations = nil

func GetConfigFromYaml() (*Configurations, error) {
	if fileCnf == nil {
		yamlFile, err := ioutil.ReadFile("config/config.yaml")
		if err != nil {
			return nil, err
		}
		temp := &Configurations{}
		err = yaml.Unmarshal(yamlFile, temp)
		if err != nil {
			return nil, err
		}
		fileCnf = temp
	}
	return fileCnf, nil
}
