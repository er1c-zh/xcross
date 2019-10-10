package lib

type Action uint8

const (
	ActionPing Action = 0
	ActionConn        = 1
	ActionBind        = 2
	ActionUDP         = 3
)

type AddressType uint8

const (
	AddrTypeIPv4   AddressType = 1
	AddrTypeIPv6               = 2
	AddrTypeDomain             = 3
)

type Port uint16

type PkgShakeHands struct {
	Action   Action
	AddrType AddressType
	Port     Port
}

func Parse(src [3]byte) *PkgShakeHands {
	return nil
}
