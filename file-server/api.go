package main

type UploadInfo struct {
	Rate int `json:"rate"`
}

func ApiGetUploadInfo(token string) (*UploadInfo, error) {
	return &UploadInfo{
		Rate: 200 * 1024,
	}, nil
}

type DownloadInfo struct {
	Rate int `json:"rate"`
}

func ApiGetDownloadInfo(token string) (*DownloadInfo, error) {
	return &DownloadInfo{
		Rate: 200 * 1024,
	}, nil
}
