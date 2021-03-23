package scaffolds

import (
	"sigs.k8s.io/kubebuilder/v3/pkg/config"
	"sigs.k8s.io/kubebuilder/v3/pkg/machinery"

	"github.com/java-operator-sdk/kubebuilder-plugin/pkg/internal/kubebuilder/cmdutil"
	"github.com/java-operator-sdk/kubebuilder-plugin/pkg/java/v1/scaffolds/internal/templates"
	"github.com/java-operator-sdk/kubebuilder-plugin/pkg/java/v1/util"
)

const (
	// kustomizeVersion is the sigs.k8s.io/kustomize version to be used in the project
	kustomizeVersion = "v3.5.4"

	imageName = "controller:latest"
)

// This file represents the scaffolding done by this init command

var _ cmdutil.Scaffolder = &initScaffolder{}

type initScaffolder struct {
	fs     machinery.Filesystem
	config config.Config
}

// NewInitScaffolder returns a new Scaffolder for project initialization operations
func NewInitScaffolder(config config.Config) cmdutil.Scaffolder {
	return &initScaffolder{
		config: config,
	}
}

// InjectFS implements Scaffolder
func (s *initScaffolder) InjectFS(fs machinery.Filesystem) {
	s.fs = fs
}

// Scaffold implements Scaffolder
func (s *initScaffolder) Scaffold() error {
	// Initialize the machinery.Scaffold that will write the files to disk
	scaffold := machinery.NewScaffold(s.fs,
		// NOTE: kubebuilder's default permissions are only for root users
		machinery.WithDirectoryPermissions(0755),
		machinery.WithFilePermissions(0644),
		machinery.WithConfig(s.config),
	)

	return scaffold.Execute(
		&templates.OperatorFile{
			Package:      util.ReverseDomain(s.config.GetDomain()),
			OperatorName: util.ToClassname(s.config.GetProjectName()),
		},
	)
}
