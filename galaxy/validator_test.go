package galaxy

import (
	"context"
	"fmt"
	"testing"
)

func TestIAPVerfiy(t *testing.T) {
	client := New()

	ctx := context.Background()
	resp, err := client.Verify(ctx, "1231237123")
	if err != nil {
		t.Errorf("%s", err)
	}
	fmt.Println(resp)
}
