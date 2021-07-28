package main

import "strings"

func isNumber(s string) bool {
    s = strings.Trim(s," ")
    if len(s) == 0 {
        return false
    }
    if s[0] == '+' || s[0] == '-' {
        s = s[1:]
    }
    if len(s) == 0 {
        return false
    }
    dot,e := 0,0
    for i := 0 ; i < len(s) ; i++ {
        b := s[i]
        if b >= '0' && b <= '9' {
            continue
        }else if b == '.' {
            dot++
            // e*.
            if dot > 1 || e > 0 {
                return false
            }
            if i == 0 && i == len(s)-1 {
                // 只有一个.
                return false
            }
            // ".e"
            if i == 0 && (s[i+1] == 'e' || s[i+1] == 'E') {
                return false
            }
        }else if b == 'e' || b == 'E' {
            e++
            if e > 1 {
                return false
            }
            // 只有一个e
            if i == 0 || i == len(s)-1 {
                return false
            }
            if s[i+1] == '+' || s[i+1] == '-' {
                i++
                // e+ 后面没了
                if i == len(s)-1 {
                    return false
                }
            }
        }else{
            return false
        }
    }
    return true
}