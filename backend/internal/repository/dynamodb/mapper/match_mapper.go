package mapper

import (
	"github.com/febriantarigan/berpadel/internal/domain"
	"github.com/febriantarigan/berpadel/internal/repository/dynamodb/model"
	"github.com/febriantarigan/berpadel/internal/util"
)

// map domain to item for match entity
func ToMatchItem(m domain.Match) *model.MatchItem {
	return &model.MatchItem{
		PK:               "TOURNAMENT#" + m.TournamentID,
		SK:               "MATCH#" + m.ID,
		MatchID:          m.ID,
		Round:            m.Round,
		Court:            m.Court,
		TeamA:            util.UsersToStrings(m.TeamA.Players),
		TeamB:            util.UsersToStrings(m.TeamB.Players),
		TeamAScore:       m.TeamA.Score,
		TeamBScore:       m.TeamB.Score,
		Status:           string(m.Status),
		WaitingPlayerIDs: m.WaitingPlayerIDs,
		CreatedAt:        util.ToString(m.CreatedAt),
		UpdatedAt:        util.ToString(m.UpdatedAt),
	}
}

// map domain to item for match entity
func ToDomainMatch(item model.MatchItem) *domain.Match {
	return &domain.Match{
		TournamentID: extractKey("TOURNAMENT#", item.PK),
		ID:           extractKey("MATCH#", item.SK),
		Round:        item.Round,
		Court:        item.Court,
		TeamA: domain.Team{
			Players: util.StringsToUsers(item.TeamA),
			Score:   item.TeamAScore,
		},
		TeamB: domain.Team{
			Players: util.StringsToUsers(item.TeamB),
			Score:   item.TeamBScore,
		},
		Status:           domain.MatchStatus(item.Status),
		WaitingPlayerIDs: item.WaitingPlayerIDs,
		CreatedAt:        util.ToTime(item.CreatedAt),
		UpdatedAt:        util.ToTime(item.UpdatedAt),
	}
}
