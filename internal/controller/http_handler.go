package controller

import (
	"context"
	"golang/internal/entity"
	"golang/internal/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
    userUC *usecase.UserUsecase
}

// NewController nhận usecase như dependency injection
func NewController(userUC *usecase.UserUsecase) *Controller {
    return &Controller{userUC: userUC}
}

// RegisterRoutes gắn tất cả các route HTTP
func (ctrl *Controller) RegisterRoutes(r *gin.Engine) {
    r.POST("/users", ctrl.CreateUser)
}

// HTTP Handlers
func (ctrl *Controller) CreateUser(c *gin.Context) {
    var req entity.User
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := ctrl.userUC.CreateUser(context.Background(), &req); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, req)
}


