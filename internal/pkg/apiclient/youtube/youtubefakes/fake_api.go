// Code generated by counterfeiter. DO NOT EDIT.
package youtubefakes

import (
	"context"
	"net/url"
	"sync"

	"github.com/hortbot/hortbot/internal/pkg/apiclient/youtube"
)

type FakeAPI struct {
	VideoTitleStub        func(context.Context, *url.URL) string
	videoTitleMutex       sync.RWMutex
	videoTitleArgsForCall []struct {
		arg1 context.Context
		arg2 *url.URL
	}
	videoTitleReturns struct {
		result1 string
	}
	videoTitleReturnsOnCall map[int]struct {
		result1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeAPI) VideoTitle(arg1 context.Context, arg2 *url.URL) string {
	fake.videoTitleMutex.Lock()
	ret, specificReturn := fake.videoTitleReturnsOnCall[len(fake.videoTitleArgsForCall)]
	fake.videoTitleArgsForCall = append(fake.videoTitleArgsForCall, struct {
		arg1 context.Context
		arg2 *url.URL
	}{arg1, arg2})
	fake.recordInvocation("VideoTitle", []interface{}{arg1, arg2})
	fake.videoTitleMutex.Unlock()
	if fake.VideoTitleStub != nil {
		return fake.VideoTitleStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.videoTitleReturns
	return fakeReturns.result1
}

func (fake *FakeAPI) VideoTitleCallCount() int {
	fake.videoTitleMutex.RLock()
	defer fake.videoTitleMutex.RUnlock()
	return len(fake.videoTitleArgsForCall)
}

func (fake *FakeAPI) VideoTitleCalls(stub func(context.Context, *url.URL) string) {
	fake.videoTitleMutex.Lock()
	defer fake.videoTitleMutex.Unlock()
	fake.VideoTitleStub = stub
}

func (fake *FakeAPI) VideoTitleArgsForCall(i int) (context.Context, *url.URL) {
	fake.videoTitleMutex.RLock()
	defer fake.videoTitleMutex.RUnlock()
	argsForCall := fake.videoTitleArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeAPI) VideoTitleReturns(result1 string) {
	fake.videoTitleMutex.Lock()
	defer fake.videoTitleMutex.Unlock()
	fake.VideoTitleStub = nil
	fake.videoTitleReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeAPI) VideoTitleReturnsOnCall(i int, result1 string) {
	fake.videoTitleMutex.Lock()
	defer fake.videoTitleMutex.Unlock()
	fake.VideoTitleStub = nil
	if fake.videoTitleReturnsOnCall == nil {
		fake.videoTitleReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.videoTitleReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeAPI) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.videoTitleMutex.RLock()
	defer fake.videoTitleMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeAPI) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ youtube.API = new(FakeAPI)
