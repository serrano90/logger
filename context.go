package log

import (
	goCtx "context"

	pkgCtx "github.com/fetchealth/backend_monorepo/pkg/ctx"
)

func GetFieldsFromCtx(ctx goCtx.Context) []*Field {
	fields, ok := ctx.Value(pkgCtx.LoggerFields).([]*Field)
	if !ok {
		return nil
	}
	return fields
}

func SetFieldsToCtx(ctx goCtx.Context, fields []*Field) goCtx.Context {
	return goCtx.WithValue(ctx, pkgCtx.LoggerFields, fields)
}
