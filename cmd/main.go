package main

import (
	"fmt"
	"time"
	"log"
	"github.com/getsentry/sentry-go"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/chaki8923/wedding-backend/pkg/adapter/http/handler"
	authMiddleware "github.com/chaki8923/wedding-backend/pkg/adapter/http/middleware"
	"github.com/chaki8923/wedding-backend/pkg/adapter/http/route"
	"github.com/chaki8923/wedding-backend/pkg/infra"
	"github.com/chaki8923/wedding-backend/pkg/lib/config"
	initSentry "github.com/chaki8923/wedding-backend/pkg/lib/sentry"
	"github.com/chaki8923/wedding-backend/pkg/usecase"
)

func main() {
	// Config
	c, cErr := config.New()

	// Sentry
	if err := initSentry.SetUp(c); err != nil {
		sentry.CaptureException(fmt.Errorf("initSentry err: %w", err))
	}

	// DB
	db, err := infra.NewDBConnector(c)
	if err != nil {
		sentry.CaptureException(fmt.Errorf("initDb err: %w", err))
	}
	log.Printf("DI手前-------------------------")
	// S3設定
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-northeast-1"),
	})
	if err != nil {
		log.Fatalf("failed to create AWS session: %v", err)
	}
	s3Client := s3.New(sess)


	// DI
	//DBを注入。
	mr := infra.NewMessageRepository(db)
	ir := infra.NewInvitationRepository(db,s3Client, "weddingnet")
	ivr := infra.NewInviteeRepository(db,s3Client, "weddingnet")
	upr := infra.NewUploadRepository(db,s3Client, "weddingnet")
	ur := infra.NewUserRepository(db, c)
	sr := infra.NewSendMailRepository(db)
	// repositoryを注入
	au := usecase.NewAuthUseCase(ur, c)
	su := usecase.NewMailUseCase(sr)
	mu := usecase.NewMsgUseCase(mr)
	uu := usecase.NewUserUseCase(ur)
	iu := usecase.NewIvtUseCase(ir)
	ivu := usecase.NewIvteeUseCase(ivr)
	upu := usecase.NewUploadUseCase(upr)


	ch := handler.NewCsrfHandler()
	mh := handler.NewSendMailHandler(su)
	lh := handler.NewLoginHandler(au)
	sh := handler.NewSignHandler(au)
	gh := handler.NewGraphHandler(mu, uu, iu, ivu, upu)
	ph := playground.Handler("GraphQL", "/query")
	am := authMiddleware.NewAuthMiddleware(au)
	fmt.Printf("IRです%+v\n", ir)
	fmt.Printf("IUです%+v\n", iu)
	// Rooting
	r := route.NewInitRoute(ch, lh, sh, mh, gh, ph, am)
	_, err = r.InitRouting(c)
	if err != nil {
		sentry.CaptureException(fmt.Errorf("InitRouting at NewInitRoute err: %w", err))
	}

	defer func() {
		// .envが存在しない場合
		if cErr != nil {
			sentry.CaptureException(fmt.Errorf("config err: %w", cErr))
		}
		// panic の場合も Sentry に通知する場合は Recover() を呼ぶ
		sentry.Recover()
		// サーバーへは非同期でバッファしつつ送信するため、未送信のものを忘れずに送る(引数はタイムアウト時間)
		sentry.Flush(2 * time.Second)
	}()
}
