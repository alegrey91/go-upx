package goupx

import (
	"fmt"
	"io"
	"os/exec"
)

// Read description for Options struct.
type BackupOptions struct {
	// Backup keep backup files
	Backup bool
}

// Read description for Options struct.
type CompressionTuningOptions struct {
	// Brute try all available compression methods, filters and variants
	Brute int
}

// Options describe the command args provided by the upx utility.
// Here's an example:
// ---
// Options:
//  -q     be quiet                          -v    be verbose
//  -oFILE write output to 'FILE'
//  -f     force compression of suspicious files
//  --no-color, --mono, --color, --no-progress   change look
//
// Compression tuning options:
//  --brute             try all available compression methods & filters [slow]
//  --ultra-brute       try even more compression variants [very slow]
//
// Backup options:
//  -k, --backup        keep backup files
//  --no-backup         no backup files [default]
type Options struct {
	// Output write output to 'FILE'
	Output string

	// Force force compression of suspicious files
	Force bool

	// Verbose enable/disable verbosity
	Verbose bool

	// CompressionTuningOpt keep track of CompressionTuningOptions
	CompressionTuningOpt CompressionTuningOptions

	// BackupOpt keep track of BackupOptions
	BackupOpt BackupOptions
}

// generateCommandArgs generate the command args for upx starting from the given Options struct.
// It return the list of args.
func (opts *Options) generateCommandArgs() []string {
	var commandOptions []string

	if opts.Output != "" {
		commandOptions = append(commandOptions, "-o", opts.Output)
	}

	if opts.Force {
		commandOptions = append(commandOptions, "--force")
	}

	if !opts.Verbose {
		commandOptions = append(commandOptions, "--quiet")
	}

	switch opts.CompressionTuningOpt.Brute {
	case 0:
	case 1:
		commandOptions = append(commandOptions, "--brute")
	case 2:
		commandOptions = append(commandOptions, "--ultra-brute")
	default:
	}

	if opts.BackupOpt.Backup {
		commandOptions = append(commandOptions, "--backup")
	}

	return commandOptions
}

type CmdExecution struct {
	// stderr, stdout keep track of outputs
	stdout, stderr []byte

	// parsed output
	parsedOutput

	// parsed error
	parsedError

	// exit status code
	exitStatus int
}

// UPX describe the execution of upx utility.
type UPX struct {
	// Binary file for upx (it is visible in case you have different binary name)
	Binary string

	// Args contains the command arguments
	args []string

	// UPX command execution results
	CmdExecution
}

func NewUPX() *UPX {
	return &UPX{
		Binary: "upx",
	}
}

// GetArgs return the list of command args.
func (upx *UPX) GetArgs() []string {
	return upx.args
}

// run execute the upx command under the hood.
// It return an error in case of fail.
func (upx *UPX) run(file string, cmdArgs []string) error {
	binPath, err := exec.LookPath(upx.Binary)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	cmdArgs = append(cmdArgs, file)
	upx.args = cmdArgs

	cmd := exec.Command(binPath, cmdArgs...)

	// preparing pipe to collect stderr
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	// preparing pipe to collect stdout
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	if err := cmd.Start(); err != nil {
		return fmt.Errorf("%v", err)
	}

	// collect stderr
	upx.stderr, _ = io.ReadAll(stderr)

	// collect stdout
	upx.stdout, _ = io.ReadAll(stdout)

	if err := cmd.Wait(); err != nil {
		upx.exitStatus = cmd.ProcessState.ExitCode()
		return fmt.Errorf("%v", err)
	}

	return nil
}

// Compress execute a compression with upx.
// It return false with an error message in case of fail, true with nil otherwhise.
func (upx *UPX) Compress(file string, intensity int, options Options) (bool, error) {
	var command []string

	// input sanitizing
	if intensity < 1 {
		intensity = 1
	}
	if intensity > 9 {
		intensity = 9
	}
	command = append(command, fmt.Sprint("-", intensity))
	command = append(options.generateCommandArgs(), command[0])

	err := upx.run(file, command)
	if err != nil {
		upx.parsedError = parseError(upx.stderr)
		return false, fmt.Errorf("%v", upx.parsedError.GetErrorMessage())
	}
	upx.parsedOutput = parseOutput(upx.stdout)
	return true, nil
}

// Decompress execute a decompression with upx.
// It return false with an error message in case of fail, true with nil otherwhise.
func (upx *UPX) Decompress(file string, options Options) (bool, error) {
	var command []string
	command = append(command, "-d")
	command = append(options.generateCommandArgs(), command[0])

	err := upx.run(file, command)
	if err != nil {
		upx.parsedError = parseError(upx.stderr)
		return false, fmt.Errorf("%v", upx.parsedError.GetErrorMessage())
	}
	upx.parsedOutput = parseOutput(upx.stdout)
	return true, nil
}

// TestCompressedFile execute test with upx.
// It return false with an error message in case it fail, true with nil otherwise.
func (upx *UPX) TestCompressedFile(file string) (bool, error) {
	var command []string
	command = append(command, "-t")

	err := upx.run(file, command)
	if err != nil {
		upx.parsedError = parseError(upx.stderr)
		return false, fmt.Errorf("%v", upx.parsedError.GetErrorMessage())
	}
	return true, nil
}

// TODO
//func (upx *UPX) ListCompressedFile(file string, options Options) ([]string, error) {}
