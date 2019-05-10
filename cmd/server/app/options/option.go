package options

type VpnServerOptions struct {
	ListenAddress string
}

func NewVpnServerOptions() *VpnServerOptions {
	return &VpnServerOptions{
		ListenAddress: "0.0.0.0:5555",
	}
}