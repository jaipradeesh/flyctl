package presenters

import (
	"github.com/superfly/flyctl/api"
)

type DeploymentStatus struct {
	Status *api.DeploymentStatus
}

func (p *DeploymentStatus) FieldNames() []string {
	return []string{"Status", "Description", "Timestamp"}
}

func (p *DeploymentStatus) Records() []map[string]string {
	out := []map[string]string{}

	out = append(out, map[string]string{
		"Status":      p.Status.Status,
		"Description": p.Status.Description,
		"Timestamp":   formatRelativeTime(p.Status.CreatedAt),
	})

	return out
}
