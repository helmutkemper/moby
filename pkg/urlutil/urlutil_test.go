package urlutil // import "github.com/helmutkemper/moby/pkg/urlutil"

import "testing"

var (
	gitUrls = []string{
		"git://github.com/helmutkemper/moby",
		"git@github.com:docker/docker.git",
		"git@bitbucket.org:atlassianlabs/atlassian-docker.git",
		"https://github.com/helmutkemper/moby.git",
		"http://github.com/helmutkemper/moby.git",
		"http://github.com/helmutkemper/moby.git#branch",
		"http://github.com/helmutkemper/moby.git#:dir",
	}
	incompleteGitUrls = []string{
		"github.com/helmutkemper/moby",
	}
	invalidGitUrls = []string{
		"http://github.com/helmutkemper/moby.git:#branch",
	}
	transportUrls = []string{
		"tcp://example.com",
		"tcp+tls://example.com",
		"udp://example.com",
		"unix:///example",
		"unixgram:///example",
	}
)

func TestIsGIT(t *testing.T) {
	for _, url := range gitUrls {
		if !IsGitURL(url) {
			t.Fatalf("%q should be detected as valid Git url", url)
		}
	}

	for _, url := range incompleteGitUrls {
		if !IsGitURL(url) {
			t.Fatalf("%q should be detected as valid Git url", url)
		}
	}

	for _, url := range invalidGitUrls {
		if IsGitURL(url) {
			t.Fatalf("%q should not be detected as valid Git prefix", url)
		}
	}
}

func TestIsTransport(t *testing.T) {
	for _, url := range transportUrls {
		if !IsTransportURL(url) {
			t.Fatalf("%q should be detected as valid Transport url", url)
		}
	}
}
