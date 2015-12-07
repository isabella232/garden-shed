package quota_manager

import (
	"fmt"
	"path"

	"github.com/cloudfoundry-incubator/garden-shed/layercake"
	"github.com/pivotal-golang/lager"
)

type AUFSBaseSizer struct {
	cake layercake.Cake
}

func NewAUFSBaseSizer(cake layercake.Cake) *AUFSBaseSizer {
	return &AUFSBaseSizer{cake: cake}
}

func (a *AUFSBaseSizer) BaseSize(logger lager.Logger, containerRootFSPath string) (uint64, error) {
	var size uint64
	graphID := path.Base(containerRootFSPath)

	for graphID != "" {
		img, err := a.cake.Get(layercake.DockerImageID(graphID))
		if err != nil {
			return 0, fmt.Errorf("base-size %s: %s", graphID, err)
		}

		logger.Info("base-size", lager.Data{
			"layer":      graphID,
			"size":       img.Size,
			"total-size": size,
		})

		size += uint64(img.Size)
		graphID = img.Parent
	}

	return size, nil
}