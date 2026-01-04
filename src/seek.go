package src

import (
	"net"
	"strconv"
	"strings"
	"sync"
	"time"
)

// Seek 扫描指定网段和端口，返回结果
func (m *SRC) Seek(network, ports string) []interface{} {
	hosts := []string{}
	for i := 2; i < 254; i++ {
		hosts = append(hosts, network+"."+strconv.Itoa(i))
	}
	timeOut := 3 * time.Second
	var wg sync.WaitGroup
	results := make(chan []string, len(hosts))
	for _, host := range hosts {
		wg.Add(1)
		go m.FindHost(host, ports, timeOut, &wg, results)
	}
	wg.Wait()
	close(results)
	var scanResults []interface{}
	for result := range results {
		scanResults = append(scanResults, result)
	}
	return scanResults
}

func (m *SRC) FindHost(host, ports string, timeOut time.Duration, wg *sync.WaitGroup, results chan<- []string) {
	defer wg.Done()
	openPorts := []string{}
	for _, port := range strings.Split(ports, ",") {
		address := net.JoinHostPort(host, port)
		conn, err := net.DialTimeout("tcp", address, timeOut)
		if err != nil {
			continue
		}
		conn.Close()
		openPorts = append(openPorts, port)
	}
	if len(openPorts) > 0 {
		results <- []string{host, strings.Join(openPorts, ","), time.Now().Format("2006-01-02 15:04:05")}
	}

}
