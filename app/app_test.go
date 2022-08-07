package app_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Fajar-Islami/ais_code_test/dto"
	"github.com/Fajar-Islami/ais_code_test/entity"
	"github.com/Fajar-Islami/ais_code_test/helper"
	"github.com/Fajar-Islami/ais_code_test/infrastructure/db"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

var handlerArticle, pgsql = createTestArticleApp()

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

func TestCreateArticle(t *testing.T) {
	r := SetUpRouter()
	r.POST("/api/article", handlerArticle.NewArticle)

	defer db.CloseDatabaseConnection(pgsql)

	createArticle := dto.CreateArticleDTO{
		Author: "Testing Author",
		Title:  "Testing Title",
		Body:   "Testing Body",
	}

	reqBody, _ := json.Marshal(createArticle)
	req := httptest.NewRequest("POST", "/api/article", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	responBody, _ := ioutil.ReadAll(w.Body)

	respHelper := helper.Response{}
	json.Unmarshal(responBody, &respHelper)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, true, respHelper.Status)

	jsonArticle, _ := json.Marshal(respHelper.Data)
	createdArticleData := entity.Article{}
	json.Unmarshal(jsonArticle, &createdArticleData)
	assert.NotNil(t, createdArticleData.ID)
	assert.Equal(t, createArticle.Author, createdArticleData.Author)
	assert.Equal(t, createArticle.Title, createdArticleData.Title)
	assert.Equal(t, createArticle.Body, createdArticleData.Body)
}

func TestGetArticle(t *testing.T) {
	r := SetUpRouter()
	r.GET("/api/article", handlerArticle.GetListArticle)

	defer db.CloseDatabaseConnection(pgsql)

	req := httptest.NewRequest("GET", "/api/article", nil)

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	responBody, _ := ioutil.ReadAll(w.Body)

	respHelper := helper.Response{}
	json.Unmarshal(responBody, &respHelper)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, true, respHelper.Status)

	jsonArticle, _ := json.Marshal(respHelper.Data)
	listArticleData := []entity.Article{}
	json.Unmarshal(jsonArticle, &listArticleData)
	assert.NotEmpty(t, len(listArticleData))

	for _, data := range listArticleData {
		assert.NotNil(t, data.ID)
		assert.NotNil(t, data.Author)
		assert.NotNil(t, data.Title)
		assert.NotNil(t, data.Body)
	}
}

func TestGetArticleFIlter(t *testing.T) {
	r := SetUpRouter()
	r.GET("/api/article", handlerArticle.GetListArticle)

	defer db.CloseDatabaseConnection(pgsql)

	dataTest := []struct {
		url       string
		searchReq string
		authorReq string
	}{
		{
			url:       "/api/article?search=Testing",
			searchReq: "Testing",
			authorReq: "",
		},
		{
			url:       "/api/article?author=Testing%20Author",
			searchReq: "",
			authorReq: "Testing Author",
		},
		{
			url:       "/api/article?search=Testing&author=Testing%20Author",
			searchReq: "Testing",
			authorReq: "Testing Author",
		},
	}

	for _, test := range dataTest {

		req := httptest.NewRequest("GET", test.url, nil)

		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		responBody, _ := ioutil.ReadAll(w.Body)

		respHelper := helper.Response{}
		json.Unmarshal(responBody, &respHelper)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, true, respHelper.Status)

		jsonArticle, _ := json.Marshal(respHelper.Data)
		listArticleData := []entity.Article{}
		json.Unmarshal(jsonArticle, &listArticleData)
		assert.NotEmpty(t, len(listArticleData))

		for _, data := range listArticleData {
			assert.NotNil(t, data.ID)
			assert.NotNil(t, data.Author)
			assert.NotNil(t, data.Title)
			assert.NotNil(t, data.Body)

			if test.authorReq != "" {
				assert.Equal(t, true, strings.Contains(data.Author, test.authorReq))
			}

			if test.searchReq != "" {
				assert.Equal(t, true, (strings.Contains(data.Body, test.searchReq) || strings.Contains(data.Title, test.searchReq)))
			}

		}
	}
}
