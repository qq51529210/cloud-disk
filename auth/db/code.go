package db

func GetVerificationCode(account string) (string, error) {
	value, err := session.Cmd("GET", account)
	if err != nil {
		return "", err
	}
	if s, ok := value.(string); ok {
		return s, nil
	}
	return "", nil
}

func SetVerificationCode(account, code string, expire int64) error {
	_, err := session.Cmd("SET", account, code)
	if err != nil {
		return err
	}
	_, err = session.Cmd("EXPIRE", account, expire)
	return err
}
