package usecase

import (
	"context"
	"github.com/Batyr2024/todoGolang/domain"
	mocks "github.com/Batyr2024/todoGolang/mocks/handl"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
)

func Test_TaskFindAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	tests := []struct {
		name       string
		want       []*domain.Task
		beforeTest func(ctx context.Context, ctrl *gomock.Controller, res []*domain.Task, err error) *mocks.MockinterfaceUseCase
	}{
		{name: "OK",
			want: []*domain.Task{{ID: 0, Text: "asd", Checked: false}},
			beforeTest: func(ctx context.Context, ctrl *gomock.Controller, res []*domain.Task, err error) *mocks.MockinterfaceUseCase {
				mockService := mocks.NewMockinterfaceUseCase(ctrl)
				mockService.EXPECT().FindAll(ctx).Return(res, err).AnyTimes()
				return mockService
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := New(tt.beforeTest(context.TODO(), ctrl, []*domain.Task{{ID: 0, Text: "asd", Checked: false}}, nil))
			got, _ := uc.FindAll(context.TODO())
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindAll() got = %v, want %v", got, tt.want)
			}
		})
	}
}
