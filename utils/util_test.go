package utils_test

import (
	"log"
	"testing"

	"git.solusiteknologi.co.id/go-labs/learn-unit-test/utils"
)

func TestUtil(test *testing.T) {
	// case 1  all lower
	if utils.UcFirst("hallo") != "Hallo" {
		test.Error("Error on case 1 all lower")
	}

	// case 2 all upper
	if utils.UcFirst("KEBAB") != "KEBAB" {
		test.Error("Error on case 2 all upper")
	}

	// case 3 empty
	if utils.UcFirst("") != "" {
		test.Error("Error on case 3 all upper")
	}

	// all good
	log.Println("All test case passed")
}

func TestUtil2(test *testing.T) {
	// case 1  all lower
	if utils.UcFirst("hallo") != "Hallo" {
		test.Error("Error on case 1 all lower")
	}

	// case 2 all upper
	if utils.UcFirst("KEBAB") != "KEBAB" {
		test.Error("Error on case 2 all upper")
	}

	// case 3 empty
	if utils.UcFirst("") != "" {
		test.Error("Error on case 3 all upper")
	}

	// all good
	log.Println("All test util 2 case passed")
}
