package main

import (
	"context"
	"net/http"

	"github.com/NixM0nk3y/golang-lambda/pkg/chilogger"
	"github.com/NixM0nk3y/golang-lambda/pkg/log"
	"github.com/NixM0nk3y/golang-lambda/pkg/version"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	chiadapter "github.com/awslabs/aws-lambda-go-api-proxy/chi"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	renderPkg "github.com/unrolled/render"
)

var chiLambda *chiadapter.ChiLambda

var render *renderPkg.Render

func init() {

	logger := log.Logger(context.TODO())

	render = renderPkg.New()

	// stdout and stderr are sent to AWS CloudWatch Logs
	logger.Info("lambda cold start")

	r := chi.NewRouter()

	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(chilogger.Logger())

	r.Route("/v1.0.0", func(r chi.Router) {
		r.Get("/version", getVersion)

		r.Get("/panic", func(w http.ResponseWriter, r *http.Request) {
			panic("test")
		})

	})

	chiLambda = chiadapter.New(r)
}

// VersionResponse is the
type VersionResponse struct {
	Version string `json:"version"`

	BuildHash string `json:"buildhash"`

	BuildDate string `json:"builddate"`
}

func getVersion(w http.ResponseWriter, r *http.Request) {

	render.JSON(w, 200, &VersionResponse{
		Version:   version.Version,
		BuildHash: version.BuildHash,
		BuildDate: version.BuildDate,
	})
}

// Handler is
func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// If no name is provided in the HTTP request body, throw an error
	return chiLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}
