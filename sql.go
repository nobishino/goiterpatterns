package goiterpatterns

import (
	"context"
	"database/sql"
	"iter"
)

type Scanable interface {
	Pointers() []any
}

func QueryToSQLRows[T Scanable](ctx context.Context, db *sql.DB, query string, args ...any) iter.Seq2[T, error] {
	return func(yield func(T, error) bool) {
		var zero T
		rows, err := db.QueryContext(ctx, query, args...)
		if err != nil {
			yield(zero, err)
			return
		}
		defer rows.Close()
		for t, err := range SQLRowsToSeq2[T](rows) {
			if !yield(t, err) {
				break
			}
		}
	}
}

func SQLRowsToSeq2[T Scanable](rows *sql.Rows) iter.Seq2[T, error] {
	return func(yield func(T, error) bool) {
		for rows.Next() {
			var dest T
			err := rows.Scan(dest.Pointers()...)
			if !yield(dest, err) {
				break
			}
		}
		if err := rows.Err(); err != nil {
			var zero T
			yield(zero, err)
			return
		}
	}
}
