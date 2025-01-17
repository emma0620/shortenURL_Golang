

---

# Short URL Generator by Golang

這是一個簡單的短網址生成器，允許用戶輸入一個長網址並生成對應的短網址。

## 啟動伺服器

執行以下指令來啟動伺服器：

```bash
go run cmd/main.go
```

## 測試功能

### 1. 測試首頁

瀏覽器訪問：

```bash
http://localhost:8080/
```

應返回：

```css
Welcome to the URL shortener!
```

### 2. 測試生成短網址

發送 POST 請求：

```bash
curl -X POST -H "Content-Type: application/json" -d '{"url": "https://example.com"}' http://localhost:8080/shorten
```

應返回：

```json
{
  "short_url": "http://localhost:8080/abc123"
}
```

### 3. 測試短網址重定向

瀏覽器訪問：

```bash
http://localhost:8080/abc123
```

應重定向到：

```arduino
https://example.com
```

---




```markdown
shortenURL_Golang/
├── cmd/                # 主程序入口
│   └── main.go         # 主入口文件
├── internal/           # 後端核心邏輯（業務代碼）
│   ├── handlers/       # HTTP 處理器（路由和控制器）
│   │   └── shorten.go  # 短網址處理器
│   ├── services/       # 業務邏輯服務
│   │   └── shorten.go  # 短網址生成業務邏輯
│   ├── models/         # 數據模型層
│   │   └── shorten.go  # 短網址模型
│   └── utils/          # 工具函數
│       └── validator.go# 工具方法
├── web/                # 前端資源
│   ├── static/         # 靜態資源（CSS/JS）
│   │   ├── app.js      # 前端邏輯
│   │   └── style.css   # 樣式文件
│   ├── templates/      # 模板文件（HTML）
│   │   └── index.html  # 主頁模板
├── config/             # 配置文件
│   └── config.go       # 配置信息
├── pkg/                # 可重用的庫和工具
│   └── logger/         # 日誌功能
│       └── logger.go   # 日誌初始化
├── go.mod              # Go 模塊文件
└── README.md           # 專案文檔
```


