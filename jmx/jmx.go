/*
 * Copyright 2018-2019 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package jmx

import (
	"github.com/cloudfoundry/libcfbuildpack/build"
	"github.com/cloudfoundry/libcfbuildpack/buildpack"
	"github.com/cloudfoundry/libcfbuildpack/layers"
)

// Dependency indicates that a JVM application should be run with JMX enabled.
const Dependency = "jmx"

// JMX represents the JMX configuration for a JVM application.
type JMX struct {
	info  buildpack.Info
	layer layers.HelperLayer
}

// Contribute makes the contribution to launch.
func (j JMX) Contribute() error {
	return j.layer.Contribute(func(artifact string, layer layers.HelperLayer) error {
		return layer.WriteProfile("jmx", `PORT=${BPL_JMX_PORT:=5000}

printf "JMX enabled on port ${PORT}\n"

export JAVA_OPTS="${JAVA_OPTS} \
  -Djava.rmi.server.hostname=127.0.0.1 \
  -Dcom.sun.management.jmxremote.authenticate=false \
  -Dcom.sun.management.jmxremote.ssl=false \
  -Dcom.sun.management.jmxremote.port=${PORT} \
  -Dcom.sun.management.jmxremote.rmi.port=${PORT}"
`)
	}, layers.Launch)
}

// NewJMX creates a new JMX instance. OK is true if build plan contains "jmx" dependency, otherwise false.
func NewJMX(build build.Build) (JMX, bool) {
	if !build.Plans.Has(Dependency) {
		return JMX{}, false
	}

	return JMX{build.Buildpack.Info, build.Layers.HelperLayer(Dependency, "JMX")}, true
}
