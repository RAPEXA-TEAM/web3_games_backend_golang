package test

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"os/exec"
	"time"
	"web3game/contracts/model"
)

var jsonBytes = []byte(`{"addresses":{"0x524974a11d50228aaedfce6f5dd7bd41160198ac":"0x524974a11d50228aaedfce6f5dd7bd41160198ac","0x025beebf7b8d4e7a23e4c1ae7440d945a3353495":"0x025beebf7b8d4e7a23e4c1ae7440d945a3353495","0xf8dab13f97c8cd6e1c5c205f6e46358823b1bf57":"0xf8dab13f97c8cd6e1c5c205f6e46358823b1bf57","0x73d7d7a597246521e7b440752b200b12e30eca2a":"0x73d7d7a597246521e7b440752b200b12e30eca2a","0xdb1080ed09ff4405d9abc550619f40341073d12c":"0xdb1080ed09ff4405d9abc550619f40341073d12c","0x328fc2a3a266a4a9d2ae0f62ac710cf9522f6ab3":"0x328fc2a3a266a4a9d2ae0f62ac710cf9522f6ab3","0x25fe79d56c0ec39183b37fbbbbee03fcb3e29c18":"0x25fe79d56c0ec39183b37fbbbbee03fcb3e29c18","0x8da2ff45e5ab776e775f7df122856503ac37f5e2":"0x8da2ff45e5ab776e775f7df122856503ac37f5e2","0x5ce0b726c8f3abc4b46f3f04a157830babb60248":"0x5ce0b726c8f3abc4b46f3f04a157830babb60248","0xa12c2b1cb5ac67f5fa336b1645c6e88cdeca13a0":"0xa12c2b1cb5ac67f5fa336b1645c6e88cdeca13a0"},"private_keys":{"0x524974a11d50228aaedfce6f5dd7bd41160198ac":"0x800cfbeddbd76c9c8708785a6ce537ab6e989001dd73e950d302c9f54fde98b5","0x025beebf7b8d4e7a23e4c1ae7440d945a3353495":"0x928de5b59bffd4a293cc1034ef81c59e935c81aee5330de39ee73a904aa70ce5","0xf8dab13f97c8cd6e1c5c205f6e46358823b1bf57":"0xcf28f86b0f76f368a6bd521a9b63d8545de12445fde21452351197c39b6c82af","0x73d7d7a597246521e7b440752b200b12e30eca2a":"0x75bb43a43090f2738cdc1e7daf9197a80d77188c8bb79228e01edc604226898f","0xdb1080ed09ff4405d9abc550619f40341073d12c":"0x2884cd4edd5cc78201455e837483964ed0a0201b590f5afe6a2d7e766741018f","0x328fc2a3a266a4a9d2ae0f62ac710cf9522f6ab3":"0x7f3b9fa8b2709d343d5b2b26a0ff7b4d0497e93b45f24034310ce00d44ea121a","0x25fe79d56c0ec39183b37fbbbbee03fcb3e29c18":"0x965e860f743c8d6609ae47134436c29a9b5170ee4576fec07127b33da0708c09","0x8da2ff45e5ab776e775f7df122856503ac37f5e2":"0xe5d42fef1125927fa0d5aa93b3ac53d63b873da09f40426c4bd70f2925c63001","0x5ce0b726c8f3abc4b46f3f04a157830babb60248":"0x6260f8c1a7a294b59a2e12b24c2803d728be403c88dbba9e49ce1dad095483b6","0xa12c2b1cb5ac67f5fa336b1645c6e88cdeca13a0":"0x8cff30147c81424a5a17e580c5762efce8eb7d4a40b5c2314e0de24a394343e2"}}`)

func ProvideGanacheAccounts() model.GanacheAccounts {
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
		return model.GanacheAccounts{}
	}
	itemsMap := f.(map[string]interface{})
	model := model.GanacheAccounts{}
	for k, v := range itemsMap {
		for _, v1 := range v.(map[string]interface{}) {
			if k == "addresses" {
				model.Addresses = append(model.Addresses, v1.(string))
			} else {
				model.PrivateKeys = append(model.PrivateKeys, v1.(string)[2:])
			}
		}
	}
	if err != nil {
		panic(err)
	} else {
		println("Providing accounts successfully")
	}
	return model
}
