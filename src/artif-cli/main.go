package artifcli

import (
	artifactory "artifactory.v491"
	"fmt"
)

func main() {
	client := artifactory.NewClientFromEnv()

	fmt.Printf("%#v\n", client)
}
