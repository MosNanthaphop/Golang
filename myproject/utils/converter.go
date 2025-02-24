package utils

func NumberToThai(a int) string {
	thaiNumbers := []string{"ศูนย์", "หนึ่ง", "สอง", "สาม", "สี่", "ห้า", "หก", "เจ็ด", "แปด", "เก้า"}
	if a >= 0 && a <= 9 {
		return thaiNumbers[a]
	}
	return "error convert"
}
