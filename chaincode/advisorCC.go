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
func (advisor *AdvisorChaincode) addAvisor(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	//获取待存储的定制师的信息
	if len(args != 2) {
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
		return shim.Error("保存信息时发生错误"）
	}
	err := stub.SetEvent(args[1], []byte{})
	if err != nil{
		return shim.Error("设置链码事件失败")
	}
	
	return shim.Success([]byte("定制师信息添加成功"))
}
