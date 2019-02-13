# 删除 原有文件 
rm -rf ./channel-artifacts/*
echo "-----------------------------del text is ok"
# 生成创始块文件
echo "---------------- Create ThreeOrgsOrdererGenesis.block file BEGIN -------------------"
configtxgen -profile ThreeOrgsOrdererGenesis -outputBlock ./channel-artifacts/genesis.block
echo "---------------- Create ThreeOrgsOrdererGenesis.block file END -------------------"

# 生成 创始通道
echo "---------------- Create sjOrgchannel.tx file BEGIN -------------------"
configtxgen -profile sjOrgchannel -outputCreateChannelTx ./channel-artifacts/sjchannel.tx -channelID sjchannel
echo "---------------- Create sjOrgchannel.tx file END -------------------"

 
