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

func SetPhoneCode(phone, code string, expire int64) error {
	_, err := phoneRedis.Cmd("SET", phone, code)
	if err != nil {
		return err
	}
	_, err = phoneRedis.Cmd("EXPIRE", phone, expire)
	return err
}
