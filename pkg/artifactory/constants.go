package artifactory

import "errors"

const basicAuthType = "basic"
const tokenAuthType = "token"

// MaxArtifactoryVersion is the max version of artifactory this library has been tested with
var MaxArtifactoryVersion = versionMustParse("5.5")

var errDecoding = errors.New("Could not parse response from the Artifactory server")
var errEncoding = errors.New("Could not marshal request to JSON")
