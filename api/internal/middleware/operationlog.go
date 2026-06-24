package middleware

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/DemoLiang/hridoc/api/internal/svc"
	"github.com/DemoLiang/hridoc/api/model"
	"github.com/zeromicro/go-zero/core/logx"
)

var moduleMap = map[string]string{
	"/api/user/add":        "用户管理",
	"/api/user/update":     "用户管理",
	"/api/user/delete":     "用户管理",
	"/api/category/add":    "证件类型",
	"/api/category/update": "证件类型",
	"/api/category/delete": "证件类型",
	"/api/cert/add":        "证件管理",
	"/api/cert/update":     "证件管理",
	"/api/cert/delete":     "证件管理",
	"/api/export":          "导出任务",
	"/api/clean":           "导出任务",
	"/api/excel/import":    "导入任务",
	"/api/upload":          "文件上传",
}

var actionMap = map[string]string{
	"POST":   "新增/执行",
	"GET":    "查询",
	"PUT":    "更新",
	"DELETE": "删除",
}

func OperationLog(svcCtx *svc.ServiceContext) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			var bodyBytes []byte
			if r.Body != nil {
				bodyBytes, _ = io.ReadAll(r.Body)
				r.Body = io.NopCloser(bytes.NewReader(bodyBytes))
			}

			next(w, r)

			module := moduleMap[r.URL.Path]
			if module == "" {
				return
			}

			ctx := r.Context()
			userId := GetUserId(ctx)
			username := GetUsername(ctx)
			if userId == 0 {
				return
			}

			action := actionMap[r.Method]
			if action == "" {
				action = r.Method
			}

			target := r.URL.Path
			detail := extractDetail(r, bodyBytes)

			go func() {
				logCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
				defer cancel()

				log := &model.OperationLog{
					OperatorId:   userId,
					OperatorName: username,
					Module:       module,
					Action:       action,
					Target:       sql.NullString{String: target, Valid: true},
					Detail:       sql.NullString{String: detail, Valid: detail != ""},
					Ip:           sql.NullString{String: getClientIP(r), Valid: true},
					CreatedAt:    time.Now(),
				}
				_, err := svcCtx.OperationLogModel.Insert(logCtx, log)
				if err != nil {
					logx.Errorf("operation log insert failed: %v", err)
				}
			}()
		}
	}
}

func getClientIP(r *http.Request) string {
	ip := r.Header.Get("X-Forwarded-For")
	if ip != "" {
		parts := strings.Split(ip, ",")
		if len(parts) > 0 {
			return strings.TrimSpace(parts[0])
		}
	}
	ip = r.Header.Get("X-Real-Ip")
	if ip != "" {
		return ip
	}
	return r.RemoteAddr
}

func extractDetail(r *http.Request, bodyBytes []byte) string {
	if r.Method == http.MethodGet {
		return r.URL.RawQuery
	}
	if (r.Method == http.MethodPost || r.Method == http.MethodPut) && len(bodyBytes) > 0 {
		ct := r.Header.Get("Content-Type")
		if strings.Contains(ct, "application/json") {
			var body map[string]any
			if err := json.Unmarshal(bodyBytes, &body); err == nil {
				if _, ok := body["password"]; ok {
					body["password"] = "***"
				}
				b, _ := json.Marshal(body)
				s := string(b)
				if len(s) > 1024 {
					s = s[:1024] + "..."
				}
				return s
			}
		}
	}
	return ""
}
