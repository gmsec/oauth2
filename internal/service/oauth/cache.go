package oauth

/**
缓存：支持redis，内存缓存
*/
import (
	"oauth2/internal/config"
	"oauth2/internal/model"
	"time"

	"github.com/xxjwxc/public/mycache"
	"github.com/xxjwxc/public/mylog"
	"github.com/xxjwxc/public/myredis"
)

const (
	_prefixOauth2Tbl = "Oauth2Tbl:"
	_prefixAccess    = "access:"
	_prefixRefresh   = "refresh:"
)

// var _once sync.Once
var _cache mycache.CacheIFS

func init() {
	conf := config.GetRedisDbInfo()
	if len(conf.Addrs) == 0 { // 使用内存缓存
		_cache = mycache.NewCache(conf.GroupName)
		return
	}

	// 使用redis缓存
	cn := myredis.InitRedis(myredis.WithAddr(conf.Addrs...), myredis.WithClientName("oauth2"),
		// WithPool(2, 2),
		myredis.WithTimeout(3*time.Second), myredis.WithReadTimeout(3*time.Second), myredis.WithWriteTimeout(3*time.Second),
		myredis.WithPwd(conf.Password), myredis.WithGroupName(conf.GroupName), myredis.WithDB(conf.DB))
	var err error
	_cache, err = myredis.NewRedis(cn)
	if err != nil {
		mylog.Error(err)
	}
}

// GetCacheOauth2TblByKey 获取一个key
func GetCacheOauth2TblByKey(appID string) *model.Oauth2Tbl {
	key := _prefixOauth2Tbl + appID
	if _cache.IsExist(key) {
		var tmp model.Oauth2Tbl
		err := _cache.Value(key, &tmp)
		if err == nil { // 没有错误
			return &tmp
		}
		mylog.Error(err)
	}

	return nil
}

// DeleteCacheOauth2TblByKey 获取一个key
func DeleteCacheOauth2TblByKey(appID string) error {
	key := _prefixOauth2Tbl + appID
	return _cache.Delete(key)
}

// AddCacheOauth2TblByKey 添加一个key
func AddCacheOauth2TblByKey(appID string, v *model.Oauth2Tbl) error {
	key := _prefixOauth2Tbl + appID
	return _cache.Add(key, v, time.Duration(v.TokenExpireTime)*time.Second)
}

// AddCacheToken 添加一个cache缓存
func AddCacheToken(token TokenCache, tag string) error {
	key := _prefixAccess + token.Token
	if tag != _prefixAccess {
		key = _prefixRefresh + token.Token
	}

	now := time.Now().Unix()
	off := token.ExpireTime - now
	if off <= 0 { // 已过期
		return nil
	}

	return _cache.Add(key, token, time.Duration(off)*time.Second)
}

// GetCacheToken 添加一个cache缓存
func GetCacheToken(token, tag string) *TokenCache {
	key := _prefixAccess + token
	if tag != _prefixAccess {
		key = _prefixRefresh + token
	}

	out := &TokenCache{}

	if _cache.IsExist(key) {
		err := _cache.Value(key, out)
		if err != nil {
			mylog.Error(err)
		}
		return out
	}
	return nil
}
