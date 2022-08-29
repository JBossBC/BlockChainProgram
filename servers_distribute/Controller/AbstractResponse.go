package controller

type AbstractRepsonse struct {
	Result  bool   `json:"result"`
	Message string `json:"message"`
}
