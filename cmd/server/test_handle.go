package main

import (
	"fmt"
	"s1-chat/pkg/structs"
)

type TestHandle struct {
}

func (t *TestHandle) Processing(msg *structs.Message) bool {
	fmt.Println(msg)
	return true
}
