package configutil

import "os"

func AppendLine(config, line string) error {
	f, err := os.OpenFile(config,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := f.WriteString(line + "\n"); err != nil {
		return err
	}

	return nil
}
