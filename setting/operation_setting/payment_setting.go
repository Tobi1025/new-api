package operation_setting

import "github.com/QuantumNous/new-api/setting/config"

type PaymentSetting struct {
	// AmountOptions are the preset amounts shown when the site displays USD.
	AmountOptions []int `json:"amount_options"`
	// CNYAmountOptions are the preset amounts shown when the site displays CNY.
	// They are converted to the internal USD amount before a top-up is created.
	CNYAmountOptions []int `json:"cny_amount_options"`
	// CNYMinTopUp is the minimum amount a user can enter when the site displays CNY.
	// It is converted to the internal USD amount before validating a top-up.
	CNYMinTopUp    float64         `json:"cny_min_topup"`
	AmountDiscount map[int]float64 `json:"amount_discount"` // 充值金额对应的折扣，例如 100 元 0.9 表示 100 元充值享受 9 折优惠

	ComplianceConfirmed    bool   `json:"compliance_confirmed"`
	ComplianceTermsVersion string `json:"compliance_terms_version"`
	ComplianceConfirmedAt  int64  `json:"compliance_confirmed_at"`
	ComplianceConfirmedBy  int    `json:"compliance_confirmed_by"`
	ComplianceConfirmedIP  string `json:"compliance_confirmed_ip"`
}

const CurrentComplianceTermsVersion = "v1"

// 默认配置
var paymentSetting = PaymentSetting{
	AmountOptions:    []int{10, 20, 50, 100, 200, 500},
	CNYAmountOptions: []int{10, 20, 50, 100, 200, 500},
	CNYMinTopUp:      10,
	AmountDiscount:   map[int]float64{},
}

func init() {
	// 注册到全局配置管理器
	config.GlobalConfig.Register("payment_setting", &paymentSetting)
}

func GetPaymentSetting() *PaymentSetting {
	return &paymentSetting
}

// GetTopupAmountOptions returns preset amounts in the internal USD unit.
// In CNY display mode, the administrator configures human-readable CNY values
// and this helper converts them with the current exchange rate.
func GetTopupAmountOptions() []float64 {
	if IsCNYDisplay() && len(paymentSetting.CNYAmountOptions) > 0 && USDExchangeRate > 0 {
		amounts := make([]float64, 0, len(paymentSetting.CNYAmountOptions))
		for _, amount := range paymentSetting.CNYAmountOptions {
			amounts = append(amounts, float64(amount)/USDExchangeRate)
		}
		return amounts
	}

	amounts := make([]float64, 0, len(paymentSetting.AmountOptions))
	for _, amount := range paymentSetting.AmountOptions {
		amounts = append(amounts, float64(amount))
	}
	return amounts
}

// GetTopupMinAmount returns the minimum amount in the internal USD unit used
// by the payment APIs. CNY input is converted before it reaches the backend.
func GetTopupMinAmount() float64 {
	if IsCNYDisplay() && paymentSetting.CNYMinTopUp > 0 && USDExchangeRate > 0 {
		return paymentSetting.CNYMinTopUp / USDExchangeRate
	}
	return MinTopUp
}

// GetTopupDisplayMinAmount returns the amount shown to users in the active
// display currency. It is intentionally separate from GetTopupMinAmount.
func GetTopupDisplayMinAmount() float64 {
	if IsCNYDisplay() && paymentSetting.CNYMinTopUp > 0 {
		return paymentSetting.CNYMinTopUp
	}
	return MinTopUp
}

func IsPaymentComplianceConfirmed() bool {
	return paymentSetting.ComplianceConfirmed &&
		paymentSetting.ComplianceTermsVersion == CurrentComplianceTermsVersion
}
