package perform

import (
	"fmt"
	"github.com/fatih/color"
	"os/exec"
)

type Option func(o *options)

type options struct {
	dry bool
	dir string
}

func newOptions() options {
	return options{}
}

func Dry(o *options) Option {
	return func(o *options) {
		o.dry = true
	}
}

func Dir(dir string) Option {
	return func(o *options) {
		o.dir = dir
	}
}

func Command(command string, args []string, oo ...Option) error {
	// make the options first
	var opts = newOptions()
	for _, o := range oo {
		o(&opts)
	}

	cmd := exec.Command(command, args...)
	if opts.dir != "" {
		cmd.Dir = opts.dir
	}

	fmt.Println("Running: " + color.BlueString(cmd.String()))
	// return early if we only want dry run
	if opts.dry {
		return nil
	}

	bb, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	fmt.Println(color.YellowString(string(bb)))

	return nil
}
