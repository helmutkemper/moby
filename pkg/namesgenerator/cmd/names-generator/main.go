package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/helmutkemper/moby/pkg/namesgenerator"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(namesgenerator.GetRandomName(0))
}
