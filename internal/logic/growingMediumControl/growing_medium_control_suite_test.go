// Copyright 2022 Fox Design Agency. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package growingmediumcontrol_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestGrowingMediumControl(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Growing Medium Control Suite")
}
