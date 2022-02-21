package planets

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
	"github.com/sptGabriel/starwars/app/domain/planets"
	"github.com/sptGabriel/starwars/app/gateway/api/middlewares"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_handlers_DeletePlanet(t *testing.T) {
	t.Parallel()

	const path = "/api/v1/planets/%s"

	type args struct {
		ctx      context.Context
		planetID string
	}

	type fields struct {
		useCases planets.UseCases
	}

	type want struct {
		statusCode int
		response   json.RawMessage
	}
	type test struct {
		name   string
		args   args
		fields fields
		want   want
	}

	tests := []test{
		{
			name: "Should delete an planet successfully",
			args: args{
				ctx:      context.Background(),
				planetID: `planet-001`,
			},
			fields: fields{
				useCases: &planets.UseCasesMock{
					DeleteFunc: func(_ context.Context, _ planets.ID) error {
						return nil
					},
				},
			},
			want: want{
				statusCode: http.StatusOK,
				response:   json.RawMessage(`{"active":true}`),
			},
		},
		{
			name: "Should returns not found err when planet not found",
			args: args{
				ctx:      context.Background(),
				planetID: `planet-001`,
			},
			fields: fields{
				useCases: &planets.UseCasesMock{
					DeleteFunc: func(_ context.Context, _ planets.ID) error {
						return planets.ErrPlanetNotFound
					},
				},
			},
			want: want{
				statusCode: http.StatusNotFound,
				response:   json.RawMessage(`{"message": "the planet was not found"}`),
			},
		},
		{
			name: "Should returns internal err",
			args: args{
				ctx:      context.Background(),
				planetID: `planet-001`,
			},
			fields: fields{
				useCases: &planets.UseCasesMock{
					DeleteFunc: func(_ context.Context, _ planets.ID) error {
						return errors.New(`oh noh`)
					},
				},
			},
			want: want{
				statusCode: http.StatusInternalServerError,
				response:   json.RawMessage(`{"message": "internal server error"}`),
			},
		},
	}

	for _, tt := range tests {
		tt := tt // capture range variable
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			h := NewHandler(tt.fields.useCases)
			req := createRequest(tt.args.ctx, t, http.MethodDelete, fmt.Sprintf(path, tt.args.planetID), nil)

			router := chi.NewRouter()
			router.Delete("/api/v1/planets/{id}", middlewares.Handle(h.DeletePlanet))

			res := httptest.NewRecorder()
			router.ServeHTTP(res, req)

			if res.Code != http.StatusOK {
				got, err := ioutil.ReadAll(res.Body)
				require.NoError(t, err)

				assert.JSONEq(t, string(tt.want.response), string(got))
			}
			assert.Equal(t, tt.want.statusCode, res.Code)
		})
	}
}
