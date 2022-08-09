package core

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
	"time"

	"github.com/julienschmidt/httprouter"
	"go.uber.org/zap"
)

type httpConf struct {
	Host string
	Port int
}

var wg sync.WaitGroup

func (a *app) initHttp(router *httprouter.Router) {
	srv := &http.Server{
		Addr:         fmt.Sprintf("%v:%v", a.HttpConf.Host, a.HttpConf.Port),
		Handler:      a.recoverPanic(a.rateLimit(router)),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	shutdownError := make(chan error)
	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		s := <-quit
		Logger.Info("process", zap.String("type", "signal"), zap.String("source", "signal"), zap.String("act", s.String()), zap.String("status", "delegated"))
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		err := srv.Shutdown(ctx)
		if err != nil {
			shutdownError <- err
		}
		Logger.Info("process", zap.String("type", "background"), zap.String("source", "httprouter"), zap.String("act", "closing"), zap.String("addr", srv.Addr), zap.String("status", "done"))

		wg.Wait()
		shutdownError <- nil
	}()

	Logger.Info("process", zap.String("type", "server"), zap.String("source", "httprouter"), zap.String("act", "serve"), zap.String("addr", srv.Addr), zap.String("status", "running"))
	err := srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		Logger.Fatal(err.Error())
	}

	err = <-shutdownError
	if err != nil {
		Logger.Fatal(err.Error())
	}

	Logger.Info("process", zap.String("type", "server"), zap.String("source", "httprouter"), zap.String("act", "shutdown"), zap.String("addr", srv.Addr), zap.String("status", "done"))
}

func GetParam(param string, r *http.Request) string {
	return httprouter.ParamsFromContext(r.Context()).ByName(param)
}

func GetIntParam(param string, r *http.Request) (int64, error) {
	result, err := strconv.ParseInt(httprouter.ParamsFromContext(r.Context()).ByName(param), 10, 64)
	if err != nil || result < 1 {
		return 0, errors.New("invalid id parameter")
	}
	return result, nil
}
