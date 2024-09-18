package usecase

import (
	"context"
	"fmt"
	"github.com/Batyr2024/todoGolang/domain"
	"github.com/Batyr2024/todoGolang/mocks"
	"github.com/golang/mock/gomock"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func Test_TaskFindAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name           string
		args           args
		beforeTest     func(findAllUseCase *mocks.MockinterfaceUseCase)
		wantStatusCode int
		wantBody       string
	}{
		{
			name: "OK",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(
					http.MethodPost,
					"http://localhost:3000/tasks", nil,
				),
			},
			beforeTest: func(findAllUseCase *mocks.MockinterfaceUseCase) {
				findAllUseCase.EXPECT().FindAll(context.TODO()).
					Return(
						[]*domain.Task{{
							ID:      0,
							Text:    "sdgsdgsd",
							Checked: false,
						}},
						nil,
					)
			},
			wantStatusCode: http.StatusOK,
			wantBody:       `[{"id":"0","text":"sdgsdgsd","checked":"false"}]`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRegisterTaskUC := mocks.NewMockinterfaceUseCase(ctrl)

			u := &UseCase{
				taskRepo: mockRegisterTaskUC,
			}

			if tt.beforeTest != nil {
				tt.beforeTest(mockRegisterTaskUC)
			}

			_, err := u.FindAll(context.TODO())
			fmt.Println(err, tt.wantStatusCode)

			rec := tt.args.w.(*httptest.ResponseRecorder)
			res := rec.Result()

			if !reflect.DeepEqual(res.StatusCode, tt.wantStatusCode) {
				t.Errorf("taskUseCase.FindAll() = %v, want %v", res.StatusCode, tt.wantStatusCode)
			}

			//bodyBuffer := new(bytes.Buffer)
			//n, err := bodyBuffer.ReadFrom(res.Body)
			bodyBuffer := make([]byte, rec.Body.Len())
			n, err := res.Body.Read(bodyBuffer)
			fmt.Println("convert", n, err)

			//body := strings.TrimSpace(bodyBuffer.String())

			if !reflect.DeepEqual(bodyBuffer, tt.wantBody) {
				fmt.Println("body", res.Body)
				t.Errorf("taskUseCase.FindAll() = %s, want %s", bodyBuffer, tt.wantBody)
			}
		})
	}
}
