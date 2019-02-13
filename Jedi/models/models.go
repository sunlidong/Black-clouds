package models

import (
	"github.com/astaxie/beego"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/event"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

// FabricSetup implementation
type FabricSetup struct {
	ConfigFile string
	OrgID string
	OrdererID string
	ChannelID string
	ChainCodeID string
	ChainCodeVersion string
	IsInitializedSDK bool
	IsResMgmt bool
	IsChannel bool
	IsJoinChannel bool
	IsChainCodeInstall bool
	IsChainCodeInit bool
	IsChannelCli bool
	ChannelConfig string
	ChaincodeGoPath string
	ChaincodePath string
	OrgAdmin string
	OrgName string
	UserName string
	eventID string
	client *channel.Client
	admin *resmgmt.Client
	sdk *fabsdk.FabricSDK
	event *event.Client
}

type Application struct {
	beego.Controller
	FabricSetup *FabricSetup
}

//
type AreaInfo struct {
	AreaID string `json:"area_id"`
	AreaAddress string `json:"area_address"`
	BasicNetWork string `json:"basic_net_work"`
	CPoliceName string `json:"c_police_name"`
	CPoliceNum string `json:"c_police_num"`
}

//  ofgj
type HouseInfo struct {
	HouseID    string `json:"house_id"`
	HouseOwner string `json:"house_owner"`
	RegDate    string `json:"reg_date"`
	HouseArea  string `json:"house_area"`
	HouseUsed  string `json:"house_used"`
	IsMortgage string `json:"is_mortgage"`
}

//  	oagency
type OrderInfo struct {
	DocHash   string `json:"doc_hash"`
	OrderID   string `json:"order_id"`
	RenterID  string `json:"renter_id"`
	RentMoney string `json:"rent_money"`
	BeginDate string `json:"begin_date"`
	EndDate   string `json:"end_date"`
	Note      string `json:"note"`
}

//   ozxzx
type RenterInfo struct {
	RenterID   string `json:"renter_id"`
	RenterDesc string `json:"renter_desc"`
}
