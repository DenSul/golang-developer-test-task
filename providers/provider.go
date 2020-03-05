package providers

import (
	"app/models"
	"app/redis"
	"encoding/json"
	"log"
	"os"
)

type Provider interface {
	save(pathFile string)
}

type FileProvider struct{}
type HttpProvider struct{}

func (provider FileProvider) save(pathFile string) {
	var counter int = 0
	file, _ := os.Open(pathFile)
	defer file.Close()
	dec := json.NewDecoder(file)

	// read open bracket
	_, err := dec.Token()
	if err != nil {
		log.Fatal(err)
	}

	// while the array contains values
	for dec.More() {
		var m models.ParkingTaxi
		// decode an array value (ParkingTaxi)
		err := dec.Decode(&m)
		if err != nil {
			log.Fatal(err)
		}

		client := redis.GetDB()
		pipeline := client.Pipeline()
		//pipeline.s

		counter++
	}

	// read closing bracket
	_, err = dec.Token()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Count imported items: %d", counter)
}

// Select strategy download file url or local file.
func SelectStrategyDownloadFile(urlToFile string) {

	//info, err := os.Stat(urlToFile)
	//
	//if os.IsNotExist(err) && !info.IsDir() {
	//	// use http strategy
	//}

	var provider = new(FileProvider)
	provider.save("./resources/data/data.json")
}
