package models
import (
	"fmt"
	"os"
)
var App Application
func Finit1() {
	fmt.Println("-----kaishi ")
	fSetup := FabricSetup{
		// Network parameters
		OrdererID: "orderer.sanji.com",
		// Channel parameters
		ChannelID: "sjchannel",
		ChannelConfig: os.Getenv("GOPATH") +
			"/src/project_wm/Jedi/conf/channel-artifacts/sjchannel.tx",
		// Chaincode parameters
		ChainCodeID: "mycc",
		ChaincodeGoPath: os.Getenv("GOPATH"),
		ChaincodePath: "project_wm/Jedi/chaincode",
		ChainCodeVersion:"1.0",
		OrgAdmin: "Admin",
		OrgName: "osj",
		ConfigFile: "conf/config.yaml",
		// User parameters
		UserName: "User1",
		eventID:"eventInvoke",
	} // 初始化SDK
	err := fSetup.InitializeSDK()
	if err != nil {
		fmt.Printf("SDK 初始化err: %v\n", err)
		return
	}
	// Close SDK
	// defer fSetup.CloseSDK()
	// 创建资源管理
	err = fSetup.CreateResMgmtClient()
	if err != nil {
		fmt.Printf("ResMgmtCli 创建err: %v\n", err)
		return
	}
	// 创建 Channel
	err = fSetup.CreateChannel()
	if err != nil {
		fmt.Printf("Create Channel err : %v\n", err)
		return
	}
	// 加⼊ Channel
	err = fSetup.JoinChannel()
	if err != nil {
		fmt.Printf("Join Channel err : %v\n", err)
		return
	}
	// 安装ChainCode
	err = fSetup.ChainCodeInstall()
	if err != nil {
		fmt.Printf("Install ChainCode err : %v\n", err)
		return
	}
	// 初始化ChainCode
	err = fSetup.ChainCodeInit()
	if err != nil {
		fmt.Printf("Init ChainCode err : %v\n", err)
		return
	}
	// 创建Channel Cli
	err = fSetup.CreateChannelCli()
	if err != nil {
		fmt.Printf("Create Channel Cli err : %v\n", err)
		return
	}
	App = Application{
		FabricSetup: &fSetup,
	}
}