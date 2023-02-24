package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sugimoto-ne/go_sample_app.git/entity"
	"github.com/sugimoto-ne/go_sample_app.git/testutil"
	"github.com/sugimoto-ne/go_sample_app.git/utils"
)

func TestLoginUser(t *testing.T) {
	t.Parallel()

	type want struct {
		status   int
		rspFiles []string
	}

	tests := map[string]struct {
		reqFiles []string
		want     want
	}{
		"ok": {
			reqFiles: []string{
				"testdata/login_user/ok_req.json.golden",
			},
			want: want{
				status: http.StatusOK,
				rspFiles: []string{
					"testdata/login_user/ok_rsp.json.golden",
				},
			},
		},
		"invalidPassword": {
			reqFiles: []string{
				"testdata/login_user/invalid_password_length_reqs.json.golden",
				"testdata/login_user/invalid_password_upper_reqs.json.golden",
				"testdata/login_user/invalid_password_number_reqs.json.golden",
				"testdata/login_user/invalid_password_symbol_reqs.json.golden",
			},
			want: want{
				status: http.StatusBadRequest,
				rspFiles: []string{
					"testdata/login_user/invalid_password_length_rsp.json.golden",
					"testdata/login_user/invalid_password_upper_rsp.json.golden",
					"testdata/login_user/invalid_password_number_rsp.json.golden",
					"testdata/login_user/invalid_password_symbol_rsp.json.golden",
				},
			},
		},
	}

	customValidator := utils.CustomValidator()
	for n, tt := range tests {
		tt := tt
		n := n
		wantStatus := tt.want.status
		for n2, tt2 := range tt.reqFiles {
			tt2 := tt2
			n2 := n2
			file := testutil.LoadFile(t, tt2)

			usersData := make([]*entity.User, 0)

			err := json.Unmarshal(file, &usersData)
			if err != nil {
				t.Fatalf("cannnot unmarchal json data")
			}

			for uIndex, u := range usersData {
				u := u
				index := fmt.Sprintf("%s-%d-%d", n, n2, uIndex)
				data := &struct {
					Name     string
					Password string
					Role     string
				}{
					Name:     u.Name,
					Password: u.Password,
					Role:     string(u.Role),
				}
				requestMockBody, err := json.Marshal(&data)
				if err != nil {
					t.Fatalf("cannnot marchal json data")
				}

				t.Run(
					index,
					func(t *testing.T) {
						// t.Parallel()

						w := httptest.NewRecorder()
						r := httptest.NewRequest(
							http.MethodPost,
							"/login",
							bytes.NewReader(requestMockBody),
						)

						moq := &LoginServiceMock{}
						moq.LoginFunc = func(ctx context.Context, name, password string) (string, error) {
							if wantStatus == http.StatusOK {
								return "hogehogefugafuga", nil
							}

							return "", errors.New("error from mock")
						}

						sut := Login{
							Service:   moq,
							Validator: customValidator,
						}

						sut.ServeHTTP(w, r)
						println("###########")
						resp := w.Result()

						testutil.AssertResponse(t, resp, wantStatus, testutil.LoadFile(t, tt.want.rspFiles[n2]))
					},
				)
			}
		}
	}
}
