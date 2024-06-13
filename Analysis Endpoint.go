package handlers

import (
    "encoding/json"
    "net/http"
    "options-analysis/models"
    "options-analysis/utils"
)

func AnalysisEndpoint(w http.ResponseWriter, r *http.Request) {
    var contracts []models.OptionsContract
    if err := json.NewDecoder(r.Body).Decode(&contracts); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    if len(contracts) > 4 {
        http.Error(w, "Maximum of 4 contracts allowed", http.StatusBadRequest)
        return
    }

    xValues, yValues, maxProfit, maxLoss, breakEvenPoints := utils.AnalyzeContracts(contracts)

    result := map[string]interface{}{
        "x_values": xValues,
        "y_values": yValues,
        "max_profit": maxProfit,
        "max_loss": maxLoss,
        "break_even_points": breakEvenPoints,
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(result)
}
