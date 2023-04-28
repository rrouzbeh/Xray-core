package conf_test

import (
	"testing"

	"github.com/rrouzbeh/xray-core/common/net"
	"github.com/rrouzbeh/xray-core/common/protocol"
	"github.com/rrouzbeh/xray-core/common/serial"
	. "github.com/rrouzbeh/xray-core/infra/conf"
	"github.com/rrouzbeh/xray-core/proxy/shadowsocks"
)

func TestShadowsocksServerConfigParsing(t *testing.T) {
	creator := func() Buildable {
		return new(ShadowsocksServerConfig)
	}

	runMultiTestCase(t, []TestCase{
		{
			Input: `{
				"method": "aes-256-GCM",
				"password": "xray-password"
			}`,
			Parser: loadJSON(creator),
			Output: &shadowsocks.ServerConfig{
				Users: []*protocol.User{{
					Account: serial.ToTypedMessage(&shadowsocks.Account{
						CipherType: shadowsocks.CipherType_AES_256_GCM,
						Password:   "xray-password",
					}),
				}},
				Network: []net.Network{net.Network_TCP},
			},
		},
	})
}
