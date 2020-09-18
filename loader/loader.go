package loader

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/nivida/eth-rpc-tester/job"
)

// TODO: Move Job struct type to loader

// Loader struct
type Loader struct {
	path string
}

// New initiates a new Loader
func New(p string) (loader *Loader) {
	loader = new(Loader)
	loader.path = p

	return loader
}

// GetTasks returns all the defined tests from a JSON file
func (l *Loader) GetTasks() []job.Job {
	path, _ := filepath.Abs("assets/test.json")
	jsonFile, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	bytesArray, _ := ioutil.ReadAll(jsonFile)
	var jobs job.Jobs
	json.Unmarshal(bytesArray, &jobs)

	return jobs.Jobs
}
