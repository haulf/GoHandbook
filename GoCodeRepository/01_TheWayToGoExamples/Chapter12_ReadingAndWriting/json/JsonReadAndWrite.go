// author:     haulf
// date:       2017.08.22
// brief:      Json data read/write test.

package main

import (
	"encoding/json"
	"fmt"
)

type DrugDetail struct {
	ID   string
	Name string
}

type DrugMaster struct {
	DrugData []DrugDetail
}

func main() {
	var drugMasterJson DrugMaster
	drugMasterJson.DrugData = append(drugMasterJson.DrugData, DrugDetail{ID: "0001", Name: "云南白药"})
	drugMasterJson.DrugData = append(drugMasterJson.DrugData, DrugDetail{ID: "0002", Name: "999感冒灵"})
	drugMasterJson.DrugData = append(drugMasterJson.DrugData, DrugDetail{ID: "0003", Name: "创口贴"}, DrugDetail{ID: "0004", Name: "西洋参"})

	byteArray, err := json.Marshal(drugMasterJson)
	if err != nil {
		fmt.Println("Json marshal error", err)
	}

	var jsonData string
	jsonData = string(byteArray)
	fmt.Println(">>>>>生成Json格式数据:")
	//fmt.Println(string(byteArray))
	fmt.Println(jsonData)
	fmt.Println("")

	fmt.Println(">>>>>解析Json到struct:")
	jsonUnmarshal(jsonData)
	fmt.Println("")
	fmt.Println("")
}

func jsonUnmarshal(jsonData string) {
	var drugMasterData DrugMaster
	json.Unmarshal([]byte(jsonData), &drugMasterData)
	fmt.Println(drugMasterData)
	fmt.Println("")
}
