// Code generated by counterfeiter. DO NOT EDIT.
package operationsfakes

import (
	"sync"
)

type FakeReader struct {
	ReadStub        func(p []byte) (n int, err error)
	readMutex       sync.RWMutex
	readArgsForCall []struct {
		p []byte
	}
	readReturns struct {
		result1 int
		result2 error
	}
	readReturnsOnCall map[int]struct {
		result1 int
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeReader) Read(p []byte) (n int, err error) {
	var pCopy []byte
	if p != nil {
		pCopy = make([]byte, len(p))
		copy(pCopy, p)
	}
	fake.readMutex.Lock()
	ret, specificReturn := fake.readReturnsOnCall[len(fake.readArgsForCall)]
	fake.readArgsForCall = append(fake.readArgsForCall, struct {
		p []byte
	}{pCopy})
	fake.recordInvocation("Read", []interface{}{pCopy})
	fake.readMutex.Unlock()
	if fake.ReadStub != nil {
		return fake.ReadStub(p)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.readReturns.result1, fake.readReturns.result2
}

func (fake *FakeReader) ReadCallCount() int {
	fake.readMutex.RLock()
	defer fake.readMutex.RUnlock()
	return len(fake.readArgsForCall)
}

func (fake *FakeReader) ReadArgsForCall(i int) []byte {
	fake.readMutex.RLock()
	defer fake.readMutex.RUnlock()
	return fake.readArgsForCall[i].p
}

func (fake *FakeReader) ReadReturns(result1 int, result2 error) {
	fake.ReadStub = nil
	fake.readReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeReader) ReadReturnsOnCall(i int, result1 int, result2 error) {
	fake.ReadStub = nil
	if fake.readReturnsOnCall == nil {
		fake.readReturnsOnCall = make(map[int]struct {
			result1 int
			result2 error
		})
	}
	fake.readReturnsOnCall[i] = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeReader) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.readMutex.RLock()
	defer fake.readMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeReader) recordInvocation(key string, args []interface{}) {
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