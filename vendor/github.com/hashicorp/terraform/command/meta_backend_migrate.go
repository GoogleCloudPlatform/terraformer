package command

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/hashicorp/terraform/backend"
	"github.com/hashicorp/terraform/command/clistate"
	"github.com/hashicorp/terraform/state"
	"github.com/hashicorp/terraform/states"
	"github.com/hashicorp/terraform/states/statemgr"
	"github.com/hashicorp/terraform/terraform"
)

type backendMigrateOpts struct {
	OneType, TwoType string
	One, Two         backend.Backend

	// Fields below are set internally when migrate is called

	oneEnv string // source env
	twoEnv string // dest env
	force  bool   // if true, won't ask for confirmation
}

// backendMigrateState handles migrating (copying) state from one backend
// to another. This function handles asking the user for confirmation
// as well as the copy itself.
//
// This function can handle all scenarios of state migration regardless
// of the existence of state in either backend.
//
// After migrating the state, the existing state in the first backend
// remains untouched.
//
// This will attempt to lock both states for the migration.
func (m *Meta) backendMigrateState(opts *backendMigrateOpts) error {
	log.Printf("[TRACE] backendMigrateState: need to migrate from %q to %q backend config", opts.OneType, opts.TwoType)
	// We need to check what the named state status is. If we're converting
	// from multi-state to single-state for example, we need to handle that.
	var oneSingle, twoSingle bool
	oneStates, err := opts.One.Workspaces()
	if err == backend.ErrWorkspacesNotSupported {
		oneSingle = true
		err = nil
	}
	if err != nil {
		return fmt.Errorf(strings.TrimSpace(
			errMigrateLoadStates), opts.OneType, err)
	}

	_, err = opts.Two.Workspaces()
	if err == backend.ErrWorkspacesNotSupported {
		twoSingle = true
		err = nil
	}
	if err != nil {
		return fmt.Errorf(strings.TrimSpace(
			errMigrateLoadStates), opts.TwoType, err)
	}

	// Setup defaults
	opts.oneEnv = backend.DefaultStateName
	opts.twoEnv = backend.DefaultStateName
	opts.force = m.forceInitCopy

	// Determine migration behavior based on whether the source/destination
	// supports multi-state.
	switch {
	// Single-state to single-state. This is the easiest case: we just
	// copy the default state directly.
	case oneSingle && twoSingle:
		return m.backendMigrateState_s_s(opts)

	// Single-state to multi-state. This is easy since we just copy
	// the default state and ignore the rest in the destination.
	case oneSingle && !twoSingle:
		return m.backendMigrateState_s_s(opts)

	// Multi-state to single-state. If the source has more than the default
	// state this is complicated since we have to ask the user what to do.
	case !oneSingle && twoSingle:
		// If the source only has one state and it is the default,
		// treat it as if it doesn't support multi-state.
		if len(oneStates) == 1 && oneStates[0] == backend.DefaultStateName {
			return m.backendMigrateState_s_s(opts)
		}

		return m.backendMigrateState_S_s(opts)

	// Multi-state to multi-state. We merge the states together (migrating
	// each from the source to the destination one by one).
	case !oneSingle && !twoSingle:
		// If the source only has one state and it is the default,
		// treat it as if it doesn't support multi-state.
		if len(oneStates) == 1 && oneStates[0] == backend.DefaultStateName {
			return m.backendMigrateState_s_s(opts)
		}

		return m.backendMigrateState_S_S(opts)
	}

	return nil
}

//-------------------------------------------------------------------
// State Migration Scenarios
//
// The functions below cover handling all the various scenarios that
// can exist when migrating state. They are named in an immediately not
// obvious format but is simple:
//
// Format: backendMigrateState_s1_s2[_suffix]
//
// When s1 or s2 is lower case, it means that it is a single state backend.
// When either is uppercase, it means that state is a multi-state backend.
// The suffix is used to disambiguate multiple cases with the same type of
// states.
//
//-------------------------------------------------------------------

// Multi-state to multi-state.
func (m *Meta) backendMigrateState_S_S(opts *backendMigrateOpts) error {
	log.Print("[TRACE] backendMigrateState: migrating all named workspaces")

	// Ask the user if they want to migrate their existing remote state
	migrate, err := m.confirm(&terraform.InputOpts{
		Id: "backend-migrate-multistate-to-multistate",
		Query: fmt.Sprintf(
			"Do you want to migrate all workspaces to %q?",
			opts.TwoType),
		Description: fmt.Sprintf(
			strings.TrimSpace(inputBackendMigrateMultiToMulti),
			opts.OneType, opts.TwoType),
	})
	if err != nil {
		return fmt.Errorf(
			"Error asking for state migration action: %s", err)
	}
	if !migrate {
		return fmt.Errorf("Migration aborted by user.")
	}

	// Read all the states
	oneStates, err := opts.One.Workspaces()
	if err != nil {
		return fmt.Errorf(strings.TrimSpace(
			errMigrateLoadStates), opts.OneType, err)
	}

	// Sort the states so they're always copied alphabetically
	sort.Strings(oneStates)

	// Go through each and migrate
	for _, name := range oneStates {
		// Copy the same names
		opts.oneEnv = name
		opts.twoEnv = name

		// Force it, we confirmed above
		opts.force = true

		// Perform the migration
		if err := m.backendMigrateState_s_s(opts); err != nil {
			return fmt.Errorf(strings.TrimSpace(
				errMigrateMulti), name, opts.OneType, opts.TwoType, err)
		}
	}

	return nil
}

// Multi-state to single state.
func (m *Meta) backendMigrateState_S_s(opts *backendMigrateOpts) error {
	log.Printf("[TRACE] backendMigrateState: target backend type %q does not support named workspaces", opts.TwoType)

	currentEnv := m.Workspace()

	migrate := opts.force
	if !migrate {
		var err error
		// Ask the user if they want to migrate their existing remote state
		migrate, err = m.confirm(&terraform.InputOpts{
			Id: "backend-migrate-multistate-to-single",
			Query: fmt.Sprintf(
				"Destination state %q doesn't support workspaces.\n"+
					"Do you want to copy only your current workspace?",
				opts.TwoType),
			Description: fmt.Sprintf(
				strings.TrimSpace(inputBackendMigrateMultiToSingle),
				opts.OneType, opts.TwoType, currentEnv),
		})
		if err != nil {
			return fmt.Errorf(
				"Error asking for state migration action: %s", err)
		}
	}

	if !migrate {
		return fmt.Errorf("Migration aborted by user.")
	}

	// Copy the default state
	opts.oneEnv = currentEnv

	// now switch back to the default env so we can acccess the new backend
	m.SetWorkspace(backend.DefaultStateName)

	return m.backendMigrateState_s_s(opts)
}

// Single state to single state, assumed default state name.
func (m *Meta) backendMigrateState_s_s(opts *backendMigrateOpts) error {
	log.Printf("[TRACE] backendMigrateState: migrating %q workspace to %q workspace", opts.oneEnv, opts.twoEnv)

	stateOne, err := opts.One.StateMgr(opts.oneEnv)
	if err != nil {
		return fmt.Errorf(strings.TrimSpace(
			errMigrateSingleLoadDefault), opts.OneType, err)
	}
	if err := stateOne.RefreshState(); err != nil {
		return fmt.Errorf(strings.TrimSpace(
			errMigrateSingleLoadDefault), opts.OneType, err)
	}

	// Do not migrate workspaces without state.
	if stateOne.State().Empty() {
		log.Print("[TRACE] backendMigrateState: source workspace has empty state, so nothing to migrate")
		return nil
	}

	stateTwo, err := opts.Two.StateMgr(opts.twoEnv)
	if err == backend.ErrDefaultWorkspaceNotSupported {
		// If the backend doesn't support using the default state, we ask the user
		// for a new name and migrate the default state to the given named state.
		stateTwo, err = func() (statemgr.Full, error) {
			log.Print("[TRACE] backendMigrateState: target doesn't support a default workspace, so we must prompt for a new name")
			name, err := m.UIInput().Input(context.Background(), &terraform.InputOpts{
				Id: "new-state-name",
				Query: fmt.Sprintf(
					"[reset][bold][yellow]The %q backend configuration only allows "+
						"named workspaces![reset]",
					opts.TwoType),
				Description: strings.TrimSpace(inputBackendNewWorkspaceName),
			})
			if err != nil {
				return nil, fmt.Errorf("Error asking for new state name: %s", err)
			}

			// Update the name of the target state.
			opts.twoEnv = name

			stateTwo, err := opts.Two.StateMgr(opts.twoEnv)
			if err != nil {
				return nil, err
			}

			// If the currently selected workspace is the default workspace, then set
			// the named workspace as the new selected workspace.
			if m.Workspace() == backend.DefaultStateName {
				if err := m.SetWorkspace(opts.twoEnv); err != nil {
					return nil, fmt.Errorf("Failed to set new workspace: %s", err)
				}
			}

			return stateTwo, nil
		}()
	}
	if err != nil {
		return fmt.Errorf(strings.TrimSpace(
			errMigrateSingleLoadDefault), opts.TwoType, err)
	}
	if err := stateTwo.RefreshState(); err != nil {
		return fmt.Errorf(strings.TrimSpace(
			errMigrateSingleLoadDefault), opts.TwoType, err)
	}

	// Check if we need migration at all.
	// This is before taking a lock, because they may also correspond to the same lock.
	one := stateOne.State()
	two := stateTwo.State()

	// no reason to migrate if the state is already there
	if one.Equal(two) {
		// Equal isn't identical; it doesn't check lineage.
		sm1, _ := stateOne.(statemgr.PersistentMeta)
		sm2, _ := stateTwo.(statemgr.PersistentMeta)
		if one != nil && two != nil {
			if sm1 == nil || sm2 == nil {
				log.Print("[TRACE] backendMigrateState: both source and destination workspaces have no state, so no migration is needed")
				return nil
			}
			if sm1.StateSnapshotMeta().Lineage == sm2.StateSnapshotMeta().Lineage {
				log.Printf("[TRACE] backendMigrateState: both source and destination workspaces have equal state with lineage %q, so no migration is needed", sm1.StateSnapshotMeta().Lineage)
				return nil
			}
		}
	}

	if m.stateLock {
		lockCtx := context.Background()

		lockerOne := clistate.NewLocker(lockCtx, m.stateLockTimeout, m.Ui, m.Colorize())
		if err := lockerOne.Lock(stateOne, "migration source state"); err != nil {
			return fmt.Errorf("Error locking source state: %s", err)
		}
		defer lockerOne.Unlock(nil)

		lockerTwo := clistate.NewLocker(lockCtx, m.stateLockTimeout, m.Ui, m.Colorize())
		if err := lockerTwo.Lock(stateTwo, "migration destination state"); err != nil {
			return fmt.Errorf("Error locking destination state: %s", err)
		}
		defer lockerTwo.Unlock(nil)

		// We now own a lock, so double check that we have the version
		// corresponding to the lock.
		log.Print("[TRACE] backendMigrateState: refreshing source workspace state")
		if err := stateOne.RefreshState(); err != nil {
			return fmt.Errorf(strings.TrimSpace(
				errMigrateSingleLoadDefault), opts.OneType, err)
		}
		log.Print("[TRACE] backendMigrateState: refreshing target workspace state")
		if err := stateTwo.RefreshState(); err != nil {
			return fmt.Errorf(strings.TrimSpace(
				errMigrateSingleLoadDefault), opts.OneType, err)
		}

		one = stateOne.State()
		two = stateTwo.State()
	}

	var confirmFunc func(state.State, state.State, *backendMigrateOpts) (bool, error)
	switch {
	// No migration necessary
	case one.Empty() && two.Empty():
		log.Print("[TRACE] backendMigrateState: both source and destination workspaces have empty state, so no migration is required")
		return nil

	// No migration necessary if we're inheriting state.
	case one.Empty() && !two.Empty():
		log.Print("[TRACE] backendMigrateState: source workspace has empty state, so no migration is required")
		return nil

	// We have existing state moving into no state. Ask the user if
	// they'd like to do this.
	case !one.Empty() && two.Empty():
		log.Print("[TRACE] backendMigrateState: target workspace has empty state, so might copy source workspace state")
		confirmFunc = m.backendMigrateEmptyConfirm

	// Both states are non-empty, meaning we need to determine which
	// state should be used and update accordingly.
	case !one.Empty() && !two.Empty():
		log.Print("[TRACE] backendMigrateState: both source and destination workspaces have states, so might overwrite destination with source")
		confirmFunc = m.backendMigrateNonEmptyConfirm
	}

	if confirmFunc == nil {
		panic("confirmFunc must not be nil")
	}

	if !opts.force {
		// Abort if we can't ask for input.
		if !m.input {
			log.Print("[TRACE] backendMigrateState: can't prompt for input, so aborting migration")
			return errors.New("error asking for state migration action: input disabled")
		}

		// Confirm with the user whether we want to copy state over
		confirm, err := confirmFunc(stateOne, stateTwo, opts)
		if err != nil {
			log.Print("[TRACE] backendMigrateState: error reading input, so aborting migration")
			return err
		}
		if !confirm {
			log.Print("[TRACE] backendMigrateState: user cancelled at confirmation prompt, so aborting migration")
			return nil
		}
	}

	// Confirmed! We'll have the statemgr package handle the migration, which
	// includes preserving any lineage/serial information where possible, if
	// both managers support such metadata.
	log.Print("[TRACE] backendMigrateState: migration confirmed, so migrating")
	if err := statemgr.Migrate(stateTwo, stateOne); err != nil {
		return fmt.Errorf(strings.TrimSpace(errBackendStateCopy),
			opts.OneType, opts.TwoType, err)
	}
	if err := stateTwo.PersistState(); err != nil {
		return fmt.Errorf(strings.TrimSpace(errBackendStateCopy),
			opts.OneType, opts.TwoType, err)
	}

	// And we're done.
	return nil
}

func (m *Meta) backendMigrateEmptyConfirm(one, two state.State, opts *backendMigrateOpts) (bool, error) {
	inputOpts := &terraform.InputOpts{
		Id:    "backend-migrate-copy-to-empty",
		Query: "Do you want to copy existing state to the new backend?",
		Description: fmt.Sprintf(
			strings.TrimSpace(inputBackendMigrateEmpty),
			opts.OneType, opts.TwoType),
	}

	return m.confirm(inputOpts)
}

func (m *Meta) backendMigrateNonEmptyConfirm(
	stateOne, stateTwo state.State, opts *backendMigrateOpts) (bool, error) {
	// We need to grab both states so we can write them to a file
	one := stateOne.State()
	two := stateTwo.State()

	// Save both to a temporary
	td, err := ioutil.TempDir("", "terraform")
	if err != nil {
		return false, fmt.Errorf("Error creating temporary directory: %s", err)
	}
	defer os.RemoveAll(td)

	// Helper to write the state
	saveHelper := func(n, path string, s *states.State) error {
		mgr := statemgr.NewFilesystem(path)
		return mgr.WriteState(s)
	}

	// Write the states
	onePath := filepath.Join(td, fmt.Sprintf("1-%s.tfstate", opts.OneType))
	twoPath := filepath.Join(td, fmt.Sprintf("2-%s.tfstate", opts.TwoType))
	if err := saveHelper(opts.OneType, onePath, one); err != nil {
		return false, fmt.Errorf("Error saving temporary state: %s", err)
	}
	if err := saveHelper(opts.TwoType, twoPath, two); err != nil {
		return false, fmt.Errorf("Error saving temporary state: %s", err)
	}

	// Ask for confirmation
	inputOpts := &terraform.InputOpts{
		Id:    "backend-migrate-to-backend",
		Query: "Do you want to copy existing state to the new backend?",
		Description: fmt.Sprintf(
			strings.TrimSpace(inputBackendMigrateNonEmpty),
			opts.OneType, opts.TwoType, onePath, twoPath),
	}

	// Confirm with the user that the copy should occur
	return m.confirm(inputOpts)
}

const errMigrateLoadStates = `
Error inspecting states in the %q backend:
    %s

Prior to changing backends, Terraform inspects the source and destination
states to determine what kind of migration steps need to be taken, if any.
Terraform failed to load the states. The data in both the source and the
destination remain unmodified. Please resolve the above error and try again.
`

const errMigrateSingleLoadDefault = `
Error loading state:
    %[2]s

Terraform failed to load the default state from the %[1]q backend.
State migration cannot occur unless the state can be loaded. Backend
modification and state migration has been aborted. The state in both the
source and the destination remain unmodified. Please resolve the
above error and try again.
`

const errMigrateMulti = `
Error migrating the workspace %q from the previous %q backend
to the newly configured %q backend:
    %s

Terraform copies workspaces in alphabetical order. Any workspaces
alphabetically earlier than this one have been copied. Any workspaces
later than this haven't been modified in the destination. No workspaces
in the source state have been modified.

Please resolve the error above and run the initialization command again.
This will attempt to copy (with permission) all workspaces again.
`

const errBackendStateCopy = `
Error copying state from the previous %q backend to the newly configured 
%q backend:
    %s

The state in the previous backend remains intact and unmodified. Please resolve
the error above and try again.
`

const inputBackendMigrateEmpty = `
Pre-existing state was found while migrating the previous %q backend to the
newly configured %q backend. No existing state was found in the newly
configured %[2]q backend. Do you want to copy this state to the new %[2]q
backend? Enter "yes" to copy and "no" to start with an empty state.
`

const inputBackendMigrateNonEmpty = `
Pre-existing state was found while migrating the previous %q backend to the
newly configured %q backend. An existing non-empty state already exists in
the new backend. The two states have been saved to temporary files that will be
removed after responding to this query.

Previous (type %[1]q): %[3]s
New      (type %[2]q): %[4]s

Do you want to overwrite the state in the new backend with the previous state?
Enter "yes" to copy and "no" to start with the existing state in the newly
configured %[2]q backend.
`

const inputBackendMigrateMultiToSingle = `
The existing %[1]q backend supports workspaces and you currently are
using more than one. The newly configured %[2]q backend doesn't support
workspaces. If you continue, Terraform will copy your current workspace %[3]q
to the default workspace in the target backend. Your existing workspaces in the
source backend won't be modified. If you want to switch workspaces, back them
up, or cancel altogether, answer "no" and Terraform will abort.
`

const inputBackendMigrateMultiToMulti = `
Both the existing %[1]q backend and the newly configured %[2]q backend
support workspaces. When migrating between backends, Terraform will copy
all workspaces (with the same names). THIS WILL OVERWRITE any conflicting
states in the destination.

Terraform initialization doesn't currently migrate only select workspaces.
If you want to migrate a select number of workspaces, you must manually
pull and push those states.

If you answer "yes", Terraform will migrate all states. If you answer
"no", Terraform will abort.
`

const inputBackendNewWorkspaceName = `
Please provide a new workspace name (e.g. dev, test) that will be used
to migrate the existing default workspace. 
`

const inputBackendSelectWorkspace = `
This is expected behavior when the selected workspace did not have an
existing non-empty state. Please enter a number to select a workspace:

%s
`
