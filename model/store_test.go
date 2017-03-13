package model_test

import (
	. "github.com/Alexendoo/sync/model"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func testStore(name string) (*Store, error) {
	path := "file:" + name + "?cache=shared&mode=memory"
	return OpenStore(path)
}

var _ = Describe("Model/Store", func() {
	It("Opens a store", func() {
		Expect(testStore("open")).NotTo(BeNil())
	})
})
