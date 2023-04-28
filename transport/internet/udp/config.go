package udp

import (
	"github.com/rrouzbeh/xray-core/common"
	"github.com/rrouzbeh/xray-core/transport/internet"
)

func init() {
	common.Must(internet.RegisterProtocolConfigCreator(protocolName, func() interface{} {
		return new(Config)
	}))
}
