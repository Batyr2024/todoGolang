package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/Batyr2024/todoGolang/domain"
	"github.com/Batyr2024/todoGolang/internal/usecase"
	mocks "github.com/Batyr2024/todoGolang/mocks"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func mockDataNew() []*domain.Task {
	data := []*domain.Task{}
	var i int32 = 0
	for ; i < 20; i++ {
		if i%2 == 0 {
			data = append(data, &domain.Task{ID: i, Text: "example text", Checked: true})
			continue
		}
		data = append(data, &domain.Task{ID: i, Text: "example text", Checked: false})
	}
	return data
}

func Test_TaskCreate(t *testing.T) {
	type mockBehavior func(s *mocks.MockinterfaceUseCase, task domain.Task, ctx context.Context)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	testTable := []struct {
		name         string
		inputBody    string
		inputTask    domain.Task
		mockBehavior mockBehavior
		expectedCode int
		expectedBody string
	}{
		{
			name:      "OK",
			inputBody: `{"Text":"Kyhksjdhsf"}`,
			inputTask: domain.Task{
				Text: "Kyhksjdhsf",
			},
			mockBehavior: func(s *mocks.MockinterfaceUseCase, task domain.Task, ctx context.Context) {
				s.EXPECT().Create(ctx, task).Return(&domain.Task{ID: 0, Text: "Kyhksjdhsf", Checked: false}, nil)
			},
			expectedCode: http.StatusCreated,
			expectedBody: `{"id":0,"text":"Kyhksjdhsf","checked":false}`,
		},
		{
			name:         "Empty body",
			inputBody:    `{"Texted":"Kyhksjdhsf"}`,
			mockBehavior: func(s *mocks.MockinterfaceUseCase, task domain.Task, ctx context.Context) {},
			expectedCode: http.StatusBadRequest,
			expectedBody: `{"error":"Text field is missing"}`,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			task := mocks.NewMockinterfaceUseCase(ctrl)

			req, _ := http.NewRequest("POST", "/", bytes.NewBufferString(testCase.inputBody))

			testCase.mockBehavior(task, testCase.inputTask, req.Context())

			useCase := usecase.New(task)
			handler := New(useCase)

			r := gin.New()
			r.POST("/", handler.Create)
			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedCode, w.Code)
			assert.Equal(t, testCase.expectedBody, w.Body.String())
		})
	}
}

func Test_TaskFindAll(t *testing.T) {
	type mockBehavior func(s *mocks.MockinterfaceUseCase, ctx context.Context)
	mockData := mockDataNew()
	jsonMockDataByte, _ := json.Marshal(mockData)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	testTable := []struct {
		name         string
		mockBehavior mockBehavior
		expectedCode int
		expectedBody string
	}{
		{
			name: "OK",
			mockBehavior: func(s *mocks.MockinterfaceUseCase, ctx context.Context) {
				s.EXPECT().FindAll(ctx).Return(mockData, nil)
			},
			expectedCode: http.StatusOK,
			expectedBody: string(jsonMockDataByte),
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			task := mocks.NewMockinterfaceUseCase(ctrl)

			req, _ := http.NewRequest("GET", "/", nil)

			testCase.mockBehavior(task, req.Context())

			useCase := usecase.New(task)
			handler := New(useCase)

			r := gin.New()
			r.GET("/", handler.FindAll)
			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedCode, w.Code)
			assert.Equal(t, testCase.expectedBody, w.Body.String())
		})
	}
}

func Test_TaskDeleteByID(t *testing.T) {
	type mockBehavior func(s *mocks.MockinterfaceUseCase, id string, ctx context.Context)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	testTable := []struct {
		name         string
		inputQuery   string
		mockBehavior mockBehavior
		expectedCode int
		expectedBody string
	}{
		{
			name:       "OK",
			inputQuery: "1",
			mockBehavior: func(s *mocks.MockinterfaceUseCase, id string, ctx context.Context) {
				queryId, _ := strconv.Atoi(id)
				s.EXPECT().DeleteByID(ctx, queryId).Return(nil)
			},
			expectedCode: http.StatusNoContent,
		},
		{
			name:         "Empty QUERY",
			inputQuery:   "bad",
			mockBehavior: func(s *mocks.MockinterfaceUseCase, id string, ctx context.Context) {},
			expectedCode: http.StatusInternalServerError,
		},
		{
			name:       "Error repo",
			inputQuery: "133",
			mockBehavior: func(s *mocks.MockinterfaceUseCase, id string, ctx context.Context) {
				queryId, _ := strconv.Atoi(id)
				err := errors.New("Not Found")
				s.EXPECT().DeleteByID(ctx, queryId).Return(err)
			},
			expectedCode: http.StatusNotFound,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			task := mocks.NewMockinterfaceUseCase(ctrl)

			req, _ := http.NewRequest("DELETE", "/"+testCase.inputQuery, nil)

			testCase.mockBehavior(task, testCase.inputQuery, req.Context())

			useCase := usecase.New(task)
			handler := New(useCase)

			r := gin.New()
			r.DELETE("/:id", handler.DeleteByID)
			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedCode, w.Code)
		})
	}
}

func Test_TaskDeleteAll(t *testing.T) {
	type mockBehavior func(s *mocks.MockinterfaceUseCase, ctx context.Context)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	testTable := []struct {
		name         string
		mockBehavior mockBehavior
		expectedCode int
	}{
		{
			name: "OK",
			mockBehavior: func(s *mocks.MockinterfaceUseCase, ctx context.Context) {
				s.EXPECT().DeleteAll(ctx).Return(nil)
			},
			expectedCode: http.StatusNoContent,
		},
		{
			name: "Error repo",
			mockBehavior: func(s *mocks.MockinterfaceUseCase, ctx context.Context) {
				err := errors.New("Not Found")
				s.EXPECT().DeleteAll(ctx).Return(err)
			},
			expectedCode: http.StatusNotFound,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			task := mocks.NewMockinterfaceUseCase(ctrl)

			req, _ := http.NewRequest("DELETE", "/", nil)

			testCase.mockBehavior(task, req.Context())

			useCase := usecase.New(task)
			handler := New(useCase)

			r := gin.New()
			r.DELETE("/", handler.DeleteAll)
			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedCode, w.Code)
		})
	}
}

func Test_TaskChangeCheckedByID(t *testing.T) {
	type mockBehavior func(s *mocks.MockinterfaceUseCase, id string, checked string, ctx context.Context)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	testTable := []struct {
		name            string
		inputQueryId    string
		inputQueryCheck string
		mockBehavior    mockBehavior
		expectedCode    int
	}{
		{
			name:            "OK",
			inputQueryId:    "1",
			inputQueryCheck: "true",
			mockBehavior: func(s *mocks.MockinterfaceUseCase, id string, checked string, ctx context.Context) {
				queryId, _ := strconv.Atoi(id)
				queryChecked, _ := strconv.ParseBool(checked)
				s.EXPECT().ChangeCheckedByID(ctx, queryId, queryChecked).Return(nil)
			},
			expectedCode: http.StatusNoContent,
		},
		{
			name:            "Error QUERY",
			inputQueryId:    "ad",
			inputQueryCheck: "uead",
			mockBehavior:    func(s *mocks.MockinterfaceUseCase, id string, checked string, ctx context.Context) {},
			expectedCode:    http.StatusInternalServerError,
		},
		{
			name:            "Empty QUERY",
			inputQueryId:    "",
			inputQueryCheck: "",
			mockBehavior:    func(s *mocks.MockinterfaceUseCase, id string, checked string, ctx context.Context) {},
			expectedCode:    http.StatusBadRequest,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			task := mocks.NewMockinterfaceUseCase(ctrl)

			req, _ := http.NewRequest("PATCH", "/?id="+testCase.inputQueryId+"&&check="+testCase.inputQueryCheck, nil)

			testCase.mockBehavior(task, testCase.inputQueryId, testCase.inputQueryCheck, req.Context())

			useCase := usecase.New(task)
			handler := New(useCase)

			r := gin.New()
			r.PATCH("/", handler.ChangeCheckedByID)
			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedCode, w.Code)
		})
	}
}

func Test_TaskChangeCheckedAll(t *testing.T) {
	type mockBehavior func(s *mocks.MockinterfaceUseCase, checked string, ctx context.Context)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	testTable := []struct {
		name            string
		inputQueryCheck string
		mockBehavior    mockBehavior
		expectedCode    int
	}{
		{
			name:            "OK",
			inputQueryCheck: "true",
			mockBehavior: func(s *mocks.MockinterfaceUseCase, checked string, ctx context.Context) {
				queryChecked, _ := strconv.ParseBool(checked)
				s.EXPECT().ChangeCheckedAll(ctx, queryChecked).Return(nil)
			},
			expectedCode: http.StatusNoContent,
		},
		{
			name:            "Error QUERY",
			inputQueryCheck: "falsedq",
			mockBehavior:    func(s *mocks.MockinterfaceUseCase, checked string, ctx context.Context) {},
			expectedCode:    http.StatusInternalServerError,
		},
		{
			name:            "Error Repo",
			inputQueryCheck: "false",
			mockBehavior: func(s *mocks.MockinterfaceUseCase, checked string, ctx context.Context) {
				err := errors.New("Not Found")
				queryChecked, _ := strconv.ParseBool(checked)
				s.EXPECT().ChangeCheckedAll(ctx, queryChecked).Return(err)
			},
			expectedCode: http.StatusNotFound,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			task := mocks.NewMockinterfaceUseCase(ctrl)

			req, _ := http.NewRequest("PATCH", "/"+testCase.inputQueryCheck, nil)

			testCase.mockBehavior(task, testCase.inputQueryCheck, req.Context())

			useCase := usecase.New(task)
			handler := New(useCase)

			r := gin.New()
			r.PATCH("/:check", handler.ChangeCheckedAll)
			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedCode, w.Code)
		})
	}
}

func Test_TaskChangeText(t *testing.T) {
	type mockBehavior func(s *mocks.MockinterfaceUseCase, task domain.Task, ctx context.Context)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	testTable := []struct {
		name         string
		inputBody    string
		inputTask    domain.Task
		mockBehavior mockBehavior
		expectedCode int
	}{
		{
			name:      "OK",
			inputBody: `{"id":1,"text":"Kyhksjdhsf"}`,
			inputTask: domain.Task{
				ID:   1,
				Text: "Kyhksjdhsf",
			},
			mockBehavior: func(s *mocks.MockinterfaceUseCase, task domain.Task, ctx context.Context) {
				s.EXPECT().ChangeText(ctx, int(task.ID), task.Text).Return(nil)
			},
			expectedCode: http.StatusOK,
		},
		{
			name:         "Empty body",
			inputBody:    `{"id":0,"text":""}`,
			mockBehavior: func(s *mocks.MockinterfaceUseCase, task domain.Task, ctx context.Context) {},
			expectedCode: http.StatusBadRequest,
		},
		{
			name:      "Error repo",
			inputBody: `{"id":3,"text":"dsf"}`,
			inputTask: domain.Task{
				ID:   3,
				Text: "dsf",
			},
			mockBehavior: func(s *mocks.MockinterfaceUseCase, task domain.Task, ctx context.Context) {
				err := errors.New("Not Found")
				s.EXPECT().ChangeText(ctx, int(task.ID), task.Text).Return(err)
			},
			expectedCode: http.StatusNotFound,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			task := mocks.NewMockinterfaceUseCase(ctrl)

			req, _ := http.NewRequest("PUT", "/", bytes.NewBufferString(testCase.inputBody))

			testCase.mockBehavior(task, testCase.inputTask, req.Context())

			useCase := usecase.New(task)
			handler := New(useCase)

			r := gin.New()
			r.PUT("/", handler.ChangeText)
			w := httptest.NewRecorder()

			r.ServeHTTP(w, req)

			assert.Equal(t, testCase.expectedCode, w.Code)
		})
	}
}
