package store

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/sugimoto-ne/go_sample_app.git/entity"
	"github.com/sugimoto-ne/go_sample_app.git/testutil"
)

func TestKVS_Save(t *testing.T) {
	t.Parallel()

	cli := testutil.OpenRedisForTest(t)
	sut := &KVS{Cli: cli}

	t.Run("ok", func(t *testing.T) {
		t.Parallel()

		key := "TestKVS_Load_ok"
		uid := entity.UserID(1234)

		ctx := context.Background()
		cli.Set(ctx, key, int64(uid), 30*time.Minute)
		t.Cleanup(func() {
			cli.Del(ctx, key)
		})

		got, err := sut.Load(ctx, key)
		if err != nil {
			t.Fatalf("want no error, but got %v", err)
		}

		if got != uid {
			t.Errorf("want %d, but got %d", uid, got)
		}
	})

	t.Run("notFound", func(t *testing.T) {
		t.Parallel()

		// TODO ErrNotFoundの出どころを調べる

		var ErrNotFound = errors.New("failed to get by \"TestKVS_Save/notFound\": redis: nil")
		key := "TestKVS_Save/notFound"
		ctx := context.Background()
		got, err := sut.Load(ctx, key)
		if err == nil || err == ErrNotFound {
			t.Errorf("want <%v>, but got <%v>(value = %d)", ErrNotFound, err, got)
		}
	})
}
