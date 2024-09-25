package usecase

import (
	"context"
	"fmt"
	"github.com/Batyr2024/todoGolang/domain"
	mocks "github.com/Batyr2024/todoGolang/mocks"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
)

func Test_TaskCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	tests := []struct {
		name       string
		want       *domain.Task
		beforeTest func(ctx context.Context, ctrl *gomock.Controller, req domain.Task) *mocks.MockinterfaceUseCase
	}{
		{name: "OK",
			want: &domain.Task{ID: 0, Text: "asd", Checked: false},
			beforeTest: func(ctx context.Context, ctrl *gomock.Controller, req domain.Task) *mocks.MockinterfaceUseCase {
				mockService := mocks.NewMockinterfaceUseCase(ctrl)
				mockService.EXPECT().Create(ctx, req).Return(&domain.Task{ID: 0, Text: "asd", Checked: false}, error(nil)).AnyTimes()
				return mockService
			},
		},
		{name: "EmptyText",
			want: &domain.Task{},
			beforeTest: func(ctx context.Context, ctrl *gomock.Controller, req domain.Task) *mocks.MockinterfaceUseCase {
				mockService := mocks.NewMockinterfaceUseCase(ctrl)
				mockService.EXPECT().Create(ctx, req).Return(&domain.Task{}, error(nil)).AnyTimes()
				return mockService
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := New(tt.beforeTest(context.TODO(), ctrl, domain.Task{}))
			got, err := uc.Create(context.TODO(), domain.Task{})
			fmt.Println(err)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() got = %v, want %v", got, tt.want)
			}
		})
	}
}

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

func Test_TaskChangeCheckedByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	tests := []struct {
		name       string
		want       error
		beforeTest func(ctx context.Context, ctrl *gomock.Controller, id int, checked bool) *mocks.MockinterfaceUseCase
	}{
		{name: "OK",
			want: nil,
			beforeTest: func(ctx context.Context, ctrl *gomock.Controller, id int, checked bool) *mocks.MockinterfaceUseCase {
				mockService := mocks.NewMockinterfaceUseCase(ctrl)
				mockService.EXPECT().ChangeCheckedByID(ctx, id, checked).Return(nil).AnyTimes()
				return mockService
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := New(tt.beforeTest(context.TODO(), ctrl, 0, true))
			got := uc.ChangeCheckedByID(context.TODO(), 0, true)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindAll() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_TaskChangeCheckedAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	tests := []struct {
		name       string
		want       error
		beforeTest func(ctx context.Context, ctrl *gomock.Controller, checked bool) *mocks.MockinterfaceUseCase
	}{
		{name: "OK",
			want: nil,
			beforeTest: func(ctx context.Context, ctrl *gomock.Controller, checked bool) *mocks.MockinterfaceUseCase {
				mockService := mocks.NewMockinterfaceUseCase(ctrl)
				mockService.EXPECT().ChangeCheckedAll(ctx, checked).Return(nil).AnyTimes()
				return mockService
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := New(tt.beforeTest(context.TODO(), ctrl, false))
			got := uc.ChangeCheckedAll(context.TODO(), false)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindAll() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_TaskDeleteByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	tests := []struct {
		name       string
		want       error
		beforeTest func(ctx context.Context, ctrl *gomock.Controller, id int) *mocks.MockinterfaceUseCase
	}{
		{name: "OK",
			want: nil,
			beforeTest: func(ctx context.Context, ctrl *gomock.Controller, id int) *mocks.MockinterfaceUseCase {
				mockService := mocks.NewMockinterfaceUseCase(ctrl)
				mockService.EXPECT().DeleteByID(ctx, id).Return(nil).AnyTimes()
				return mockService
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := New(tt.beforeTest(context.TODO(), ctrl, 0))
			got := uc.DeleteByID(context.TODO(), 0)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindAll() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_TaskdeleteAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	tests := []struct {
		name       string
		want       error
		beforeTest func(ctx context.Context, ctrl *gomock.Controller) *mocks.MockinterfaceUseCase
	}{
		{name: "OK",
			want: nil,
			beforeTest: func(ctx context.Context, ctrl *gomock.Controller) *mocks.MockinterfaceUseCase {
				mockService := mocks.NewMockinterfaceUseCase(ctrl)
				mockService.EXPECT().DeleteAll(ctx).Return(nil).AnyTimes()
				return mockService
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := New(tt.beforeTest(context.TODO(), ctrl))
			got := uc.DeleteAll(context.TODO())
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindAll() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_TaskChangeText(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	tests := []struct {
		name       string
		want       error
		beforeTest func(ctx context.Context, ctrl *gomock.Controller, id int, text string) *mocks.MockinterfaceUseCase
	}{
		{name: "OK",
			want: nil,
			beforeTest: func(ctx context.Context, ctrl *gomock.Controller, id int, text string) *mocks.MockinterfaceUseCase {
				mockService := mocks.NewMockinterfaceUseCase(ctrl)
				mockService.EXPECT().ChangeText(ctx, id, text).Return(nil).AnyTimes()
				return mockService
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := New(tt.beforeTest(context.TODO(), ctrl, 0, "asd"))
			got := uc.ChangeText(context.TODO(), 0, "asd")
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindAll() got = %v, want %v", got, tt.want)
			}
		})
	}
}
