package database

import (
	"database/sql"

	"github.com/PretendoNetwork/nex-go/v2/types"
	ranking_types "github.com/PretendoNetwork/nex-protocols-go/v2/ranking/types"
	"github.com/PretendoNetwork/pikmin-3/globals"
)

func GetRankingsAndCountByCategoryAndRankingOrderParam(category *types.PrimitiveU32, rankingOrderParam *ranking_types.RankingOrderParam) (*types.List[*ranking_types.RankingRankData], uint32, error) {
	rankings := types.NewList[*ranking_types.RankingRankData]()

	rows, err := Postgres.Query(`
		SELECT
		owner_pid,
		score,
		groups,
		param,
		common_data
		FROM rankings WHERE category=$1 ORDER BY score DESC LIMIT $2 OFFSET $3`,
		category,
		rankingOrderParam.Length,
		rankingOrderParam.Offset,
	)
	if err != nil {
		return nil, 0, err
	}

	row := 1
	for rows.Next() {
		rankingRankData := ranking_types.NewRankingRankData()
		rankingRankData.UniqueID.Value = 0
		rankingRankData.Order.Value = uint32(row)
		rankingRankData.Category = category

		err := rows.Scan(
			&rankingRankData.PrincipalID,
			&rankingRankData.Score,
			&rankingRankData.Groups,
			&rankingRankData.Param,
			&rankingRankData.CommonData,
		)

		if err != nil && err != sql.ErrNoRows {
			globals.Logger.Critical(err.Error())
		}

		if err == nil {
			rankings.Append(rankingRankData)

			row += 1
		}
	}

	return rankings, uint32(rankings.Length()), nil
}
