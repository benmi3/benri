package benriconfig

import (
	"fmt"
	"log"
	"os"

	"gitlab.com/benmi/benri/modules/ddns"
	"gopkg.in/yaml.v3"
)

var data = `
a: Easy!
b:
  c: 2
  d: [3, 4]
`

// Note: struct fields must be public in order for unmarshal to
// correctly populate the data.
type T struct {
	A string
	B struct {
		RenamedC int   `yaml:"c"`
		D        []int `yaml:",flow"`
	}
}

type Conf struct {
	Ddns []ddns.DdnsSettings
}

func readFile(filename string) ([]byte, error) {
	filedata, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}
	return filedata, nil
}

// ParseYAML parses a YAML file into a map[string]interface{}
func (c *Conf) ParseYAML(filename string) error {
	// Read the YAML file
	filedata, err := readFile(filename)
	if err != nil {
		return fmt.Errorf("error reading YAML file: %w", err)
	}
	// Parse the YAML data
	err = yaml.Unmarshal(filedata, c)
	if err != nil {
		return fmt.Errorf("error parsing YAML: %w", err)
	}

	return nil
}

func (c *Conf) init() error {
	err := c.ParseYAML("benriconfig.yaml")
	if err != nil {
		return err
	}
	return nil
}

func main() {
	t := T{}

	err := yaml.Unmarshal([]byte(data), &t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- t:\n%v\n\n", t)

	d, err := yaml.Marshal(&t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- t dump:\n%s\n\n", string(d))

	m := make(map[interface{}]interface{})

	err = yaml.Unmarshal([]byte(data), &m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- m:\n%v\n\n", m)

	d, err = yaml.Marshal(&m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- m dump:\n%s\n\n", string(d))
}
