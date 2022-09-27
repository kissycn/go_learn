package book_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"ielee.com/go_learn/chapter12_test/ginkgo/book"
)

var _ = Describe("Books", func() {

	var foxInSocks, lesMis *book.Book

	BeforeEach(func() {
		lesMis = &book.Book{
			Title:  "Les Miserables",
			Author: "Victor Hugo",
			Pages:  2783,
		}
		foxInSocks = &book.Book{
			Title:  "Fox In Socks",
			Author: "Dr. Seuss",
			Pages:  24,
		}
	})

	Describe("Categorizing books", func() {
		When("with more than 300 pages", func() {
			It("should be a novel", func() {
				Expect(lesMis.Category()).To(Equal(book.CategoryNovel))
			})
		})

		Context("with fewer than 300 pages", func() {
			It("should be a short story", func() {
				Expect(foxInSocks.Category()).To(Equal(book.CategoryShortStory))
			})
		})
	})
})
