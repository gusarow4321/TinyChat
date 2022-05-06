package paseto

import "github.com/google/wire"

// ProviderSet is paseto providers.
var ProviderSet = wire.NewSet(NewPasetoMaker)
