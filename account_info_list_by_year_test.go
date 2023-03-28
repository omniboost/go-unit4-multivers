package multivers_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestAccountInfoListByYearGet(t *testing.T) {
	req := client.NewAccountInfoListByYearGetRequest()
	req.PathParams().Year = 2022
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
