package dns

type DNSConfig struct {
	ServerAA string `json:"server_aa" yaml:"server_aa" mapstructure:"server_aa"` // server for Authority name server
	ServerAB string `json:"server_ab" yaml:"server_ab" mapstructure:"server_ab"` // server for Authority name server
	ServerCC string `json:"server_cc" yaml:"server_cc" mapstructure:"server_cc"` //server for Recursor name server
	IPAA     string `json:"ipaa" yaml:"ipaa" mapstructure:"ipaa"`
	IPAB     string `json:"ipab" yaml:"ipab" mapstructure:"ipab"`
	IPCC     string `json:"ipcc" yaml:"ipcc" mapstructure:"ipcc"`
	UserZone string `json:"user_zone" yaml:"user_zone" mapstructure:"user_zone"`
	PODZone  string `json:"pod_zone" yaml:"pod_zone" mapstructure:"pod_zone"`
	TopZone  string `json:"top_zone" yaml:"top_zone" mapstructure:"top_zone"`
}

const (
	DefaultUserZone = "user.fuxi"
	DefaultPODZone  = "pod.fuxi"
	DefaultTopZone  = "fuxi"
)

func DefaultConfig() DNSConfig {
	return DNSConfig{
		ServerAA: "http://127.0.0.1:8888",
		ServerAB: "http://127.0.0.1:8888",
		ServerCC: "http://127.0.0.1:8889",
		IPAA:     "127.0.0.1",
		IPAB:     "127.0.0.1",
		IPCC:     "127.0.0.1",
		UserZone: DefaultUserZone,
		PODZone:  DefaultPODZone,
		TopZone:  DefaultTopZone,
	}
}
