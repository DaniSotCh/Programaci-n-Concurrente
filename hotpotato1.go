package main

import (
    "bufio"
    "fmt"
    "net"
    "strconv"
    "strings"
    "time"
)

func server(ch, end chan int) { // de pie
    ln, _ := net.Listen("tcp", "MYIP:8000")
    defer ln.Close()
    con, _ := ln.Accept()
    defer con.Close()
    r := bufio.NewReader(con)
    for {
        msg, _ := r.ReadString('\n')
        msg = strings.TrimSpace(msg)
        fmt.Println("Recibido: ", msg)
        if n, err := strconv.Atoi(msg); err == nil {
            if n == 0 {
                fmt.Println("Recáspita, perdí! 😔")
            } else {
                time.Sleep(time.Second)
            }
            ch<- n - 1
            if n < 0 {
                break
            }
        }
    }
    end<- 0
}
func client(ch chan int) {
    var con net.Conn
    created := false
    for {
        n := <-ch
        if !created {
            created = true
            con, _ = net.Dial("tcp", "NEXTIP:8000")
        }
        fmt.Println("Enviando: ", n)
        fmt.Fprintf(con, "%d\n", n)
        if n < -1 {
            fmt.Println("Uf! no perdí!")
            break
        }
    }
    con.Close()
}
func start(ch chan int) {
    var n int
    fmt.Scanf("%d\n", &n)
    ch<- n
}

func main() {
    ch := make(chan int)
    end := make(chan int)
    go server(ch, end)
    go client(ch)
    go start(ch)

    <-end
}
