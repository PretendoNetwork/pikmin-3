package database

import (
	"time"

	"github.com/PretendoNetwork/nex-go/v2/types"
	ranking_types "github.com/PretendoNetwork/nex-protocols-go/v2/ranking/types"
)

func InsertRankingByPIDAndRankingScoreData(pid *types.PID, rankingScoreData *ranking_types.RankingScoreData, uniqueID *types.PrimitiveU64) error {
	now := time.Now().UnixNano()

	_, err := Postgres.Exec(`
		INSERT INTO rankings (
			owner_pid,
			category,
			score,
			order_by,
			update_mode,
			groups,
			param,
			common_data,
			created_at
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`,
		pid.Value(),
		rankingScoreData.Category.Value,
		rankingScoreData.Score.Value,
		rankingScoreData.OrderBy.Value,
		rankingScoreData.UpdateMode.Value,
		rankingScoreData.Groups.Value,
		rankingScoreData.Param.Value,
		make([]byte, 0),
		now,
	)

	return err
}
