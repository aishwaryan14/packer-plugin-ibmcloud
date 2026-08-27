package vpc

import (
	"fmt"

	"github.com/hashicorp/packer-plugin-sdk/packer"
)

// emitStage writes a machine-readable stage marker to the Packer UI.
// The marker format is:  STAGE:<stage>:<status>
// status is one of: START, END, FAIL
//
// When packer is run with -machine-readable every ui.Say line is emitted as a
// structured CSV token that automation can grep for without Packer knowledge:
//
//	1746750000,ibmcloud-vpc.centos,ui,say,STAGE:create_instance:START
func emitStage(ui packer.Ui, stage, status string) {
	ui.Say(fmt.Sprintf("STAGE:%s:%s", stage, status))
}
