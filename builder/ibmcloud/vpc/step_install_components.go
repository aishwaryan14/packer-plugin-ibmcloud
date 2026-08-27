package vpc

import (
	"context"

	"github.com/hashicorp/packer-plugin-sdk/multistep"
	"github.com/hashicorp/packer-plugin-sdk/multistep/commonsteps"
	"github.com/hashicorp/packer-plugin-sdk/packer"
)

type stepInstallComponents struct {
	provision commonsteps.StepProvision
}

func (s *stepInstallComponents) Run(ctx context.Context, state multistep.StateBag) multistep.StepAction {
	ui := state.Get("ui").(packer.Ui)

	emitStage(ui, "installing_components", "START")
	action := s.provision.Run(ctx, state)
	if action == multistep.ActionHalt {
		emitStage(ui, "installing_components", "FAIL")
		return multistep.ActionHalt
	}
	emitStage(ui, "installing_components", "END")
	return multistep.ActionContinue
}

func (s *stepInstallComponents) Cleanup(state multistep.StateBag) {
	s.provision.Cleanup(state)
}
