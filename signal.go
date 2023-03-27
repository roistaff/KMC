package main

import (
    "fmt"
    "os"
    "os/signal"
    "syscall"
)

func main() {

    sigs := make(chan os.Signal, 1)
    
    // signal.Notify は、指定されたシグナル通知を受信するために、 与えられたチャネルを登録します。
    signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

    done := make(chan bool, 1)

    go func() {

        sig := <-sigs
        fmt.Println()
        fmt.Println(sig)
        done <- true
    }()
   
   // プログラムはシグナルを受信するまで
   // (前述の done に値を送信するゴルーチンで知らされる) 待機した後、終了します。
    fmt.Println("awaiting signal")
    <-done
    fmt.Println("exiting")
}
