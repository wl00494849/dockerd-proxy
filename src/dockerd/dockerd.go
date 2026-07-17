package dockerd

import (
	"context"
	"docker-proxy/src/core"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
)

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

type Dockerd struct {
	UnixCli *http.Client
}

func NewDockerdCli() *Dockerd {
	config := core.LoadConfig()
	client := &http.Client{
		Transport: &http.Transport{
			DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
				return net.Dial("unix", config.DockerHost)
			},
		},
	}

	return &Dockerd{
		UnixCli: client,
	}
}

func (dock *Dockerd) Containers() []ContainerInfo {

	resp, err := dock.UnixCli.Get("http://docker/containers/json?all=true")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var data []ContainerInfo
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		panic(err)
	}

	return data
}

func (dock *Dockerd) Images() []ImageInfo {

	resp, err := dock.UnixCli.Get("http://docker/images/json?all=true")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var data []ImageInfo
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		panic(err)
	}

	return data
}

func (dock *Dockerd) Start(id string) error {

	url := fmt.Sprintf("http://docker/containers/%s/start", id)
	resp, err := dock.UnixCli.Post(url, "", nil)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("docker error: %s", body)
	}

	return nil
}

func (dock *Dockerd) Stop(id string) error {

	url := fmt.Sprintf("http://docker/containers/%s/stop", id)
	resp, err := dock.UnixCli.Post(url, "", nil)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("docker error: %s", body)
	}

	return nil
}
