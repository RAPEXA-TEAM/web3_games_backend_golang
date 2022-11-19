package web3

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"os/exec"
	"time"
	"web3game/contracts/model"
)

func ProvideGanacheAccounts() (model model.GanacheAccounts) {
	var err error
	err = exec.Command("ganache", "--account_keys_path", "/home/khoujani/keys.json").Start()
	time.Sleep(time.Second * 2)
	homeDir, err := os.UserHomeDir()
	accountFilepath := homeDir + "/keys.json"
	file, err := os.Open(accountFilepath)
	defer file.Close()
	if err != nil {
		time.Sleep(2 * time.Second)
		println("ganache-cli is not running, try again \n" + err.Error())
		ProvideGanacheAccounts()
	}
	var f interface{}
	byteResult, _ := ioutil.ReadAll(file)
	err = json.Unmarshal(byteResult, &f)
	if err != nil {
		println("Error parsing json \n" + err.Error())
		return model
	}
	itemsMap := f.(map[string]interface{})
	for k, v := range itemsMap {
		for i, v1 := range v.(map[string]interface{}) {
			if k == "addresses" {
				model.Addresses = append(model.Addresses, v1.(string))
			} else {
				model.PrivateKeys = append(model.PrivateKeys, v1.(string)[2:])
				println("PublicKey: " + string(i) + " <--->" + " PrivateKey: " + v1.(string)[2:])
			}
		}
	}
	if err != nil {
		panic(err)
	} else {
		println("Providing accounts successfully")
	}
	return
}
