package nex

import (
	"fmt"
	"os"

	nex "github.com/PretendoNetwork/nex-go"
	"github.com/PretendoNetwork/pikmin-3/globals"
)

var serverBuildString string

func StartAuthenticationServer() {
	globals.AuthenticationServer = nex.NewServer()
	globals.AuthenticationServer.SetPRUDPVersion(1)
	globals.AuthenticationServer.SetPRUDPProtocolMinorVersion(3)
	globals.AuthenticationServer.SetDefaultNEXVersion(&nex.NEXVersion{
		Major: 3,
		Minor: 3,
		Patch: 0,
	})
	globals.AuthenticationServer.SetKerberosPassword(globals.KerberosPassword)
	globals.AuthenticationServer.SetAccessKey("f6accfc1")

	globals.AuthenticationServer.On("Data", func(packet *nex.PacketV1) {
		request := packet.RMCRequest()

		fmt.Println("==PIKMIN 3 - Auth==")
		fmt.Printf("Protocol ID: %#v\n", request.ProtocolID())
		fmt.Printf("Method ID: %#v\n", request.MethodID())
		fmt.Println("===============")
	})

	registerCommonAuthenticationServerProtocols()

	globals.AuthenticationServer.Listen(fmt.Sprintf(":%s", os.Getenv("PN_PIKMIN3_AUTHENTICATION_SERVER_PORT")))
}
