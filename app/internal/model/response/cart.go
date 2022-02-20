package response

type CartProductRes struct {
	Num             int     `json:"num"`
	CurrentAllPrice float64 `json:"currentAllPrice"`
	AllPrice        float64 `json:"allPrice"`
}
