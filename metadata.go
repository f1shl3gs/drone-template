package template

// prefix: DRONE_REPO
type Repo struct {
	Owner string `envconfig:"DRONE_REPO_OWNER"`
	Name  string `envconfig:"DRONE_REPO_NAME"`
}

// prefix: DRONE_COMMIT
type Commit struct {
	Sha          string `envconfig:"DRONE_COMMIT_SHA"`
	Before       string `envconfig:"DRONE_COMMIT_BEFORE"`
	After        string `envconfig:"DRONE_COMMIT_AFTER"`
	Ref          string `envconfig:"DRONE_COMMIT_REF"`
	Branch       string `envconfig:"DRONE_COMMIT_BRANCH"`
	Link         string `envconfig:"DRONE_COMMIT_LINK"`
	Message      string `envconfig:"DRONE_COMMIT_MESSAGE"`
	Author       string `envconfig:"DRONE_COMMIT_AUTHOR"`
	AuthorEmail  string `envconfig:"DRONE_COMMIT_AUTHOR_EMAIL"`
	AuthorAvatar string `envconfig:"DRONE_COMMIT_AUTHOR_AVATAR"`
	AuthorName   string `envconfig:"DRONE_COMMIT_AUTHOR_NAME"`
	BuildLink    string `envconfig:"DRONE_COMMIT_BUILD_LINK"`
}

// prefix: DRONE_BUILD
type Build struct {
	// msic
	Tag          string `envconfig:"DRONE_TAG"`
	Branch       string `envconfig:"DRONE_BRANCH"`
	SourceBranch string `envconfig:"DRONE_SOURCE_BRANCH"`
	TargetBranch string `envconfig:"DRONE_TARGET_BRANCH"`
	DeployTo     string `envconfig:"DRONE_DEPLOY_TO"`

	// build info
	Link     string `envconfig:"DRONE_BUILD_LINK"`
	Number   string `envconfig:"DRONE_BUILD_NUMBER"`
	Event    string `envconfig:"DRONE_BUILD_EVENT"`
	Status   string `envconfig:"DRONE_BUILD_STATUS"`
	Action   string `envconfig:"DRONE_BUILD_ACTION"`
	Created  int64  `envconfig:"DRONE_BUILD_CREATED"`
	Started  int64  `envconfig:"DRONE_BUILD_STARTED"`
	Finished int64  `envconfig:"DRONE_BUILD_FINISHED"`
}

// prefix: DRONE_STAGE
type Stage struct {
	Kind      string   `envconfig:"DRONE_STAGE_KIND"`
	Name      string   `envconfig:"DRONE_STAGE_NAME"`
	Number    int      `envconfig:"DRONE_STAGE_NUMBER"`
	Machine   string   `envconfig:"DRONE_STAGE_MACHINE"`
	OS        string   `envconfig:"DRONE_STAGE_OS"`
	Arch      string   `envconfig:"DRONE_STAGE_ARCH"`
	Variant   string   `envconfig:"DRONE_STAGE_VARIANT"`
	DependsOn []string `envconfig:"DRONE_STAGE_DEPENDS_ON"`
}

// prefix: DRONE_SYSTEM
type System struct {
	Proto    string `envconfig:"DRONE_SYSTEM_PROTO"`
	Host     string `envconfig:"DRONE_SYSTEM_HOST"`
	Hostname string `envconfig:"DRONE_SYSTEM_HOSTNAME"`
	Version  string `envconfig:"DRONE_SYSTEM_VERSION"`
}

type Runner struct {
	Machine  string `envconfig:"DRONE_MACHINE"`
	Host     string `envconfig:"DRONE_RUNNER_HOST"`
	Hostname string `envconfig:"DRONE_RUNNER_HOSTNAME"`
	Platform string `envconfig:"DRONE_RUNNER_PLATFORM"`
}

type Metadata struct {
	Repo
	Build
	Commit
	Stage
	System
}
