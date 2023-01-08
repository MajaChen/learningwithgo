package golang

import (
	"context"
	"fmt"
	"testing"
)

func TestContext(t *testing.T) {
	type keytypea string
	type keytypeb string
	var keyA keytypea = "keyA"
	ctx := context.Background()
	ctxA := context.WithValue(ctx, keyA, "valA")

	var keyC keytypeb = "keyA"
	ctxC := context.WithValue(ctxA, keyC, "valC")
	fmt.Println("keyA = ", ctxC.Value(keyA))
	fmt.Println("keyC = ", ctxC.Value(keyC))
}
