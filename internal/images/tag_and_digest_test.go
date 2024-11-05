package images

import (
	"testing"
)

// test that the remove tag if digest exists works as expected when we have both present.
func TestBothExist(t *testing.T) {
	image := "testingImage"
	tag := "v1.0.1"
	digest := "sha123:123649118273"
	url := image + ":" + tag + "@" + digest

	expectedResult := image + "@" + digest
	result := removeTagIfDigestExists(url)

	if expectedResult != result {
		t.Fatalf("removeTagIfDigestExists(%s) = %s but we want %s\n", url, result, expectedResult)
	}
}

// test that the remove tag if digest exists works as expected when we have only digest.
func TestOnlyDigest(t *testing.T) {
	image := "testingImage"
	digest := "sha123:123649118273"
	url := image + "@" + digest

	expectedResult := image + "@" + digest
	result := removeTagIfDigestExists(url)

	if expectedResult != result {
		t.Fatalf("removeTagIfDigestExists(%s) = %s but we want %s\n", url, result, expectedResult)
	}
}

// test that the remove tag if digest exists works as expected when we have tag only.
func TestOnlyTag(t *testing.T) {
	image := "testingImage"
	tag := "v1.0.1"
	url := image + ":" + tag
	expectedResult := image + ":" + tag
	result := removeTagIfDigestExists(url)

	if expectedResult != result {
		t.Fatalf("removeTagIfDigestExists(%s) = %s but we want %s\n", url, result, expectedResult)
	}
}
