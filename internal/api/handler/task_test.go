package handler

import (
	"bytes"
	"context"
	"github.com/Batyr2024/todoGolang/domain"
	"github.com/Batyr2024/todoGolang/internal/usecase"
	mocks "github.com/Batyr2024/todoGolang/mocks/handl"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTask_Create(t *testing.T) {
	type mockBehavior func(s *mocks.MockinterfaceUseCase, task domain.Task, ctx context.Context)

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
				ID:      0,
				Text:    "Kyhksjdhsf",
				Checked: false,
			},
			mockBehavior: func(s *mocks.MockinterfaceUseCase, task domain.Task, ctx context.Context) {
				s.EXPECT().Create(ctx, task).Return(nil)
			},
			expectedCode: http.StatusCreated,
			expectedBody: `1`,
		},
		{
			name:         "Empty body",
			inputBody:    `{"Texted":"Kyhksjdhsf"}`,
			mockBehavior: func(s *mocks.MockinterfaceUseCaseMock, task domain.Task, ctx context.Context) {},
			expectedCode: http.StatusBadRequest,
			expectedBody: `{"error":"Text field is missing"}`,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			task := mocks.NewMockinterfaceHandler(ctrl)

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

func TestTask_FindAll(t *testing.T) {
	type mockBehavior func(s *mocks.MockinterfaceHandler, ctx context.Context)
	type bodyType interface{}
	testTable := []struct {
		name         string
		inputTasks   []domain.Task
		mockBehavior mockBehavior
		expectedCode int
		expectedBody bodyType
	}{
		{
			name: "OK",
			mockBehavior: func(s *mocks.MockinterfaceHandler, ctx context.Context) {

				s.EXPECT().FindAll(ctx).Return([]*domain.Task{}, nil)
			},
			inputTasks:   []domain.Task{},
			expectedCode: http.StatusOK,
			expectedBody: []domain.Task{},
		},
		{
			name:         "Empty body",
			mockBehavior: func(s *mocks.MockinterfaceHandler, ctx context.Context) {},
			expectedCode: http.StatusBadRequest,
			expectedBody: `{"error":"Text field is missing"}`,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			task := mocks.NewMockinterfaceHandler(ctrl)

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

func TestTask_ChangeCheckedByID(t *testing.T) {
	type mockBehavior func(s *mocks.MockinterfaceHandler, ctx context.Context, id int, checked bool)
	testTable := []struct {
		name         string
		inputId      int
		inputCheck   bool
		mockBehavior mockBehavior
		expectedCode int
		expectedBody interface{}
	}{
		{
			name: "OK",
			mockBehavior: func(s *mocks.MockinterfaceHandler, ctx context.Context, id int, checked bool) {
				s.EXPECT().ChangeCheckedByID(ctx, id, checked).Return(nil)
			},
			expectedCode: http.StatusOK,
			expectedBody: 1,
		},
		{
			name:         "Empty id",
			mockBehavior: func(s *mocks.MockinterfaceHandler, ctx context.Context, id int, checked bool) {},
			expectedCode: http.StatusBadRequest,
			expectedBody: `{"error":"Id field is missing"}`,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			task := mocks.NewMockinterfaceHandler(ctrl)

			req, _ := http.NewRequest("GET", "/", nil)

			testCase.mockBehavior(task, req.Context(), 1, false)

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
