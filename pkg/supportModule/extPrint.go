package supportModule

import (
	"fmt"
	"time"
)

func PrintlnWithTimeShtamp(text string) {
	fmt.Println(time.Now().Format(time.RFC3339), " : ", text)
}
