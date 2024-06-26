package call

type Chain string

type ChainConfig struct {
	MultiCallAddress string
	Url              string
}

const (
	Arbitrum  Chain = "arbitrum"
	Aurora    Chain = "aurora"
	Avalanche Chain = "avalanche"
	Bsc       Chain = "bsc"
	Ethereum  Chain = "ethereum"
	Fantom    Chain = "fantom"
	Moonbeam  Chain = "moonbeam"
	Moonriver Chain = "moonriver"
	Celo      Chain = "celo"
)

var DefaultChainConfigs = map[Chain]ChainConfig{
	Arbitrum: {
		MultiCallAddress: "0x7a7443f8c577d537f1d8cd4a629d40a3148dd7ee",
		Url:              "https://arb1.arbitrum.io/rpc",
	},
	Aurora: {
		MultiCallAddress: "0x88b373B83166E72FD55648Ce114712633f1782E2",
		Url:              "https://mainnet.aurora.dev",
	},
	Avalanche: {
		MultiCallAddress: "0xa00FB557AA68d2e98A830642DBbFA534E8512E5f",
		Url:              "https://api.avax.network/ext/bc/C/rpc",
	},
	Bsc: {
		MultiCallAddress: "0xAD38F5025CBe01c8637BF683B3B68096A1c882CA",
		Url:              "https://bsc-dataseed1.ninicoin.io",
	},
	Ethereum: {
		MultiCallAddress: "0x5BA1e12693Dc8F9c48aAD8770482f4739bEeD696",
		Url:              "https://mainnet.infura.io/v3/9aa3d95b3bc440fa88ea12eaa4456161",
	},
	Fantom: {
		MultiCallAddress: "0x7F4e475462A0fA0F1e2C69d50866D54505F99D72",
		Url:              "https://rpcapi.fantom.network",
	},
	Moonbeam: {
		MultiCallAddress: "0x6477204E12A7236b9619385ea453F370aD897bb2",
		Url:              "https://moonbeam.public.blastapi.io",
	},
	Moonriver: {
		MultiCallAddress: "0xEae947bF407A4a4f1c5a6A73312734A2863e3855",
		Url:              "https://moonriver.public.blastapi.io",
	},
	Celo: {
		MultiCallAddress: "0x7F4e475462A0fA0F1e2C69d50866D54505F99D72",
		Url:              "https://forno.celo.org",
	},
}
