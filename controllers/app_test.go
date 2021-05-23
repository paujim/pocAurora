package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/rdsdataservice"
	"github.com/gin-gonic/gin"
	"github.com/paujim/pocAurora/repositories"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockDataService struct {
	mock.Mock
}

func (m *MockDataService) ExecuteStatement(input *rdsdataservice.ExecuteStatementInput) (*rdsdataservice.ExecuteStatementOutput, error) {
	args := m.Called(input)
	var resp *rdsdataservice.ExecuteStatementOutput
	if args.Get(0) != nil {
		resp = args.Get(0).(*rdsdataservice.ExecuteStatementOutput)
	}
	return resp, args.Error(1)
}

func performRequest(r http.Handler, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(http.MethodGet, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestHappyPath(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockDS := &MockDataService{}
	output := &rdsdataservice.ExecuteStatementOutput{Records: [][]*rdsdataservice.Field{
		{
			{StringValue: aws.String("759d7dea-e763-4d41-9351-95da0f7fbac3")},
			{StringValue: aws.String("Tab Download The App (Bm69)")},
			{LongValue: aws.Int64(2)},
			{StringValue: aws.String("da20428b-bd21-412e-bea4-0a7b625a0778")},
			{StringValue: aws.String("Townsville")},
			{StringValue: aws.String("4a2788f8-e825-4d36-9894-efd4baf1cfae")},
			{StringValue: aws.String("1579236600")},
		},
		{
			{StringValue: aws.String("7d9aaf4e-556d-4ed6-932f-c3c5b77eb1ec")},
			{StringValue: aws.String("Sizzle Here Feb 7 - \"The Great Kiwi Bbq\" Trot")},
			{LongValue: aws.Int64(1)},
			{StringValue: aws.String("a959f9bf-2f49-4089-a8a1-ee0acd823886")},
			{StringValue: aws.String("Alexandra Park")},
			{StringValue: aws.String("161d9be2-e909-4326-8c2c-35ed71fb460b")},
			{StringValue: aws.String("1579237140")},
		},
		{
			{StringValue: aws.String("e9d10a80-f73f-4357-90c8-b0abe100be7d")},
			{StringValue: aws.String("Racv Summer Pacing Cup")},
			{LongValue: aws.Int64(7)},
			{StringValue: aws.String("91cf31c9-d1db-422f-b35a-8290498029ed")},
			{StringValue: aws.String("Cobram")},
			{StringValue: aws.String("161d9be2-e909-4326-8c2c-35ed71fb460b")},
			{StringValue: aws.String("1579236780")},
		},
	}}
	mockDS.On("ExecuteStatement", mock.Anything).Return(output, nil)
	repo := repositories.NewSQLClient(mockDS, aws.String("arn"), aws.String("secret"))
	app := NewApp(repo, gin.Default())

	router := app.SetupServer()

	w := performRequest(router, "/rest/v1/racing?method=nextraces-categorygroup&count=5&include_categories=%5B%224a2788f8-e825-4d36-9894-efd4baf1cfae%22%2C%229daef0d7-bf3c-4f50-921d-8e818c60fe61%22%2C%22161d9be2-e909-4326-8c2c-35ed71fb460b%22%5D")
	assert.Equal(t, http.StatusOK, w.Code)
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)

	assert.Nil(t, err)

	raceIds := response["data"].(map[string]interface{})["category_race_map"].(map[string]interface{})["161d9be2-e909-4326-8c2c-35ed71fb460b"]
	assert.Equal(t, map[string]interface{}{
		"race_ids": []interface{}{
			"7d9aaf4e-556d-4ed6-932f-c3c5b77eb1ec",
			"e9d10a80-f73f-4357-90c8-b0abe100be7d",
		},
	}, raceIds)
}

func TestBadRequestPath(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockDS := &MockDataService{}
	output := &rdsdataservice.ExecuteStatementOutput{}
	mockDS.On("ExecuteStatement", mock.Anything).Return(output, nil)
	repo := repositories.NewSQLClient(mockDS, aws.String("arn"), aws.String("secret"))
	app := NewApp(repo, gin.Default())

	router := app.SetupServer()

	w := performRequest(router, "/rest/v1/racing?count=5&include_categories=%5B%224a2788f8-e825-4d36-9894-efd4baf1cfae%22%2C%229daef0d7-bf3c-4f50-921d-8e818c60fe61%22%2C%22161d9be2-e909-4326-8c2c-35ed71fb460b%22%5D")
	assert.Equal(t, http.StatusBadRequest, w.Code)

	w = performRequest(router, "/rest/v1/racing?method=nextraces-categorygroup&include_categories=%5B%224a2788f8-e825-4d36-9894-efd4baf1cfae%22%2C%229daef0d7-bf3c-4f50-921d-8e818c60fe61%22%2C%22161d9be2-e909-4326-8c2c-35ed71fb460b%22%5D")
	assert.Equal(t, http.StatusBadRequest, w.Code)

	w = performRequest(router, "/rest/v1/racing?method=nextraces-categorygroup&count=5")
	assert.Equal(t, http.StatusBadRequest, w.Code)
}
