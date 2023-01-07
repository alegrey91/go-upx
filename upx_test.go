package goupx_test

import (
	"fmt"
	"testing"

	goupx "github.com/alegrey91/go-upx"
)

func TestCompress(t *testing.T) {
	testCases := []struct {
		file           string
		intensity      int
		options        goupx.Options
		expectedResult bool
	}{
		{
			"samples/sample",
			0,
			goupx.Options{
				Output:  "/tmp/file",
				Force:   true,
				Verbose: false,
				CompressionTuningOpt: goupx.CompressionTuningOptions{
					Brute: 0,
				},
				BackupOpt: goupx.BackupOptions{
					Backup: false,
				},
			},
			true,
		},
		{
			"samples/sample",
			1,
			goupx.Options{
				Output:  "/tmp/file",
				Force:   true,
				Verbose: false,
				CompressionTuningOpt: goupx.CompressionTuningOptions{
					Brute: 0,
				},
				BackupOpt: goupx.BackupOptions{
					Backup: false,
				},
			},
			true,
		},
		{
			"samples/sample",
			19,
			goupx.Options{
				Output:  "/tmp/file",
				Force:   true,
				Verbose: false,
				CompressionTuningOpt: goupx.CompressionTuningOptions{
					Brute: 0,
				},
				BackupOpt: goupx.BackupOptions{
					Backup: false,
				},
			},
			true,
		},
		{
			"samples/sample",
			1,
			goupx.Options{
				Output:  "/tmp/file",
				Force:   false,
				Verbose: false,
				CompressionTuningOpt: goupx.CompressionTuningOptions{
					Brute: 0,
				},
				BackupOpt: goupx.BackupOptions{
					Backup: false,
				},
			},
			false,
		},
		{
			"samples/sample",
			1,
			goupx.Options{
				Output:  "/tmp/file",
				Force:   true,
				Verbose: false,
				CompressionTuningOpt: goupx.CompressionTuningOptions{
					Brute: 1,
				},
				BackupOpt: goupx.BackupOptions{
					Backup: false,
				},
			},
			true,
		},
		{
			"samples/sample",
			1,
			goupx.Options{
				Output:  "/tmp/file",
				Force:   true,
				Verbose: false,
				CompressionTuningOpt: goupx.CompressionTuningOptions{
					Brute: 2,
				},
				BackupOpt: goupx.BackupOptions{
					Backup: false,
				},
			},
			true,
		},
		{
			"samples/sample",
			1,
			goupx.Options{
				Output:  "/tmp/file",
				Force:   true,
				Verbose: false,
				CompressionTuningOpt: goupx.CompressionTuningOptions{
					Brute: 0,
				},
				BackupOpt: goupx.BackupOptions{
					Backup: true,
				},
			},
			true,
		},
		{
			"samples/sample",
			1,
			goupx.Options{
				Output:  "/this/path/does/not/exists",
				Force:   true,
				Verbose: false,
				CompressionTuningOpt: goupx.CompressionTuningOptions{
					Brute: 0,
				},
				BackupOpt: goupx.BackupOptions{
					Backup: true,
				},
			},
			false,
		},
		{
			"",
			1,
			goupx.Options{
				Output:  "/tmp/file",
				Force:   true,
				Verbose: false,
				CompressionTuningOpt: goupx.CompressionTuningOptions{
					Brute: 0,
				},
				BackupOpt: goupx.BackupOptions{
					Backup: true,
				},
			},
			false,
		},
		{
			"samples/sample",
			1,
			goupx.Options{
				Output:  "",
				Force:   true,
				Verbose: false,
				CompressionTuningOpt: goupx.CompressionTuningOptions{
					Brute: 0,
				},
				BackupOpt: goupx.BackupOptions{
					Backup: true,
				},
			},
			true,
		},
	}

	upx := goupx.NewUPX()

	for id, tt := range testCases {
		t.Run(fmt.Sprintf("Checking rule with id %d", id), func(t *testing.T) {
			testResult, testErr := upx.Compress(tt.file, tt.intensity, tt.options)
			t.Logf("args: %s", upx.GetArgs())
			if testErr != nil {
				t.Logf("%v", testErr)
			}

			if testResult != tt.expectedResult {
				t.Fatal("Test failed")
			}
		})
	}
}
