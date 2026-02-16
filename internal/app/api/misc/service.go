package misc

import (
	"context"
	"fmt"
	"time"

	"github.com/lvjp/womblock/pkg/buildinfo"
)

type Service interface {
	Version(context.Context) (*VersionResponse, error)
}

func NewService() Service {
	return &service{}
}

type service struct{}

func (*service) Version(ctx context.Context) (*VersionResponse, error) {
	bi := buildinfo.Get()

	ret := &VersionResponse{
		Go:       bi.GoVersion,
		Modified: bi.Modified,
		Platform: bi.GoOS + "/" + bi.GoArch,
	}

	if bi.Revision != "-" {
		ret.Revision = bi.Revision
	}

	if bi.RevisionTime != "-" {
		revisionTime, err := time.Parse(time.RFC3339, bi.RevisionTime)
		if err != nil {
			return nil, fmt.Errorf("misc_version: failed to parse revision time: %q", bi.RevisionTime)
		}

		ret.Time = revisionTime
	}

	return ret, nil
}
