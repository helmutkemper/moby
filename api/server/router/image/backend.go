package image // import "github.com/helmutkemper/moby/api/server/router/image"

import (
	"context"
	"io"

	"github.com/helmutkemper/moby/api/types"
	"github.com/helmutkemper/moby/api/types/filters"
	"github.com/helmutkemper/moby/api/types/image"
	"github.com/helmutkemper/moby/api/types/registry"
	specs "github.com/opencontainers/image-spec/specs-go/v1"
)

// Backend is all the methods that need to be implemented
// to provide image specific functionality.
type Backend interface {
	imageBackend
	importExportBackend
	registryBackend
}

type imageBackend interface {
	ImageDelete(imageRef string, force, prune bool) ([]types.ImageDeleteResponseItem, error)
	ImageHistory(imageName string) ([]*image.HistoryResponseItem, error)
	Images(imageFilters filters.Args, all bool, withExtraAttrs bool) ([]*types.ImageSummary, error)
	LookupImage(name string) (*types.ImageInspect, error)
	TagImage(imageName, repository, tag string) (string, error)
	ImagesPrune(ctx context.Context, pruneFilters filters.Args) (*types.ImagesPruneReport, error)
}

type importExportBackend interface {
	LoadImage(inTar io.ReadCloser, outStream io.Writer, quiet bool) error
	ImportImage(src string, repository, platform string, tag string, msg string, inConfig io.ReadCloser, outStream io.Writer, changes []string) error
	ExportImage(names []string, outStream io.Writer) error
}

type registryBackend interface {
	PullImage(ctx context.Context, image, tag string, platform *specs.Platform, metaHeaders map[string][]string, authConfig *types.AuthConfig, outStream io.Writer) error
	PushImage(ctx context.Context, image, tag string, metaHeaders map[string][]string, authConfig *types.AuthConfig, outStream io.Writer) error
	SearchRegistryForImages(ctx context.Context, filtersArgs string, term string, limit int, authConfig *types.AuthConfig, metaHeaders map[string][]string) (*registry.SearchResults, error)
}
