package internal

type Response struct {
	Code        int    `json:"code"`
	Guo         string `json:"guo"`
	Sheng       string `json:"sheng"`
	Shi         string `json:"shi"`
	Name        string `json:"name"`
	Weather1    string `json:"weather1"`
	Weather2    string `json:"weather2"`
	Wd1         string `json:"wd1"`
	Wd2         string `json:"wd2"`
	WindDir1    string `json:"winddirection1"`
	WindDir2    string `json:"winddirection2"`
	WindLevel1  string `json:"windleve1"`
	WindLevel2  string `json:"windleve2"`
	Weather1Img string `json:"weather1img"`
	Weather2Img string `json:"weather2img"`
	Lon         string `json:"lon"`
	Lat         string `json:"lat"`
	Uptime      string `json:"uptime"`
}
