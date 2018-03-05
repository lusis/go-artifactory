package artifactory_test

import (
	"fmt"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/lusis/go-artifactory/pkg/artifactory"
)

func testShouldRunIntegrationTests(t *testing.T) bool {
	// are the env vars set?
	c, err := artifactory.NewClientFromEnv()
	if err != nil {
		log.Printf("error creating client: %s", err.Error())
		return false
	}
	// are we pointing to our test instance?

	return !strings.Contains(c.Config.BaseURL, "lusis")
}

func TestShouldRun(t *testing.T) {
	spew.Dump(os.Environ())
	fmt.Printf("%t\n", testShouldRunIntegrationTests(t))
}
