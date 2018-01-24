// Code generated by counterfeiter. DO NOT EDIT.
package imagepullerfakes

import (
	"io"
	"net/url"
	"sync"

	"code.cloudfoundry.org/groot/imagepuller"
	"code.cloudfoundry.org/lager"
)

type FakeFetcher struct {
	ImageInfoStub        func(logger lager.Logger, imageURL *url.URL) (imagepuller.ImageInfo, error)
	imageInfoMutex       sync.RWMutex
	imageInfoArgsForCall []struct {
		logger   lager.Logger
		imageURL *url.URL
	}
	imageInfoReturns struct {
		result1 imagepuller.ImageInfo
		result2 error
	}
	imageInfoReturnsOnCall map[int]struct {
		result1 imagepuller.ImageInfo
		result2 error
	}
	StreamBlobStub        func(logger lager.Logger, imageURL *url.URL, layerInfo imagepuller.LayerInfo) (io.ReadCloser, int64, error)
	streamBlobMutex       sync.RWMutex
	streamBlobArgsForCall []struct {
		logger    lager.Logger
		imageURL  *url.URL
		layerInfo imagepuller.LayerInfo
	}
	streamBlobReturns struct {
		result1 io.ReadCloser
		result2 int64
		result3 error
	}
	streamBlobReturnsOnCall map[int]struct {
		result1 io.ReadCloser
		result2 int64
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeFetcher) ImageInfo(logger lager.Logger, imageURL *url.URL) (imagepuller.ImageInfo, error) {
	fake.imageInfoMutex.Lock()
	ret, specificReturn := fake.imageInfoReturnsOnCall[len(fake.imageInfoArgsForCall)]
	fake.imageInfoArgsForCall = append(fake.imageInfoArgsForCall, struct {
		logger   lager.Logger
		imageURL *url.URL
	}{logger, imageURL})
	fake.recordInvocation("ImageInfo", []interface{}{logger, imageURL})
	fake.imageInfoMutex.Unlock()
	if fake.ImageInfoStub != nil {
		return fake.ImageInfoStub(logger, imageURL)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.imageInfoReturns.result1, fake.imageInfoReturns.result2
}

func (fake *FakeFetcher) ImageInfoCallCount() int {
	fake.imageInfoMutex.RLock()
	defer fake.imageInfoMutex.RUnlock()
	return len(fake.imageInfoArgsForCall)
}

func (fake *FakeFetcher) ImageInfoArgsForCall(i int) (lager.Logger, *url.URL) {
	fake.imageInfoMutex.RLock()
	defer fake.imageInfoMutex.RUnlock()
	return fake.imageInfoArgsForCall[i].logger, fake.imageInfoArgsForCall[i].imageURL
}

func (fake *FakeFetcher) ImageInfoReturns(result1 imagepuller.ImageInfo, result2 error) {
	fake.ImageInfoStub = nil
	fake.imageInfoReturns = struct {
		result1 imagepuller.ImageInfo
		result2 error
	}{result1, result2}
}

func (fake *FakeFetcher) ImageInfoReturnsOnCall(i int, result1 imagepuller.ImageInfo, result2 error) {
	fake.ImageInfoStub = nil
	if fake.imageInfoReturnsOnCall == nil {
		fake.imageInfoReturnsOnCall = make(map[int]struct {
			result1 imagepuller.ImageInfo
			result2 error
		})
	}
	fake.imageInfoReturnsOnCall[i] = struct {
		result1 imagepuller.ImageInfo
		result2 error
	}{result1, result2}
}

func (fake *FakeFetcher) StreamBlob(logger lager.Logger, imageURL *url.URL, layerInfo imagepuller.LayerInfo) (io.ReadCloser, int64, error) {
	fake.streamBlobMutex.Lock()
	ret, specificReturn := fake.streamBlobReturnsOnCall[len(fake.streamBlobArgsForCall)]
	fake.streamBlobArgsForCall = append(fake.streamBlobArgsForCall, struct {
		logger    lager.Logger
		imageURL  *url.URL
		layerInfo imagepuller.LayerInfo
	}{logger, imageURL, layerInfo})
	fake.recordInvocation("StreamBlob", []interface{}{logger, imageURL, layerInfo})
	fake.streamBlobMutex.Unlock()
	if fake.StreamBlobStub != nil {
		return fake.StreamBlobStub(logger, imageURL, layerInfo)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fake.streamBlobReturns.result1, fake.streamBlobReturns.result2, fake.streamBlobReturns.result3
}

func (fake *FakeFetcher) StreamBlobCallCount() int {
	fake.streamBlobMutex.RLock()
	defer fake.streamBlobMutex.RUnlock()
	return len(fake.streamBlobArgsForCall)
}

func (fake *FakeFetcher) StreamBlobArgsForCall(i int) (lager.Logger, *url.URL, imagepuller.LayerInfo) {
	fake.streamBlobMutex.RLock()
	defer fake.streamBlobMutex.RUnlock()
	return fake.streamBlobArgsForCall[i].logger, fake.streamBlobArgsForCall[i].imageURL, fake.streamBlobArgsForCall[i].layerInfo
}

func (fake *FakeFetcher) StreamBlobReturns(result1 io.ReadCloser, result2 int64, result3 error) {
	fake.StreamBlobStub = nil
	fake.streamBlobReturns = struct {
		result1 io.ReadCloser
		result2 int64
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeFetcher) StreamBlobReturnsOnCall(i int, result1 io.ReadCloser, result2 int64, result3 error) {
	fake.StreamBlobStub = nil
	if fake.streamBlobReturnsOnCall == nil {
		fake.streamBlobReturnsOnCall = make(map[int]struct {
			result1 io.ReadCloser
			result2 int64
			result3 error
		})
	}
	fake.streamBlobReturnsOnCall[i] = struct {
		result1 io.ReadCloser
		result2 int64
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeFetcher) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.imageInfoMutex.RLock()
	defer fake.imageInfoMutex.RUnlock()
	fake.streamBlobMutex.RLock()
	defer fake.streamBlobMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeFetcher) recordInvocation(key string, args []interface{}) {
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

var _ imagepuller.Fetcher = new(FakeFetcher)