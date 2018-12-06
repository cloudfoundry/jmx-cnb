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

package jmx_test

import (
	"testing"

	"github.com/buildpack/libbuildpack/buildplan"
	"github.com/cloudfoundry/jmx-buildpack/jmx"
	"github.com/cloudfoundry/libcfbuildpack/test"
	"github.com/sclevine/spec"
	"github.com/sclevine/spec/report"
)

func TestJMX(t *testing.T) {
	spec.Run(t, "JMX", testJMX, spec.Report(report.Terminal{}))
}

func testJMX(t *testing.T, when spec.G, it spec.S) {

	it("returns true if build plan does exist", func() {
		f := test.NewBuildFactory(t)
		f.AddBuildPlan(t, jmx.Dependency, buildplan.Dependency{})

		_, ok := jmx.NewJMX(f.Build)
		if !ok {
			t.Errorf("NewJMX = %t, expected true", ok)
		}
	})

	it("returns false if build plan does not exist", func() {
		f := test.NewBuildFactory(t)

		_, ok := jmx.NewJMX(f.Build)
		if ok {
			t.Errorf("NewJMX = %t, expected false", ok)
		}
	})

	it("contributes JMX configuration", func() {
		f := test.NewBuildFactory(t)
		f.AddBuildPlan(t, jmx.Dependency, buildplan.Dependency{})

		d, _ := jmx.NewJMX(f.Build)
		if err := d.Contribute(); err != nil {
			t.Fatal(err)
		}

		layer := f.Build.Layers.Layer("jmx")
		test.BeLayerLike(t, layer, false, false, true)
		test.BeProfileLike(t, layer, "jmx", `PORT=${BPL_JMX_PORT:=5000}

printf "JMX enabled on port ${PORT}"

export JAVA_OPTS="${JAVA_OPTS} \ 
  -Djava.rmi.server.hostname=127.0.0.1 \
  -Dcom.sun.management.jmxremote.authenticate=false \
  -Dcom.sun.management.jmxremote.ssl=false \
  -Dcom.sun.management.jmxremote.port=${PORT} \
  -Dcom.sun.management.jmxremote.rmi.port=${PORT}"
`)
	})
}
