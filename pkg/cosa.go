package pkg

import (
	"fmt"

	"github.com/jestebanrods/dummy-golang/internal"
)

func getCosaError() string {
	fmt.Println("call get cosas error")
	return internal.CosaError
}
