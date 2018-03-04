package artifactory_test

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/davecgh/go-spew/spew"
)

func testShouldRunIntegrationTests() bool {
	// are the env vars set?
	if os.Getenv("ARTIFACTORY_TOKEN") == "" || os.Getenv("ARTIFACTORY_URL") == "" {
		return false
	}
	// are we pointing to our test instance?
	return strings.Contains(os.Getenv("ARTIFACTORY_URL"), "lusis.jfrog.io")
}

func TestShouldRun(t *testing.T) {
	spew.Dump(os.Environ)
	fmt.Printf("%t\n", testShouldRunIntegrationTests())
}
