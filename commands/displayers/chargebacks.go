package displayers

import "github.com/VictorAvelar/mollie-api-go/v2/mollie"

// MollieChargeback wrapper for displaying
type MollieChargeback struct {
	*mollie.Chargeback
}

// KV is a displayable group of key value
func (cb *MollieChargeback) KV() []map[string]interface{} {
	var out []map[string]interface{}

	x := map[string]interface{}{
		"ID":         cb.ID,
		"Payment":    cb.PaymentID,
		"Amount":     stringCombinator(" ", cb.Amount.Value, cb.Amount.Currency),
		"Settlement": stringCombinator(" ", cb.SettlementAmount.Value, cb.SettlementAmount.Currency),
		"Created at": cb.CreatedAt.Format("02-01-2006"),
	}

	out = append(out, x)

	return out
}

// MollieChargebackList wrapper for displaying
type MollieChargebackList struct {
	*mollie.ChargebackList
}

// KV is a displayable group of key value
func (lp *MollieChargebackList) KV() []map[string]interface{} {
	var out []map[string]interface{}

	for _, p := range lp.Embedded.Chargebacks {
		x := map[string]interface{}{
			"ID":         p.ID,
			"Payment":    p.PaymentID,
			"Amount":     stringCombinator(" ", p.Amount.Value, p.Amount.Currency),
			"Settlement": stringCombinator(" ", p.SettlementAmount.Value, p.SettlementAmount.Currency),
			"Created at": p.CreatedAt.Format("02-01-2006"),
		}

		out = append(out, x)
	}

	return out
}
