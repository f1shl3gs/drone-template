package template

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExecute(t *testing.T) {
	cases := []struct {
		Environment map[string]string
		Text        string
		Result      string
	}{
		{
			Environment: map[string]string{
				"DRONE_REPO_OWNER": "foo",
				"DRONE_REPO_NAME":  "bar",
			},
			Text:   "{{ .Repo.Owner }}/{{ .Repo.Name }}",
			Result: "foo/bar",
		},
		{
			Environment: map[string]string{
				"foo": "bar",
			},
			Text:   `{{ env "foo" }}`,
			Result: "bar",
		},
		{
			Environment: map[string]string{
				"DRONE_BUILD_CREATED": "1570291340",
			},
			Text:   "{{ date .Build.Created }}",
			Result: "2019-10-06T00:02:20+08:00",
		},
		{
			Environment: map[string]string{
				"DRONE_BUILD_CREATED": "1570291340",
			},
			Text:   `{{ date .Build.Created "America/Havana" }}`,
			Result: "2019-10-05T12:02:20-04:00",
		},
		{
			Environment: map[string]string{
				"DRONE_BUILD_CREATED": "1570291340",
			},
			Text:   `{{ date .Build.Created "America/Havana" "2006-01-02 15:04:05" }}`,
			Result: "2019-10-05 12:02:20",
		},
	}

	for _, c := range cases {
		for key, value := range c.Environment {
			err := os.Setenv(key, value)
			require.NoError(t, err)
		}

		actual, err := Execute(c.Text)
		require.NoError(t, err)

		assert.Equal(t, c.Result, actual)
	}
}
