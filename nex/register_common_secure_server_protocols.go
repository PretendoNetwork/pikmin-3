package nex

import (
	secureconnection "github.com/PretendoNetwork/nex-protocols-common-go/secure-connection"
	"github.com/PretendoNetwork/pikmin-3/globals"
)

func registerCommonSecureServerProtocols() {
	secureconnection.NewCommonSecureConnectionProtocol(globals.SecureServer)
}
