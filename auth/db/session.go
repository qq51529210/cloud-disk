package db

func HasToken(token string) (bool, error) {
	value, err := session.Cmd("EXISTS", token)
	if err != nil {
		return false, err
	}
	if n, ok := value.(int64); ok {
		return n == 1, nil
	}
	return false, nil
}

func SetToken(token string) error {
	_, err := session.Cmd("SET", token, "1")
	return err
}
