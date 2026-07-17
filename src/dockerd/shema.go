package dockerd

type ImageInfo struct {
	ID          string   `json:"ID"`
	RepoTags    []string `json:"RepoTags"`
	RepoDigests []string `json:"RepoDigests"`
}

type ContainerInfo struct {
	ID      string   `json:"Id"`
	Names   []string `json:"Names"`
	Image   string   `json:"Image"`
	ImageID string   `json:"Imageid"`
	Status  string   `json:"Status"`
}
