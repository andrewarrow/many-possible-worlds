package redis

func SetAuth(uuid string) {
}
func GetAuth(email, password string) string {
	m, _ := nc().HGetAll(ctx, email).Result()
	if len(m) == 0 {
		return ""
	}
	if m["pass"] != password {
		return ""
	}
	if m["all"] == "1" {
		return "all"
	}
	return m["worlds"]
}
