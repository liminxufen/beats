// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package http

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "http", Asset); err != nil {
		panic(err)
	}
}

// Asset returns asset data
func Asset() string {
	return "eJzMlD1vszAUhXd+xRHzG+kdMjF07lR1yFZ1cPBJ4gQw9b2k5d9XhEIDdSqlH1LvaJvnPL42XuDANsNOtU4AdVowQ3q7Wt2nCWApeXC1Ol9luEkAoJtC6W1TMAECCxphhq1JAKGqq7aS4SEVKdJ/SDtw+pgAG8fCSnZiLFCZkmNqV9rWHSX4ZhiJZE8p56TAp4ai43gMeBE61GlrbyS4auNDabqVZ8vm+ZPd0FiGGbPX8Os9c51NfeoCrHbsjXquQFhpNHjtbRuNPbB99sF+Mbc2beGNnea+N1xqXwl/pOM96g+2PDCnO9JGw3Nv+QttFzXaSIzOF1PW3f+5/L+MGtW7YOQap5F45xUb31TflJ4JDF57mZzo1Rel+x4lNbhceH6m0+dnqEuPhDAcJ7flapEPhFFgTY0pvAYAAP//U2hb1w=="
}
