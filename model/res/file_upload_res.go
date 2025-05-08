package res

// FileUploadResponse 回复上传文件模板
type FileUploadResponse struct {
	FileName  string `json:"file_name" `
	Url       string `json:"url"`
	IsSuccess bool   `json:"is_success"`
	Msg       string `json:"msg"`
}
