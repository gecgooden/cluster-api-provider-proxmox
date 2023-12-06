package instance

import (
	"context"
	"fmt"
	"strings"

	"github.com/k8s-proxmox/proxmox-go/api"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

// make sure storage exists and supports "images" type of content
func (s *Service) ensureStorageAvailable(ctx context.Context) error {
	log := log.FromContext(ctx)
	log.Info("ensuring storage is available")
	storageName := s.scope.GetStorage()
	if storageName == "" { // no storage specified, find available storage
		storage, err := s.findVMStorage(ctx)
		if err != nil {
			return err
		}
		storageName = storage.Storage
		s.scope.SetStorage(storageName)
	} else { // storage specified, check if it supports "images" type of content
		log.Info("checking if specified storage supports image type of content")
		storage, err := s.client.RESTClient().GetStorage(ctx, storageName)
		if err != nil {
			return err
		}
		if !supportsImage(storage) {
			return fmt.Errorf("storage %s does not support \"images\" type of content", storageName)
		}
	}
	return nil
}

// get one storage supporting "images" type of content
func (s *Service) findVMStorage(ctx context.Context) (*api.Storage, error) {
	log := log.FromContext(ctx)
	log.Info("finding available storage")
	storages, err := s.client.RESTClient().GetStorages(ctx)
	if err != nil {
		return nil, err
	}
	for _, storage := range storages {
		if supportsImage(storage) {
			return storage, nil
		}
	}
	return nil, fmt.Errorf("no available storage")
}

func supportsImage(storage *api.Storage) bool {
	return strings.Contains(storage.Content, "images")
}
