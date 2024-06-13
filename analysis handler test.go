package tests

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"
    "options-analysis/handlers"
    "options-analysis/models"
)

func TestAnalysisEndpoint(t *testing.T) {
    contracts := []models.OptionsContract{
        {
            Type:          models.Call,
            StrikePrice:   100.0,
            Bid:           1.0,
            Ask:           1.2,
            ExpirationDate: "2024-12-31",
            LongShort:     models.Long,
        },
        // Add more test contracts as needed
    }

    reqBody, _ := json.Marshal(contracts)
    req, _ := http.NewRequest("POST", "/analyze", bytes.NewBuffer(reqBody))
    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(handlers.AnalysisEndpoint)

    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
    }

    var response map[string]interface{}
    if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
        t.Errorf("error parsing response: %v", err)
    }

    // Add assertions for response fields
}
