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
//
// Canonical stage names and their mapping to the VPC image-job API reason codes:
//
//	create_instance      → provisioning_worker  (failed_worker_provisioning)
//	wait_instance        → provisioning_worker  (failed_worker_provisioning)
//	installing_components → installing_components (failed_component_installation)
//	image_validating     → image_validating      (failed_component_validation)
//	capture_image        → image_capturing       (failed_image_capture)
//	image_importing      → image_importing       (failed_image_import)
func emitStage(ui packer.Ui, stage, status string) {
	ui.Say(fmt.Sprintf("STAGE:%s:%s", stage, status))
}
