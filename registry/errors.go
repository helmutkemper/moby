package registry // import "github.com/helmutkemper/moby/registry"

import (
	"net/url"

	"github.com/docker/distribution/registry/api/errcode"
	"github.com/helmutkemper/moby/errdefs"
)

type notFoundError string

func (e notFoundError) Error() string {
	return string(e)
}

func (notFoundError) NotFound() {}

func translateV2AuthError(err error) error {
	switch e := err.(type) {
	case *url.Error:
		switch e2 := e.Err.(type) {
		case errcode.Error:
			switch e2.Code {
			case errcode.ErrorCodeUnauthorized:
				return errdefs.Unauthorized(err)
			}
		}
	}

	return err
}
