//go:build tools
// +build tools

package tools

import (
	_ "memory_golang/api/pkg/httpserv/gql/scalar"

	_ "github.com/99designs/gqlgen"
	_ "github.com/99designs/gqlgen/graphql"
	_ "github.com/99designs/gqlgen/graphql/introspection"
)
