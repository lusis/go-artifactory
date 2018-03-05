package artifactory_test

import (
	"log"
	"os"
	"testing"

	"github.com/lusis/go-artifactory/pkg/artifactory"
)

func testShouldRunIntegrationTests(t *testing.T) bool {
	if os.Getenv("ARTIFACTORY_TEST_INTEGRATION") == "" {
		return false
	}
	// are the env vars set?
	_, err := artifactory.NewClientFromEnv()
	if err != nil {
		log.Printf("error creating client: %s", err.Error())
		return false
	}
	return true
}

func TestShouldRun(t *testing.T) {
	if !testShouldRunIntegrationTests(t) {
		t.Skip("skipping integration tests")
	}
}
