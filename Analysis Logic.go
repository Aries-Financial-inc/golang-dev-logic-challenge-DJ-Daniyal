package utils

import (
    "math"
    "options-analysis/models"
)

func AnalyzeContracts(contracts []models.OptionsContract) ([]float64, []float64, float64, float64, []float64) {
    // Assume a range for the underlying price
    minPrice, maxPrice := 0.0, 1000.0
    step := 1.0

    var xValues, yValues []float64
    var maxProfit, maxLoss float64 = -math.MaxFloat64, math.MaxFloat64
    breakEvenPoints := make([]float64, 0)

    for price := minPrice; price <= maxPrice; price += step {
        profitLoss := 0.0
        for _, contract := range contracts {
            switch contract.Type {
            case models.Call:
                if contract.LongShort == models.Long {
                    profitLoss += math.Max(0, price-contract.StrikePrice) - contract.Ask
                } else {
                    profitLoss += contract.Bid - math.Max(0, price-contract.StrikePrice)
                }
            case models.Put:
                if contract.LongShort == models.Long {
                    profitLoss += math.Max(0, contract.StrikePrice-price) - contract.Ask
                } else {
                    profitLoss += contract.Bid - math.Max(0, contract.StrikePrice-price)
                }
            }
        }
        xValues = append(xValues, price)
        yValues = append(yValues, profitLoss)
        if profitLoss > maxProfit {
            maxProfit = profitLoss
        }
        if profitLoss < maxLoss {
            maxLoss = profitLoss
        }
        if math.Abs(profitLoss) < 0.01 { // Consider as break-even point
            breakEvenPoints = append(breakEvenPoints, price)
        }
    }

    return xValues, yValues, maxProfit, maxLoss, breakEvenPoints
}
