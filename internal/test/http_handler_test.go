package controller_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"golang/internal/controller"
	"golang/internal/entity"
	mocks "golang/internal/mock"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateUser_Success(t *testing.T) {
	mockUC := new(mocks.UserUC)
	mockUC.On("CreateUser", mock.Anything, mock.AnythingOfType("*entity.User")).Return(nil)

	router := gin.Default()
	ctrl := controller.NewController(mockUC)
	ctrl.RegisterRoutes(router)

	user := &entity.User{
		Name:  "Nguyen Quy",
		Email: "quy@example.com",
	}
	body, _ := json.Marshal(user)
	req, _ := http.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	fmt.Print(w.Code)
	mockUC.AssertExpectations(t)
}
