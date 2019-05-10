package app


import (
	"fmt"
	"github.com/spf13/cobra"
	"lll.github.com/websocket-vpn/pkg/websocket/server"
	"os"
	"lll.github.com/websocket-vpn/cmd/server/app/options"
)

func NewVpnServerCommand(stopCh <-chan struct{}) *cobra.Command {
	s := options.NewVpnServerOptions()

	cmd := &cobra.Command{
		Use: "VpnServer",
		Long: "Vpn server through websocket",
		RunE: func (cmd *cobra.Command, args []string) error {
			fmt.Print("start websocket vpn server!")

			return server.Run()
		},
	}

	cmd.SetArgs(os.Args)
	fs := cmd.Flags()
	fs.StringP("listen-address", "l", "0.0.0.0:5555", "The vpn server listen address")

	fs.Parse(os.Args)
	s.ListenAddress, _ = fs.GetString("listen-address")
	return cmd
}