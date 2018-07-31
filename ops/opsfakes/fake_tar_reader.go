// Code generated by counterfeiter. DO NOT EDIT.
package opsfakes

import (
	"sync"
)

type FakeTarReader struct {
	ReadFileStub        func(string) ([]byte, error)
	readFileMutex       sync.RWMutex
	readFileArgsForCall []struct {
		arg1 string
	}
	readFileReturns struct {
		result1 []byte
		result2 error
	}
	readFileReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	FileMd5sStub        func() (map[string]string, error)
	fileMd5sMutex       sync.RWMutex
	fileMd5sArgsForCall []struct{}
	fileMd5sReturns     struct {
		result1 map[string]string
		result2 error
	}
	fileMd5sReturnsOnCall map[int]struct {
		result1 map[string]string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeTarReader) ReadFile(arg1 string) ([]byte, error) {
	fake.readFileMutex.Lock()
	ret, specificReturn := fake.readFileReturnsOnCall[len(fake.readFileArgsForCall)]
	fake.readFileArgsForCall = append(fake.readFileArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("ReadFile", []interface{}{arg1})
	fake.readFileMutex.Unlock()
	if fake.ReadFileStub != nil {
		return fake.ReadFileStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.readFileReturns.result1, fake.readFileReturns.result2
}

func (fake *FakeTarReader) ReadFileCallCount() int {
	fake.readFileMutex.RLock()
	defer fake.readFileMutex.RUnlock()
	return len(fake.readFileArgsForCall)
}

func (fake *FakeTarReader) ReadFileArgsForCall(i int) string {
	fake.readFileMutex.RLock()
	defer fake.readFileMutex.RUnlock()
	return fake.readFileArgsForCall[i].arg1
}

func (fake *FakeTarReader) ReadFileReturns(result1 []byte, result2 error) {
	fake.ReadFileStub = nil
	fake.readFileReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeTarReader) ReadFileReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.ReadFileStub = nil
	if fake.readFileReturnsOnCall == nil {
		fake.readFileReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.readFileReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeTarReader) FileMd5s() (map[string]string, error) {
	fake.fileMd5sMutex.Lock()
	ret, specificReturn := fake.fileMd5sReturnsOnCall[len(fake.fileMd5sArgsForCall)]
	fake.fileMd5sArgsForCall = append(fake.fileMd5sArgsForCall, struct{}{})
	fake.recordInvocation("FileMd5s", []interface{}{})
	fake.fileMd5sMutex.Unlock()
	if fake.FileMd5sStub != nil {
		return fake.FileMd5sStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.fileMd5sReturns.result1, fake.fileMd5sReturns.result2
}

func (fake *FakeTarReader) FileMd5sCallCount() int {
	fake.fileMd5sMutex.RLock()
	defer fake.fileMd5sMutex.RUnlock()
	return len(fake.fileMd5sArgsForCall)
}

func (fake *FakeTarReader) FileMd5sReturns(result1 map[string]string, result2 error) {
	fake.FileMd5sStub = nil
	fake.fileMd5sReturns = struct {
		result1 map[string]string
		result2 error
	}{result1, result2}
}

func (fake *FakeTarReader) FileMd5sReturnsOnCall(i int, result1 map[string]string, result2 error) {
	fake.FileMd5sStub = nil
	if fake.fileMd5sReturnsOnCall == nil {
		fake.fileMd5sReturnsOnCall = make(map[int]struct {
			result1 map[string]string
			result2 error
		})
	}
	fake.fileMd5sReturnsOnCall[i] = struct {
		result1 map[string]string
		result2 error
	}{result1, result2}
}

func (fake *FakeTarReader) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.readFileMutex.RLock()
	defer fake.readFileMutex.RUnlock()
	fake.fileMd5sMutex.RLock()
	defer fake.fileMd5sMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeTarReader) recordInvocation(key string, args []interface{}) {
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
