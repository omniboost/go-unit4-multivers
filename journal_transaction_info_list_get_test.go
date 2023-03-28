package multivers_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestJournalTransactionInfoListGet(t *testing.T) {
	req := client.NewJournalTransactionInfoListGetRequest()
	req.PathParams().FiscalYear = 2022
	req.PathParams().JournalID = "19"
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
