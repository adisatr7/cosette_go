package responses

import (
	"time"
	"math/rand"
)


// Returns a random response from a list of responses
func GetRandomResponse(listOfResponses []string) string {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomIndex := rng.Intn(len(listOfResponses))

	return listOfResponses[randomIndex]
}