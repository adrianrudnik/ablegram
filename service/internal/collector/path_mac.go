//go:build darwin

package collector

import (
	"fmt"
	"os"
)

func enrichPlatformExcludePaths() {
	h, err := os.UserHomeDir()
	if err != nil {
		Logger.Warn().Err(err).Msg("Could not get user home dir on darwin")
	}

	excludePaths = append(excludePaths, []string{
		fmt.Sprintf("%s/Library", h),
		fmt.Sprintf("%s/Applications", h),
		fmt.Sprintf("%s/Pictures/Photos Library.photoslibrary", h),
	}...)
}
