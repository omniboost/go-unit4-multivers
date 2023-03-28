package multivers_test

import (
	"encoding/json"
	"log"
	"testing"
)

func TestCustomerInvoiceInfoListByJournalGet(t *testing.T) {
	req := client.NewCustomerInvoiceInfoListByJournalGetRequest()
	req.PathParams().FiscalYear = 2022
	req.PathParams().JournalID = "19"
	req.PathParams().InvoiceState = 0
	req.QueryParams().Filter.Set("InvoiceReference eq '19181'")
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	log.Println(string(b))
}
