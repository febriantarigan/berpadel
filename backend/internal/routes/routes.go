package routes

import (
	"github.com/febriantarigan/berpadel/internal/handler"
	"github.com/gin-gonic/gin"
)

type Handlers struct {
	User        *handler.UserHandler
	Tournament  *handler.TournamentHandler
	Match       *handler.MatchHandler
	Leaderboard *handler.LeaderboardHandler
}

// SetupRouter organizes all routes and groups
func SetupRouter(r *gin.Engine, h Handlers) {
	// 1. Versioned Group
	v1 := r.Group("/api/v1")
	{
		// Public: Anyone can view tournament results
		//v1.GET("/tournaments/:id", tournamentHandler.GetTournament)
		v1.POST("/tournaments", h.Tournament.CreateTournament)
		//v1.GET("/tournaments/:tournamentId/matches", matchHandler.GetScores)
		v1.PUT("/tournaments/:tournamentId/matches/:match_id/score", h.Match.SubmitScore)
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})
}
