package log

import (
	"context"

	pkgCtx "github.com/fetchealth/backend_monorepo/pkg/ctx"
)

type Field struct {
	Key   string
	Value interface{}
}

func AddFieldToCtx(ctx context.Context, key string, value interface{}) context.Context {
	fields := pkgCtx.GetFieldsFromCtx(ctx).([]*Field)
	if fields == nil {
		fields = make([]*Field, 0)
	}

	fields = append(fields, &Field{
		Key:   key,
		Value: value,
	})

	return pkgCtx.SetFieldsToCtx(ctx, fields)
}

func AddFieldsArrayToCtx(ctx context.Context, newFields ...*Field) context.Context {
	fields := pkgCtx.GetFieldsFromCtx(ctx).([]*Field)
	if fields == nil {
		fields = make([]*Field, 0)
	}

	fields = append(fields, newFields...)

	return pkgCtx.SetFieldsToCtx(ctx, fields)
}
