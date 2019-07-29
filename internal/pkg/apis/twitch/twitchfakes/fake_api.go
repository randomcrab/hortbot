// Code generated by counterfeiter. DO NOT EDIT.
package twitchfakes

import (
	"context"
	"sync"

	"github.com/hortbot/hortbot/internal/pkg/apis/twitch"
	"golang.org/x/oauth2"
)

type FakeAPI struct {
	GetChannelStub        func(context.Context, *oauth2.Token) (*twitch.Channel, *oauth2.Token, error)
	getChannelMutex       sync.RWMutex
	getChannelArgsForCall []struct {
		arg1 context.Context
		arg2 *oauth2.Token
	}
	getChannelReturns struct {
		result1 *twitch.Channel
		result2 *oauth2.Token
		result3 error
	}
	getChannelReturnsOnCall map[int]struct {
		result1 *twitch.Channel
		result2 *oauth2.Token
		result3 error
	}
	GetChannelByIDStub        func(context.Context, int64) (*twitch.Channel, error)
	getChannelByIDMutex       sync.RWMutex
	getChannelByIDArgsForCall []struct {
		arg1 context.Context
		arg2 int64
	}
	getChannelByIDReturns struct {
		result1 *twitch.Channel
		result2 error
	}
	getChannelByIDReturnsOnCall map[int]struct {
		result1 *twitch.Channel
		result2 error
	}
	SetChannelGameStub        func(context.Context, int64, *oauth2.Token, string) (*oauth2.Token, error)
	setChannelGameMutex       sync.RWMutex
	setChannelGameArgsForCall []struct {
		arg1 context.Context
		arg2 int64
		arg3 *oauth2.Token
		arg4 string
	}
	setChannelGameReturns struct {
		result1 *oauth2.Token
		result2 error
	}
	setChannelGameReturnsOnCall map[int]struct {
		result1 *oauth2.Token
		result2 error
	}
	SetChannelStatusStub        func(context.Context, int64, *oauth2.Token, string) (*oauth2.Token, error)
	setChannelStatusMutex       sync.RWMutex
	setChannelStatusArgsForCall []struct {
		arg1 context.Context
		arg2 int64
		arg3 *oauth2.Token
		arg4 string
	}
	setChannelStatusReturns struct {
		result1 *oauth2.Token
		result2 error
	}
	setChannelStatusReturnsOnCall map[int]struct {
		result1 *oauth2.Token
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeAPI) GetChannel(arg1 context.Context, arg2 *oauth2.Token) (*twitch.Channel, *oauth2.Token, error) {
	fake.getChannelMutex.Lock()
	ret, specificReturn := fake.getChannelReturnsOnCall[len(fake.getChannelArgsForCall)]
	fake.getChannelArgsForCall = append(fake.getChannelArgsForCall, struct {
		arg1 context.Context
		arg2 *oauth2.Token
	}{arg1, arg2})
	fake.recordInvocation("GetChannel", []interface{}{arg1, arg2})
	fake.getChannelMutex.Unlock()
	if fake.GetChannelStub != nil {
		return fake.GetChannelStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getChannelReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeAPI) GetChannelCallCount() int {
	fake.getChannelMutex.RLock()
	defer fake.getChannelMutex.RUnlock()
	return len(fake.getChannelArgsForCall)
}

func (fake *FakeAPI) GetChannelCalls(stub func(context.Context, *oauth2.Token) (*twitch.Channel, *oauth2.Token, error)) {
	fake.getChannelMutex.Lock()
	defer fake.getChannelMutex.Unlock()
	fake.GetChannelStub = stub
}

func (fake *FakeAPI) GetChannelArgsForCall(i int) (context.Context, *oauth2.Token) {
	fake.getChannelMutex.RLock()
	defer fake.getChannelMutex.RUnlock()
	argsForCall := fake.getChannelArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeAPI) GetChannelReturns(result1 *twitch.Channel, result2 *oauth2.Token, result3 error) {
	fake.getChannelMutex.Lock()
	defer fake.getChannelMutex.Unlock()
	fake.GetChannelStub = nil
	fake.getChannelReturns = struct {
		result1 *twitch.Channel
		result2 *oauth2.Token
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeAPI) GetChannelReturnsOnCall(i int, result1 *twitch.Channel, result2 *oauth2.Token, result3 error) {
	fake.getChannelMutex.Lock()
	defer fake.getChannelMutex.Unlock()
	fake.GetChannelStub = nil
	if fake.getChannelReturnsOnCall == nil {
		fake.getChannelReturnsOnCall = make(map[int]struct {
			result1 *twitch.Channel
			result2 *oauth2.Token
			result3 error
		})
	}
	fake.getChannelReturnsOnCall[i] = struct {
		result1 *twitch.Channel
		result2 *oauth2.Token
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeAPI) GetChannelByID(arg1 context.Context, arg2 int64) (*twitch.Channel, error) {
	fake.getChannelByIDMutex.Lock()
	ret, specificReturn := fake.getChannelByIDReturnsOnCall[len(fake.getChannelByIDArgsForCall)]
	fake.getChannelByIDArgsForCall = append(fake.getChannelByIDArgsForCall, struct {
		arg1 context.Context
		arg2 int64
	}{arg1, arg2})
	fake.recordInvocation("GetChannelByID", []interface{}{arg1, arg2})
	fake.getChannelByIDMutex.Unlock()
	if fake.GetChannelByIDStub != nil {
		return fake.GetChannelByIDStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getChannelByIDReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeAPI) GetChannelByIDCallCount() int {
	fake.getChannelByIDMutex.RLock()
	defer fake.getChannelByIDMutex.RUnlock()
	return len(fake.getChannelByIDArgsForCall)
}

func (fake *FakeAPI) GetChannelByIDCalls(stub func(context.Context, int64) (*twitch.Channel, error)) {
	fake.getChannelByIDMutex.Lock()
	defer fake.getChannelByIDMutex.Unlock()
	fake.GetChannelByIDStub = stub
}

func (fake *FakeAPI) GetChannelByIDArgsForCall(i int) (context.Context, int64) {
	fake.getChannelByIDMutex.RLock()
	defer fake.getChannelByIDMutex.RUnlock()
	argsForCall := fake.getChannelByIDArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeAPI) GetChannelByIDReturns(result1 *twitch.Channel, result2 error) {
	fake.getChannelByIDMutex.Lock()
	defer fake.getChannelByIDMutex.Unlock()
	fake.GetChannelByIDStub = nil
	fake.getChannelByIDReturns = struct {
		result1 *twitch.Channel
		result2 error
	}{result1, result2}
}

func (fake *FakeAPI) GetChannelByIDReturnsOnCall(i int, result1 *twitch.Channel, result2 error) {
	fake.getChannelByIDMutex.Lock()
	defer fake.getChannelByIDMutex.Unlock()
	fake.GetChannelByIDStub = nil
	if fake.getChannelByIDReturnsOnCall == nil {
		fake.getChannelByIDReturnsOnCall = make(map[int]struct {
			result1 *twitch.Channel
			result2 error
		})
	}
	fake.getChannelByIDReturnsOnCall[i] = struct {
		result1 *twitch.Channel
		result2 error
	}{result1, result2}
}

func (fake *FakeAPI) SetChannelGame(arg1 context.Context, arg2 int64, arg3 *oauth2.Token, arg4 string) (*oauth2.Token, error) {
	fake.setChannelGameMutex.Lock()
	ret, specificReturn := fake.setChannelGameReturnsOnCall[len(fake.setChannelGameArgsForCall)]
	fake.setChannelGameArgsForCall = append(fake.setChannelGameArgsForCall, struct {
		arg1 context.Context
		arg2 int64
		arg3 *oauth2.Token
		arg4 string
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("SetChannelGame", []interface{}{arg1, arg2, arg3, arg4})
	fake.setChannelGameMutex.Unlock()
	if fake.SetChannelGameStub != nil {
		return fake.SetChannelGameStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.setChannelGameReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeAPI) SetChannelGameCallCount() int {
	fake.setChannelGameMutex.RLock()
	defer fake.setChannelGameMutex.RUnlock()
	return len(fake.setChannelGameArgsForCall)
}

func (fake *FakeAPI) SetChannelGameCalls(stub func(context.Context, int64, *oauth2.Token, string) (*oauth2.Token, error)) {
	fake.setChannelGameMutex.Lock()
	defer fake.setChannelGameMutex.Unlock()
	fake.SetChannelGameStub = stub
}

func (fake *FakeAPI) SetChannelGameArgsForCall(i int) (context.Context, int64, *oauth2.Token, string) {
	fake.setChannelGameMutex.RLock()
	defer fake.setChannelGameMutex.RUnlock()
	argsForCall := fake.setChannelGameArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeAPI) SetChannelGameReturns(result1 *oauth2.Token, result2 error) {
	fake.setChannelGameMutex.Lock()
	defer fake.setChannelGameMutex.Unlock()
	fake.SetChannelGameStub = nil
	fake.setChannelGameReturns = struct {
		result1 *oauth2.Token
		result2 error
	}{result1, result2}
}

func (fake *FakeAPI) SetChannelGameReturnsOnCall(i int, result1 *oauth2.Token, result2 error) {
	fake.setChannelGameMutex.Lock()
	defer fake.setChannelGameMutex.Unlock()
	fake.SetChannelGameStub = nil
	if fake.setChannelGameReturnsOnCall == nil {
		fake.setChannelGameReturnsOnCall = make(map[int]struct {
			result1 *oauth2.Token
			result2 error
		})
	}
	fake.setChannelGameReturnsOnCall[i] = struct {
		result1 *oauth2.Token
		result2 error
	}{result1, result2}
}

func (fake *FakeAPI) SetChannelStatus(arg1 context.Context, arg2 int64, arg3 *oauth2.Token, arg4 string) (*oauth2.Token, error) {
	fake.setChannelStatusMutex.Lock()
	ret, specificReturn := fake.setChannelStatusReturnsOnCall[len(fake.setChannelStatusArgsForCall)]
	fake.setChannelStatusArgsForCall = append(fake.setChannelStatusArgsForCall, struct {
		arg1 context.Context
		arg2 int64
		arg3 *oauth2.Token
		arg4 string
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("SetChannelStatus", []interface{}{arg1, arg2, arg3, arg4})
	fake.setChannelStatusMutex.Unlock()
	if fake.SetChannelStatusStub != nil {
		return fake.SetChannelStatusStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.setChannelStatusReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeAPI) SetChannelStatusCallCount() int {
	fake.setChannelStatusMutex.RLock()
	defer fake.setChannelStatusMutex.RUnlock()
	return len(fake.setChannelStatusArgsForCall)
}

func (fake *FakeAPI) SetChannelStatusCalls(stub func(context.Context, int64, *oauth2.Token, string) (*oauth2.Token, error)) {
	fake.setChannelStatusMutex.Lock()
	defer fake.setChannelStatusMutex.Unlock()
	fake.SetChannelStatusStub = stub
}

func (fake *FakeAPI) SetChannelStatusArgsForCall(i int) (context.Context, int64, *oauth2.Token, string) {
	fake.setChannelStatusMutex.RLock()
	defer fake.setChannelStatusMutex.RUnlock()
	argsForCall := fake.setChannelStatusArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeAPI) SetChannelStatusReturns(result1 *oauth2.Token, result2 error) {
	fake.setChannelStatusMutex.Lock()
	defer fake.setChannelStatusMutex.Unlock()
	fake.SetChannelStatusStub = nil
	fake.setChannelStatusReturns = struct {
		result1 *oauth2.Token
		result2 error
	}{result1, result2}
}

func (fake *FakeAPI) SetChannelStatusReturnsOnCall(i int, result1 *oauth2.Token, result2 error) {
	fake.setChannelStatusMutex.Lock()
	defer fake.setChannelStatusMutex.Unlock()
	fake.SetChannelStatusStub = nil
	if fake.setChannelStatusReturnsOnCall == nil {
		fake.setChannelStatusReturnsOnCall = make(map[int]struct {
			result1 *oauth2.Token
			result2 error
		})
	}
	fake.setChannelStatusReturnsOnCall[i] = struct {
		result1 *oauth2.Token
		result2 error
	}{result1, result2}
}

func (fake *FakeAPI) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getChannelMutex.RLock()
	defer fake.getChannelMutex.RUnlock()
	fake.getChannelByIDMutex.RLock()
	defer fake.getChannelByIDMutex.RUnlock()
	fake.setChannelGameMutex.RLock()
	defer fake.setChannelGameMutex.RUnlock()
	fake.setChannelStatusMutex.RLock()
	defer fake.setChannelStatusMutex.RUnlock()
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

var _ twitch.API = new(FakeAPI)
