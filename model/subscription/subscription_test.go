package subscription

import (
	"encoding/json"
	"testing"
)

func TestSubscriptionDetailMarshalTTOAccountID(t *testing.T) {
	for _, detail := range []SubscriptionDetail{{TTOTCMAccountID: "new"}, {TTOTMCAccountID: "legacy"}} {
		body, err := json.Marshal(detail)
		if err != nil {
			t.Fatal(err)
		}
		var values map[string]any
		if err := json.Unmarshal(body, &values); err != nil {
			t.Fatal(err)
		}
		if _, ok := values["tto_tmc_account_id"]; ok {
			t.Fatal("misspelled field was encoded")
		}
		if _, ok := values["tto_tcm_account_id"]; !ok {
			t.Fatal("tto_tcm_account_id missing")
		}
	}
}
