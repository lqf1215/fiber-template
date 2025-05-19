package enum

// PayType 支付方式
type PayType int

const (
	PayTypeStripe PayType = 1 // Stripe
)

type Currency string

const (
	CurrencyUSDT Currency = "USDT" // USDT
	CurrencyUSD  Currency = "USD"  // 美元
	CurrencyCNH  Currency = "CNY"  // 人民币
	CurrencyVND  Currency = "VND"  // 越南盾
	CurrencyHKD  Currency = "HKD"  // 港元

)
