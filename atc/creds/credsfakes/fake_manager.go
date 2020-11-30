// Code generated by counterfeiter. DO NOT EDIT.
package credsfakes

import (
	"sync"

	"code.cloudfoundry.org/lager"
	"github.com/concourse/concourse/atc/creds"
)

type FakeManager struct {
	CloseStub        func(lager.Logger)
	closeMutex       sync.RWMutex
	closeArgsForCall []struct {
		arg1 lager.Logger
	}
	HealthStub        func() (*creds.HealthResponse, error)
	healthMutex       sync.RWMutex
	healthArgsForCall []struct {
	}
	healthReturns struct {
		result1 *creds.HealthResponse
		result2 error
	}
	healthReturnsOnCall map[int]struct {
		result1 *creds.HealthResponse
		result2 error
	}
	InitStub        func(lager.Logger) error
	initMutex       sync.RWMutex
	initArgsForCall []struct {
		arg1 lager.Logger
	}
	initReturns struct {
		result1 error
	}
	initReturnsOnCall map[int]struct {
		result1 error
	}
	IsConfiguredStub        func() bool
	isConfiguredMutex       sync.RWMutex
	isConfiguredArgsForCall []struct {
	}
	isConfiguredReturns struct {
		result1 bool
	}
	isConfiguredReturnsOnCall map[int]struct {
		result1 bool
	}
	NewSecretsFactoryStub        func(lager.Logger) (creds.SecretsFactory, error)
	newSecretsFactoryMutex       sync.RWMutex
	newSecretsFactoryArgsForCall []struct {
		arg1 lager.Logger
	}
	newSecretsFactoryReturns struct {
		result1 creds.SecretsFactory
		result2 error
	}
	newSecretsFactoryReturnsOnCall map[int]struct {
		result1 creds.SecretsFactory
		result2 error
	}
	ValidateStub        func() error
	validateMutex       sync.RWMutex
	validateArgsForCall []struct {
	}
	validateReturns struct {
		result1 error
	}
	validateReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeManager) Close(arg1 lager.Logger) {
	fake.closeMutex.Lock()
	fake.closeArgsForCall = append(fake.closeArgsForCall, struct {
		arg1 lager.Logger
	}{arg1})
	fake.recordInvocation("Close", []interface{}{arg1})
	fake.closeMutex.Unlock()
	if fake.CloseStub != nil {
		fake.CloseStub(arg1)
	}
}

func (fake *FakeManager) CloseCallCount() int {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	return len(fake.closeArgsForCall)
}

func (fake *FakeManager) CloseCalls(stub func(lager.Logger)) {
	fake.closeMutex.Lock()
	defer fake.closeMutex.Unlock()
	fake.CloseStub = stub
}

func (fake *FakeManager) CloseArgsForCall(i int) lager.Logger {
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	argsForCall := fake.closeArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeManager) Health() (*creds.HealthResponse, error) {
	fake.healthMutex.Lock()
	ret, specificReturn := fake.healthReturnsOnCall[len(fake.healthArgsForCall)]
	fake.healthArgsForCall = append(fake.healthArgsForCall, struct {
	}{})
	fake.recordInvocation("Health", []interface{}{})
	fake.healthMutex.Unlock()
	if fake.HealthStub != nil {
		return fake.HealthStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.healthReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeManager) HealthCallCount() int {
	fake.healthMutex.RLock()
	defer fake.healthMutex.RUnlock()
	return len(fake.healthArgsForCall)
}

func (fake *FakeManager) HealthCalls(stub func() (*creds.HealthResponse, error)) {
	fake.healthMutex.Lock()
	defer fake.healthMutex.Unlock()
	fake.HealthStub = stub
}

func (fake *FakeManager) HealthReturns(result1 *creds.HealthResponse, result2 error) {
	fake.healthMutex.Lock()
	defer fake.healthMutex.Unlock()
	fake.HealthStub = nil
	fake.healthReturns = struct {
		result1 *creds.HealthResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeManager) HealthReturnsOnCall(i int, result1 *creds.HealthResponse, result2 error) {
	fake.healthMutex.Lock()
	defer fake.healthMutex.Unlock()
	fake.HealthStub = nil
	if fake.healthReturnsOnCall == nil {
		fake.healthReturnsOnCall = make(map[int]struct {
			result1 *creds.HealthResponse
			result2 error
		})
	}
	fake.healthReturnsOnCall[i] = struct {
		result1 *creds.HealthResponse
		result2 error
	}{result1, result2}
}

func (fake *FakeManager) Init(arg1 lager.Logger) error {
	fake.initMutex.Lock()
	ret, specificReturn := fake.initReturnsOnCall[len(fake.initArgsForCall)]
	fake.initArgsForCall = append(fake.initArgsForCall, struct {
		arg1 lager.Logger
	}{arg1})
	fake.recordInvocation("Init", []interface{}{arg1})
	fake.initMutex.Unlock()
	if fake.InitStub != nil {
		return fake.InitStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.initReturns
	return fakeReturns.result1
}

func (fake *FakeManager) InitCallCount() int {
	fake.initMutex.RLock()
	defer fake.initMutex.RUnlock()
	return len(fake.initArgsForCall)
}

func (fake *FakeManager) InitCalls(stub func(lager.Logger) error) {
	fake.initMutex.Lock()
	defer fake.initMutex.Unlock()
	fake.InitStub = stub
}

func (fake *FakeManager) InitArgsForCall(i int) lager.Logger {
	fake.initMutex.RLock()
	defer fake.initMutex.RUnlock()
	argsForCall := fake.initArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeManager) InitReturns(result1 error) {
	fake.initMutex.Lock()
	defer fake.initMutex.Unlock()
	fake.InitStub = nil
	fake.initReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeManager) InitReturnsOnCall(i int, result1 error) {
	fake.initMutex.Lock()
	defer fake.initMutex.Unlock()
	fake.InitStub = nil
	if fake.initReturnsOnCall == nil {
		fake.initReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.initReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeManager) IsConfigured() bool {
	fake.isConfiguredMutex.Lock()
	ret, specificReturn := fake.isConfiguredReturnsOnCall[len(fake.isConfiguredArgsForCall)]
	fake.isConfiguredArgsForCall = append(fake.isConfiguredArgsForCall, struct {
	}{})
	fake.recordInvocation("IsConfigured", []interface{}{})
	fake.isConfiguredMutex.Unlock()
	if fake.IsConfiguredStub != nil {
		return fake.IsConfiguredStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.isConfiguredReturns
	return fakeReturns.result1
}

func (fake *FakeManager) IsConfiguredCallCount() int {
	fake.isConfiguredMutex.RLock()
	defer fake.isConfiguredMutex.RUnlock()
	return len(fake.isConfiguredArgsForCall)
}

func (fake *FakeManager) IsConfiguredCalls(stub func() bool) {
	fake.isConfiguredMutex.Lock()
	defer fake.isConfiguredMutex.Unlock()
	fake.IsConfiguredStub = stub
}

func (fake *FakeManager) IsConfiguredReturns(result1 bool) {
	fake.isConfiguredMutex.Lock()
	defer fake.isConfiguredMutex.Unlock()
	fake.IsConfiguredStub = nil
	fake.isConfiguredReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeManager) IsConfiguredReturnsOnCall(i int, result1 bool) {
	fake.isConfiguredMutex.Lock()
	defer fake.isConfiguredMutex.Unlock()
	fake.IsConfiguredStub = nil
	if fake.isConfiguredReturnsOnCall == nil {
		fake.isConfiguredReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.isConfiguredReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeManager) NewSecretsFactory(arg1 lager.Logger) (creds.SecretsFactory, error) {
	fake.newSecretsFactoryMutex.Lock()
	ret, specificReturn := fake.newSecretsFactoryReturnsOnCall[len(fake.newSecretsFactoryArgsForCall)]
	fake.newSecretsFactoryArgsForCall = append(fake.newSecretsFactoryArgsForCall, struct {
		arg1 lager.Logger
	}{arg1})
	fake.recordInvocation("NewSecretsFactory", []interface{}{arg1})
	fake.newSecretsFactoryMutex.Unlock()
	if fake.NewSecretsFactoryStub != nil {
		return fake.NewSecretsFactoryStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.newSecretsFactoryReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeManager) NewSecretsFactoryCallCount() int {
	fake.newSecretsFactoryMutex.RLock()
	defer fake.newSecretsFactoryMutex.RUnlock()
	return len(fake.newSecretsFactoryArgsForCall)
}

func (fake *FakeManager) NewSecretsFactoryCalls(stub func(lager.Logger) (creds.SecretsFactory, error)) {
	fake.newSecretsFactoryMutex.Lock()
	defer fake.newSecretsFactoryMutex.Unlock()
	fake.NewSecretsFactoryStub = stub
}

func (fake *FakeManager) NewSecretsFactoryArgsForCall(i int) lager.Logger {
	fake.newSecretsFactoryMutex.RLock()
	defer fake.newSecretsFactoryMutex.RUnlock()
	argsForCall := fake.newSecretsFactoryArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeManager) NewSecretsFactoryReturns(result1 creds.SecretsFactory, result2 error) {
	fake.newSecretsFactoryMutex.Lock()
	defer fake.newSecretsFactoryMutex.Unlock()
	fake.NewSecretsFactoryStub = nil
	fake.newSecretsFactoryReturns = struct {
		result1 creds.SecretsFactory
		result2 error
	}{result1, result2}
}

func (fake *FakeManager) NewSecretsFactoryReturnsOnCall(i int, result1 creds.SecretsFactory, result2 error) {
	fake.newSecretsFactoryMutex.Lock()
	defer fake.newSecretsFactoryMutex.Unlock()
	fake.NewSecretsFactoryStub = nil
	if fake.newSecretsFactoryReturnsOnCall == nil {
		fake.newSecretsFactoryReturnsOnCall = make(map[int]struct {
			result1 creds.SecretsFactory
			result2 error
		})
	}
	fake.newSecretsFactoryReturnsOnCall[i] = struct {
		result1 creds.SecretsFactory
		result2 error
	}{result1, result2}
}

func (fake *FakeManager) Validate() error {
	fake.validateMutex.Lock()
	ret, specificReturn := fake.validateReturnsOnCall[len(fake.validateArgsForCall)]
	fake.validateArgsForCall = append(fake.validateArgsForCall, struct {
	}{})
	fake.recordInvocation("Validate", []interface{}{})
	fake.validateMutex.Unlock()
	if fake.ValidateStub != nil {
		return fake.ValidateStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.validateReturns
	return fakeReturns.result1
}

func (fake *FakeManager) ValidateCallCount() int {
	fake.validateMutex.RLock()
	defer fake.validateMutex.RUnlock()
	return len(fake.validateArgsForCall)
}

func (fake *FakeManager) ValidateCalls(stub func() error) {
	fake.validateMutex.Lock()
	defer fake.validateMutex.Unlock()
	fake.ValidateStub = stub
}

func (fake *FakeManager) ValidateReturns(result1 error) {
	fake.validateMutex.Lock()
	defer fake.validateMutex.Unlock()
	fake.ValidateStub = nil
	fake.validateReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeManager) ValidateReturnsOnCall(i int, result1 error) {
	fake.validateMutex.Lock()
	defer fake.validateMutex.Unlock()
	fake.ValidateStub = nil
	if fake.validateReturnsOnCall == nil {
		fake.validateReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.validateReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeManager) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.closeMutex.RLock()
	defer fake.closeMutex.RUnlock()
	fake.healthMutex.RLock()
	defer fake.healthMutex.RUnlock()
	fake.initMutex.RLock()
	defer fake.initMutex.RUnlock()
	fake.isConfiguredMutex.RLock()
	defer fake.isConfiguredMutex.RUnlock()
	fake.newSecretsFactoryMutex.RLock()
	defer fake.newSecretsFactoryMutex.RUnlock()
	fake.validateMutex.RLock()
	defer fake.validateMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeManager) recordInvocation(key string, args []interface{}) {
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

var _ creds.Manager = new(FakeManager)
