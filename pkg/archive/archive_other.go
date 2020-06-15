// +build !linux

package archive // import "github.com/helmutkemper/moby/pkg/archive"

func getWhiteoutConverter(format WhiteoutFormat, inUserNS bool) tarWhiteoutConverter {
	return nil
}
