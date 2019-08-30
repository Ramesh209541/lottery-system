package main

import (
	"encoding/json"
	"io/ioutil"
	log "logger"
)

/*GethttpConfigFileValues is used for parse the values from config file*/
func parseConfigFileValues(configFile string, lotteryConfig *lotteryConfigType) error {
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Error.Println("Unable to Fetch Configuration File due to Error: ", err)
		return err
	}

	err = json.Unmarshal(data, lotteryConfig)
	if err != nil {
		log.Error.Println("Unable to Fetch Config File Data. Please Check file data format. Error: ", err)
		return err
	}
	return nil
}
