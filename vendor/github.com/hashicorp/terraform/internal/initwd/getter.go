package initwd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	cleanhttp "github.com/hashicorp/go-cleanhttp"
	getter "github.com/hashicorp/go-getter"
	"github.com/hashicorp/terraform/registry/regsrc"
)

// We configure our own go-getter detector and getter sets here, because
// the set of sources we support is part of Terraform's documentation and
// so we don't want any new sources introduced in go-getter to sneak in here
// and work even though they aren't documented. This also insulates us from
// any meddling that might be done by other go-getter callers linked into our
// executable.

var goGetterDetectors = []getter.Detector{
	new(getter.GitHubDetector),
	new(getter.GitDetector),
	new(getter.BitBucketDetector),
	new(getter.GCSDetector),
	new(getter.S3Detector),
	new(getter.FileDetector),
}

var goGetterNoDetectors = []getter.Detector{}

var goGetterDecompressors = map[string]getter.Decompressor{
	"bz2": new(getter.Bzip2Decompressor),
	"gz":  new(getter.GzipDecompressor),
	"xz":  new(getter.XzDecompressor),
	"zip": new(getter.ZipDecompressor),

	"tar.bz2":  new(getter.TarBzip2Decompressor),
	"tar.tbz2": new(getter.TarBzip2Decompressor),

	"tar.gz": new(getter.TarGzipDecompressor),
	"tgz":    new(getter.TarGzipDecompressor),

	"tar.xz": new(getter.TarXzDecompressor),
	"txz":    new(getter.TarXzDecompressor),
}

var goGetterGetters = map[string]getter.Getter{
	"file":  new(getter.FileGetter),
	"gcs":   new(getter.GCSGetter),
	"git":   new(getter.GitGetter),
	"hg":    new(getter.HgGetter),
	"s3":    new(getter.S3Getter),
	"http":  getterHTTPGetter,
	"https": getterHTTPGetter,
}

var getterHTTPClient = cleanhttp.DefaultClient()

var getterHTTPGetter = &getter.HttpGetter{
	Client: getterHTTPClient,
	Netrc:  true,
}

// A reusingGetter is a helper for the module installer that remembers
// the final resolved addresses of all of the sources it has already been
// asked to install, and will copy from a prior installation directory if
// it has the same resolved source address.
//
// The keys in a reusingGetter are resolved and trimmed source addresses
// (with a scheme always present, and without any "subdir" component),
// and the values are the paths where each source was previously installed.
type reusingGetter map[string]string

// getWithGoGetter retrieves the package referenced in the given address
// into the installation path and then returns the full path to any subdir
// indicated in the address.
//
// The errors returned by this function are those surfaced by the underlying
// go-getter library, which have very inconsistent quality as
// end-user-actionable error messages. At this time we do not have any
// reasonable way to improve these error messages at this layer because
// the underlying errors are not separately recognizable.
func (g reusingGetter) getWithGoGetter(instPath, addr string) (string, error) {
	packageAddr, subDir := splitAddrSubdir(addr)

	log.Printf("[DEBUG] will download %q to %s", packageAddr, instPath)

	realAddr, err := getter.Detect(packageAddr, instPath, goGetterDetectors)
	if err != nil {
		return "", err
	}

	if isMaybeRelativeLocalPath(realAddr) {
		return "", &MaybeRelativePathErr{addr}
	}

	var realSubDir string
	realAddr, realSubDir = splitAddrSubdir(realAddr)
	if realSubDir != "" {
		subDir = filepath.Join(realSubDir, subDir)
	}

	if realAddr != packageAddr {
		log.Printf("[TRACE] go-getter detectors rewrote %q to %q", packageAddr, realAddr)
	}

	if prevDir, exists := g[realAddr]; exists {
		log.Printf("[TRACE] copying previous install %s to %s", prevDir, instPath)
		err := os.Mkdir(instPath, os.ModePerm)
		if err != nil {
			return "", fmt.Errorf("failed to create directory %s: %s", instPath, err)
		}
		err = copyDir(instPath, prevDir)
		if err != nil {
			return "", fmt.Errorf("failed to copy from %s to %s: %s", prevDir, instPath, err)
		}
	} else {
		log.Printf("[TRACE] fetching %q to %q", realAddr, instPath)
		client := getter.Client{
			Src: realAddr,
			Dst: instPath,
			Pwd: instPath,

			Mode: getter.ClientModeDir,

			Detectors:     goGetterNoDetectors, // we already did detection above
			Decompressors: goGetterDecompressors,
			Getters:       goGetterGetters,
		}
		err = client.Get()
		if err != nil {
			return "", err
		}
		// Remember where we installed this so we might reuse this directory
		// on subsequent calls to avoid re-downloading.
		g[realAddr] = instPath
	}

	// Our subDir string can contain wildcards until this point, so that
	// e.g. a subDir of * can expand to one top-level directory in a .tar.gz
	// archive. Now that we've expanded the archive successfully we must
	// resolve that into a concrete path.
	var finalDir string
	if subDir != "" {
		finalDir, err = getter.SubdirGlob(instPath, subDir)
		log.Printf("[TRACE] expanded %q to %q", subDir, finalDir)
		if err != nil {
			return "", err
		}
	} else {
		finalDir = instPath
	}

	// If we got this far then we have apparently succeeded in downloading
	// the requested object!
	return filepath.Clean(finalDir), nil
}

// splitAddrSubdir splits the given address (which is assumed to be a
// registry address or go-getter-style address) into a package portion
// and a sub-directory portion.
//
// The package portion defines what should be downloaded and then the
// sub-directory portion, if present, specifies a sub-directory within
// the downloaded object (an archive, VCS repository, etc) that contains
// the module's configuration files.
//
// The subDir portion will be returned as empty if no subdir separator
// ("//") is present in the address.
func splitAddrSubdir(addr string) (packageAddr, subDir string) {
	return getter.SourceDirSubdir(addr)
}

var localSourcePrefixes = []string{
	"./",
	"../",
	".\\",
	"..\\",
}

func isLocalSourceAddr(addr string) bool {
	for _, prefix := range localSourcePrefixes {
		if strings.HasPrefix(addr, prefix) {
			return true
		}
	}
	return false
}

func isRegistrySourceAddr(addr string) bool {
	_, err := regsrc.ParseModuleSource(addr)
	return err == nil
}

type MaybeRelativePathErr struct {
	Addr string
}

func (e *MaybeRelativePathErr) Error() string {
	return fmt.Sprintf("Terraform cannot determine the module source for %s", e.Addr)
}

func isMaybeRelativeLocalPath(addr string) bool {
	if strings.HasPrefix(addr, "file://") {
		_, err := os.Stat(addr[7:])
		if err != nil {
			return true
		}
	}
	return false
}
