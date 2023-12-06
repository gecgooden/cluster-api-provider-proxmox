package random

import (
	"context"
	"math/rand"
	"time"

	"github.com/k8s-proxmox/proxmox-go/api"

	"github.com/k8s-proxmox/cluster-api-provider-proxmox/cloud/scheduler/framework"
	"github.com/k8s-proxmox/cluster-api-provider-proxmox/cloud/scheduler/plugins/names"
)

type Random struct{}

var _ framework.NodeScorePlugin = &Random{}

const (
	Name = names.Random
)

func (pl *Random) Name() string {
	return Name
}

// return random score: 0 <= n < 100.
// just a sample plugin
func (pl *Random) Score(ctx context.Context, state *framework.CycleState, config api.VirtualMachineCreateOptions, nodeInfo *framework.NodeInfo) (int64, *framework.Status) {
	src := rand.NewSource(time.Now().Unix())
	r := rand.New(src)
	score := r.Int63n(100)
	return score, nil
}
