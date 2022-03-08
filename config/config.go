package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var Cards *CardsCrawl

type CardsCrawl struct {
	CardsList []string `yaml:"cards"`
}

func LoadYamlConfig() {
	t := CardsCrawl{}
	data, err := ioutil.ReadFile("card.yml")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = yaml.Unmarshal(data, &t)
	Cards = &t
}
