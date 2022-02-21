package planets

import (
	"context"
	"encoding/json"
	"errors"
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

func Test_handlers_ListPlanets(t *testing.T) {
	t.Parallel()

	const path = "/api/v1/planets"

	type args struct {
		ctx context.Context
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
			name: "Should list planets successfully",
			args: args{
				ctx: context.Background(),
			},
			fields: fields{
				useCases: &planets.UseCasesMock{
					ListFunc: func(_ context.Context) ([]planets.Planet, error) {
						return []planets.Planet{
							{
								ID:                      `001`,
								Name:                    `test_name`,
								Climate:                 `test_clim`,
								Terrain:                 `test_terr`,
								QuantityFilmAppearances: 200,
							},
						}, nil
					},
				},
			},
			want: want{
				statusCode: http.StatusOK,
				response:   json.RawMessage(`{"message":the planet name cannot be empty}`),
			},
		},
		{
			name: "Should returns internal err",
			args: args{
				ctx: context.Background(),
			},
			fields: fields{
				useCases: &planets.UseCasesMock{
					ListFunc: func(contextMoqParam context.Context) ([]planets.Planet, error) {
						return nil, errors.New(`oh noh`)
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
			req := createRequest(tt.args.ctx, t, http.MethodGet, path, nil)

			router := chi.NewRouter()
			router.Get(path, middlewares.Handle(h.ListPlanets))

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
