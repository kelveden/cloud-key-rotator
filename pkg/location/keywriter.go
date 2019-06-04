package location

//keyWriter defines the function signature for writing key to a location, e.g. CircleCI, K8S cluster or GitHub.
import "github.com/ovotech/cloud-key-rotator/pkg/cred"

//KeyWriter interface
type KeyWriter interface {
	Write(serviceAccountName, keyWrapper keyWrapper, creds cred.Credentials) (UpdatedLocation, error)
}
