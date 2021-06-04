package main

type uploadInfo struct {
	Rate int `json:"rate"`
}

func apiGetUploadInfo(token string, info *uploadInfo) error {
	return nil
}
