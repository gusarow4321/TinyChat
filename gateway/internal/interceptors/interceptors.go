package interceptors

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewAuthInterceptor)
