package repository

import (
	"context"
	"github.com/Batyr2024/todoGolang/config"
	"github.com/Batyr2024/todoGolang/db"
	"github.com/go-playground/assert/v2"

	//DB "github.com/Batyr2024/todoGolang/db/sqlc"
	"github.com/Batyr2024/todoGolang/domain"
	"github.com/Batyr2024/todoGolang/mocks"
	"github.com/golang/mock/gomock"

	"testing"
)

//type testRepository struct {
//	db *DB.Queries
//}
//
//func New() *testRepository {
//	connect, _ := pgx.Connect(context.Background(), "postgres://dunice:dunice@localhost:5432/todo_test?sslmode=disable")
//	return &testRepository{DB.New(connect)}
//}
//
//var tRepo *testRepository = New();

func Test_TaskCreate(t *testing.T) {
	type mockBehavior func(s *mocks.MockQuerier, task domain.Task, ctx context.Context)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	testTable := []struct {
		name         string
		inputTask    domain.Task
		mockBehavior mockBehavior
		expectedBody *domain.Task
		expectedErr  error
	}{
		{
			name:      "OK",
			inputTask: domain.Task{Text: "Text15"},
			mockBehavior: func(s *mocks.MockQuerier, task domain.Task, ctx context.Context) {
				s.EXPECT().Create(ctx, task.Text).Return(&domain.Task{ID: task.ID, Text: "Text15", Checked: false}, nil)
			},
			expectedBody: &domain.Task{ID: 0, Text: "Text15", Checked: false},
			expectedErr:  nil,
		},
		{
			name:         "Empty body",
			mockBehavior: func(s *mocks.MockQuerier, task domain.Task, ctx context.Context) {},
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			task := mocks.NewMockQuerier(ctrl)
			dataBase := db.Connect(config.Config{DBUrl: "postgres://dunice:dunice@localhost:5432/todo"})
			repo := New(dataBase)
			allData, _ := repo.FindAll(context.Background())
			testCase.expectedBody.ID = allData[len(allData)-1].ID + 1

			w, err := repo.Create(context.Background(), testCase.inputTask)

			testCase.mockBehavior(task, *testCase.expectedBody, context.Background())
			assert.Equal(t, testCase.expectedBody, w)
			assert.Equal(t, testCase.expectedErr, err)
		})
	}
}
