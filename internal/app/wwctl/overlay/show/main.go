package show

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/hpcng/warewulf/internal/pkg/config"
	"github.com/hpcng/warewulf/internal/pkg/util"
	"github.com/hpcng/warewulf/internal/pkg/wwlog"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func CobraRunE(cmd *cobra.Command, args []string) error {
	var overlaySourceDir string

	overlayKind := args[0]
	overlayName := args[1]
	fileName := args[2]

	if overlayKind != "system" && overlayKind != "runtime" {
		return errors.New("overlay kind must be of type 'system' or 'runtime'")
	}

	if overlayKind == "system" {
		overlaySourceDir = config.SystemOverlaySource(overlayName)
	} else if overlayKind == "runtime" {
		overlaySourceDir = config.RuntimeOverlaySource(overlayName)
	}

	if !util.IsDir(overlaySourceDir) {
		wwlog.Printf(wwlog.ERROR, "Overlay does not exist: %s\n", overlayName)
		os.Exit(1)
	}

	overlayFile := path.Join(overlaySourceDir, fileName)

	if !util.IsFile(overlayFile) {
		wwlog.Printf(wwlog.ERROR, "File does not exist within overlay: %s:%s\n", overlayName, fileName)
		os.Exit(1)
	}

	f, err := ioutil.ReadFile(overlayFile)
	if err != nil {
		wwlog.Printf(wwlog.ERROR, "Could not read file: %s\n", err)
		os.Exit(1)
	}

	fmt.Print(string(f))

	return nil
}
