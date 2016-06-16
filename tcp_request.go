package utils

import (
	"io"
	"fmt"
	"net"
	"time"
	"bytes"
	"strconv"
	"strings"
)

func readDone(str string) bool {
    if len(str) < 2 || str[len(str) - 1] != '\n' {
		return false
	}

    /**
     * TODO for ssdb
     * Request:  3\nget3\nkey\n\n
     * Response: 2\nok\n3\nval\n\n
     **/
    if len(str) >2 && str[len(str) - 1] == '\n' && str[len(str) - 2] == '\n' {
        return true
    }

	switch str[0] {
	case '+':
        return true
	case '-':
        return true
	case ':':
		return true
	case '$':
		/**
		 * $3\r\nabc\r\n
		 **/
        if str == "$-1" { //$-1
            return true
        }
		line := strings.Split(str, "\r\n")
		if len(line) < 2 {
			return false
		} else {
			length, _ := strconv.Atoi(line[0][1:])
			strLen := len(str)
			tmp := length
			for {
				if tmp == 0 {
					break
				}
				length++  //$1024\r\nblablabla\r\n => ++ is for 1024's length
				tmp = tmp / 10
			}
			length = length + 5 //5 is for $\r\n\r\n
			if strLen >= length {
				return true
			} else {
				return false
			}
		}
	case '*':
		/**
		 * *2\r\n$1\r\na\r\n$1\r\nb\r\n
		 **/
		length := 0
		for ind := 1; ind < len(str) && str[ind] != '\r'; ind++ {
			if str[ind] < '0' || str[ind] > '9' {
				fmt.Println("协议解析错误\n", str)
				return false
			}
			length *= 10
			length += int(str[ind] - '0')
		}

        for _, substr := range strings.Split(str, "\r\n") {
            if substr != "" && substr[0] == '$' {
                length--
            }
        }
		if length != 0 {
			return false
		} else {
			return true
		}
	default:
		return true
	}
	return false
}

//tcp请求后端并返回结果
func callCmd(addr string, request []byte) ([]string, error) {
	//var (
	//	retry  int = 0
	//)

	//连接后端
	conn, err := net.DialTimeout("tcp", addr, 3 * time.Second)
	//for err != nil && retry < 1 { //重试
	//	retry++
	//	conn, err = net.DialTimeout("tcp", addr, 3 * time.Second)
	//}
	if err != nil {
		fmt.Println("Connect faild: ", err)
		return nil, err
	}
	defer conn.Close()
    err = conn.SetDeadline(time.Now().Add(10 * time.Second))
    if err != nil {
        fmt.Println("Set deadline failed: ", err)
        return nil, err
    }

	//向后端发请求
	_, err = conn.Write(request)
	if err != nil {
		fmt.Println("Request failed: ", err)
		return nil, err
	}

	var buf [512]byte
	result := bytes.NewBuffer(nil)

	//读取后端数据
	for {
		n, err := conn.Read(buf[0:])
		result.Write(buf[0:n])
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("Read connection failed: ", err)
			return nil, err
		}

		//解析协议 TODO
		if readDone(string(result.Bytes())) {
			break
		}
	}

	//分割\r\n
	resultFilter := strings.Split(string(result.Bytes()), "\n")
	for i, _ := range resultFilter {
		resultFilter[i] = strings.Split(resultFilter[i], "\r")[0]
	}
	return resultFilter, nil
}
