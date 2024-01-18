package Models

type Data struct {
	Medication string `json:"medication"`
	Time       string `json:"time"`
	Frequency  string `json:"frequency"`
}

type Med struct {
	Id string `json:"medication"`
}
