package main

import (
	"fmt"
	"strings"
)

var str = "xdfdfdsfsdfsdfdsfdsfdferfqrwdfdsfdsfsdfdsfdfsdfdsdffdsdsfdfsdfdsfxcvfdsfdsfdsfdsfdvvcxvcv"

func join() string {
	return fmt.Sprintf("go=%s", str)
}

func join1() string {
	return "go=" + str
}

func join2() string {
	return strings.Join([]string{"go=", str}, "")
}

func join3() string {
	b := strings.Builder{}
	b.WriteString("go=")
	b.WriteString(str)
	return b.String()
}
