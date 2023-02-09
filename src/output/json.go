package output

import (
	"encoding/json"
	"fmt"
	"log"
)

func AsJson(value any) {
	jsonResult, err := json.Marshal(value)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(string(jsonResult))
}
