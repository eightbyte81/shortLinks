package cache

import "fmt"

type LinksCache struct {
	linksCache *Cache
}

func NewLinksCache(linksCache *Cache) *LinksCache {
	return &LinksCache{linksCache: linksCache}
}

func (r *LinksCache) SetLinksInCache(shortLink string, defaultLink string) {
	r.linksCache.Mutex.Lock()
	defer r.linksCache.Mutex.Unlock()

	r.linksCache.Data[shortLink] = defaultLink
}

func (r *LinksCache) GetDefaultLinkFromCacheByShortLink(shortLink string) (string, error) {
	r.linksCache.Mutex.Lock()
	defer r.linksCache.Mutex.Unlock()

	if cacheData, found := r.linksCache.Data[shortLink]; found {
		return cacheData, nil
	}

	return "", fmt.Errorf("failed to get default link from cache by short link %s: link not found", shortLink)
}
