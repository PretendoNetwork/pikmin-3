package database

import (
	"time"

	"github.com/PretendoNetwork/nex-go/v2/types"
	ranking_types "github.com/PretendoNetwork/nex-protocols-go/v2/ranking/types"
)

func InsertRankingByPIDAndRankingScoreData(pid types.PID, rankingScoreData ranking_types.RankingScoreData, uniqueID types.UInt64) error {
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
		uint32(pid),
		uint32(rankingScoreData.Category),
		uint32(rankingScoreData.Score),
		uint8(rankingScoreData.OrderBy),
		uint8(rankingScoreData.UpdateMode),
		[]byte(rankingScoreData.Groups),
		uint64(rankingScoreData.Param),
		make([]byte, 0),
		now,
	)

	return err
}
