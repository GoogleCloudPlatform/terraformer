package discovery

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/hashicorp/errwrap"
	getter "github.com/hashicorp/go-getter"
	multierror "github.com/hashicorp/go-multierror"
	"github.com/hashicorp/terraform/httpclient"
	"github.com/hashicorp/terraform/registry"
	"github.com/hashicorp/terraform/registry/regsrc"
	"github.com/hashicorp/terraform/registry/response"
	"github.com/hashicorp/terraform/svchost/disco"
	"github.com/hashicorp/terraform/tfdiags"
	tfversion "github.com/hashicorp/terraform/version"
	"github.com/mitchellh/cli"
)

// Releases are located by querying the terraform registry.

const protocolVersionHeader = "x-terraform-protocol-version"

var httpClient *http.Client

var errVersionNotFound = errors.New("version not found")

func init() {
	httpClient = httpclient.New()

	httpGetter := &getter.HttpGetter{
		Client: httpClient,
		Netrc:  true,
	}

	getter.Getters["http"] = httpGetter
	getter.Getters["https"] = httpGetter
}

// An Installer maintains a local cache of plugins by downloading plugins
// from an online repository.
type Installer interface {
	Get(name string, req Constraints) (PluginMeta, tfdiags.Diagnostics, error)
	PurgeUnused(used map[string]PluginMeta) (removed PluginMetaSet, err error)
}

// ProviderInstaller is an Installer implementation that knows how to
// download Terraform providers from the official HashiCorp releases service
// into a local directory. The files downloaded are compliant with the
// naming scheme expected by FindPlugins, so the target directory of a
// provider installer can be used as one of several plugin discovery sources.
type ProviderInstaller struct {
	Dir string

	// Cache is used to access and update a local cache of plugins if non-nil.
	// Can be nil to disable caching.
	Cache PluginCache

	PluginProtocolVersion uint

	// OS and Arch specify the OS and architecture that should be used when
	// installing plugins. These use the same labels as the runtime.GOOS and
	// runtime.GOARCH variables respectively, and indeed the values of these
	// are used as defaults if either of these is the empty string.
	OS   string
	Arch string

	// Skip checksum and signature verification
	SkipVerify bool

	Ui cli.Ui // Ui for output

	// Services is a required *disco.Disco, which may have services and
	// credentials pre-loaded.
	Services *disco.Disco

	// registry client
	registry *registry.Client
}

// Get is part of an implementation of type Installer, and attempts to download
// and install a Terraform provider matching the given constraints.
//
// This method may return one of a number of sentinel errors from this
// package to indicate issues that are likely to be resolvable via user action:
//
//     ErrorNoSuchProvider: no provider with the given name exists in the repository.
//     ErrorNoSuitableVersion: the provider exists but no available version matches constraints.
//     ErrorNoVersionCompatible: a plugin was found within the constraints but it is
//                               incompatible with the current Terraform version.
//
// These errors should be recognized and handled as special cases by the caller
// to present a suitable user-oriented error message.
//
// All other errors indicate an internal problem that is likely _not_ solvable
// through user action, or at least not within Terraform's scope. Error messages
// are produced under the assumption that if presented to the user they will
// be presented alongside context about what is being installed, and thus the
// error messages do not redundantly include such information.
func (i *ProviderInstaller) Get(provider string, req Constraints) (PluginMeta, tfdiags.Diagnostics, error) {
	var diags tfdiags.Diagnostics

	// a little bit of initialization.
	if i.OS == "" {
		i.OS = runtime.GOOS
	}
	if i.Arch == "" {
		i.Arch = runtime.GOARCH
	}
	if i.registry == nil {
		i.registry = registry.NewClient(i.Services, nil)
	}

	// get a full listing of versions for the requested provider
	allVersions, err := i.listProviderVersions(provider)

	// TODO: return multiple errors
	if err != nil {
		log.Printf("[DEBUG] %s", err)
		if registry.IsServiceUnreachable(err) {
			registryHost, err := i.hostname()
			if err == nil && registryHost == regsrc.PublicRegistryHost.Raw {
				return PluginMeta{}, diags, ErrorPublicRegistryUnreachable
			}
			return PluginMeta{}, diags, ErrorServiceUnreachable
		}
		if registry.IsServiceNotProvided(err) {
			return PluginMeta{}, diags, err
		}
		return PluginMeta{}, diags, ErrorNoSuchProvider
	}

	// Add any warnings from the response to diags
	for _, warning := range allVersions.Warnings {
		hostname, err := i.hostname()
		if err != nil {
			return PluginMeta{}, diags, err
		}
		diag := tfdiags.SimpleWarning(fmt.Sprintf("%s: %s", hostname, warning))
		diags = diags.Append(diag)
	}

	if len(allVersions.Versions) == 0 {
		return PluginMeta{}, diags, ErrorNoSuitableVersion
	}
	providerSource := allVersions.ID

	// Filter the list of plugin versions to those which meet the version constraints
	versions := allowedVersions(allVersions, req)
	if len(versions) == 0 {
		return PluginMeta{}, diags, ErrorNoSuitableVersion
	}

	// sort them newest to oldest. The newest version wins!
	response.ProviderVersionCollection(versions).Sort()

	// if the chosen provider version does not support the requested platform,
	// filter the list of acceptable versions to those that support that platform
	if err := i.checkPlatformCompatibility(versions[0]); err != nil {
		versions = i.platformCompatibleVersions(versions)
		if len(versions) == 0 {
			return PluginMeta{}, diags, ErrorNoVersionCompatibleWithPlatform
		}
	}

	// we now have a winning platform-compatible version
	versionMeta := versions[0]
	v := VersionStr(versionMeta.Version).MustParse()

	// check protocol compatibility
	if err := i.checkPluginProtocol(versionMeta); err != nil {
		closestMatch, err := i.findClosestProtocolCompatibleVersion(allVersions.Versions)
		if err != nil {
			// No operation here if we can't find a version with compatible protocol
			return PluginMeta{}, diags, err
		}

		// Prompt version suggestion to UI based on closest protocol match
		var errMsg string
		closestVersion := VersionStr(closestMatch.Version).MustParse()
		if v.NewerThan(closestVersion) {
			errMsg = providerProtocolTooNew
		} else {
			errMsg = providerProtocolTooOld
		}

		constraintStr := req.String()
		if constraintStr == "" {
			constraintStr = "(any version)"
		}

		return PluginMeta{}, diags, errwrap.Wrap(ErrorVersionIncompatible, fmt.Errorf(fmt.Sprintf(
			errMsg, provider, v.String(), tfversion.String(),
			closestVersion.String(), closestVersion.MinorUpgradeConstraintStr(), constraintStr)))
	}

	downloadURLs, err := i.listProviderDownloadURLs(providerSource, versionMeta.Version)
	if err != nil {
		return PluginMeta{}, diags, err
	}
	providerURL := downloadURLs.DownloadURL

	if !i.SkipVerify {
		// Terraform verifies the integrity of a provider release before downloading
		// the plugin binary. The digital signature (SHA256SUMS.sig) on the
		// release distribution (SHA256SUMS) is verified with the public key of the
		// publisher provided in the Terraform Registry response, ensuring that
		// everything is as intended by the publisher. The checksum of the provider
		// plugin is expected in the SHA256SUMS file and is double checked to match
		// the checksum of the original published release to the Registry. This
		// enforces immutability of releases between the Registry and the plugin's
		// host location. Lastly, the integrity of the binary is verified upon
		// download matches the Registry and signed checksum.
		sha256, err := i.getProviderChecksum(downloadURLs)
		if err != nil {
			return PluginMeta{}, diags, err
		}

		// add the checksum parameter for go-getter to verify the download for us.
		if sha256 != "" {
			providerURL = providerURL + "?checksum=sha256:" + sha256
		}
	}

	printedProviderName := fmt.Sprintf("%q (%s)", provider, providerSource)
	i.Ui.Info(fmt.Sprintf("- Downloading plugin for provider %s %s...", printedProviderName, versionMeta.Version))
	log.Printf("[DEBUG] getting provider %s version %q", printedProviderName, versionMeta.Version)
	err = i.install(provider, v, providerURL)
	if err != nil {
		return PluginMeta{}, diags, err
	}

	// Find what we just installed
	// (This is weird, because go-getter doesn't directly return
	//  information about what was extracted, and we just extracted
	//  the archive directly into a shared dir here.)
	log.Printf("[DEBUG] looking for the %s %s plugin we just installed", provider, versionMeta.Version)
	metas := FindPlugins("provider", []string{i.Dir})
	log.Printf("[DEBUG] all plugins found %#v", metas)
	metas, _ = metas.ValidateVersions()
	metas = metas.WithName(provider).WithVersion(v)
	log.Printf("[DEBUG] filtered plugins %#v", metas)
	if metas.Count() == 0 {
		// This should never happen. Suggests that the release archive
		// contains an executable file whose name doesn't match the
		// expected convention.
		return PluginMeta{}, diags, fmt.Errorf(
			"failed to find installed plugin version %s; this is a bug in Terraform and should be reported",
			versionMeta.Version,
		)
	}

	if metas.Count() > 1 {
		// This should also never happen, and suggests that a
		// particular version was re-released with a different
		// executable filename. We consider releases as immutable, so
		// this is an error.
		return PluginMeta{}, diags, fmt.Errorf(
			"multiple plugins installed for version %s; this is a bug in Terraform and should be reported",
			versionMeta.Version,
		)
	}

	// By now we know we have exactly one meta, and so "Newest" will
	// return that one.
	return metas.Newest(), diags, nil
}

func (i *ProviderInstaller) install(provider string, version Version, url string) error {
	if i.Cache != nil {
		log.Printf("[DEBUG] looking for provider %s %s in plugin cache", provider, version)
		cached := i.Cache.CachedPluginPath("provider", provider, version)
		if cached == "" {
			log.Printf("[DEBUG] %s %s not yet in cache, so downloading %s", provider, version, url)
			err := getter.Get(i.Cache.InstallDir(), url)
			if err != nil {
				return err
			}
			// should now be in cache
			cached = i.Cache.CachedPluginPath("provider", provider, version)
			if cached == "" {
				// should never happen if the getter is behaving properly
				// and the plugins are packaged properly.
				return fmt.Errorf("failed to find downloaded plugin in cache %s", i.Cache.InstallDir())
			}
		}

		// Link or copy the cached binary into our install dir so the
		// normal resolution machinery can find it.
		filename := filepath.Base(cached)
		targetPath := filepath.Join(i.Dir, filename)
		// check if the target dir exists, and create it if not
		var err error
		if _, StatErr := os.Stat(i.Dir); os.IsNotExist(StatErr) {
			err = os.MkdirAll(i.Dir, 0700)
		}
		if err != nil {
			return err
		}

		log.Printf("[DEBUG] installing %s %s to %s from local cache %s", provider, version, targetPath, cached)

		// Delete if we can. If there's nothing there already then no harm done.
		// This is important because we can't create a link if there's
		// already a file of the same name present.
		// (any other error here we'll catch below when we try to write here)
		os.Remove(targetPath)

		// We don't attempt linking on Windows because links are not
		// comprehensively supported by all tools/apps in Windows and
		// so we choose to be conservative to avoid creating any
		// weird issues for Windows users.
		linkErr := errors.New("link not supported for Windows") // placeholder error, never actually returned
		if runtime.GOOS != "windows" {
			// Try hard linking first. Hard links are preferable because this
			// creates a self-contained directory that doesn't depend on the
			// cache after install.
			linkErr = os.Link(cached, targetPath)

			// If that failed, try a symlink. This _does_ depend on the cache
			// after install, so the user must manage the cache more carefully
			// in this case, but avoids creating redundant copies of the
			// plugins on disk.
			if linkErr != nil {
				linkErr = os.Symlink(cached, targetPath)
			}
		}

		// If we still have an error then we'll try a copy as a fallback.
		// In this case either the OS is Windows or the target filesystem
		// can't support symlinks.
		if linkErr != nil {
			srcFile, err := os.Open(cached)
			if err != nil {
				return fmt.Errorf("failed to open cached plugin %s: %s", cached, err)
			}
			defer srcFile.Close()

			destFile, err := os.OpenFile(targetPath, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, os.ModePerm)
			if err != nil {
				return fmt.Errorf("failed to create %s: %s", targetPath, err)
			}

			_, err = io.Copy(destFile, srcFile)
			if err != nil {
				destFile.Close()
				return fmt.Errorf("failed to copy cached plugin from %s to %s: %s", cached, targetPath, err)
			}

			err = destFile.Close()
			if err != nil {
				return fmt.Errorf("error creating %s: %s", targetPath, err)
			}
		}

		// One way or another, by the time we get here we should have either
		// a link or a copy of the cached plugin within i.Dir, as expected.
	} else {
		log.Printf("[DEBUG] plugin cache is disabled, so downloading %s %s from %s", provider, version, url)
		err := getter.Get(i.Dir, url)
		if err != nil {
			return err
		}
	}
	return nil
}

func (i *ProviderInstaller) PurgeUnused(used map[string]PluginMeta) (PluginMetaSet, error) {
	purge := make(PluginMetaSet)

	present := FindPlugins("provider", []string{i.Dir})
	for meta := range present {
		chosen, ok := used[meta.Name]
		if !ok {
			purge.Add(meta)
		}
		if chosen.Path != meta.Path {
			purge.Add(meta)
		}
	}

	removed := make(PluginMetaSet)
	var errs error
	for meta := range purge {
		path := meta.Path
		err := os.Remove(path)
		if err != nil {
			errs = multierror.Append(errs, fmt.Errorf(
				"failed to remove unused provider plugin %s: %s",
				path, err,
			))
		} else {
			removed.Add(meta)
		}
	}

	return removed, errs
}

func (i *ProviderInstaller) getProviderChecksum(resp *response.TerraformProviderPlatformLocation) (string, error) {
	// Get SHA256SUMS file.
	shasums, err := getFile(resp.ShasumsURL)
	if err != nil {
		log.Printf("[ERROR] error fetching checksums from %q: %s", resp.ShasumsURL, err)
		return "", ErrorMissingChecksumVerification
	}

	// Get SHA256SUMS.sig file.
	signature, err := getFile(resp.ShasumsSignatureURL)
	if err != nil {
		log.Printf("[ERROR] error fetching checksums signature from %q: %s", resp.ShasumsSignatureURL, err)
		return "", ErrorSignatureVerification
	}

	// Verify the GPG signature returned from the Registry.
	asciiArmor := resp.SigningKeys.GPGASCIIArmor()
	signer, err := verifySig(shasums, signature, asciiArmor)
	if err != nil {
		log.Printf("[ERROR] error verifying signature: %s", err)
		return "", ErrorSignatureVerification
	}

	// Also verify the GPG signature against the HashiCorp public key. This is
	// a temporary additional check until a more robust key verification
	// process is added in a future release.
	_, err = verifySig(shasums, signature, HashicorpPublicKey)
	if err != nil {
		log.Printf("[ERROR] error verifying signature against HashiCorp public key: %s", err)
		return "", ErrorSignatureVerification
	}

	// Display identity for GPG key which succeeded verifying the signature.
	// This could also be used to display to the user with i.Ui.Info().
	identities := []string{}
	for k := range signer.Identities {
		identities = append(identities, k)
	}
	identity := strings.Join(identities, ", ")
	log.Printf("[DEBUG] verified GPG signature with key from %s", identity)

	// Extract checksum for this os/arch platform binary and verify against Registry
	checksum := checksumForFile(shasums, resp.Filename)
	if checksum == "" {
		log.Printf("[ERROR] missing checksum for %s from source %s", resp.Filename, resp.ShasumsURL)
		return "", ErrorMissingChecksumVerification
	} else if checksum != resp.Shasum {
		log.Printf("[ERROR] unexpected checksum for %s from source %q", resp.Filename, resp.ShasumsURL)
		return "", ErrorChecksumVerification
	}

	return checksum, nil
}

func (i *ProviderInstaller) hostname() (string, error) {
	provider := regsrc.NewTerraformProvider("", i.OS, i.Arch)
	svchost, err := provider.SvcHost()
	if err != nil {
		return "", err
	}

	return svchost.ForDisplay(), nil
}

// list all versions available for the named provider
func (i *ProviderInstaller) listProviderVersions(name string) (*response.TerraformProviderVersions, error) {
	provider := regsrc.NewTerraformProvider(name, i.OS, i.Arch)
	versions, err := i.registry.TerraformProviderVersions(provider)
	return versions, err
}

func (i *ProviderInstaller) listProviderDownloadURLs(name, version string) (*response.TerraformProviderPlatformLocation, error) {
	urls, err := i.registry.TerraformProviderLocation(regsrc.NewTerraformProvider(name, i.OS, i.Arch), version)
	if urls == nil {
		return nil, fmt.Errorf("No download urls found for provider %s", name)
	}
	return urls, err
}

// findClosestProtocolCompatibleVersion searches for the provider version with the closest protocol match.
// Prerelease versions are filtered.
func (i *ProviderInstaller) findClosestProtocolCompatibleVersion(versions []*response.TerraformProviderVersion) (*response.TerraformProviderVersion, error) {
	// Loop through all the provider versions to find the earliest and latest
	// versions that match the installer protocol to then select the closest of the two
	var latest, earliest *response.TerraformProviderVersion
	for _, version := range versions {
		// Prereleases are filtered and will not be suggested
		v, err := VersionStr(version.Version).Parse()
		if err != nil || v.IsPrerelease() {
			continue
		}

		if err := i.checkPluginProtocol(version); err == nil {
			if earliest == nil {
				// Found the first provider version with compatible protocol
				earliest = version
			}
			// Update the latest protocol compatible version
			latest = version
		}
	}
	if earliest == nil {
		// No compatible protocol was found for any version
		return nil, ErrorNoVersionCompatible
	}

	// Convert protocols to comparable types
	protoString := strconv.Itoa(int(i.PluginProtocolVersion))
	protocolVersion, err := VersionStr(protoString).Parse()
	if err != nil {
		return nil, fmt.Errorf("invalid plugin protocol version: %q", i.PluginProtocolVersion)
	}

	earliestVersionProtocol, err := VersionStr(earliest.Protocols[0]).Parse()
	if err != nil {
		return nil, err
	}

	// Compare installer protocol version with the first protocol listed of the earliest match
	// [A, B] where A is assumed the earliest compatible major version of the protocol pair
	if protocolVersion.NewerThan(earliestVersionProtocol) {
		// Provider protocols are too old, the closest version is the earliest compatible version
		return earliest, nil
	}

	// Provider protocols are too new, the closest version is the latest compatible version
	return latest, nil
}

func (i *ProviderInstaller) checkPluginProtocol(versionMeta *response.TerraformProviderVersion) error {
	// TODO: should this be a different error? We should probably differentiate between
	// no compatible versions and no protocol versions listed at all
	if len(versionMeta.Protocols) == 0 {
		return fmt.Errorf("no plugin protocol versions listed")
	}

	protoString := strconv.Itoa(int(i.PluginProtocolVersion))
	protocolVersion, err := VersionStr(protoString).Parse()
	if err != nil {
		return fmt.Errorf("invalid plugin protocol version: %q", i.PluginProtocolVersion)
	}
	protocolConstraint, err := protocolVersion.MinorUpgradeConstraintStr().Parse()
	if err != nil {
		// This should not fail if the preceding function succeeded.
		return fmt.Errorf("invalid plugin protocol version: %q", protocolVersion.String())
	}

	for _, p := range versionMeta.Protocols {
		proPro, err := VersionStr(p).Parse()
		if err != nil {
			// invalid protocol reported by the registry. Move along.
			log.Printf("[WARN] invalid provider protocol version %q found in the registry", versionMeta.Version)
			continue
		}
		// success!
		if protocolConstraint.Allows(proPro) {
			return nil
		}
	}

	return ErrorNoVersionCompatible
}

// REVIEWER QUESTION (again): this ends up swallowing a bunch of errors from
// checkPluginProtocol. Do they need to be percolated up better, or would
// debug messages would suffice in these situations?
func (i *ProviderInstaller) findPlatformCompatibleVersion(versions []*response.TerraformProviderVersion) (*response.TerraformProviderVersion, error) {
	for _, version := range versions {
		if err := i.checkPlatformCompatibility(version); err == nil {
			return version, nil
		}
	}

	return nil, ErrorNoVersionCompatibleWithPlatform
}

// platformCompatibleVersions returns a list of provider versions that are
// compatible with the requested platform.
func (i *ProviderInstaller) platformCompatibleVersions(versions []*response.TerraformProviderVersion) []*response.TerraformProviderVersion {
	var v []*response.TerraformProviderVersion
	for _, version := range versions {
		if err := i.checkPlatformCompatibility(version); err == nil {
			v = append(v, version)
		}
	}
	return v
}

func (i *ProviderInstaller) checkPlatformCompatibility(versionMeta *response.TerraformProviderVersion) error {
	if len(versionMeta.Platforms) == 0 {
		return fmt.Errorf("no supported provider platforms listed")
	}
	for _, p := range versionMeta.Platforms {
		if p.Arch == i.Arch && p.OS == i.OS {
			return nil
		}
	}
	return fmt.Errorf("version %s does not support the requested platform %s_%s", versionMeta.Version, i.OS, i.Arch)
}

// take the list of available versions for a plugin, and filter out those that
// don't fit the constraints.
func allowedVersions(available *response.TerraformProviderVersions, required Constraints) []*response.TerraformProviderVersion {
	var allowed []*response.TerraformProviderVersion

	for _, v := range available.Versions {
		version, err := VersionStr(v.Version).Parse()
		if err != nil {
			log.Printf("[WARN] invalid version found for %q: %s", available.ID, err)
			continue
		}
		if required.Allows(version) {
			allowed = append(allowed, v)
		}
	}
	return allowed
}

func checksumForFile(sums []byte, name string) string {
	for _, line := range strings.Split(string(sums), "\n") {
		parts := strings.Fields(line)
		if len(parts) > 1 && parts[1] == name {
			return parts[0]
		}
	}
	return ""
}

func getFile(url string) ([]byte, error) {
	resp, err := httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s", resp.Status)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return data, err
	}
	return data, nil
}

// providerProtocolTooOld is a message sent to the CLI UI if the provider's
// supported protocol versions are too old for the user's version of terraform,
// but an older version of the provider is compatible.
const providerProtocolTooOld = `
[reset][bold][red]Provider %q v%s is not compatible with Terraform %s.[reset][red]

Provider version %s is the earliest compatible version. Select it with 
the following version constraint:

	version = %q

Terraform checked all of the plugin versions matching the given constraint:
    %s

Consult the documentation for this provider for more information on
compatibility between provider and Terraform versions.
`

// providerProtocolTooNew is a message sent to the CLI UI if the provider's
// supported protocol versions are too new for the user's version of terraform,
// and the user could either upgrade terraform or choose an older version of the
// provider
const providerProtocolTooNew = `
[reset][bold][red]Provider %q v%s is not compatible with Terraform %s.[reset][red]

Provider version %s is the latest compatible version. Select it with 
the following constraint:

    version = %q

Terraform checked all of the plugin versions matching the given constraint:
    %s

Consult the documentation for this provider for more information on
compatibility between provider and Terraform versions.

Alternatively, upgrade to the latest version of Terraform for compatibility with newer provider releases.
`
