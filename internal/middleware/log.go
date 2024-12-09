package middleware

import (
	"bytes"
	"crmeb_go/pkg/logs"
	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io"
	"time"
)

type LogM struct {
}

func NewLogM() *LogM {
	return &LogM{}
}

func (m *LogM) RequestLogMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// The configuration is initialized once per request
		trace := uuid.NewString()
		logs.Log.WithValue(ctx, zap.String("trace", trace))
		logs.Log.WithValue(ctx, zap.String("request_method", ctx.Request.Method))
		logs.Log.WithValue(ctx, zap.Any("request_headers", ctx.Request.Header))
		logs.Log.WithValue(ctx, zap.String("request_url", ctx.Request.URL.String()))
		if ctx.Request.Body != nil {
			bodyBytes, _ := ctx.GetRawData()
			ctx.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes)) // 关键点
			logs.Log.WithValue(ctx, zap.String("request_params", string(bodyBytes)))
		}
		logs.Log.WithContext(ctx).Info("Request")
		ctx.Next()
	}
}
func (m *LogM) ResponseLogMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: ctx.Writer}
		ctx.Writer = blw
		startTime := time.Now()
		ctx.Next()
		duration := time.Since(startTime).String()
		logs.Log.WithContext(ctx).Info("Response", zap.Any("response_body", blw.body.String()), zap.Any("time", duration))
	}
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
