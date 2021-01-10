package accountmgr

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"google.golang.org/grpc"

	"github.com/hzhyvinskyi/BcH/accountMgr"
	pb "github.com/hzhyvinskyi/BcH/common/pb/generated"
)

func main() {
	grpcPort := ":50051"

	var logger log.Logger
	logger = log.NewJSONLogger(os.Stdout)
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
	logger = log.With(logger, "caller", log.DefaultCaller)

	accountMgrService := accountMgr.NewService(logger)
	accountMgrEndpoint := accountMgr.MakeEndpoints(accountMgrService)
	grpcServer := accountMgr.NewGRPCServer(accountMgrEndpoint, logger)

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGALRM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	grpcListener, err := net.Listen("tcp", grpcPort)
	if err != nil {
		logger.Log("during", "Listen", "err", err)
		os.Exit(1)
	}

	go func() {
		gRPCServer := grpc.NewServer()
		pb.RegisterAccountMgrServiceServer(gRPCServer, grpcServer)
		level.Info(logger).Log("msg", "Server started successfully 🚀")
		gRPCServer.Serve(grpcListener)
	}()

	level.Error(logger).Log("exit", <-errs)
}
