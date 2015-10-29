// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/cloudfoundry-incubator/garden-shed/docker_drivers/aufs"
)

type FakeBackingStoreMgr struct {
	CreateStub        func(id string, quota int64) (string, error)
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		id    string
		quota int64
	}
	createReturns struct {
		result1 string
		result2 error
	}
	DeleteStub        func(id string) error
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		id string
	}
	deleteReturns struct {
		result1 error
	}
}

func (fake *FakeBackingStoreMgr) Create(id string, quota int64) (string, error) {
	fake.createMutex.Lock()
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		id    string
		quota int64
	}{id, quota})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub(id, quota)
	} else {
		return fake.createReturns.result1, fake.createReturns.result2
	}
}

func (fake *FakeBackingStoreMgr) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeBackingStoreMgr) CreateArgsForCall(i int) (string, int64) {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return fake.createArgsForCall[i].id, fake.createArgsForCall[i].quota
}

func (fake *FakeBackingStoreMgr) CreateReturns(result1 string, result2 error) {
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeBackingStoreMgr) Delete(id string) error {
	fake.deleteMutex.Lock()
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		id string
	}{id})
	fake.deleteMutex.Unlock()
	if fake.DeleteStub != nil {
		return fake.DeleteStub(id)
	} else {
		return fake.deleteReturns.result1
	}
}

func (fake *FakeBackingStoreMgr) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeBackingStoreMgr) DeleteArgsForCall(i int) string {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return fake.deleteArgsForCall[i].id
}

func (fake *FakeBackingStoreMgr) DeleteReturns(result1 error) {
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 error
	}{result1}
}

var _ aufs.BackingStoreMgr = new(FakeBackingStoreMgr)
