# Kong Gateway Admin API Wrapper
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/captainsuchard/kong-admin-api-wrapper?label=Version&logo=go&style=for-the-badge)
![Libraries.io dependency status for latest release](https://img.shields.io/librariesio/release/go/kong-admin-api-wrapper?label=Dependencies&style=for-the-badge)
![GitHub](https://img.shields.io/github/license/captainsuchard/kong-admin-api-wrapper?style=for-the-badge)
![GitHub last commit](https://img.shields.io/github/last-commit/captainsuchard/kong-admin-api-wrapper?style=for-the-badge)
![GitHub issues](https://img.shields.io/github/issues/captainsuchard/kong-admin-api-wrapper?style=for-the-badge)
![GitHub pull requests](https://img.shields.io/github/issues-pr/captainsuchard/kong-admin-api-wrapper?style=for-the-badge)

This package is a wrapper for the Kong Gateway Admin API. It is a work in progress and is not yet complete. It is 
intended to be used with the Kong Gateway version 2.0.0 or higher.

It is based on the Kong Gateway Admin API documentation found here: https://docs.konghq.com/gateway/latest/admin-api/
and uses the `go-kong` package found here: https://github.com/Kong/go-kong as base.

## Installation
To install the package, run the following command:
```bash
go get github.com/captainsuchard/kong-admin-api-wrapper
```

## Usage
To use the package, import it into your project:
```go
import (
    "github.com/captainsuchard/kong-admin-api-wrapper"
)
```

## Contributing
If you would like to contribute to the project, please fork the repository and submit a pull request.

## License
[GPL 3.0](https://choosealicense.com/licenses/gpl-3.0/)

## Authors
Created by [captainsuchard]()