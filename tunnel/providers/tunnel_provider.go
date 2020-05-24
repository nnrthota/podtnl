package providers

import (
	"github.com/narendranathreddythota/podtnl/tunnel/types"
)

type Tunnel struct {
	Name          string             // Name of the tunnel
	Proto         string             // TCP, HTTP,
	LocalIP       string             // Local IP Address where our local service is running
	LocalPort     string             // Local Port  where our local service is running
	Auth          string             // Tunnel Authentication for secure access
	RemoteAddress string             // Result Online Address
	Inspect       bool               // Inspect transaction data tunnel that will be logged in binary file
	Status        types.TunnelStatus // Status of the tunnel
	IsCreated     types.IsCreated    // Is Tunnel Created
}

type ITunnelProvider interface {
	Start() error
	End() error
	OpenManyTunnels(t []*Tunnel) error
	CreateTunnel(t *Tunnel) error
	CloseTunnel(t *Tunnel) error
	CloseManyTunnels(t []*Tunnel) error
}
