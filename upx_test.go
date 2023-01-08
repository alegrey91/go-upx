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
			"samples/true",
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
			"samples/true",
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
			"samples/true",
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
			"samples/true",
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
			"samples/true",
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
			"samples/true",
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
			"samples/true",
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
			"samples/true",
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

func TestDecompress(t *testing.T) {

	testCases := []struct {
		file           string
		options        goupx.Options
		expectedResult bool
	}{
		{
			"samples/true_compressed",
			goupx.Options{
				Output:  "/tmp/file",
				Force:   true,
				Verbose: false,
				BackupOpt: goupx.BackupOptions{
					Backup: false,
				},
			},
			true,
		},
		{
			"samples/true_compressed",
			goupx.Options{
				Output:  "/tmp/file",
				Force:   false,
				Verbose: false,
				BackupOpt: goupx.BackupOptions{
					Backup: false,
				},
			},
			false,
		},
		{
			"samples/true_compressed",
			goupx.Options{
				Output:  "/tmp/file",
				Force:   true,
				Verbose: false,
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
			testResult, testErr := upx.Decompress(tt.file, tt.options)
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

func TestTestCompressedFile(t *testing.T) {

	testCases := []struct {
		file           string
		expectedResult bool
	}{
		{
			"samples/true_compressed",
			true,
		},
		{
			"samples/true",
			false,
		},
	}
	upx := goupx.NewUPX()

	for id, tt := range testCases {
		t.Run(fmt.Sprintf("Checking rule with id %d", id), func(t *testing.T) {
			testResult, testErr := upx.TestCompressedFile(tt.file)
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