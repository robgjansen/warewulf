package kernel

import (
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path"

	"github.com/pkg/errors"

	"github.com/hpcng/warewulf/internal/pkg/config"
	"github.com/hpcng/warewulf/internal/pkg/util"
	"github.com/hpcng/warewulf/internal/pkg/wwlog"
)

func ParentDir() string {
	return path.Join(config.LocalStateDir, "provision/kernel")
}

func KernelImage(kernelVersion string) string {
	if kernelVersion == "" {
		wwlog.Printf(wwlog.ERROR, "Kernel Version is not defined\n")
		return ""
	}

	if !util.ValidString(kernelVersion, "^[a-zA-Z0-9-._]+$") {
		wwlog.Printf(wwlog.ERROR, "Runtime overlay name contains illegal characters: %s\n", kernelVersion)
		return ""
	}

	return path.Join(ParentDir(), kernelVersion, "vmlinuz")
}

func KmodsImage(kernelVersion string) string {
	if kernelVersion == "" {
		wwlog.Printf(wwlog.ERROR, "Kernel Version is not defined\n")
		return ""
	}

	if !util.ValidString(kernelVersion, "^[a-zA-Z0-9-._]+$") {
		wwlog.Printf(wwlog.ERROR, "Runtime overlay name contains illegal characters: %s\n", kernelVersion)
		return ""
	}

	return path.Join(ParentDir(), kernelVersion, "kmods.img")
}

func ListKernels() ([]string, error) {
	var ret []string

	err := os.MkdirAll(ParentDir(), 0755)
	if err != nil {
		return ret, errors.New("Could not create Kernel parent directory: " + ParentDir())
	}

	wwlog.Printf(wwlog.DEBUG, "Searching for Kernel image directories: %s\n", ParentDir())

	kernels, err := ioutil.ReadDir(ParentDir())
	if err != nil {
		return ret, err
	}

	for _, kernel := range kernels {
		wwlog.Printf(wwlog.VERBOSE, "Found Kernel: %s\n", kernel.Name())

		ret = append(ret, kernel.Name())

	}

	return ret, nil
}

func Build(kernelVersion string, root string) (string, error) {
	kernelImage := path.Join(root, "/boot/vmlinuz-"+kernelVersion)
	kernelDrivers := path.Join(root, "/lib/modules/"+kernelVersion)
	kernelDestination := KernelImage(kernelVersion)
	driversDestination := KmodsImage(kernelVersion)

	// Create the destination paths just in case it doesn't exist
	err := os.MkdirAll(path.Dir(kernelDestination), 0755)
	if err != nil {
		return "", errors.Wrap(err, "failed to create kernel dest")
	}

	err = os.MkdirAll(path.Dir(driversDestination), 0755)
	if err != nil {
		return "", errors.Wrap(err, "failed to create driver dest")
	}

	if !util.IsFile(kernelImage) {
		if !util.IsFile(kernelImage + ".gz") {
			return "", errors.New("Could not locate kernel image")
		} else {
			kernelImage = kernelImage + ".gz"
		}
	}

	if !util.IsDir(kernelDrivers) {
		return "", errors.New("Could not locate kernel drivers")
	}

	wwlog.Printf(wwlog.VERBOSE, "Setting up Kernel\n")
	if _, err := os.Stat(kernelImage); err == nil {
		kernel, err := os.Open(kernelImage)
		if err != nil {
			return "", errors.Wrap(err, "could not open kernel")
		}
		defer kernel.Close()

		gzipreader, err := gzip.NewReader(kernel)
		if err == nil {
			defer gzipreader.Close()

			writer, err := os.Create(kernelDestination)
			if err != nil {
				return "", errors.Wrap(err, "could not decompress kernel")
			}
			defer writer.Close()

			_, err = io.Copy(writer, gzipreader)
			if err != nil {
				return "", errors.Wrap(err, "could not write decompressed kernel")
			}

		} else {

			err := util.CopyFile(kernelImage, kernelDestination)
			if err != nil {
				return "", errors.Wrap(err, "could not copy kernel")
			}
		}

	}

	wwlog.Printf(wwlog.VERBOSE, "Building Kernel driver image\n")
	if _, err := os.Stat(kernelDrivers); err == nil {
		compressor, err := exec.LookPath("pigz")
		if err != nil {
			wwlog.Printf(wwlog.VERBOSE, "Could not locate PIGZ, using GZIP\n")
			compressor = "gzip"
		} else {
			wwlog.Printf(wwlog.VERBOSE, "Using PIGZ to compress the container: %s\n", compressor)
		}

		cmd := fmt.Sprintf("cd /; find .%s | cpio --quiet -o -H newc | %s -c > \"%s\"", kernelDrivers, compressor, driversDestination)

		wwlog.Printf(wwlog.DEBUG, "RUNNING: %s\n", cmd)
		err = exec.Command("/bin/sh", "-c", cmd).Run()
		if err != nil {
			return "", err
		}
	}

	return "Done", nil
}

func DeleteKernel(name string) error {
	fullPath := path.Join(ParentDir(), name)

	wwlog.Printf(wwlog.VERBOSE, "Removing path: %s\n", fullPath)
	return os.RemoveAll(fullPath)
}
