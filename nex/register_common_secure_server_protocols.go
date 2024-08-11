package nex

import (
	secureconnection "github.com/PretendoNetwork/nex-protocols-common-go/v2/secure-connection"
	secure "github.com/PretendoNetwork/nex-protocols-go/v2/secure-connection"

	common_ranking "github.com/PretendoNetwork/nex-protocols-common-go/v2/ranking"
	ranking "github.com/PretendoNetwork/nex-protocols-go/v2/ranking"
	"github.com/PretendoNetwork/pikmin-3/database"
	"github.com/PretendoNetwork/pikmin-3/globals"
)

func registerCommonSecureServerProtocols() {
	secureProtocol := secure.NewProtocol()
	globals.SecureEndpoint.RegisterServiceProtocol(secureProtocol)
	secureconnection.NewCommonProtocol(secureProtocol)

	rankingProtocol := ranking.NewProtocol()
	globals.SecureEndpoint.RegisterServiceProtocol(rankingProtocol)
	commonRankingProtocol := common_ranking.NewCommonProtocol(rankingProtocol)

	commonRankingProtocol.GetRankingsAndCountByCategoryAndRankingOrderParam = database.GetRankingsAndCountByCategoryAndRankingOrderParam
	commonRankingProtocol.InsertRankingByPIDAndRankingScoreData = database.InsertRankingByPIDAndRankingScoreData
}
