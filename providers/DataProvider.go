package providers

import "os"

type Provider interface {
	get(pathFile string)
}

type FileProvider struct{}
type HttpProvider struct{}

func (file FileProvider) get(pathFile string) {

}

// Select strategy download file url or local file.
func SelectStrategyDownloadFile(urlToFile string) {
	info, err := os.Stat(urlToFile)
	if os.IsNotExist(err) && !info.IsDir() {
		// use http strategy
	}

	// use file strategy
}
