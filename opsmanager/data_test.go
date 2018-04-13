package opsmanager_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"strings"

	. "github.com/pivotal-cf/aqueduct-courier/opsmanager"
)

var _ = Describe("data", func() {
	It("returns a name", func() {
		d := NewData(
			strings.NewReader(""),
			"foo",
			"bar",
		)
		Expect(d.Name()).To(Equal("foo_bar"))
	})

	It("returns content for the data", func() {
		dataReader := strings.NewReader("best-data")
		d := NewData(dataReader, "", "")
		Expect(d.Content()).To(Equal(dataReader))
	})

	It("returns json as data type", func() {
		d := NewData(nil, "", "")
		Expect(d.ContentType()).To(Equal(JSONDataType))
	})
})
