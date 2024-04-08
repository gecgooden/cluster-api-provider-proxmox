/*
Copyright 2023 Teppei Sudo.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package scope

import (
	"context"
	"fmt"

	"github.com/go-logr/logr"
	"github.com/go-logr/zapr"
	"github.com/k8s-proxmox/proxmox-go/proxmox"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/cluster-api/util"
	"sigs.k8s.io/controller-runtime/pkg/client"

	infrav1 "github.com/k8s-proxmox/cluster-api-provider-proxmox/api/v1beta1"
)

type ProxmoxServices struct {
	Compute *proxmox.Service
}

func newComputeService(ctx context.Context, cluster *infrav1.ProxmoxCluster, crClient client.Client) (*proxmox.Service, error) {
	serverRef := cluster.Spec.ServerRef
	secretRef := serverRef.SecretRef
	if secretRef == nil {
		return nil, errors.New("failed to get proxmox client from nil secretRef")
	}

	var secret corev1.Secret
	key := client.ObjectKey{Namespace: secretRef.Namespace, Name: secretRef.Name}
	if err := crClient.Get(ctx, key, &secret); err != nil {
		return nil, fmt.Errorf("failed to get secret from secretRef: %w", err)
	}

	secret.SetOwnerReferences(util.EnsureOwnerRef(secret.OwnerReferences, metav1.OwnerReference{
		APIVersion: infrav1.GroupVersion.String(),
		Kind:       "ProxmoxCluster",
		Name:       cluster.Name,
		UID:        cluster.UID,
	}))
	if err := crClient.Update(ctx, &secret); err != nil {
		return nil, fmt.Errorf("failed to set ownerReference to secret: %w", err)
	}

	authConfig := proxmox.AuthConfig{
		Username: string(secret.Data["PROXMOX_USER"]),
		Password: string(secret.Data["PROXMOX_PASSWORD"]),
		TokenID:  string(secret.Data["PROXMOX_TOKENID"]),
		Secret:   string(secret.Data["PROXMOX_SECRET"]),
	}
	clientConfig := proxmox.ClientConfig{
		InsecureSkipVerify: true,
	}
	param := proxmox.NewParams(serverRef.Endpoint, authConfig, clientConfig)
	service, err := proxmox.GetOrCreateService(param)
	if err != nil {
		return nil, fmt.Errorf("failed to get or create service: %w", err)
	}
	var log logr.Logger

	// Configure logging level to the most verbose
	config := zap.NewDevelopmentConfig()
	config.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel) // Set the level to Debug

	// Create logger
	logger, err := config.Build()
	if err != nil {
		panic("failed to initialize zap logger: " + err.Error())
	}
	defer logger.Sync()
	log = zapr.NewLogger(logger)
	service.RESTClient().SetLogger(log)
	return service, nil
}
