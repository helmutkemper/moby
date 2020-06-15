// +build linux,cgo,!static_build

package devicemapper // import "github.com/helmutkemper/moby/pkg/devicemapper"

// #cgo pkg-config: devmapper
import "C"
