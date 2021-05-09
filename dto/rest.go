package dto

type MultiplierResult struct {
	Result string `json:"result,omitempty"`
}

type MultiplierRequest struct {
	First  string `json:"first"`
	Second string `json:"second"`
}
