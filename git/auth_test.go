package git

import (
	"strings"
	"testing"

	"github.com/biptec/terratest/modules/docker"
	terragit "github.com/biptec/terratest/modules/git"
	"github.com/biptec/terratest/modules/random"
)

func TestIntegrationGitAuth(t *testing.T) {
	t.Parallel()

	tag := "gruntwork-io/go-commons:" + strings.ToLower(random.UniqueId())
	ref := terragit.GetCurrentGitRef(t)
	docker.Build(t, "./test", &docker.BuildOptions{
		Tags:      []string{tag},
		BuildArgs: []string{"repo_ref=" + ref},
	})
	docker.Run(t, tag, &docker.RunOptions{
		Remove:               true,
		EnvironmentVariables: []string{"GITHUB_OAUTH_TOKEN"},
	})
}
