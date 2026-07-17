package apiv1

type ResponseMsg struct {
	IsSuccess bool   `json:"is_sucess"`
	Message   string `json:"message"`
}

type ContainerRequest struct {
	ImageID  string `json:"image_id"`
	RepoTags string `json:"tags"`
}
