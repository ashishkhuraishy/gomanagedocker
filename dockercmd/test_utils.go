package dockercmd

import (
	"context"
	"errors"
	"fmt"
	"slices"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	dimage "github.com/docker/docker/api/types/image"
	"github.com/docker/docker/api/types/volume"
	"github.com/docker/docker/client"
)

type MockApi struct {
	mockVolumes []*volume.Volume
	mockImages  []dimage.Summary
	client.CommonAPIClient
}

func (m *MockApi) VolumeList(ctx context.Context, options volume.ListOptions) (volume.ListResponse, error) {
	return volume.ListResponse{
		Volumes: m.mockVolumes,
	}, nil
}

func (m *MockApi) VolumeRemove(ctx context.Context, volumeID string, force bool) error {
	final := []*volume.Volume{}

	for _, vol := range m.mockVolumes {
		if vol.Name == volumeID {
			continue
		}

		final = append(final, vol)
	}

	m.mockVolumes = final
	return nil

}

func (mo *MockApi) VolumesPrune(ctx context.Context, pruneFilter filters.Args) (types.VolumesPruneReport, error) {
	panic("not implemented") // TODO: Implement
}

func (m *MockApi) ImageList(ctx context.Context, options dimage.ListOptions) ([]dimage.Summary, error) {
	return m.mockImages, nil
}

func (m *MockApi) ImageRemove(ctx context.Context, image string, options dimage.RemoveOptions) ([]dimage.DeleteResponse, error) {

	res := []dimage.DeleteResponse{}

	index := slices.IndexFunc(m.mockImages, func(i dimage.Summary) bool {
		if i.ID == image {
			return true
		}

		return false
	})

	if !options.Force && m.mockImages[index].Containers > 0 {
		return nil, errors.New(fmt.Sprintf("unable to delete %s (must be forced) - image is ...", m.mockImages[index].ID))
	}

	m.mockImages = slices.Delete(m.mockImages, index, index+1)

	return res, nil

}

func (te *MockApi) ImagesPrune(ctx context.Context, pruneFilter filters.Args) (types.ImagesPruneReport, error) {
	final := []dimage.Summary{}

	for _, img := range te.mockImages {
		if img.Containers == 0 {
			continue
		}

		final = append(final, img)
	}

	te.mockImages = final

	return types.ImagesPruneReport{}, nil

}
