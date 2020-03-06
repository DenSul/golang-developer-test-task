package loader

import (
	"app/provider"
	"app/storage"
	"fmt"
)

var AllowedIndex = []string{"global_id", "id", "mode"}

const Delimiter = ":"

type Loader struct {
	Storage  storage.Storage
	Provider provider.Provider
}

func (l *Loader) Run() {

	fmt.Println(l.Storage.Connect())
	//var data = l.Provider.Load()

	//var mapIndex = make(map[string]map[string]interface{})
	//for k, d := range data {
	//	key := fmt.Sprintf("id%s%d", Delimiter, d.ID)
	//	result, err := json.Marshal(d)
	//	if err != nil {
	//		panic(fmt.Sprintf("unable to marshal data with id %v: %v", d.ID, err))
	//	}
	//
	//
	//	if contains(AllowedIndex, k) {
	//		fmt.Println()
	//	}

	//}

	//fmt.Println(l.Provider.Load())
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
