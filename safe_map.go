/*
 * 纸喵软件
 * Copyright (c) 2017~2020 http://zhimiao.org All rights reserved.
 * Author: 倒霉狐狸 <mail@xiaoliu.org>
 * Date: 2020/3/3 下午4:26
 */

package gutils

import "sync"

// SafeStringMap 安全map
type SafeStringMap struct {
	sync.RWMutex
	Map map[string]string
}

// NewSafeStringMap  创建
func NewSafeStringMap() *SafeStringMap {
	st := new(SafeStringMap)
	st.Map = make(map[string]string)
	return st
}

// GET  获取
func (st *SafeStringMap) GET(key string) string {
	st.RLock()
	value := st.Map[key]
	st.RUnlock()
	return value
}

// SET 设置
func (st *SafeStringMap) SET(key string, value string) {
	st.Lock()
	st.Map[key] = value
	st.Unlock()
}

// SETNX map锁
func (st *SafeStringMap) SETNX(key string, value string) (ok bool) {
	ok = false
	st.Lock()
	if _, ok = st.Map[key]; !ok {
		st.Map[key] = value
		ok = true
	}
	st.Unlock()
	return ok
}

// DEL 删除
func (st *SafeStringMap) DEL(key string) {
	st.Lock()
	delete(st.Map, key)
	st.Unlock()
}
