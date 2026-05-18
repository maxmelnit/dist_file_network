package control

import (
	"hash/fnv"
	"net/http"
	"strings"
)

const nodeCount = 3

// Generic FNV hash function used for determining appropriate storage node
func fnv_hash(item string) int {
	hash := fnv.New32a()
	hash.Write([]byte(item))
	return int(hash.Sum32())
}

// Route parses the object request, and routes it to the appropriate storage node
func Route(writer http.ResponseWriter, request *http.Request) {

	// Get the requested object ID
	objectID := strings.TrimPrefix(request.URL.Path, "/")

	// Calculate which node to capture the object from
	node := fnv_hash(objectID) % nodeCount

}
