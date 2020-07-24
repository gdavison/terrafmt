package test2

import (
	"fmt"
)

func testExtraLines() string {
	return fmt.Sprintf(`
resource "aws_s3_bucket" "extra-lines" {
  
  bucket = "tf-test-bucket-extra-lines"
}
`)
}

// This is included to verify blocks with diffs and no diffs in the same file
func testNoFormattingErrors(randInt int) string {
	return fmt.Sprintf(`
resource "aws_s3_bucket" "no-errors" {
  bucket = "tf-test-bucket-no-errors-%d"
}
`, randInt)
}

func testExtraSpace(randInt int) string {
	return fmt.Sprintf(`
resource "aws_s3_bucket" "extra-space" {
  bucket    = "tf-test-bucket-extra-space-%d"
}
`, randInt) + testReturnSprintfSimple()
}
