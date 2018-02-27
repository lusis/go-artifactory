package requests

// EmptyTrashCanRequest represents a EmptyTrashCan request
type EmptyTrashCanRequest struct{}

func (r EmptyTrashCanRequest) minVersion() Version { return versionMustParse("4.4.3") }
func (r EmptyTrashCanRequest) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (r EmptyTrashCanRequest) deprecated() bool    { return false }

// DeleteItemFromTrashCanRequest represents a DeleteItemFromTrashCan request
type DeleteItemFromTrashCanRequest struct{}

func (r DeleteItemFromTrashCanRequest) minVersion() Version { return versionMustParse("4.4.3") }
func (r DeleteItemFromTrashCanRequest) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (r DeleteItemFromTrashCanRequest) deprecated() bool    { return false }

// RestoreItemFromTrashCanRequest represents a RestoreItemFromTrashCan request
type RestoreItemFromTrashCanRequest struct{}

func (r RestoreItemFromTrashCanRequest) minVersion() Version { return versionMustParse("4.4.3") }
func (r RestoreItemFromTrashCanRequest) maxVersion() Version { return versionMustParse(CurrentVersion) }
func (r RestoreItemFromTrashCanRequest) deprecated() bool    { return false }
