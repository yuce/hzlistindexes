package internal

import (
	"context"
	"fmt"

	"github.com/hazelcast/hazelcast-go-client"
	"github.com/hazelcast/hazelcast-go-client/types"

	"github.com/yuce/hzlistindexes/internal/proto/codec"
)

func GetIndexes(ctx context.Context, ci *hazelcast.ClientInternal, mapName string) ([]types.IndexConfig, error) {
	req := codec.EncodeMCGetMapConfigRequest(mapName)
	resp, err := ci.InvokeOnRandomTarget(ctx, req, nil)
	if err != nil {
		return nil, fmt.Errorf("invoking get map config request: %w", err)
	}
	//inMemoryFormat, backupCount, asyncBackupCount, timeToLiveSeconds, maxIdleSeconds, maxSize, maxSizePolicy, readBackupData, evictionPolicy, mergePolicy, globalIndexes := codec.DecodeMCGetMapConfigResponse(resp)
	_, _, _, _, _, _, _, _, _, _, globalIndexes := codec.DecodeMCGetMapConfigResponse(resp)
	return globalIndexes, nil
}
