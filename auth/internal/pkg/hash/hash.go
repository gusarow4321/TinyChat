package hash

import "github.com/google/wire"

// ProviderSet is hash providers.
var ProviderSet = wire.NewSet(NewPasswordHasher)
