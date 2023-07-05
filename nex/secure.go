package nex

import (
	"fmt"
	"os"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/pikmin-3/globals"
)

func StartSecureServer() {
	globals.SecureServer = nex.NewServer()
	globals.SecureServer.SetPRUDPVersion(1)
	globals.SecureServer.SetPRUDPProtocolMinorVersion(3)
	globals.SecureServer.SetDefaultNEXVersion(&nex.NEXVersion{
		Major: 3,
		Minor: 3,
		Patch: 0,
	})
	globals.SecureServer.SetKerberosPassword(globals.KerberosPassword)
	globals.SecureServer.SetAccessKey("f6accfc1")

	globals.SecureServer.On("Data", func(packet *nex.PacketV1) {
		request := packet.RMCRequest()

		fmt.Println("==PIKMIN 3 - Secure==")
		fmt.Printf("Protocol ID: %#v\n", request.ProtocolID())
		fmt.Printf("Method ID: %#v\n", request.MethodID())
		fmt.Println("===============")
	})

	registerCommonSecureServerProtocols()
	registerSecureServerNEXProtocols()

	globals.SecureServer.Listen(fmt.Sprintf(":%s", os.Getenv("PN_PIKMIN3_SECURE_SERVER_PORT")))
}
