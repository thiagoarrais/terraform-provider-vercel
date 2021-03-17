package project

import (
	"encoding/json"
	"fmt"
	"github.com/chronark/terraform-provider-vercel/pkg/vercel/httpApi"
)

type ProjectHandler struct {
	Api *httpApi.Api
}

func (p *ProjectHandler) Create(project CreateProject) (string, error) {
	res, err := p.Api.Post("/v6/projects", project)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	var createdProject Project
	err = json.NewDecoder(res.Body).Decode(&createdProject)
	if err != nil {
		return "", nil
	}

	return createdProject.ID, nil
}
func (p *ProjectHandler) Read(id string) (project Project, err error) {
	res, err := p.Api.Get(fmt.Sprintf("/v1/projects/%s", id))
	if err != nil {
		return Project{}, fmt.Errorf("Unable to fetch project from vercel: %w", err)
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&project)
	if err != nil {
		return Project{}, fmt.Errorf("Unable to unmarshal project: %w", err)
	}
	return project, nil
}
func (p *ProjectHandler) Update(id string, project UpdateProject) error {
	res, err := p.Api.Patch(fmt.Sprintf("/v2/projects/%s", id), project)
	if err != nil {
		return fmt.Errorf("Unable to update project: %w", err)
	}
	defer res.Body.Close()
	return nil
}
func (p *ProjectHandler) Delete(id string) error {
	res, err := p.Api.Delete(fmt.Sprintf("/v1/projects/%s", id))
	if err != nil {
		return fmt.Errorf("Unable to delete project: %w", err)
	}
	defer res.Body.Close()
	return nil
}