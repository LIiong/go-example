package main

import (
	"github.com/gin-gonic/gin"
	"time"
	"net"
	"errors"
	"fmt"
)


func main_web() {
	r := gin.Default()
	ip, err := externalIP()
	if err != nil {
		fmt.Println(err)
		return
    }
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello GO! 0.2",
			"time":time.Now().Format("2006-01-02 15:04:05"),
			"IP":ip.String(),
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func externalIP() (net.IP, error) {
    ifaces, err := net.Interfaces()
    if err != nil {
        return nil, err
    }
    for _, iface := range ifaces {
        if iface.Flags&net.FlagUp == 0 {
            continue // interface down
        }
        if iface.Flags&net.FlagLoopback != 0 {
            continue // loopback interface
        }
        addrs, err := iface.Addrs()
        if err != nil {
            return nil, err
        }
        for _, addr := range addrs {
            ip := getIpFromAddr(addr)
            if ip == nil {
                continue
            }
            return ip, nil
        }
    }
    return nil, errors.New("connected to the network?")
}

func getIpFromAddr(addr net.Addr) net.IP {
    var ip net.IP
    switch v := addr.(type) {
    case *net.IPNet:
        ip = v.IP
    case *net.IPAddr:
        ip = v.IP
    }
    if ip == nil || ip.IsLoopback() {
        return nil
    }
    ip = ip.To4()
    if ip == nil {
        return nil // not an ipv4 address
    }

    return ip
}
