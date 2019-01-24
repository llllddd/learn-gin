package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type AdvisorChaincode struct {
}

func (advisor *AdvisorChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

func (advisor *AdvisorChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	//得到想要运行的函数和参数
	fun, args := stub.GetFunctionAndParameters()
	if fun == "addAdvisor" {
		//添加定制师信息
		return advisor.addAdvisor(stub, args)
	} else if fun == "queryAdvisorInfoByID" {
		//由ID查询定制师信息
		return advisor.queryAdvisorInfoByID(stub, args)
	} else if fun == "updateAdvisor" {
		//更新定制师信息
		return advisor.updateAdvisor(stub, args)
	}
	//else if fun == "queryAdvisorInfoByName" {
	//由定制师姓名查询定制师信息
	//	return advisor.queryAdvisorInfoByName(stub, args)
	//}
	return shim.Error("没有对应的函数")
}

func main() {
	err := shim.Start(new(AdvisorChaincode))
	if err != nil {
		fmt.Printf("启动AdvisorChaincode时发生错误 %s", err)
	}

}
