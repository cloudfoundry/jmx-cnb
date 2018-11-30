/*
 * Copyright 2018 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"testing"

	"github.com/buildpack/libbuildpack/buildplan"
	"github.com/cloudfoundry/jmx-buildpack/jmx"
	"github.com/cloudfoundry/jvm-application-buildpack/jvmapplication"
	"github.com/cloudfoundry/libcfbuildpack/detect"
	"github.com/cloudfoundry/libcfbuildpack/test"
	"github.com/sclevine/spec"
	"github.com/sclevine/spec/report"
)

func TestDetect(t *testing.T) {
	spec.Run(t, "Detect", testDetect, spec.Report(report.Terminal{}))
}

func testDetect(t *testing.T, when spec.G, it spec.S) {

	it("fails without jvm-application", func() {
		f := test.NewDetectFactory(t)
		f.AddEnv(t, "BP_JMX", "")

		exitStatus, err := d(f.Detect)
		if err != nil {
			t.Fatal(err)
		}

		if exitStatus != detect.FailStatusCode {
			t.Errorf("os.Exit = %d, expected 100", exitStatus)
		}
	})

	it("fails without BP_JMX", func() {
		f := test.NewDetectFactory(t)
		f.AddBuildPlan(t, jvmapplication.Dependency, buildplan.Dependency{})

		exitStatus, err := d(f.Detect)
		if err != nil {
			t.Fatal(err)
		}

		if exitStatus != detect.FailStatusCode {
			t.Errorf("os.Exit = %d, expected 100", exitStatus)
		}
	})

	it("passes with jvm-application and BP_JMX", func() {
		f := test.NewDetectFactory(t)
		f.AddEnv(t, "BP_JMX", "")
		f.AddBuildPlan(t, jvmapplication.Dependency, buildplan.Dependency{})

		exitStatus, err := d(f.Detect)
		if err != nil {
			t.Fatal(err)
		}

		if exitStatus != detect.PassStatusCode {
			t.Errorf("os.Exit = %d, expected 0", exitStatus)
		}

		test.BeBuildPlanLike(t, f.Output, buildplan.BuildPlan{
			jmx.Dependency: buildplan.Dependency{},
		})
	})
}
