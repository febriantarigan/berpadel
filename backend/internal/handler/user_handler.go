package handler

import (
	"github.com/febriantarigan/berpadel/internal/domain"
	"github.com/febriantarigan/berpadel/internal/service"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(us *service.UserService) *UserHandler {
	return &UserHandler{userService: us}
}

func (h *UserHandler) GetUsers(c *gin.Context) {
	id := c.Query("id")
	search := c.Query("search")

	ctx := c.Request.Context()

	// 🔹 get single user by ID
	if id != "" {
		user, err := h.userService.GetByID(ctx, id)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		if user == nil {
			c.JSON(404, gin.H{"error": "user not found"})
			return
		}

		c.JSON(200, user)
		return
	}

	// 🔹 list or search users
	var (
		users []*domain.User
		err   error
	)

	if search != "" {
		users, err = h.userService.SearchByName(ctx, search)
	} else {
		users, err = h.userService.List(ctx)
	}

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, users)
}
