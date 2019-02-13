package models

import (
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/errors/retry"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	packager "github.com/hyperledger/fabric-sdk-go/pkg/fab/ccpackager/gopackager"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	"github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric/common/cauthdsl"
	"github.com/pkg/errors"
)

func (setup *FabricSetup) InitializeSDK() error {
	// 判断SDK是否初始化
	if setup.IsInitializedSDK == true {
		return errors.New("SDK已被初始化")
	}
	// 使⽤配置⽂件初始化SDK
	sdk, err := fabsdk.New(config.FromFile(setup.ConfigFile))
	if err != nil {
		return errors.WithMessage(err, "初始化SDK失败")
	}
	setup.sdk = sdk
	fmt.Println("SDK 初始化成功!")
	setup.IsInitializedSDK = true
	return nil
}
func (setup *FabricSetup) CreateResMgmtClient() error {
	if setup.IsResMgmt == true {
		return errors.New("资源管理客户的已经创建！")
	}
	// 创建资源管理客户端上下⽂
	resourceManagerClientContext :=
		setup.sdk.Context(fabsdk.WithUser(setup.OrgAdmin),
			fabsdk.WithOrg(setup.OrgName))
	// 创建资源管理客户端
	resMgmtClient, err := resmgmt.New(resourceManagerClientContext)
	if err != nil {
		return errors.WithMessage(err, "创建资源管理客户端失败!")
	}
	setup.admin = resMgmtClient
	fmt.Println("创建资源管理客户端成功！")
	setup.IsResMgmt = true
	return nil
}
func (setup *FabricSetup) CreateChannel() error {
	if setup.IsChannel == true {
		return errors.New("Channel 已经创建！")
	}
	// 创建Channel
	req := resmgmt.SaveChannelRequest{
		ChannelID:         setup.ChannelID,
		ChannelConfigPath: setup.ChannelConfig,
	}
	txID, err := setup.admin.SaveChannel(req)
	if err != nil || txID.TransactionID == "" {
		return errors.WithMessage(err, "Channel 创建失败！")
	}
	fmt.Println("Channel 创建成功！")
	setup.IsChannel = true
	return nil
}
func (setup *FabricSetup) JoinChannel() error {
	if setup.IsJoinChannel == true {
		return errors.New("资源管理客户的已经创建！")
	}
	// 加⼊ Channel
	err := setup.admin.JoinChannel(
		setup.ChannelID,
		resmgmt.WithRetry(retry.DefaultResMgmtOpts),
		resmgmt.WithOrdererEndpoint(setup.OrdererID),
	)
	if err != nil {
		return errors.WithMessage(err, "加⼊Channel 失败!")
	}
	fmt.Println("Channel 加⼊成功!")
	setup.IsJoinChannel = true
	return nil
}
func (setup *FabricSetup) ChainCodeInstall() error {
	if setup.IsChainCodeInstall == true {
		return errors.New("链码已经安装！")
	}
	// 链码打包
	ccPkg, err := packager.NewCCPackage(setup.ChaincodePath,
		setup.ChaincodeGoPath)
	if err != nil {
		return errors.WithMessage(err, "链码打包失败！")
	}
	fmt.Println("链码打包成功!")
	// 安装链码
	installCCReq := resmgmt.InstallCCRequest{
		Name:    setup.ChainCodeID,
		Path:    setup.ChaincodePath,
		Version: setup.ChainCodeVersion,
		Package: ccPkg,
	}
	_, err = setup.admin.InstallCC(installCCReq,
		resmgmt.WithRetry(retry.DefaultResMgmtOpts))
	if err != nil {
		return errors.WithMessage(err, "链码安装失败！")
	}
	fmt.Println("链码安装成功！")
	setup.IsChainCodeInstall = true
	return nil
}
func (setup *FabricSetup) ChainCodeInit() error {
	if setup.IsChainCodeInit == true {
		return errors.New("链码已被初始化！")
	}
	// 设置背书策略，该参数依赖 configtx.yaml ⽂件中 Organizations -> ID
	ccPolicy := cauthdsl.SignedByAnyMember([]string{"osj.sanji.com"})
	resp, err := setup.admin.InstantiateCC(
		setup.ChannelID,
		resmgmt.InstantiateCCRequest{
			Name:    setup.ChainCodeID,
			Path:    setup.ChaincodeGoPath,
			Version: setup.ChainCodeVersion,
			Args:    [][]byte{[]byte("init")},
			Policy:  ccPolicy})
	if err != nil || resp.TransactionID == "" {
		return errors.WithMessage(err, "链码初始化失败！")
	}
	fmt.Println("链码初始化成功！")
	setup.IsChainCodeInit = true
	return nil
}
func (setup *FabricSetup) CreateChannelCli() error {
	if setup.IsChannelCli == true {
		return errors.New("通道Cli已经创建")
	}
	// 创建Channle cli
	clientContext := setup.sdk.ChannelContext(
		setup.ChannelID,
		fabsdk.WithUser(setup.UserName))
	channelCli, err := channel.New(clientContext)
	if err != nil {
		return errors.WithMessage(err, "Channel Cli 创建失败！")
	}
	setup.client = channelCli
	fmt.Println("Channel Cli 创建成功")
	setup.IsChannelCli = true
	return nil
}
func (setup *FabricSetup) CloseSDK() {
	setup.sdk.Close()
}
