module github.com/irisnet/irishub-sdk-go

go 1.15

require (
	github.com/OneOfOne/xxhash v1.2.5 // indirect
	github.com/bluele/gcache v0.0.0-20190518031135-bc40bd653833
	github.com/btcsuite/btcd v0.22.0-beta
	github.com/btcsuite/btcutil v1.0.3-0.20201208143702-a53e38424cce
	github.com/cosmos/cosmos-sdk v0.44.2
	github.com/cosmos/go-bip39 v1.0.0
	github.com/gin-gonic/gin v1.7.7 // indirect
	github.com/gogo/protobuf v1.3.3
	github.com/golang/protobuf v1.5.2
	github.com/hashicorp/golang-lru v0.5.5-0.20210104140557-80c98217689d // indirect
	github.com/magiconair/properties v1.8.5
	github.com/pkg/errors v0.9.1
	github.com/regen-network/cosmos-proto v0.3.1
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/cast v1.4.0 // indirect
	github.com/stretchr/testify v1.7.0
	github.com/syndtr/goleveldb v1.0.1-0.20210819022825-2ae1ddf74ef7 // indirect
	github.com/tendermint/crypto v0.0.0-20191022145703-50d29ede1e15
	github.com/tendermint/go-amino v0.16.0
	github.com/tendermint/tendermint v0.34.13
	github.com/tendermint/tm-db v0.6.4
	golang.org/x/crypto v0.0.0-20210513164829-c07d793c2f9a
	google.golang.org/genproto v0.0.0-20210805201207-89edb61ffb67
	google.golang.org/grpc v1.40.0
	google.golang.org/protobuf v1.27.1
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/yaml.v2 v2.4.0
)

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
