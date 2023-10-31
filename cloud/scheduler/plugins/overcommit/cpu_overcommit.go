package overcommit

import (
	"context"

	"github.com/sp-yduck/proxmox-go/api"

	"github.com/sp-yduck/cluster-api-provider-proxmox/cloud/scheduler/framework"
	"github.com/sp-yduck/cluster-api-provider-proxmox/cloud/scheduler/plugins/names"
)

type CPUOvercommit struct{}

var _ framework.NodeFilterPlugin = &CPUOvercommit{}

const (
	CPUOvercommitName         = names.CPUOvercommit
	defaultCPUOvercommitRatio = 4
)

func (pl *CPUOvercommit) Name() string {
	return CPUOvercommitName
}

// filter by cpu overcommit ratio
func (pl *CPUOvercommit) Filter(ctx context.Context, _ *framework.CycleState, config api.VirtualMachineCreateOptions, nodeInfo *framework.NodeInfo) *framework.Status {
	cpu := sumCPUs(nodeInfo.QEMUs())
	maxCPU := nodeInfo.Node().MaxCpu
	sockets := config.Sockets
	if sockets == 0 {
		sockets = 1
	}
	ratio := float32(cpu+config.Cores*sockets) / float32(maxCPU)
	if ratio > defaultCPUOvercommitRatio {
		status := framework.NewStatus()
		status.SetCode(1)
		return status
	}
	return &framework.Status{}
}

func sumCPUs(qemus []*api.VirtualMachine) int {
	var result int
	for _, q := range qemus {
		result += q.Cpus
	}
	return result
}