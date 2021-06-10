package db

func GetSMSNumber(phone string) (string, error) {
	value, err := phoneRedis.Cmd("GET", phone)
	if err != nil {
		return "", err
	}
	if s, ok := value.(string); ok {
		return s, nil
	}
	return "", nil
}

func SetSMSNumber(phone, sms string) error {
	_, err := phoneRedis.Cmd("SET", phone, sms)
	return err
}
