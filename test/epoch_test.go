package test

import (
	"fmt"
	"testing"
	"time"
)

func TestEpoch(t *testing.T) {
	time := time.Now().UnixMilli()
	fmt.Println(time)
}
