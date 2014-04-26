package unqlite

/*
 #include <stdlib.h>

 #include "unqlite.h"

 extern int
 kv_fetch_callback_func(void *data, unsigned int dataLen, void *userData);

 // you cannot export xConsumer directory because a function with a const pointer argument cannot be exported.
 typedef int (*xConsumer)(const void *pData, unsigned int iDataLen, void *pUserData);

 static inline int
 _unqlite_kv_fetch_callback(unqlite *pDb, const void *pKey,int nKeyLen, void *pUserData) {
   return unqlite_kv_fetch_callback(pDb, pKey, nKeyLen, (xConsumer) kv_fetch_callback_func, pUserData);
 }
*/
import "C"

import (
	"unsafe"
)

const (
	O_READONLY        uint = C.UNQLITE_OPEN_READONLY        /* Read only mode. Ok for [unqlite_open] */
	O_READWRITE       uint = C.UNQLITE_OPEN_READWRITE       /* Ok for [unqlite_open] */
	O_CREATE          uint = C.UNQLITE_OPEN_CREATE          /* Ok for [unqlite_open] */
	O_EXCLUSIVE       uint = C.UNQLITE_OPEN_EXCLUSIVE       /* VFS only */
	O_TEMP_DB         uint = C.UNQLITE_OPEN_TEMP_DB         /* VFS only */
	O_NOMUTEX         uint = C.UNQLITE_OPEN_NOMUTEX         /* Ok for [unqlite_open] */
	O_OMIT_JOURNALING uint = C.UNQLITE_OPEN_OMIT_JOURNALING /* Omit journaling for this database. Ok for [unqlite_open] */
	O_IN_MEMORY       uint = C.UNQLITE_OPEN_IN_MEMORY       /* An in memory database. Ok for [unqlite_open]*/
	O_MMAP            uint = C.UNQLITE_OPEN_MMAP            /* Obtain a memory view of the whole file. Ok for [unqlite_open] */
)

type Unqlite struct {
	db *C.unqlite
}

// Open a database and return unqite object.
// If fileName is ":mem:", then a private, in-memory database is created for the connection.
// See: http://unqlite.org/c_api/unqlite_open.html
func OpenUnqlite(fileName string, mode uint) (*Unqlite, error) {
	cname := C.CString(fileName)
	defer C.free(unsafe.Pointer(cname))

	var db *C.unqlite
	if rc := C.unqlite_open(&db, cname, C.uint(mode)); rc != C.UNQLITE_OK {
		return nil, ErrCode(rc)
	}

	return &Unqlite{db}, nil
}

// Close the database.
// See: http://unqlite.org/c_api/unqlite_close.html
func (u *Unqlite) Close() error {
	if rc := C.unqlite_close(u.db); rc != C.UNQLITE_OK {
		return ErrCode(rc)
	}

	return nil
}

/*
 * Key-Value Store Interface
 */

// KvStore write a new record into the database. If the record does not exists, it is created. Otherwise, it is replaced.
// See: http://unqlite.org/c_api/unqlite_kv_store.html
func (u *Unqlite) KvStore(key []byte, value []byte) error {
	if rc := C.unqlite_kv_store(u.db, unsafe.Pointer(&key[0]), C.int(len(key)), unsafe.Pointer(&value[0]), C.unqlite_int64(len(value))); rc != C.UNQLITE_OK {
		return ErrCode(rc)
	}

	return nil
}

// KvStore write a new record into the database. If the record does not exists, it is created. Otherwise, the new data chunk is appended to the end of the old chunk.
// See: http://unqlite.org/c_api/unqlite_kv_append.html
func (u *Unqlite) KvAppend(key []byte, value []byte) error {
	if rc := C.unqlite_kv_append(u.db, unsafe.Pointer(&key[0]), C.int(len(key)), unsafe.Pointer(&value[0]), C.unqlite_int64(len(value))); rc != C.UNQLITE_OK {
		return ErrCode(rc)
	}

	return nil
}

// KvFetch a record from the database
// See: http://unqlite.org/c_api/unqlite_kv_fetch.html
func (u *Unqlite) KvFetch(key []byte) ([]byte, error) {
	var n C.unqlite_int64

	if rc := C.unqlite_kv_fetch(u.db, unsafe.Pointer(&key[0]), C.int(len(key)), nil, &n); rc != C.UNQLITE_OK {
		return nil, ErrCode(rc)
	}

	buf := make([]byte, int64(n))
	if rc := C.unqlite_kv_fetch(u.db, unsafe.Pointer(&key[0]), C.int(len(key)), unsafe.Pointer(&buf[0]), &n); rc != C.UNQLITE_OK {
		return nil, ErrCode(rc)
	}

	return buf, nil
}

type kvFetchCallbackFunc func([]byte) ErrCode

//export kv_fetch_callback_func
func kv_fetch_callback_func(data unsafe.Pointer, dataLen C.uint, userData unsafe.Pointer) C.int {
	f := *(*kvFetchCallbackFunc)(userData)
	d := C.GoBytes(data, C.int(int(uint(dataLen))))
	return C.int(f(d))
}

// TODO: temporary variables might be destroyed by GC?
func (u *Unqlite) KvFetchCallback(key []byte, f kvFetchCallbackFunc) error {
	if rc := C._unqlite_kv_fetch_callback(u.db, unsafe.Pointer(&key[0]), C.int(len(key)), unsafe.Pointer(&f)); rc != C.UNQLITE_OK {
		return ErrCode(rc)
	}

	return nil
}

// KvDelete remove a particular record from the database, you can use this high-level thread-safe routine to perform the deletion.
// See: http://unqlite.org/c_api/unqlite_kv_delete.html
func (u *Unqlite) KvDelete(key []byte) error {
	if rc := C.unqlite_kv_delete(u.db, unsafe.Pointer(&key[0]), C.int(len(key))); rc != C.UNQLITE_OK {
		return ErrCode(rc)
	}

	return nil
}
