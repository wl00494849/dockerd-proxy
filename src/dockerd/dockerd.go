package dockerd

import (
	"context"
	"docker-proxy/src/core"
	"encoding/json"
	"net"
	"net/http"
)

type Container struct {
	ID     string   `json:"Id"`
	Names  []string `json:"Names"`
	Image  string   `json:"Image"`
	Status string   `json:"Status"`
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

func (d *Dockerd) ContainerList() []Container {

	req, err := http.NewRequest("GET", "http://docker/containers/json?all=true", nil)
	if err != nil {
		panic(err)
	}

	resp, err := d.UnixCli.Do(req)
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
