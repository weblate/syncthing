// Code generated by counterfeiter. DO NOT EDIT.
package mocks

import (
	"context"
	"sync"

	"github.com/syncthing/syncthing/lib/model"
	"github.com/syncthing/syncthing/lib/model/types"
)

type FolderSummaryService struct {
	OnEventRequestStub        func()
	onEventRequestMutex       sync.RWMutex
	onEventRequestArgsForCall []struct {
	}
	ServeStub        func(context.Context) error
	serveMutex       sync.RWMutex
	serveArgsForCall []struct {
		arg1 context.Context
	}
	serveReturns struct {
		result1 error
	}
	serveReturnsOnCall map[int]struct {
		result1 error
	}
	SummaryStub        func(string) (*types.FolderSummary, error)
	summaryMutex       sync.RWMutex
	summaryArgsForCall []struct {
		arg1 string
	}
	summaryReturns struct {
		result1 *types.FolderSummary
		result2 error
	}
	summaryReturnsOnCall map[int]struct {
		result1 *types.FolderSummary
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FolderSummaryService) OnEventRequest() {
	fake.onEventRequestMutex.Lock()
	fake.onEventRequestArgsForCall = append(fake.onEventRequestArgsForCall, struct {
	}{})
	stub := fake.OnEventRequestStub
	fake.recordInvocation("OnEventRequest", []interface{}{})
	fake.onEventRequestMutex.Unlock()
	if stub != nil {
		fake.OnEventRequestStub()
	}
}

func (fake *FolderSummaryService) OnEventRequestCallCount() int {
	fake.onEventRequestMutex.RLock()
	defer fake.onEventRequestMutex.RUnlock()
	return len(fake.onEventRequestArgsForCall)
}

func (fake *FolderSummaryService) OnEventRequestCalls(stub func()) {
	fake.onEventRequestMutex.Lock()
	defer fake.onEventRequestMutex.Unlock()
	fake.OnEventRequestStub = stub
}

func (fake *FolderSummaryService) Serve(arg1 context.Context) error {
	fake.serveMutex.Lock()
	ret, specificReturn := fake.serveReturnsOnCall[len(fake.serveArgsForCall)]
	fake.serveArgsForCall = append(fake.serveArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	stub := fake.ServeStub
	fakeReturns := fake.serveReturns
	fake.recordInvocation("Serve", []interface{}{arg1})
	fake.serveMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FolderSummaryService) ServeCallCount() int {
	fake.serveMutex.RLock()
	defer fake.serveMutex.RUnlock()
	return len(fake.serveArgsForCall)
}

func (fake *FolderSummaryService) ServeCalls(stub func(context.Context) error) {
	fake.serveMutex.Lock()
	defer fake.serveMutex.Unlock()
	fake.ServeStub = stub
}

func (fake *FolderSummaryService) ServeArgsForCall(i int) context.Context {
	fake.serveMutex.RLock()
	defer fake.serveMutex.RUnlock()
	argsForCall := fake.serveArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FolderSummaryService) ServeReturns(result1 error) {
	fake.serveMutex.Lock()
	defer fake.serveMutex.Unlock()
	fake.ServeStub = nil
	fake.serveReturns = struct {
		result1 error
	}{result1}
}

func (fake *FolderSummaryService) ServeReturnsOnCall(i int, result1 error) {
	fake.serveMutex.Lock()
	defer fake.serveMutex.Unlock()
	fake.ServeStub = nil
	if fake.serveReturnsOnCall == nil {
		fake.serveReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.serveReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FolderSummaryService) Summary(arg1 string) (*types.FolderSummary, error) {
	fake.summaryMutex.Lock()
	ret, specificReturn := fake.summaryReturnsOnCall[len(fake.summaryArgsForCall)]
	fake.summaryArgsForCall = append(fake.summaryArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.SummaryStub
	fakeReturns := fake.summaryReturns
	fake.recordInvocation("Summary", []interface{}{arg1})
	fake.summaryMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FolderSummaryService) SummaryCallCount() int {
	fake.summaryMutex.RLock()
	defer fake.summaryMutex.RUnlock()
	return len(fake.summaryArgsForCall)
}

func (fake *FolderSummaryService) SummaryCalls(stub func(string) (*types.FolderSummary, error)) {
	fake.summaryMutex.Lock()
	defer fake.summaryMutex.Unlock()
	fake.SummaryStub = stub
}

func (fake *FolderSummaryService) SummaryArgsForCall(i int) string {
	fake.summaryMutex.RLock()
	defer fake.summaryMutex.RUnlock()
	argsForCall := fake.summaryArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FolderSummaryService) SummaryReturns(result1 *types.FolderSummary, result2 error) {
	fake.summaryMutex.Lock()
	defer fake.summaryMutex.Unlock()
	fake.SummaryStub = nil
	fake.summaryReturns = struct {
		result1 *types.FolderSummary
		result2 error
	}{result1, result2}
}

func (fake *FolderSummaryService) SummaryReturnsOnCall(i int, result1 *types.FolderSummary, result2 error) {
	fake.summaryMutex.Lock()
	defer fake.summaryMutex.Unlock()
	fake.SummaryStub = nil
	if fake.summaryReturnsOnCall == nil {
		fake.summaryReturnsOnCall = make(map[int]struct {
			result1 *types.FolderSummary
			result2 error
		})
	}
	fake.summaryReturnsOnCall[i] = struct {
		result1 *types.FolderSummary
		result2 error
	}{result1, result2}
}

func (fake *FolderSummaryService) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.onEventRequestMutex.RLock()
	defer fake.onEventRequestMutex.RUnlock()
	fake.serveMutex.RLock()
	defer fake.serveMutex.RUnlock()
	fake.summaryMutex.RLock()
	defer fake.summaryMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FolderSummaryService) recordInvocation(key string, args []interface{}) {
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

var _ model.FolderSummaryService = new(FolderSummaryService)
