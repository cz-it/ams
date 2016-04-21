/**
* Author: CZ cz.theng@gmail.com
 */

package uidmgr

import (
	"sync"
)

const (
	subBeginUID uint32 = uint32(134217728) // 2^27
)

var uidmgr UIDMgr

//NewUID create a new UID
func NewUID(platform uint8, appID uint16) (UID uint64, err error) {
	return uidmgr.NewUID(platform, appID)
}

//UIDMgr is a UID Manager
type UIDMgr struct {
	app map[uint32]uint32
	mtx sync.Mutex
}

//NewUID create a new UID
func (mgr *UIDMgr) NewUID(platform uint8, appID uint16) (UID uint64, err error) {
	mgr.mtx.Lock()
	mgr.mtx.Unlock()
	if mgr.app == nil {
		mgr.app = make(map[uint32]uint32)
	}
	appid := uint32(platform)<<16 | uint32(appID)
	println("appid:", appid)
	if subid, ok := mgr.app[appid]; ok {
		mgr.app[appid] = subid + 1
	} else {
		mgr.app[appid] = subBeginUID
	}
	err = nil
	UID = uint64(appid)<<32 | uint64(mgr.app[appid])
	return
}
