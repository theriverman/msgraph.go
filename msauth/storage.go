package msauth

import (
	"io/ioutil"
	"os"
)

// ReadLocation reads data from file
func ReadLocation(loc string) ([]byte, error) {
	return ioutil.ReadFile(loc)
}

// WriteLocation writes data to file
func WriteLocation(loc string, b []byte, m os.FileMode) error {
	return ioutil.WriteFile(loc, b, m)
}
