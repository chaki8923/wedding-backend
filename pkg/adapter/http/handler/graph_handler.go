package handler

import (
	"context"
	"log"
	"encoding/json"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"github.com/chaki8923/wedding-backend/pkg/adapter/http/resolver"
	"github.com/chaki8923/wedding-backend/pkg/lib/graph/generated"
	"github.com/chaki8923/wedding-backend/pkg/lib/graph/loader"
	"github.com/chaki8923/wedding-backend/pkg/usecase"
	"github.com/graph-gophers/dataloader"
	"github.com/labstack/echo/v4"
)

type Graph interface {
	QueryHandler() echo.HandlerFunc
}

type GraphHandler struct {
	MsgUseCase  usecase.Message
	UserUseCase usecase.User
	IvtUseCase  usecase.Invitation
	IvteeUseCase  usecase.Invitee
	UpdUseCase  usecase.Upload
	AgyUseCase  usecase.Allergy
}

func NewGraphHandler(mu usecase.Message, uc usecase.User, iu usecase.Invitation, ivu usecase.Invitee, upu usecase.Upload, agy usecase.Allergy) Graph {
	GraphHandler := GraphHandler{
		MsgUseCase:  mu,
		UserUseCase: uc,
		IvtUseCase:  iu,
		IvteeUseCase: ivu,
		UpdUseCase: upu,
		AgyUseCase: agy,
	}
	return &GraphHandler
}

func (g *GraphHandler) QueryHandler() echo.HandlerFunc {
	ldr := &loader.Loaders{
		UserLoader: dataloader.NewBatchedLoader(
			g.UserUseCase.BatchGetUsers,
			dataloader.WithCache(&dataloader.NoCache{}),
		),
	}

	rslvr := resolver.Resolver{
		MsgUseCase:  g.MsgUseCase,
		UserUseCase: g.UserUseCase,
		IvtUseCase:  g.IvtUseCase,
		IvteeUseCase: g.IvteeUseCase,
		UpdUseCase: g.UpdUseCase,
		AgyUseCase: g.AgyUseCase,
	}

	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(generated.Config{Resolvers: &rslvr}),
	)

	// Set up error presenter
	srv.SetErrorPresenter(func(ctx context.Context, e error) *gqlerror.Error {
		// Log the error
		log.Printf("GraphQL Error: %v", e)
		// Extract the request from the context
		reqCtx := graphql.GetOperationContext(ctx)
		variablesJSON, _ := json.Marshal(reqCtx.Variables)
		log.Printf("Input Data: %s", variablesJSON)
		
		// Customize the error message
		gqlErr := graphql.DefaultErrorPresenter(ctx, e)
		gqlErr.Message = "Internal server error"
		return gqlErr
	})

	// Log request variables
	srv.AroundOperations(func(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
		return next(ctx)
	})
	
	return func(c echo.Context) error {
		loader.Middleware(ldr, srv).ServeHTTP(c.Response(), c.Request())
		return nil
	}
}
