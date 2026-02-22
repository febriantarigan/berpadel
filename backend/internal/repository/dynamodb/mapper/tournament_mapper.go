package mapper

import (
	"fmt"

	"github.com/febriantarigan/berpadel/internal/domain"
	"github.com/febriantarigan/berpadel/internal/repository/dynamodb/model"
	"github.com/febriantarigan/berpadel/internal/util"
)

func ToTournamentItem(t domain.Tournament) *model.TournamentItem {
	return &model.TournamentItem{
		PK:        "TOURNAMENT#" + t.ID,
		SK:        "META",
		Name:      t.Name,
		Type:      string(t.Type),
		Status:    string(t.Status),
		Location:  t.Location,
		Season:    t.Season,
		MaxPoints: t.MaxPoints,
		Courts:    t.Courts,
		PlayerIDs: t.PlayerIDs,
		CreatedAt: util.ToString(t.CreatedAt),
		UpdatedAt: util.ToString(t.UpdatedAt),

		GSI1PK: "TOURNAMENT",
		GSI1SK: fmt.Sprintf("%s#%s", t.Season, t.ID),
	}
}

func ToDomainTournament(item model.TournamentItem) *domain.Tournament {
	return &domain.Tournament{
		ID:        extractKey("TOURNAMENT#", item.PK),
		Name:      item.Name,
		Type:      domain.TournamentType(item.Type),
		Status:    domain.TournamentStatus(item.Status),
		Location:  item.Location,
		Season:    item.Season,
		MaxPoints: item.MaxPoints,
		Courts:    item.Courts,
		PlayerIDs: item.PlayerIDs,
		CreatedAt: util.ToTime(item.CreatedAt),
		UpdatedAt: util.ToTime(item.UpdatedAt),
	}
}
