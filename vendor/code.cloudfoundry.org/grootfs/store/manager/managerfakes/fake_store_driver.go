// Code generated by counterfeiter. DO NOT EDIT.
package managerfakes

import (
	"sync"

	"code.cloudfoundry.org/grootfs/store/manager"
	"code.cloudfoundry.org/lager"
)

type FakeStoreDriver struct {
	ConfigureStoreStub        func(logger lager.Logger, storePath string, ownerUID, ownerGID int) error
	configureStoreMutex       sync.RWMutex
	configureStoreArgsForCall []struct {
		logger    lager.Logger
		storePath string
		ownerUID  int
		ownerGID  int
	}
	configureStoreReturns struct {
		result1 error
	}
	configureStoreReturnsOnCall map[int]struct {
		result1 error
	}
	ValidateFileSystemStub        func(logger lager.Logger, path string) error
	validateFileSystemMutex       sync.RWMutex
	validateFileSystemArgsForCall []struct {
		logger lager.Logger
		path   string
	}
	validateFileSystemReturns struct {
		result1 error
	}
	validateFileSystemReturnsOnCall map[int]struct {
		result1 error
	}
	InitFilesystemStub        func(logger lager.Logger, filesystemPath, storePath string) error
	initFilesystemMutex       sync.RWMutex
	initFilesystemArgsForCall []struct {
		logger         lager.Logger
		filesystemPath string
		storePath      string
	}
	initFilesystemReturns struct {
		result1 error
	}
	initFilesystemReturnsOnCall map[int]struct {
		result1 error
	}
	DeInitFilesystemStub        func(logger lager.Logger, storePath string) error
	deInitFilesystemMutex       sync.RWMutex
	deInitFilesystemArgsForCall []struct {
		logger    lager.Logger
		storePath string
	}
	deInitFilesystemReturns struct {
		result1 error
	}
	deInitFilesystemReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeStoreDriver) ConfigureStore(logger lager.Logger, storePath string, ownerUID int, ownerGID int) error {
	fake.configureStoreMutex.Lock()
	ret, specificReturn := fake.configureStoreReturnsOnCall[len(fake.configureStoreArgsForCall)]
	fake.configureStoreArgsForCall = append(fake.configureStoreArgsForCall, struct {
		logger    lager.Logger
		storePath string
		ownerUID  int
		ownerGID  int
	}{logger, storePath, ownerUID, ownerGID})
	fake.recordInvocation("ConfigureStore", []interface{}{logger, storePath, ownerUID, ownerGID})
	fake.configureStoreMutex.Unlock()
	if fake.ConfigureStoreStub != nil {
		return fake.ConfigureStoreStub(logger, storePath, ownerUID, ownerGID)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.configureStoreReturns.result1
}

func (fake *FakeStoreDriver) ConfigureStoreCallCount() int {
	fake.configureStoreMutex.RLock()
	defer fake.configureStoreMutex.RUnlock()
	return len(fake.configureStoreArgsForCall)
}

func (fake *FakeStoreDriver) ConfigureStoreArgsForCall(i int) (lager.Logger, string, int, int) {
	fake.configureStoreMutex.RLock()
	defer fake.configureStoreMutex.RUnlock()
	return fake.configureStoreArgsForCall[i].logger, fake.configureStoreArgsForCall[i].storePath, fake.configureStoreArgsForCall[i].ownerUID, fake.configureStoreArgsForCall[i].ownerGID
}

func (fake *FakeStoreDriver) ConfigureStoreReturns(result1 error) {
	fake.ConfigureStoreStub = nil
	fake.configureStoreReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStoreDriver) ConfigureStoreReturnsOnCall(i int, result1 error) {
	fake.ConfigureStoreStub = nil
	if fake.configureStoreReturnsOnCall == nil {
		fake.configureStoreReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.configureStoreReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeStoreDriver) ValidateFileSystem(logger lager.Logger, path string) error {
	fake.validateFileSystemMutex.Lock()
	ret, specificReturn := fake.validateFileSystemReturnsOnCall[len(fake.validateFileSystemArgsForCall)]
	fake.validateFileSystemArgsForCall = append(fake.validateFileSystemArgsForCall, struct {
		logger lager.Logger
		path   string
	}{logger, path})
	fake.recordInvocation("ValidateFileSystem", []interface{}{logger, path})
	fake.validateFileSystemMutex.Unlock()
	if fake.ValidateFileSystemStub != nil {
		return fake.ValidateFileSystemStub(logger, path)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.validateFileSystemReturns.result1
}

func (fake *FakeStoreDriver) ValidateFileSystemCallCount() int {
	fake.validateFileSystemMutex.RLock()
	defer fake.validateFileSystemMutex.RUnlock()
	return len(fake.validateFileSystemArgsForCall)
}

func (fake *FakeStoreDriver) ValidateFileSystemArgsForCall(i int) (lager.Logger, string) {
	fake.validateFileSystemMutex.RLock()
	defer fake.validateFileSystemMutex.RUnlock()
	return fake.validateFileSystemArgsForCall[i].logger, fake.validateFileSystemArgsForCall[i].path
}

func (fake *FakeStoreDriver) ValidateFileSystemReturns(result1 error) {
	fake.ValidateFileSystemStub = nil
	fake.validateFileSystemReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStoreDriver) ValidateFileSystemReturnsOnCall(i int, result1 error) {
	fake.ValidateFileSystemStub = nil
	if fake.validateFileSystemReturnsOnCall == nil {
		fake.validateFileSystemReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.validateFileSystemReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeStoreDriver) InitFilesystem(logger lager.Logger, filesystemPath string, storePath string) error {
	fake.initFilesystemMutex.Lock()
	ret, specificReturn := fake.initFilesystemReturnsOnCall[len(fake.initFilesystemArgsForCall)]
	fake.initFilesystemArgsForCall = append(fake.initFilesystemArgsForCall, struct {
		logger         lager.Logger
		filesystemPath string
		storePath      string
	}{logger, filesystemPath, storePath})
	fake.recordInvocation("InitFilesystem", []interface{}{logger, filesystemPath, storePath})
	fake.initFilesystemMutex.Unlock()
	if fake.InitFilesystemStub != nil {
		return fake.InitFilesystemStub(logger, filesystemPath, storePath)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.initFilesystemReturns.result1
}

func (fake *FakeStoreDriver) InitFilesystemCallCount() int {
	fake.initFilesystemMutex.RLock()
	defer fake.initFilesystemMutex.RUnlock()
	return len(fake.initFilesystemArgsForCall)
}

func (fake *FakeStoreDriver) InitFilesystemArgsForCall(i int) (lager.Logger, string, string) {
	fake.initFilesystemMutex.RLock()
	defer fake.initFilesystemMutex.RUnlock()
	return fake.initFilesystemArgsForCall[i].logger, fake.initFilesystemArgsForCall[i].filesystemPath, fake.initFilesystemArgsForCall[i].storePath
}

func (fake *FakeStoreDriver) InitFilesystemReturns(result1 error) {
	fake.InitFilesystemStub = nil
	fake.initFilesystemReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStoreDriver) InitFilesystemReturnsOnCall(i int, result1 error) {
	fake.InitFilesystemStub = nil
	if fake.initFilesystemReturnsOnCall == nil {
		fake.initFilesystemReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.initFilesystemReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeStoreDriver) DeInitFilesystem(logger lager.Logger, storePath string) error {
	fake.deInitFilesystemMutex.Lock()
	ret, specificReturn := fake.deInitFilesystemReturnsOnCall[len(fake.deInitFilesystemArgsForCall)]
	fake.deInitFilesystemArgsForCall = append(fake.deInitFilesystemArgsForCall, struct {
		logger    lager.Logger
		storePath string
	}{logger, storePath})
	fake.recordInvocation("DeInitFilesystem", []interface{}{logger, storePath})
	fake.deInitFilesystemMutex.Unlock()
	if fake.DeInitFilesystemStub != nil {
		return fake.DeInitFilesystemStub(logger, storePath)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.deInitFilesystemReturns.result1
}

func (fake *FakeStoreDriver) DeInitFilesystemCallCount() int {
	fake.deInitFilesystemMutex.RLock()
	defer fake.deInitFilesystemMutex.RUnlock()
	return len(fake.deInitFilesystemArgsForCall)
}

func (fake *FakeStoreDriver) DeInitFilesystemArgsForCall(i int) (lager.Logger, string) {
	fake.deInitFilesystemMutex.RLock()
	defer fake.deInitFilesystemMutex.RUnlock()
	return fake.deInitFilesystemArgsForCall[i].logger, fake.deInitFilesystemArgsForCall[i].storePath
}

func (fake *FakeStoreDriver) DeInitFilesystemReturns(result1 error) {
	fake.DeInitFilesystemStub = nil
	fake.deInitFilesystemReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeStoreDriver) DeInitFilesystemReturnsOnCall(i int, result1 error) {
	fake.DeInitFilesystemStub = nil
	if fake.deInitFilesystemReturnsOnCall == nil {
		fake.deInitFilesystemReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deInitFilesystemReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeStoreDriver) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.configureStoreMutex.RLock()
	defer fake.configureStoreMutex.RUnlock()
	fake.validateFileSystemMutex.RLock()
	defer fake.validateFileSystemMutex.RUnlock()
	fake.initFilesystemMutex.RLock()
	defer fake.initFilesystemMutex.RUnlock()
	fake.deInitFilesystemMutex.RLock()
	defer fake.deInitFilesystemMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeStoreDriver) recordInvocation(key string, args []interface{}) {
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

var _ manager.StoreDriver = new(FakeStoreDriver)