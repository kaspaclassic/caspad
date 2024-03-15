package rpccontext

import (
	"github.com/kaspaclassic/caspad/app/protocol"
	"github.com/kaspaclassic/caspad/domain"
	"github.com/kaspaclassic/caspad/domain/utxoindex"
	"github.com/kaspaclassic/caspad/infrastructure/config"
	"github.com/kaspaclassic/caspad/infrastructure/network/addressmanager"
	"github.com/kaspaclassic/caspad/infrastructure/network/connmanager"
	"github.com/kaspaclassic/caspad/infrastructure/network/netadapter"
)

// Context represents the RPC context
type Context struct {
	Config            *config.Config
	NetAdapter        *netadapter.NetAdapter
	Domain            domain.Domain
	ProtocolManager   *protocol.Manager
	ConnectionManager *connmanager.ConnectionManager
	AddressManager    *addressmanager.AddressManager
	UTXOIndex         *utxoindex.UTXOIndex
	ShutDownChan      chan<- struct{}

	NotificationManager *NotificationManager
}

// NewContext creates a new RPC context
func NewContext(cfg *config.Config,
	domain domain.Domain,
	netAdapter *netadapter.NetAdapter,
	protocolManager *protocol.Manager,
	connectionManager *connmanager.ConnectionManager,
	addressManager *addressmanager.AddressManager,
	utxoIndex *utxoindex.UTXOIndex,
	shutDownChan chan<- struct{}) *Context {

	context := &Context{
		Config:            cfg,
		NetAdapter:        netAdapter,
		Domain:            domain,
		ProtocolManager:   protocolManager,
		ConnectionManager: connectionManager,
		AddressManager:    addressManager,
		UTXOIndex:         utxoIndex,
		ShutDownChan:      shutDownChan,
	}
	context.NotificationManager = NewNotificationManager(cfg.ActiveNetParams)

	return context
}
