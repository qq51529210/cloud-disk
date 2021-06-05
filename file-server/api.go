package main

type UploadInfo struct {
	Rate int `json:"rate"`
}

func ApiGetUploadInfo(token string, info *UploadInfo) error {
	return nil
}
