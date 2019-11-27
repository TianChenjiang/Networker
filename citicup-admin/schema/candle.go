package schema

type Candle struct {
	Close    float32 `json:"close"`          //收盘价
	Highest  float32 `json:"high"`           //最高价
	Lowest   float32 `json:"low"`            //最低价
	Open     float32 `json:"open"`           //开盘价
	Vol      float32 `json:"vol"`            //成交量
}
