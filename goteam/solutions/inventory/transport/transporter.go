package transport

import (
	"net"

	"github.com/gopheracademy/code/goteam/solutions/inventory"
)

type InventoryTransporter interface {
	inventory.Service
	Serve(net.Listener) error
}
