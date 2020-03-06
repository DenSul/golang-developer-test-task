package loader

import (
	"app/providers"
	"app/storage"
	"encoding/json"
	"fmt"
)

var AllowedIndex = []string{"global_id", "id", "mode"}

const Delimiter = ":"

type Loader struct {
	Storage  storage.Storage
	Provider providers.Provider
}

func (l *Loader) Run() {

	var data = l.Provider.Load()
	for _, d := range data {
		key := fmt.Sprintf("id%s%d", Delimiter, d.ID)
		result, err := json.Marshal(d)
		if err != nil {
			panic(fmt.Sprintf("unable to marshal data with id %v: %v", d.ID, err))
		}

		fmt.Println(key, result)
	}

	//fmt.Println(l.Provider.Load())
}
