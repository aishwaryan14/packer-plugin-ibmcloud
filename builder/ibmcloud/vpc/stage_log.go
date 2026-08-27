package vpc

import (
	"github.com/hashicorp/packer-plugin-sdk/packer"
)

// emitStage writes a machine-readable stage marker to the Packer UI.
// The marker format is:  STAGE <stage> <status>
// status is one of: START, END, FAIL
//
// ui.Machine routes to the machine-readable output channel (log only for
// BasicUi; structured CSV when packer is run with -machine-readable) and
// does not pollute the human-readable terminal output.
//
// Canonical stage names and their mapping to the VPC image-job API reason codes:
//
//	create_instance       → provisioning_worker   (failed_worker_provisioning)
//	wait_instance         → provisioning_worker   (failed_worker_provisioning)
//	installing_components → installing_components (failed_component_installation)
//	image_validating      → image_validating      (failed_component_validation)
//	capture_image         → image_capturing       (failed_image_capture)
//	image_importing       → image_importing       (failed_image_import)
func emitStage(ui packer.Ui, stage, status string) {
	ui.Machine("stage", stage, status)
}
