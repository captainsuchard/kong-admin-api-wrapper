// Package vars stores the variables used throughout the package. While the variables are exported, they are not meant
// to be changed by the user. The variables are set by the setup package and its functions.
package vars

import "github.com/kong/go-kong/kong"

// KongClient is the internally used client used to connect to the kong admin api
var KongClient *kong.Client
