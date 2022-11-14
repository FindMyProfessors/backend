package pagination

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/99designs/gqlgen/graphql"
	"github.com/FindMyProfessors/backend/graph/model"
)

func DecodeCursor(cursor *string) (string, error) {
	if cursor == nil {
		return "0", nil
	}
	bytes, err := base64.StdEncoding.DecodeString(*cursor)
	if err != nil {
		return "", err
	}
	bytesString := string(bytes)
	return bytesString, nil
}

func GetPageInfo(firstElement string, lastElement string, lengthOfOutput int, amountRequested int) *model.PageInfo {
	encodedFirst := base64.StdEncoding.EncodeToString([]byte(firstElement))
	encodedLast := base64.StdEncoding.EncodeToString([]byte(lastElement))

	return &model.PageInfo{
		StartCursor: encodedFirst,
		EndCursor:   encodedLast,
		HasNextPage: lengthOfOutput >= amountRequested,
	}
}

func Pagination(ctx context.Context, _ interface{}, next graphql.Resolver, maxLength int) (res interface{}, err error) {
	fieldContext := graphql.GetFieldContext(ctx)
	first := fieldContext.Args["first"].(int)
	if first > maxLength {
		return nil, fmt.Errorf("you are only allowed to request non-negative amounts less than or equal to %d", maxLength)
	}
	return next(ctx)
}
