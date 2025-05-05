package helpers

import (
	"fmt"
	"io"
)

func RequestBodyLogger (body io.ReadCloser) {
	b, err := io.ReadAll(body)
	if err != nil {
    fmt.Println(err)
  }
	fmt.Println("bodyBytes:", string(b))
}
