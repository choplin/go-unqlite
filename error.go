package unqlite

/*
 #include "unqlite.h"
*/
import "C"

type ErrCode int

const (
	OK             ErrCode = C.UNQLITE_OK             /* Successful result */
	NOMEM          ErrCode = C.UNQLITE_NOMEM          /* Out of memory */
	ABORT          ErrCode = C.UNQLITE_ABORT          /* Another thread have released this instance */
	IOERR          ErrCode = C.UNQLITE_IOERR          /* IO error */
	CORRUPT        ErrCode = C.UNQLITE_CORRUPT        /* Corrupt pointer */
	LOCKED         ErrCode = C.UNQLITE_LOCKED         /* Forbidden Operation */
	BUSY           ErrCode = C.UNQLITE_BUSY           /* The database file is locked */
	DONE           ErrCode = C.UNQLITE_DONE           /* Operation done */
	PERM           ErrCode = C.UNQLITE_PERM           /* Permission error */
	NOTIMPLEMENint ErrCode = C.UNQLITE_NOTIMPLEMENTED /* Method not implemented by the underlying Key/Value storage engine */
	NOTFOUND       ErrCode = C.UNQLITE_NOTFOUND       /* No such record */
	NOOP           ErrCode = C.UNQLITE_NOOP           /* No such method */
	INVALID        ErrCode = C.UNQLITE_INVALID        /* Invalid parameter */
	EOF            ErrCode = C.UNQLITE_EOF            /* End Of Input */
	UNKNOWN        ErrCode = C.UNQLITE_UNKNOWN        /* Unknown configuration option */
	LIMIT          ErrCode = C.UNQLITE_LIMIT          /* Database limit reached */
	EXISTS         ErrCode = C.UNQLITE_EXISTS         /* Record exists */
	EMPTY          ErrCode = C.UNQLITE_EMPTY          /* Empty record */
	COMPILE_ERRint ErrCode = C.UNQLITE_COMPILE_ERR    /* Compilation error */
	VM_ERR         ErrCode = C.UNQLITE_VM_ERR         /* Virtual machine error */
	FULL           ErrCode = C.UNQLITE_FULL           /* Full database (unlikely) */
	CANTOPEN       ErrCode = C.UNQLITE_CANTOPEN       /* Unable to open the database file */
	READ_ONLY      ErrCode = C.UNQLITE_READ_ONLY      /* Read only Key/Value storage engine */
	LOCKERR        ErrCode = C.UNQLITE_LOCKERR        /* Locking protocol error */
)

var errorString = map[ErrCode]string{
	NOMEM:          "Out of memory",
	ABORT:          "Another thread have released this instance",
	IOERR:          "IO error",
	CORRUPT:        "Corrupt pointer",
	LOCKED:         "Forbidden Operation",
	BUSY:           "The database file is locked",
	DONE:           "Operation done",
	PERM:           "Permission error",
	NOTIMPLEMENint: "Method not implemented by the underlying Key/Value storage engine",
	NOTFOUND:       "No such record",
	NOOP:           "No such method",
	INVALID:        "Invalid parameter",
	EOF:            "End Of Input",
	UNKNOWN:        "Unknown configuration option",
	LIMIT:          "Database limit reached",
	EXISTS:         "Record exists",
	EMPTY:          "Empty record",
	COMPILE_ERRint: "Compilation error",
	VM_ERR:         "Virtual machine error",
	FULL:           "Full database (unlikely)",
	CANTOPEN:       "Unable to open the database file",
	READ_ONLY:      "Read only Key/Value storage engine",
	LOCKERR:        "Locking protocol error",
}

func (err ErrCode) Error() string {
	return errorString[err]
}
