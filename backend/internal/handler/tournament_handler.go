package handler

import (
	"errors"
	"net/http"

	"github.com/febriantarigan/berpadel/internal/domain"
	"github.com/febriantarigan/berpadel/internal/handler/dto"
	"github.com/febriantarigan/berpadel/internal/handler/response"
	"github.com/febriantarigan/berpadel/internal/service"
	"github.com/gin-gonic/gin"
)

type TournamentHandler struct {
	tournamentService *service.TournamentService
}

func NewTournamentHandler(ts *service.TournamentService) *TournamentHandler {
	return &TournamentHandler{tournamentService: ts}
}

func GetTournament(c *gin.Context) {

}

func (h *TournamentHandler) CreateTournament(c *gin.Context) {
	var req dto.CreateTournamentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, response.NewError("invalid_request", err.Error()))
		return
	}

	tType, err := domain.ParseTournamentType(req.Type)
	if err != nil {
		response.Error(c, http.StatusBadRequest, response.NewError("invalid_request", err.Error()))
		return
	}

	tStatus, err := domain.ParseTournamentStatus(req.Status)
	if err != nil {
		response.Error(c, http.StatusBadRequest, response.NewError("invalid_request", err.Error()))
		return
	}

	tPlayers, err := mapPlayers(req.Players)
	if err != nil {
		response.Error(c, http.StatusBadRequest, response.NewError("invalid_request", err.Error()))
		return
	}

	result, err := h.tournamentService.CreateTournament(c.Request.Context(), service.CreateTournamentInput{
		Name:      req.Name,
		Type:      tType,
		Status:    tStatus,
		Location:  req.Location,
		Season:    req.Season,
		MaxPoints: req.MaxPoints,
		Players:   tPlayers,
		Courts:    req.Courts,
	})
	if err != nil {
		response.Error(c, http.StatusInternalServerError, response.NewError("internal_server_error", err.Error()))
		return
	}

	matches := make([]*domain.Match, 0)
	for _, s := range result.Rounds {
		matches = append(matches, s.Matches...)
	}
	response.Success(c, http.StatusCreated, response.NewTournamentResponse(result.Tournament, matches))
}

func mapPlayers(req []dto.PlayerRequest) ([]service.PlayerInput, error) {
	players := make([]service.PlayerInput, 0, len(req))

	for _, p := range req {
		switch {
		case p.UserID != nil:
			players = append(players, service.PlayerInput{
				UserID: p.UserID,
			})

		case p.Name != nil && p.Gender != nil:
			gender, err := domain.ParseGender(*p.Gender)
			if err != nil {
				return nil, err
			}

			players = append(players, service.PlayerInput{
				Name:   p.Name,
				Gender: gender,
			})

		default:
			return nil, errors.New("invalid player input")
		}
	}

	return players, nil
}
