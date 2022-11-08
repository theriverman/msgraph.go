package msauth

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// ReadLocation reads data from file
func ReadLocation(loc string) ([]byte, error) {
	return ioutil.ReadFile(loc)
}

// WriteLocation writes data to file
func WriteLocation(loc string, b []byte, m os.FileMode) error {
	return ioutil.WriteFile(loc, b, m)
}
