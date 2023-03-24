package todo

import (
	"strings"
	"testing"
)

func TestExtractProjectFromTextWithSlash(t *testing.T) {
	got, _ := extractProjectFromText("#someday/maybe ride a bicycle")
	want := "someday/maybe"

	assertEqual(t, got, want)
}

func TestExtractProjectFromTextWithNumberSign(t *testing.T) {
	got, _ := extractProjectFromText("#someday#maybe ride a bicycle")
	want := "someday#maybe"

	assertEqual(t, got, want)
}

func TestExtractProjectFromText(t *testing.T) {
	got, _ := extractProjectFromText("#wishlist ride a bicycle")
	want := "wishlist"

	assertEqual(t, got, want)
}

func TestExtractProjectFromTextSignsPrefix(t *testing.T) {
	got, _ := extractProjectFromText("###////wishlist ride a bicycle")
	want := "wishlist"

	assertEqual(t, got, want)
}

func TestExtractProjectFromTextSignsPrefix2(t *testing.T) {
	got, _ := extractProjectFromText("### ////wishlist ride a bicycle")
	want := "wishlist"

	assertEqual(t, got, want)
}

func TestExtractProjectFromTextSignsPrefix3(t *testing.T) {
	got, _ := extractProjectFromText("### //// wishlist ride a bicycle")
	want := "wishlist"

	assertEqual(t, got, want)
}

func TestExtractProjectFromTextInbox(t *testing.T) {
	got, _ := extractProjectFromText("some   task")
	want := "Inbox"

	assertEqual(t, got, want)
}

func TestExtractProjectFromTextBadInput(t *testing.T) {
	got, _ := extractProjectFromText("#   ")
	want := "Inbox"

	assertEqual(t, got, want)
}

func TestExtractProjectFromTextDoubleNumberSign(t *testing.T) {
	got, _ := extractProjectFromText("##wishlist    ride a bicycle")
	want := "wishlist"

	assertEqual(t, got, want)
}

func assertEqual(t *testing.T, got, want string) {
	if !strings.EqualFold(got, want) {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestAutoCorrectProjectIdea(t *testing.T) {
	got := autoCorrectProject("idea")
	want := "ideas"

	assertEqual(t, got, want)
}

func TestAutoCorrectProjectIdeaLower(t *testing.T) {
	got := autoCorrectProject("IDEA")
	want := "ideas"

	assertEqual(t, got, want)
}

func TestAutoCorrectProjectIdeas(t *testing.T) {
	got := autoCorrectProject("ideas")
	want := "ideas"

	assertEqual(t, got, want)
}

func TestAutoCorrectProjectIdeasLower(t *testing.T) {
	got := autoCorrectProject("Ideas")
	want := "ideas"

	assertEqual(t, got, want)
}
