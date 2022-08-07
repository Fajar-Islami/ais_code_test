package app_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

var handlerArticle = createTestArticleApp()

func TestHealthHandler(t *testing.T) {
	mockResponse := `{"status":"ok"}`
	r := SetUpRouter()
	r.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, map[string]string{
			"status": "ok",
		})
	})

	req := httptest.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	responseData, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)
}

// func TestCreateArticle(t *testing.T) {
// 	r := SetUpRouter()

// 	r.POST("/api/article", handlerArticle.NewArticle)

// 	createArticle := dto.CreateArticleDTO{
// 		Author: "Testing",
// 		Title:  "Testing Title",
// 		Body:   "Testing Body",
// 	}

// 	reqBody, _ := json.Marshal(createArticle)
// 	req := httptest.NewRequest("POST", "/api/article", bytes.NewBuffer(reqBody))
// 	w := httptest.NewRecorder()

// 	r.ServeHTTP(w, req)

// 	assert.Equal(t, http.StatusCreated, w.Code)

// 	jsonData, _ := json.Marshal()

// }
