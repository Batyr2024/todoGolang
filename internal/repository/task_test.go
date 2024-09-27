package repository

import (
	"context"
	"errors"
	DB "github.com/Batyr2024/todoGolang/db/sqlc"
	"github.com/go-playground/assert/v2"

	//DB "github.com/Batyr2024/todoGolang/db/sqlc"
	"github.com/Batyr2024/todoGolang/domain"
	"github.com/Batyr2024/todoGolang/mocks"
	"github.com/golang/mock/gomock"

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

func Test_TaskChangeCheckedAll(t *testing.T) {
	type mockBehavior func(s mocks.MockQuerier, check bool, ctx context.Context)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	testTable := []struct {
		name         string
		inputChecked bool
		mockBehavior mockBehavior
		expectedErr  error
	}{
		{
			name:         "OK",
			inputChecked: true,
			mockBehavior: func(s mocks.MockQuerier, check bool, ctx context.Context) {
				s.EXPECT().ChangeCheckedAll(ctx, &DB.ChangeCheckedAllParams{Checked: &check}).Return(nil, nil)
			},
			expectedErr: nil,
		},
		{
			name:         "Empty body",
			inputChecked: true,
			mockBehavior: func(s mocks.MockQuerier, check bool, ctx context.Context) {
				s.EXPECT().ChangeCheckedAll(ctx, &DB.ChangeCheckedAllParams{Checked: &check}).Return(nil, errors.New("error db"))

			},
			expectedErr: errors.New("error db"),
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			task := mocks.NewMockQuerier(ctrl)

			testCase.mockBehavior(*task, testCase.inputChecked, context.Background())
			_, err := task.ChangeCheckedAll(context.Background(), &DB.ChangeCheckedAllParams{Checked: &testCase.inputChecked})
			assert.Equal(t, testCase.expectedErr, err)
		})
	}
}

func Test_TaskChangeCheckedByID(t *testing.T) {
	type mockBehavior func(s mocks.MockQuerier, id int, check bool, ctx context.Context)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	testTable := []struct {
		name         string
		inputID      int
		inputChecked bool
		mockBehavior mockBehavior
		expectedErr  error
	}{
		{
			name:         "OK",
			inputID:      3,
			inputChecked: true,
			mockBehavior: func(s mocks.MockQuerier, id int, check bool, ctx context.Context) {
				s.EXPECT().ChangeCheckedByID(ctx, &DB.ChangeCheckedByIDParams{ID: int32(id), Checked: &check}).Return(nil, nil)
			},
			expectedErr: nil,
		},
		{
			name:         "Empty body",
			inputID:      3,
			inputChecked: true,
			mockBehavior: func(s mocks.MockQuerier, id int, check bool, ctx context.Context) {
				s.EXPECT().ChangeCheckedByID(ctx, &DB.ChangeCheckedByIDParams{ID: int32(id), Checked: &check}).Return(nil, errors.New("error db"))

			},
			expectedErr: errors.New("error db"),
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			task := mocks.NewMockQuerier(ctrl)

			testCase.mockBehavior(*task, testCase.inputID, testCase.inputChecked, context.Background())
			_, err := task.ChangeCheckedByID(context.Background(), &DB.ChangeCheckedByIDParams{ID: int32(testCase.inputID), Checked: &testCase.inputChecked})
			assert.Equal(t, testCase.expectedErr, err)
		})
	}
}

func Test_TaskFindAll(t *testing.T) {
	type mockBehavior func(s mocks.MockQuerier, ctx context.Context)
	mockData := mockDataNew()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	testTable := []struct {
		name         string
		mockBehavior mockBehavior
		expectedBody []*domain.Task
		expectedErr  error
	}{
		{
			name: "OK",
			mockBehavior: func(s mocks.MockQuerier, ctx context.Context) {
				s.EXPECT().GetAll(ctx).Return(mockDataNew(), nil)
			},
			expectedBody: mockData,
			expectedErr:  nil,
		},
		{
			name: "Empty data",
			mockBehavior: func(s mocks.MockQuerier, ctx context.Context) {
				s.EXPECT().GetAll(ctx).Return([]*domain.Task{}, nil)

			},
			expectedBody: []*domain.Task{},
			expectedErr:  nil,
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			task := mocks.NewMockQuerier(ctrl)

			testCase.mockBehavior(*task, context.Background())
			w, err := task.GetAll(context.Background())
			assert.Equal(t, testCase.expectedBody, w)
			assert.Equal(t, testCase.expectedErr, err)
		})
	}
}

func Test_TaskCreate(t *testing.T) {
	type mockBehavior func(s mocks.MockQuerier, task domain.Task, ctx context.Context)
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
			mockBehavior: func(s mocks.MockQuerier, task domain.Task, ctx context.Context) {
				s.EXPECT().Create(ctx, &DB.CreateParams{Text: task.Text}).Return(&domain.Task{ID: task.ID, Text: task.Text, Checked: false}, nil)
			},
			expectedBody: &domain.Task{ID: 0, Text: "Text15", Checked: false},
			expectedErr:  nil,
		},
		{
			name:      "Empty body",
			inputTask: domain.Task{},
			mockBehavior: func(s mocks.MockQuerier, task domain.Task, ctx context.Context) {
				if task.Text == "" {
					s.EXPECT().Create(ctx, &DB.CreateParams{Text: task.Text}).Return(&domain.Task{}, errors.New("error db"))
					return
				}
				s.EXPECT().Create(ctx, task).Return(&domain.Task{ID: task.ID, Text: task.Text, Checked: false}, nil)
			},
			expectedBody: &domain.Task{},
			expectedErr:  errors.New("error db"),
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			task := mocks.NewMockQuerier(ctrl)

			testCase.mockBehavior(*task, *testCase.expectedBody, context.Background())

			w, err := task.Create(context.Background(), &DB.CreateParams{Text: testCase.inputTask.Text})
			assert.Equal(t, testCase.expectedBody, w)
			assert.Equal(t, testCase.expectedErr, err)
		})
	}
}

func Test_TaskDeleteByID(t *testing.T) {
	type mockBehavior func(s mocks.MockQuerier, id int, ctx context.Context)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	testTable := []struct {
		name         string
		inputID      int
		mockBehavior mockBehavior
		expectedErr  error
	}{
		{
			name:    "OK",
			inputID: 3,
			mockBehavior: func(s mocks.MockQuerier, id int, ctx context.Context) {
				s.EXPECT().DeleteByID(ctx, &DB.DeleteByIDParams{ID: int32(id)}).Return(nil, nil)
			},
			expectedErr: nil,
		},
		{
			name:    "Empty body",
			inputID: 3,
			mockBehavior: func(s mocks.MockQuerier, id int, ctx context.Context) {
				s.EXPECT().DeleteByID(ctx, &DB.DeleteByIDParams{ID: int32(id)}).Return(nil, errors.New("error db"))

			},
			expectedErr: errors.New("error db"),
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			task := mocks.NewMockQuerier(ctrl)

			testCase.mockBehavior(*task, testCase.inputID, context.Background())
			_, err := task.DeleteByID(context.Background(), &DB.DeleteByIDParams{ID: int32(testCase.inputID)})
			assert.Equal(t, testCase.expectedErr, err)
		})
	}
}

func Test_TaskDeleteAll(t *testing.T) {
	type mockBehavior func(s mocks.MockQuerier, ctx context.Context)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	testTable := []struct {
		name         string
		mockBehavior mockBehavior
		expectedErr  error
	}{
		{
			name: "OK",
			mockBehavior: func(s mocks.MockQuerier, ctx context.Context) {
				s.EXPECT().DeleteAll(ctx).Return(nil, nil)
			},
			expectedErr: nil,
		},
		{
			name: "Empty body",
			mockBehavior: func(s mocks.MockQuerier, ctx context.Context) {
				s.EXPECT().DeleteAll(ctx).Return(nil, errors.New("error db"))
			},
			expectedErr: errors.New("error db"),
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			task := mocks.NewMockQuerier(ctrl)

			testCase.mockBehavior(*task, context.Background())
			_, err := task.DeleteAll(context.Background())
			assert.Equal(t, testCase.expectedErr, err)
		})
	}
}

func Test_TaskChangeText(t *testing.T) {
	type mockBehavior func(s mocks.MockQuerier, task domain.Task, ctx context.Context)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	testTable := []struct {
		name         string
		inputTask    domain.Task
		mockBehavior mockBehavior
		expectedErr  error
	}{
		{
			name:      "OK",
			inputTask: domain.Task{ID: 1, Text: "Text15"},
			mockBehavior: func(s mocks.MockQuerier, task domain.Task, ctx context.Context) {
				s.EXPECT().ChangeTextByID(ctx, &DB.ChangeTextByIDParams{ID: task.ID, Text: task.Text}).Return(nil, nil)
			},
			expectedErr: nil,
		},
		{
			name:      "Empty body",
			inputTask: domain.Task{},
			mockBehavior: func(s mocks.MockQuerier, task domain.Task, ctx context.Context) {
				s.EXPECT().ChangeTextByID(ctx, &DB.ChangeTextByIDParams{ID: task.ID, Text: task.Text}).Return(nil, errors.New("error db"))
			},
			expectedErr: errors.New("error db"),
		},
	}
	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			task := mocks.NewMockQuerier(ctrl)
			repo := NewDBTX(mocks.NewMockDBTX(ctrl))
			testCase.mockBehavior(*task, testCase.inputTask, context.Background())

			err := repo.ChangeText(context.Background(), int(testCase.inputTask.ID), testCase.inputTask.Text)
			assert.Equal(t, testCase.expectedErr, err)
		})
	}
}
