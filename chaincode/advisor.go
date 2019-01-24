package main

type Advisor struct {
	AdvisorID    string   `json:"advisorId"`    //人员ID
	Gender       int      `json:"gender"`       //性别0女1男
	ProviderInfo string   `json:"providerInfo"` //供应商信息TODO：详细的供应商信息
	Status       int      `json:"status"`       //认证状态
	CommentCount int      `json:"commentCount"` //点评数
	CommentScore float64  `json:"commentScore"` //点评分
	Comment      string   `json:"comment"`      //评论内容 TODO：详细的评论内容
	OrderCount   int      `json:"orderCount"`   //接单量
	GroupCount   int      `json:"groupCount"`   //成团量
	ServeCount   int      `json:"serveCount"`   //历史服务次数
	DistrictList []string `json"districtList"`  //擅长目的地 TODO:详细的目的地信息

	History []HistoryItem //当前定制师的历史记录
}

type HistoryItem struct {
	TxId    string
	Advisor Advisor
}

//TODO:详细的供应商信息
//type ProviderInfoDto struct {
//}

//TODO:擅长目的地
//type AdvisorDistrictRelDto struct {
//}
