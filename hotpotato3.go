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

var hosts []string = []string{"10.142.232.189:8001",
                            "10.142.232.190:8003"}
var idPropio string
var idAnterior string
var lib []string
const (
    PROT  = "tcp"
    LOCAL = "10.142.232.189:8001"
)

func send(n string, lib []string) {
    msg  := fmt.Sprintf("%d", n)
    host := hosts[rand.Intn(len(hosts))]
    fmt.Printf("Enviando %d a %s\n", n, host)
    con, _ := net.Dial(PROT, host)
    defer con.Close()
    fmt.Fprintln(con, msg)
}
func start() {
    var num string
    for {
        fmt.Scanf("%d\n", &num)
        send(num,lib)
    }
}


func handle(con net.Conn) {
    defer con.Close()
    r := bufio.NewReader(con)
    msg, _ := r.ReadString('\n')
    msg = strings.TrimSpace(msg)
    if n, err := strconv.Atoi(msg); err == nil {
        fmt.Println("Recibido: ", n)
        if n == 0 {
            fmt.Println(n)
        } else {
                send(n,lib)
        }
    }
}


func main() {
    idPropio = "10.142.232.189:8001"
    idAnterior = "10.142.232.189:8001"
    rand.Seed(time.Now().UTC().UnixNano())
    go send(idPropio, lib)
    ln, _ := net.Listen(PROT, LOCAL)
    defer ln.Close()
    for {
        con, _ := ln.Accept()
        go handle(con)
    }
}
