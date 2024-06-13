package tests

import (
    "testing"
    "options-analysis/models"
)

func TestOptionsContractValidation(t *testing.T) {
    contract := models.OptionsContract{
        Type:          models.Call,
        StrikePrice:   100.0,
        Bid:           1.0,
        Ask:           1.2,
        ExpirationDate: "2024-12-31",
        LongShort:     models.Long,
    }

    if contract.Type != models.Call {
        t.Errorf("Expected Type to be 'call', got '%s'", contract.Type)
    }
    // Additional field validation checks can be added here
}
