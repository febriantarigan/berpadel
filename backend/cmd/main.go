package main

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	ginadapter "github.com/awslabs/aws-lambda-go-api-proxy/gin"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	iconfig "github.com/febriantarigan/berpadel/internal/config"
	"github.com/febriantarigan/berpadel/internal/handler"
	repo "github.com/febriantarigan/berpadel/internal/repository/dynamodb"
	"github.com/febriantarigan/berpadel/internal/routes"
	"github.com/febriantarigan/berpadel/internal/service"
)

var ginLambda *ginadapter.GinLambdaV2

/*func init() {
	logger, _ := zap.NewProduction()
	undo := zap.ReplaceGlobals(logger)
	defer undo()

	if os.Getenv("APP_ENV") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()

	r.Use(middleware.StructuredLogger(logger))
	r.Use(middleware.ZapRecovery(logger))
	r.Use(middleware.ErrorHandler())
	r.Use(middleware.CacheRequestBody())

	routes.SetupRouter(r)

	// 5. Setup Proxy for AWS Lambda
	// NewV2 handles the 'Payload Format 2.0' used by AWS HTTP APIs
	ginLambda = ginadapter.NewV2(r)
}

// Handler is the function AWS Lambda calls on every request
func Handler(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	return ginLambda.ProxyWithContext(ctx, req)
}*/

func main() {
	if os.Getenv("AWS_LAMBDA_FUNCTION_NAME") != "" {
		//lambda.Start(Handler)
	} else {
		cfg, err := config.LoadDefaultConfig(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		if err := iconfig.InitConfig(); err != nil {
			log.Fatal(err)
		}

		db := dynamodb.NewFromConfig(cfg)
		baseRepo := repo.NewBaseRepository(db)

		tournamentRepo := repo.NewTournamentRepository(baseRepo)
		matchRepo := repo.NewMatchRepository(baseRepo)
		leaderboardRepo := repo.NewLeaderboardRepository(baseRepo)
		userRepo := repo.NewUserRepository(baseRepo)

		userService := service.NewUserService(userRepo)
		tournamentService := service.NewTournamentService(userRepo, tournamentRepo, matchRepo, leaderboardRepo)
		matchService := service.NewMatchService(db, matchRepo, leaderboardRepo, userRepo)
		leaderboardService := service.NewLeaderboardService(leaderboardRepo)

		userHandler := handler.NewUserHandler(userService)
		tournamentHandler := handler.NewTournamentHandler(tournamentService)
		matchHandler := handler.NewMatchHandler(matchService)
		leaderboardHandler := handler.NewLeaderboardHandler(leaderboardService)

		r := gin.Default()
		config := cors.DefaultConfig()
		config.AllowAllOrigins = true
		config.AllowMethods = []string{"GET", "POST", "OPTIONS"}
		config.AllowHeaders = []string{"Origin", "Content-Type", "Accept"}
		r.Use(cors.New(config))
		routes.SetupRouter(r, routes.Handlers{
			User:        userHandler,
			Tournament:  tournamentHandler,
			Match:       matchHandler,
			Leaderboard: leaderboardHandler,
		})
		r.Run(":8080")
	}
}
