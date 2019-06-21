package main

import (
    "bufio"
    "fmt"
    "math/rand"
    "net"
    "strconv"
    "strings"
    "time"
)

var hosts []string = []string{"10.142.232.190:8003",
                              "10.142.232.182:8000"}

const (
    PROT  = "tcp"
    LOCAL = "10.142.232.190:8003"
)

func send(n int) {
    msg  := fmt.Sprintf("%d", n)
    host := hosts[rand.Intn(len(hosts))]
    fmt.Printf("Enviando %d a %s\n", n, host)
    con, _ := net.Dial(PROT, host)
    defer con.Close()
    fmt.Fprintln(con, msg)
}

func handle(con net.Conn) {
    defer con.Close()
    r := bufio.NewReader(con)
    msg, _ := r.ReadString('\n')
    msg = strings.TrimSpace(msg)
    if n, err := strconv.Atoi(msg); err == nil {
        fmt.Println("Recibido: ", n)
        if n == 0 {
            fmt.Println("Me tocó perder 😞")
        } else {
                send(n - 1)
        }
    }
}

func start() {
    var num int
    for {
        fmt.Scanf("%d\n", &num)
        send(num)
    }
}

func main() {
    rand.Seed(time.Now().UTC().UnixNano())
    go start()
    ln, _ := net.Listen(PROT, LOCAL)
    defer ln.Close()
    for {
        con, _ := ln.Accept()
        go handle(con)
    }
}
