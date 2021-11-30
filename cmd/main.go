package main

import (
	"gin-essential/inject"
	"net/http"
	"time"
)

func main() {
	// 初始化依赖注入器
	injector, injectorCleanFunc, err := inject.GenInjector()
	if err != nil {
		panic(err)
	}
	srv := &http.Server{
		Addr:         ":8080",
		Handler:      injector.Engine,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	err = srv.ListenAndServe()
	if err != nil {
		panic(err)
	}
	injectorCleanFunc()
}

// // v - 需要读取的数据对象
// // key - 缓存key
// // query - 用来从DB读取完整数据的方法
// // cacheVal - 用来写缓存的方法
// func (c cacheNode) doTake(v interface{}, key string, query func(v interface{}) error,
// 	cacheVal func(v interface{}) error) error {
// 	// 用barrier来防止缓存击穿，确保一个进程内只有一个请求去加载key对应的数据
// 	val, fresh, err := c.barrier.DoEx(key, func() (interface{}, error) {
// 		// 从cache里读取数据
// 		if err := c.doGetCache(key, v); err != nil {
// 			// 如果是预先放进来的placeholder（用来防止缓存穿透）的，那么就返回预设的errNotFound
// 			// 如果是未知错误，那么就直接返回，因为我们不能放弃缓存出错而直接把所有请求去请求DB，
// 			// 这样在高并发的场景下会把DB打挂掉的
// 			if err == errPlaceholder {
// 				return nil, c.errNotFound
// 			} else if err != c.errNotFound {
// 				// why we just return the error instead of query from db,
// 				// because we don't allow the disaster pass to the DBs.
// 				// fail fast, in case we bring down the dbs.
// 				return nil, err
// 			}

// 			// 请求DB
// 			// 如果返回的error是errNotFound，那么我们就需要在缓存里设置placeholder，防止缓存穿透
// 			if err = query(v); err == c.errNotFound {
// 				if err = c.setCacheWithNotFound(key); err != nil {
// 					logx.Error(err)
// 				}

// 				return nil, c.errNotFound
// 			} else if err != nil {
// 				// 统计DB失败
// 				c.stat.IncrementDbFails()
// 				return nil, err
// 			}

// 			// 把数据写入缓存
// 			if err = cacheVal(v); err != nil {
// 				logx.Error(err)
// 			}
// 		}

// 		// 返回json序列化的数据
// 		return jsonx.Marshal(v)
// 	})
// 	if err != nil {
// 		return err
// 	}
// 	if fresh {
// 		return nil
// 	}

// 	// got the result from previous ongoing query
// 	c.stat.IncrementTotal()
// 	c.stat.IncrementHit()

// 	// 把数据写入到传入的v对象里
// 	return jsonx.Unmarshal(val.([]byte), v)
// }
