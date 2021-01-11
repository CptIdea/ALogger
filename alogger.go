package ALogger

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"
)

//Only go run
func ALogger(port string,filename string) {
	Ls, err := net.Listen("tcp", port)
	if err != nil {
		log.Panic(err)
	}
	for {
		file, err := os.Open(filename)
		if err != nil {
			log.Panic(err)
		}
		conn, err := Ls.Accept()
		if err != nil {
			log.Panic(err)
		}
		ans, err := bufio.NewReader(conn).ReadString('\n')
		if ans == "getall\n" {
			io.Copy(conn, file)
			conn.Close()
			file.Close()
		} else {
			go func(conn net.Conn, file *os.File) {
				ioutil.ReadAll(file)
				for {
					newLog, err := ioutil.ReadAll(file)
					if err != nil {
						log.Panic(err)
					}
					_, err = fmt.Fprint(conn, string(newLog))
					if err != nil {
						conn.Close()
						break
					}
				}
			}(conn, file)
		}
	}
}