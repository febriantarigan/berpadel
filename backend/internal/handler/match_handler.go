package handler

import (
	"net/http"

	"github.com/febriantarigan/berpadel/internal/handler/dto"
	"github.com/febriantarigan/berpadel/internal/service"
	"github.com/gin-gonic/gin"
)

type MatchHandler struct {
	matchService *service.MatchService
}

func NewMatchHandler(ms *service.MatchService) *MatchHandler {
	return &MatchHandler{matchService: ms}
}

func (h *MatchHandler) SubmitScore(c *gin.Context) {
	//matchID := c.Param("matchId")
	//tournamentID := c.Param("tournamentId")

	var req dto.SubmitMatchScoreRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	/*match, err := h.matchService.GetMatch(c.Request.Context(), tournamentID, matchID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "match not found"})
		return
	}

	score := domain.MatchScore{
		TeamAScore: req.TeamAScore,
		TeamBScore: req.TeamBScore,
	}

	/*err = h.matchService.SubmitScore(c.Request.Context(), service.PutMatchScoreInput{
		Match: match,
		Score: score,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}*/

	c.Status(http.StatusNoContent)
}
