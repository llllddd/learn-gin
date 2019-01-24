package main

import (
	"encoding/json"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

const Advisor_Prefix = "ADV"

//保存定制师信息
func PutAdvisor(stub shim.ChaincodeStubInterface, adv Advisor) ([]byte, bool) {
	//先得到定制师信息序列化
	info, err := json.Marshal(adv)
	if err != nil {
		return nil, false
	}
	//将序列化后的信息存储到couchdb中考虑需要前缀0x00的情形
	key := Advisor_Prefix + adv.AdvisorID
	err = stub.PutState(key, info)
	if err != nil {
		return nil, false
	}

	return info, true
}

//根据定制师Id查询定制师信息状态
func GetAdvisorInfo(stub shim.ChaincodeStubInterface, advisorId string) (Advisor, bool) {
	var adv Advisor
	//由定制师ID查询对应的数据库中的序列化信息
	info, err := stub.GetState(advisorId)
	if err != nil {
		return adv, false
	}
	if info == nil {
		return adv, false
	}
	//反序列化
	err = json.Unmarshal(info, &adv)
	if err != nil {
		return adv, false
	}

	return adv, true

}

//添加定制师信息，定制师的ID为Key， Advisor为value
func (advisor *AdvisorChaincode) addAdvisor(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	//获取待存储的定制师的信息
	if len(args) != 2 {
		return shim.Error("参数个数不正确")
	}

	var adv Advisor
	//key := Advisor_Prefix + args[0]
	//将信息反序列化
	err := json.Unmarshal([]byte(args[0]), &adv)

	if err != nil {
		return shim.Error("反序列化信息时发生错误")
	}
	//检验定制师ID是否已存在
	_, exist := GetAdvisorInfo(stub, adv.AdvisorID)

	if exist {
		return shim.Error("定制师信息已存在")
	}
	//保存定制师信息到数据库中
	_, ok := PutAdvisor(stub, adv)
	if !ok {
		return shim.Error("保存信息时发生错误")
	}
	err = stub.SetEvent(args[1], []byte{})
	if err != nil {
		return shim.Error("设置链码事件失败")
	}

	return shim.Success([]byte("定制师信息添加成功"))
}

//由定制师的ID查询定制师信息可溯源args AdvisorID
func (advisor *AdvisorChaincode) queryAdvisorInfoByID(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	//判断参数
	if len(args) != 1 {
		return shim.Error("参数个数不正确")
	}
	//由给定的ID查询定制师状态
	key := Advisor_Prefix + args[0]

	info, err := stub.GetState(key)
	if err != nil {
		return shim.Error("由定制师ID查询数据失败")
	}
	if info == nil {
		return shim.Error("由定制师ID未查询到相关信息")
	}
	//反序列化查询到的信息
	var adv Advisor
	err = json.Unmarshal(info, &adv)

	if err != nil {
		return shim.Error("反序列化信息时发生错误")
	}

	//获取历史变更数据
	iterator, err := stub.GetHistoryForKey(adv.AdvisorID)
	for err != nil {
		return shim.Error("由定制师ID查询历史变更数据失败")
	}
	defer iterator.Close()

	//得到定制师的历史数据
	var history []HistoryItem
	var historyAdvisor Advisor

	for iterator.HasNext() {
		historyData, err := iterator.Next()
		if err != nil {
			return shim.Error("获取adv历史数据失败")
		}

		var historyItem HistoryItem
		historyItem.TxId = historyData.TxId
		//反序列化
		if historyData.Value == nil {
			var empty Advisor
			historyItem.Advisor = empty
		} else {
			json.Unmarshal(historyData.Value, &historyAdvisor)
			historyItem.Advisor = historyAdvisor
		}

		history = append(history, historyItem)
	}

	adv.History = history

	//返回序列化后的结果
	result, err := json.Marshal(adv)
	if err != nil {
		return shim.Error("序列化定制师信息时错误")
	}

	return shim.Success(result)
}

//更新定制师信息参数需要定制师ID和定制时信息
func (advisor *AdvisorChaincode) updateAdvisor(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	//判断参数
	if len(args) != 2 {
		return shim.Error("参数个数不正确")
	}

	var info Advisor
	//将定制师信息反序列化
	err := json.Unmarshal([]byte(args[0]), &info)
	if err != nil {
		return shim.Error("反序列化失败")
	}
	//由定制师ID查询到数据库中对应的信息
	key := Advisor_Prefix + info.AdvisorID
	result, ok := GetAdvisorInfo(stub, key)
	if !ok {
		return shim.Error("由定制师ID查询信息时错误")
	}

	//更新相关信息
	result.ProviderInfo = info.ProviderInfo
	result.Status = info.Status
	result.CommentCount = info.CommentCount
	result.CommentScore = info.CommentScore
	result.Comment = info.Comment
	result.OrderCount = info.OrderCount
	result.GroupCount = info.GroupCount
	result.ServeCount = info.ServeCount
	result.DistrictList = info.DistrictList

	_, ok = PutAdvisor(stub, result)
	if !ok {
		return shim.Error("保存更新后的定制师信息错误")
	}

	err = stub.SetEvent(args[1], []byte{})
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success([]byte("信息更新成功"))
}
