package multivers_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestAdministrationGet(t *testing.T) {
	req := client.NewFinTransGetRequest()
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
