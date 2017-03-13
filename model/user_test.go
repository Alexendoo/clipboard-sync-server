package model_test

import (
	. "github.com/Alexendoo/sync/model"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Model/User", func() {
	var testUser User

	BeforeEach(func() {
		testUser = User{
			ID:          1,
			Name:        "foo",
			PassKey:     []byte("key"),
			PassSalt:    []byte("salt"),
			PassVersion: -1,
		}
	})

	It("Adds users", func() {
		store, _ := testStore("adduser")

		Expect(store.AddUser(&testUser)).
			NotTo(HaveOccurred())
	})

	It("Recalls users", func() {
		store, _ := testStore("recalluser")

		Expect(store.AddUser(&testUser)).
			NotTo(HaveOccurred())

		Expect(store.GetUser(testUser.ID)).
			To(Equal(&testUser))
	})
})
