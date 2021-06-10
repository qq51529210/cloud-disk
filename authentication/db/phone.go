package db

func GetPhoneCode(phone string) (string, error) {
	value, err := phoneRedis.Cmd("GET", phone)
	if err != nil {
		return "", err
	}
	if s, ok := value.(string); ok {
		return s, nil
	}
	return "", nil
}

func SetPhoneCode(phone, code string) error {
	_, err := phoneRedis.Cmd("SET", phone, code)
	return err
}
