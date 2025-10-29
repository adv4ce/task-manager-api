package handlers

import (
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"
    "task_manager/internal/models"
		"task_manager/internal/config"
    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
)

var (
	testContainer *models.TaskLib
	testRouter *gin.Engine
)

func TestMain(m *testing.M) {
  setupTest()
  m.Run()
}

func setupTest() {
	cfg := config.Load()
  gin.SetMode(gin.TestMode)
  testContainer = models.CreateContainer()
  testRouter = CreateRouter(testContainer, cfg)
}

func TestHealth(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/health", nil)
	w := httptest.NewRecorder()

	testRouter.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "healthy")
}

func TestCreateFirst(t *testing.T) {
	json := `{"title":"Первая задача","description":"Тестовое описание","status":"in progress","priority":"medium"}`
	req, _ := http.NewRequest("POST", "/api/tasks", strings.NewReader(json))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	testRouter.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "successful")
	assert.Contains(t, w.Body.String(), "Первая задача")
	assert.Contains(t, w.Body.String(), "Task created")
}

func TestCreateSecond(t *testing.T) {
	json := `{"title":"Вторая задача","description":"Тестовое описание","status":"in progress","priority":"medium"}`
	req, _ := http.NewRequest("POST", "/api/tasks", strings.NewReader(json))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	testRouter.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "successful")
	assert.Contains(t, w.Body.String(), "Вторая задача")
	assert.Contains(t, w.Body.String(), "Task created")
}

func TestGetOne(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/tasks/1", nil)
	w := httptest.NewRecorder()

	testRouter.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "successful")
	assert.Contains(t, w.Body.String(), `"id":1`)
	assert.Contains(t, w.Body.String(), "Первая задача")
}

func TestGetAll(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/tasks", nil)
	w := httptest.NewRecorder()

	testRouter.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "successful")
	assert.Contains(t, w.Body.String(), "Первая задача")
}

func TestDelete(t *testing.T) {
	req, _ := http.NewRequest("DELETE", "/api/tasks/2", nil)
	w := httptest.NewRecorder()

	testRouter.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "successful")
	assert.Contains(t, w.Body.String(), `"deletedID":2`)
}

func TestUpdate(t *testing.T) {
	json := `{"title":"Изменение первой задачи","description":"Изменение тестового описания","status":"In progress","priority":"medium"}`
	req, _ := http.NewRequest("PUT", "/api/tasks/1", strings.NewReader(json))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	testRouter.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "successful")
	assert.Contains(t, w.Body.String(), "Updated successfully")
	assert.Contains(t, w.Body.String(), "Изменение первой задачи")
}

func TestPatch(t *testing.T) {
	json := `{"title":"Патч первой задачи"}`
	req, _ := http.NewRequest("PATCH", "/api/tasks/1", strings.NewReader(json))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	testRouter.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "successful")
	assert.Contains(t, w.Body.String(), "Updated successfully")
	assert.Contains(t, w.Body.String(), "Патч первой задачи")
}

func TestClose(t *testing.T) {
	req, _ := http.NewRequest("POST", "/api/tasks/1", nil)

	w := httptest.NewRecorder()

	testRouter.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "successful")
	assert.Contains(t, w.Body.String(), "Task closed")
	assert.Contains(t, w.Body.String(), `"id":1`)
}

func TestLastGetAll(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/tasks", nil)
	w := httptest.NewRecorder()

	testRouter.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "successful")
	assert.Contains(t, w.Body.String(), "Патч первой задачи")
	assert.Contains(t, w.Body.String(), "Completed")
}