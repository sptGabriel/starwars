package planets

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
	"github.com/sptGabriel/starwars/app/domain/planets"
	"github.com/sptGabriel/starwars/app/gateway/api/middlewares"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_handlers_CreatePlanet(t *testing.T) {
	t.Parallel()

	const path = "/api/v1/planets"

	type args struct {
		ctx  context.Context
		body json.RawMessage
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
			name: "Should create an planet successfully",
			args: args{
				ctx: context.Background(),
				body: json.RawMessage(`{
					"name": 	"testName",
					"climate": 	"testClim",
					"terrain":	"testTerr"
				}`),
			},
			fields: fields{
				useCases: &planets.UseCasesMock{
					CreateFunc: func(contextMoqParam context.Context, planet planets.Planet) error {
						return nil
					},
				},
			},
			want: want{
				statusCode: http.StatusCreated,
				response:   json.RawMessage(`{"active":true}`),
			},
		},
		// {
		// 	name: "Should return bad request when request body is invalid",
		// 	args: args{
		// 		ctx:  context.Background(),
		// 		body: json.RawMessage(`wrong`),
		// 	},
		// 	fields: fields{},
		// 	want: want{
		// 		statusCode: http.StatusBadRequest,
		// 		response:   json.RawMessage(`{"active":true}`),
		// 	},
		// },
		// {
		// 	name: "Should return unprocessable entity when request body does not have a valid semantics",
		// 	args: args{
		// 		ctx: context.Background(),
		// 		body: json.RawMessage(`{
		// 			"name": 	"4",
		// 			"climate": 	"3",
		// 			"terrain":	"2"
		// 		}`),
		// 	},
		// 	fields: fields{},
		// 	want: want{
		// 		statusCode: http.StatusUnprocessableEntity,
		// 		response:   json.RawMessage(`{"active":true}`),
		// 	},
		// },
	}

	for _, tt := range tests {
		tt := tt // capture range variable
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			h := NewHandler(tt.fields.useCases)
			req := createRequest(tt.args.ctx, t, http.MethodPost, path, tt.args.body)

			router := chi.NewRouter()
			router.Post(path, middlewares.Handle(h.CreatePlanet))

			res := httptest.NewRecorder()
			router.ServeHTTP(res, req)

			assert.Equal(t, tt.want.statusCode, res.Code)
		})
	}
}

func createRequest(ctx context.Context, t *testing.T, method, path string, body json.RawMessage) *http.Request {
	t.Helper()

	reader := bytes.NewReader(body)
	request, err := http.NewRequestWithContext(ctx, method, path, reader)
	require.NoError(t, err)

	return request
}
