/*
   Copyright 2020 Docker Hub Tool authors

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package account

import (
	"bytes"
	"testing"
	"time"

	"gotest.tools/v3/assert"
	"gotest.tools/v3/golden"

	"github.com/docker/hub-tool/internal/hub"
)

func TestInfoOutput(t *testing.T) {
	account := account{
		Account: &hub.Account{
			ID:       "id",
			Name:     "my-user-name",
			FullName: "My Full Name",
			Location: "MyLocation",
			Company:  "My Company",
			Joined:   time.Now(),
		},
		Plan: &hub.Plan{
			Name: "free",
			Limits: hub.Limits{
				Seats:          1,
				PrivateRepos:   2,
				Teams:          9999,
				Collaborators:  9999,
				ParallelBuilds: 3,
			},
		},
		Consumption: &hub.Consumption{
			Seats:               0,
			PrivateRepositories: 1,
			Teams:               2,
		},
	}
	buf := bytes.NewBuffer(nil)
	err := printAccount(buf, account)
	assert.NilError(t, err)
	golden.Assert(t, buf.String(), "info.golden")
}
