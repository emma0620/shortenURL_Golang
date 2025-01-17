package services

import "sync"

var (
	urlMap = make(map[string]string)
	mu     sync.RWMutex
)

// GenerateShortURL 生成短網址
func GenerateShortURL(originalURL string) string {
	shortCode := "abc123" // 模擬短網址的生成邏輯
	mu.Lock()
	urlMap[shortCode] = originalURL
	mu.Unlock()
	return "http://localhost:8080/" + shortCode
}

// GetOriginalURL 根據短網址獲取原始網址
func GetOriginalURL(shortCode string) string {
	mu.RLock()
	defer mu.RUnlock()
	return urlMap[shortCode]
}
