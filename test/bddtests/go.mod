// Copyright SecureKey Technologies Inc. All Rights Reserved.
//
// SPDX-License-Identifier: Apache-2.0

module github.com/trustbloc/sidetree-mock/test/bddtests

require (
	github.com/cucumber/godog v0.8.1
	github.com/fsouza/go-dockerclient v1.3.0
	github.com/mr-tron/base58 v1.1.3
	github.com/pkg/errors v0.9.1
	github.com/sirupsen/logrus v1.4.2
	github.com/trustbloc/sidetree-core-go v0.6.0
	github.com/trustbloc/sidetree-mock v0.0.0
)

replace (
	github.com/trustbloc/sidetree-mock => ../../
)

go 1.13
