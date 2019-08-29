package schema

//k线图需要的数据
type Candle struct {
	Close   float32 `json:"close"`   //收盘价
	Highest float32 `json:"highest"` //最高价
	Lowest  float32 `json:"lowest"`  //最低价
	Open    float32 `json:"open"`    //开盘价
	Vol     float32 `json:"vol"`     //成交量
}

//TODO: 三个网络的数据
