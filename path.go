package utils

import (
    "fmt"
)

func validChar(c byte) bool {
    if c >= '0' && c <= '9' {
        return true
    }

    if c >= 'a' && c <= 'z' {
        return true
    }

    if c >= 'A' && c <= 'Z' {
        return true
    }

    return false
}

func GetPath(path string) ([]string, error ) {
    rec := 0
    tmp := ""
    ret := make([]string, 0)

    for ind := 0; ind < len(path); ind++ {
        switch {
        case path[ind] == '.':
            if tmp == "" {
                rec++
            } else {
                tmp += "."
            }
        case path[ind] == '/':
            if tmp == "" && rec != 2 {
                rec = 0
                continue
            }
            switch rec {
            case 0:
                ret = append(ret, tmp)
            case 1:
                tmp = "." + tmp
                ret = append(ret, tmp)
            case 2:
                if len(ret) == 0 {
                    return nil, fmt.Errorf("invalid path(%s)", path)
                }
                fmt.Println("..", ret, ret[:len(ret) - 1])
                ret = ret[:len(ret) - 1]
            default:
                return nil, fmt.Errorf("invalid path(%s)", path)
            }

            rec = 0
            tmp = ""
        case validChar(path[ind]):
            if tmp == "" {
                for rec > 0 {
                    tmp += "."
                    rec--
                }
            }
            tmp += string([]byte{path[ind]})
        default:
            return nil, fmt.Errorf("invalid path(%s)", path)
        }
    }

    if tmp != "" {
        ret = append(ret, tmp)
    }

    return ret, nil
}
