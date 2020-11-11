package types

type Config struct {
	Project      string   `json:"project"`
	APIs         []string `json:"apis"`
	ProtoModules []string `json:"protomodules"`
}
