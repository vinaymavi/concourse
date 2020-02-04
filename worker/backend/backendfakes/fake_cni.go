// Code generated by counterfeiter. DO NOT EDIT.
package backendfakes

import (
	"context"
	"sync"

	cni "github.com/containerd/go-cni"
)

type FakeCNI struct {
	GetConfigStub        func() *cni.ConfigResult
	getConfigMutex       sync.RWMutex
	getConfigArgsForCall []struct {
	}
	getConfigReturns struct {
		result1 *cni.ConfigResult
	}
	getConfigReturnsOnCall map[int]struct {
		result1 *cni.ConfigResult
	}
	LoadStub        func(...cni.CNIOpt) error
	loadMutex       sync.RWMutex
	loadArgsForCall []struct {
		arg1 []cni.CNIOpt
	}
	loadReturns struct {
		result1 error
	}
	loadReturnsOnCall map[int]struct {
		result1 error
	}
	RemoveStub        func(context.Context, string, string, ...cni.NamespaceOpts) error
	removeMutex       sync.RWMutex
	removeArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 string
		arg4 []cni.NamespaceOpts
	}
	removeReturns struct {
		result1 error
	}
	removeReturnsOnCall map[int]struct {
		result1 error
	}
	SetupStub        func(context.Context, string, string, ...cni.NamespaceOpts) (*cni.CNIResult, error)
	setupMutex       sync.RWMutex
	setupArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 string
		arg4 []cni.NamespaceOpts
	}
	setupReturns struct {
		result1 *cni.CNIResult
		result2 error
	}
	setupReturnsOnCall map[int]struct {
		result1 *cni.CNIResult
		result2 error
	}
	StatusStub        func() error
	statusMutex       sync.RWMutex
	statusArgsForCall []struct {
	}
	statusReturns struct {
		result1 error
	}
	statusReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCNI) GetConfig() *cni.ConfigResult {
	fake.getConfigMutex.Lock()
	ret, specificReturn := fake.getConfigReturnsOnCall[len(fake.getConfigArgsForCall)]
	fake.getConfigArgsForCall = append(fake.getConfigArgsForCall, struct {
	}{})
	fake.recordInvocation("GetConfig", []interface{}{})
	fake.getConfigMutex.Unlock()
	if fake.GetConfigStub != nil {
		return fake.GetConfigStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.getConfigReturns
	return fakeReturns.result1
}

func (fake *FakeCNI) GetConfigCallCount() int {
	fake.getConfigMutex.RLock()
	defer fake.getConfigMutex.RUnlock()
	return len(fake.getConfigArgsForCall)
}

func (fake *FakeCNI) GetConfigCalls(stub func() *cni.ConfigResult) {
	fake.getConfigMutex.Lock()
	defer fake.getConfigMutex.Unlock()
	fake.GetConfigStub = stub
}

func (fake *FakeCNI) GetConfigReturns(result1 *cni.ConfigResult) {
	fake.getConfigMutex.Lock()
	defer fake.getConfigMutex.Unlock()
	fake.GetConfigStub = nil
	fake.getConfigReturns = struct {
		result1 *cni.ConfigResult
	}{result1}
}

func (fake *FakeCNI) GetConfigReturnsOnCall(i int, result1 *cni.ConfigResult) {
	fake.getConfigMutex.Lock()
	defer fake.getConfigMutex.Unlock()
	fake.GetConfigStub = nil
	if fake.getConfigReturnsOnCall == nil {
		fake.getConfigReturnsOnCall = make(map[int]struct {
			result1 *cni.ConfigResult
		})
	}
	fake.getConfigReturnsOnCall[i] = struct {
		result1 *cni.ConfigResult
	}{result1}
}

func (fake *FakeCNI) Load(arg1 ...cni.CNIOpt) error {
	fake.loadMutex.Lock()
	ret, specificReturn := fake.loadReturnsOnCall[len(fake.loadArgsForCall)]
	fake.loadArgsForCall = append(fake.loadArgsForCall, struct {
		arg1 []cni.CNIOpt
	}{arg1})
	fake.recordInvocation("Load", []interface{}{arg1})
	fake.loadMutex.Unlock()
	if fake.LoadStub != nil {
		return fake.LoadStub(arg1...)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.loadReturns
	return fakeReturns.result1
}

func (fake *FakeCNI) LoadCallCount() int {
	fake.loadMutex.RLock()
	defer fake.loadMutex.RUnlock()
	return len(fake.loadArgsForCall)
}

func (fake *FakeCNI) LoadCalls(stub func(...cni.CNIOpt) error) {
	fake.loadMutex.Lock()
	defer fake.loadMutex.Unlock()
	fake.LoadStub = stub
}

func (fake *FakeCNI) LoadArgsForCall(i int) []cni.CNIOpt {
	fake.loadMutex.RLock()
	defer fake.loadMutex.RUnlock()
	argsForCall := fake.loadArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCNI) LoadReturns(result1 error) {
	fake.loadMutex.Lock()
	defer fake.loadMutex.Unlock()
	fake.LoadStub = nil
	fake.loadReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeCNI) LoadReturnsOnCall(i int, result1 error) {
	fake.loadMutex.Lock()
	defer fake.loadMutex.Unlock()
	fake.LoadStub = nil
	if fake.loadReturnsOnCall == nil {
		fake.loadReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.loadReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeCNI) Remove(arg1 context.Context, arg2 string, arg3 string, arg4 ...cni.NamespaceOpts) error {
	fake.removeMutex.Lock()
	ret, specificReturn := fake.removeReturnsOnCall[len(fake.removeArgsForCall)]
	fake.removeArgsForCall = append(fake.removeArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 string
		arg4 []cni.NamespaceOpts
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("Remove", []interface{}{arg1, arg2, arg3, arg4})
	fake.removeMutex.Unlock()
	if fake.RemoveStub != nil {
		return fake.RemoveStub(arg1, arg2, arg3, arg4...)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.removeReturns
	return fakeReturns.result1
}

func (fake *FakeCNI) RemoveCallCount() int {
	fake.removeMutex.RLock()
	defer fake.removeMutex.RUnlock()
	return len(fake.removeArgsForCall)
}

func (fake *FakeCNI) RemoveCalls(stub func(context.Context, string, string, ...cni.NamespaceOpts) error) {
	fake.removeMutex.Lock()
	defer fake.removeMutex.Unlock()
	fake.RemoveStub = stub
}

func (fake *FakeCNI) RemoveArgsForCall(i int) (context.Context, string, string, []cni.NamespaceOpts) {
	fake.removeMutex.RLock()
	defer fake.removeMutex.RUnlock()
	argsForCall := fake.removeArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeCNI) RemoveReturns(result1 error) {
	fake.removeMutex.Lock()
	defer fake.removeMutex.Unlock()
	fake.RemoveStub = nil
	fake.removeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeCNI) RemoveReturnsOnCall(i int, result1 error) {
	fake.removeMutex.Lock()
	defer fake.removeMutex.Unlock()
	fake.RemoveStub = nil
	if fake.removeReturnsOnCall == nil {
		fake.removeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.removeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeCNI) Setup(arg1 context.Context, arg2 string, arg3 string, arg4 ...cni.NamespaceOpts) (*cni.CNIResult, error) {
	fake.setupMutex.Lock()
	ret, specificReturn := fake.setupReturnsOnCall[len(fake.setupArgsForCall)]
	fake.setupArgsForCall = append(fake.setupArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 string
		arg4 []cni.NamespaceOpts
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("Setup", []interface{}{arg1, arg2, arg3, arg4})
	fake.setupMutex.Unlock()
	if fake.SetupStub != nil {
		return fake.SetupStub(arg1, arg2, arg3, arg4...)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.setupReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeCNI) SetupCallCount() int {
	fake.setupMutex.RLock()
	defer fake.setupMutex.RUnlock()
	return len(fake.setupArgsForCall)
}

func (fake *FakeCNI) SetupCalls(stub func(context.Context, string, string, ...cni.NamespaceOpts) (*cni.CNIResult, error)) {
	fake.setupMutex.Lock()
	defer fake.setupMutex.Unlock()
	fake.SetupStub = stub
}

func (fake *FakeCNI) SetupArgsForCall(i int) (context.Context, string, string, []cni.NamespaceOpts) {
	fake.setupMutex.RLock()
	defer fake.setupMutex.RUnlock()
	argsForCall := fake.setupArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeCNI) SetupReturns(result1 *cni.CNIResult, result2 error) {
	fake.setupMutex.Lock()
	defer fake.setupMutex.Unlock()
	fake.SetupStub = nil
	fake.setupReturns = struct {
		result1 *cni.CNIResult
		result2 error
	}{result1, result2}
}

func (fake *FakeCNI) SetupReturnsOnCall(i int, result1 *cni.CNIResult, result2 error) {
	fake.setupMutex.Lock()
	defer fake.setupMutex.Unlock()
	fake.SetupStub = nil
	if fake.setupReturnsOnCall == nil {
		fake.setupReturnsOnCall = make(map[int]struct {
			result1 *cni.CNIResult
			result2 error
		})
	}
	fake.setupReturnsOnCall[i] = struct {
		result1 *cni.CNIResult
		result2 error
	}{result1, result2}
}

func (fake *FakeCNI) Status() error {
	fake.statusMutex.Lock()
	ret, specificReturn := fake.statusReturnsOnCall[len(fake.statusArgsForCall)]
	fake.statusArgsForCall = append(fake.statusArgsForCall, struct {
	}{})
	fake.recordInvocation("Status", []interface{}{})
	fake.statusMutex.Unlock()
	if fake.StatusStub != nil {
		return fake.StatusStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.statusReturns
	return fakeReturns.result1
}

func (fake *FakeCNI) StatusCallCount() int {
	fake.statusMutex.RLock()
	defer fake.statusMutex.RUnlock()
	return len(fake.statusArgsForCall)
}

func (fake *FakeCNI) StatusCalls(stub func() error) {
	fake.statusMutex.Lock()
	defer fake.statusMutex.Unlock()
	fake.StatusStub = stub
}

func (fake *FakeCNI) StatusReturns(result1 error) {
	fake.statusMutex.Lock()
	defer fake.statusMutex.Unlock()
	fake.StatusStub = nil
	fake.statusReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeCNI) StatusReturnsOnCall(i int, result1 error) {
	fake.statusMutex.Lock()
	defer fake.statusMutex.Unlock()
	fake.StatusStub = nil
	if fake.statusReturnsOnCall == nil {
		fake.statusReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.statusReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeCNI) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getConfigMutex.RLock()
	defer fake.getConfigMutex.RUnlock()
	fake.loadMutex.RLock()
	defer fake.loadMutex.RUnlock()
	fake.removeMutex.RLock()
	defer fake.removeMutex.RUnlock()
	fake.setupMutex.RLock()
	defer fake.setupMutex.RUnlock()
	fake.statusMutex.RLock()
	defer fake.statusMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCNI) recordInvocation(key string, args []interface{}) {
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

var _ cni.CNI = new(FakeCNI)
