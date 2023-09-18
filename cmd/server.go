package cmd

import (
	"context"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"mall/global"
	"mall/initialize"
	"mall/router"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func init() {
	rootCmd.AddCommand(serverCmd)
}

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start server",
	Run: func(cmd *cobra.Command, args []string) {
		runServer()
	},
}

func runServer() {
	initialize.Init()

	engine := gin.New()
	engine.Use(ginzap.Ginzap(global.GVA_LOG, time.RFC3339, true))
	engine.Use(ginzap.RecoveryWithZap(global.GVA_LOG, true))
	router.SetupRouter(engine)

	defer global.GVA_LOG.Sync()

	server := &http.Server{
		Addr:    ":" + strconv.Itoa(global.GVA_CONFIG.App.Port),
		Handler: engine,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			global.GVA_LOG.Fatal("Server start error:", zap.Error(err))
		}
	}()

	// 服务优雅退出
	quic := make(chan os.Signal)
	signal.Notify(quic, syscall.SIGINT, syscall.SIGTERM)
	<-quic
	global.GVA_LOG.Info("Shutdown Server ...")

	// 创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	if err := server.Shutdown(ctx); err != nil {
		global.GVA_LOG.Fatal("Server Shutdown:", zap.Error(err))
	}
}
