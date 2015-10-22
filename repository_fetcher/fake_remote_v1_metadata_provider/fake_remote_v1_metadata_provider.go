// This file was generated by counterfeiter
package fake_remote_v1_metadata_provider

import (
	"sync"

	"github.com/cloudfoundry-incubator/garden-shed/repository_fetcher"
)

type FakeRemoteV1MetadataProvider struct {
	ProvideMetadataStub        func(*repository_fetcher.FetchRequest) (*repository_fetcher.ImageV1Metadata, error)
	provideMetadataMutex       sync.RWMutex
	provideMetadataArgsForCall []struct {
		arg1 *repository_fetcher.FetchRequest
	}
	provideMetadataReturns struct {
		result1 *repository_fetcher.ImageV1Metadata
		result2 error
	}
}

func (fake *FakeRemoteV1MetadataProvider) ProvideMetadata(arg1 *repository_fetcher.FetchRequest) (*repository_fetcher.ImageV1Metadata, error) {
	fake.provideMetadataMutex.Lock()
	fake.provideMetadataArgsForCall = append(fake.provideMetadataArgsForCall, struct {
		arg1 *repository_fetcher.FetchRequest
	}{arg1})
	fake.provideMetadataMutex.Unlock()
	if fake.ProvideMetadataStub != nil {
		return fake.ProvideMetadataStub(arg1)
	} else {
		return fake.provideMetadataReturns.result1, fake.provideMetadataReturns.result2
	}
}

func (fake *FakeRemoteV1MetadataProvider) ProvideMetadataCallCount() int {
	fake.provideMetadataMutex.RLock()
	defer fake.provideMetadataMutex.RUnlock()
	return len(fake.provideMetadataArgsForCall)
}

func (fake *FakeRemoteV1MetadataProvider) ProvideMetadataArgsForCall(i int) *repository_fetcher.FetchRequest {
	fake.provideMetadataMutex.RLock()
	defer fake.provideMetadataMutex.RUnlock()
	return fake.provideMetadataArgsForCall[i].arg1
}

func (fake *FakeRemoteV1MetadataProvider) ProvideMetadataReturns(result1 *repository_fetcher.ImageV1Metadata, result2 error) {
	fake.ProvideMetadataStub = nil
	fake.provideMetadataReturns = struct {
		result1 *repository_fetcher.ImageV1Metadata
		result2 error
	}{result1, result2}
}

var _ repository_fetcher.RemoteV1MetadataProvider = new(FakeRemoteV1MetadataProvider)