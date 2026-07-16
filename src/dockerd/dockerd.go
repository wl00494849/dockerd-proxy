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

type Container struct {
	ID      string   `json:"id"`
	Names   []string `json:"names"`
	Image   string   `json:"image"`
	ImageID string   `json:"image_id"`
	Status  string   `json:"status"`
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

func (d *Dockerd) Containers() []Container {

	resp, err := d.UnixCli.Get("http://docker/containers/json?all=true")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var data []Container
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		panic(err)
	}

	return data
}

func (d *Dockerd) Images() {

}

func (d *Dockerd) Start(id string) error {
	url := fmt.Sprintf("http://docker/containers/%s/start", id)
	resp, err := d.UnixCli.Post(url, "", nil)
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

func (d *Dockerd) Stop(id string) {

}
