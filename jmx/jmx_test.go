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
	. "github.com/onsi/gomega"
	"github.com/sclevine/spec"
	"github.com/sclevine/spec/report"
)

func TestJMX(t *testing.T) {
	spec.Run(t, "JMX", func(t *testing.T, _ spec.G, it spec.S) {

		g := NewGomegaWithT(t)

		var f *test.BuildFactory

		it.Before(func() {
			f = test.NewBuildFactory(t)
		})

		it("returns true if build plan does exist", func() {
			f.AddBuildPlan(jmx.Dependency, buildplan.Dependency{})

			_, ok := jmx.NewJMX(f.Build)
			g.Expect(ok).To(BeTrue())
		})

		it("returns false if build plan does not exist", func() {
			_, ok := jmx.NewJMX(f.Build)
			g.Expect(ok).To(BeFalse())
		})

		it("contributes JMX configuration", func() {
			f.AddBuildPlan(jmx.Dependency, buildplan.Dependency{})

			d, _ := jmx.NewJMX(f.Build)
			g.Expect(d.Contribute()).To(Succeed())

			layer := f.Build.Layers.Layer("jmx")
			g.Expect(layer).To(test.HaveLayerMetadata(false, false, true))
			g.Expect(layer).To(test.HaveProfile("jmx", `PORT=${BPL_JMX_PORT:=5000}

printf "JMX enabled on port ${PORT}\n"

export JAVA_OPTS="${JAVA_OPTS} \
  -Djava.rmi.server.hostname=127.0.0.1 \
  -Dcom.sun.management.jmxremote.authenticate=false \
  -Dcom.sun.management.jmxremote.ssl=false \
  -Dcom.sun.management.jmxremote.port=${PORT} \
  -Dcom.sun.management.jmxremote.rmi.port=${PORT}"
`))
		})
	}, spec.Report(report.Terminal{}))
}
