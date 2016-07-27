package mysql

const (
	MinProtocolVersion byte   = 10
	MaxPayloadLen      int    = 1<<24 - 1
	TimeFormat         string = "2006-01-02 15:04:05"
	ServerVersion      string = "5.5.40-九华山-0.1"
)

const (
	OK_HEADER          byte = 0x00
	ERR_HEADER         byte = 0xff
	EOF_HEADER         byte = 0xfe
	LocalInFile_HEADER byte = 0xfb
)

const (
	SERVER_STATUS_IN_TRANS             uint16 = 0x0001
	SERVER_STATUS_AUTOCOMMIT           uint16 = 0x0002
	SERVER_MORE_RESULTS_EXISTS         uint16 = 0x0008
	SERVER_STATUS_NO_GOOD_INDEX_USED   uint16 = 0x0010
	SERVER_STATUS_NO_INDEX_USED        uint16 = 0x0020
	SERVER_STATUS_CURSOR_EXISTS        uint16 = 0x0040
	SERVER_STATUS_LAST_ROW_SEND        uint16 = 0x0080
	SERVER_STATUS_DB_DROPPED           uint16 = 0x0100
	SERVER_STATUS_NO_BACKSLASH_ESCAPED uint16 = 0x0200
	SERVER_STATUS_METADATA_CHANGED     uint16 = 0x0400
	SERVER_QUERY_WAS_SLOW              uint16 = 0x0800
	SERVER_PS_OUT_PARAMS               uint16 = 0x1000
)

const (
	COM_SLEEP byte = iota
	COM_QUIT
	COM_INIT_DB
	COM_QUERY
	COM_FIELD_LIST
	COM_CREATE_DB
	COM_DROP_DB
	COM_REFRESH
	COM_SHUTDOWN
	COM_STATISTICS
	COM_PROCESS_INFO
	COM_CONNECT
	COM_PROCESS_KILL
	COM_DEBUG
	COM_PING
	COM_TIME
	COM_DELAYED_INSERT
	COM_CHANGE_USER
	COM_BINLOG_DUMP
	COM_TABLE_DUMP
	COM_CONNECT_OUT
	COM_REGISTER_SLAVE
	COM_STMT_PREPARE
	COM_STMT_EXECUTE
	COM_STMT_SEND_LONG_DATA
	COM_STMT_CLOSE
	COM_STMT_RESET
	COM_SET_OPTION
	COM_STMT_FETCH
	COM_DAEMON
	COM_BINLOG_DUMP_GTID
	COM_RESET_CONNECTION
)

const (
	CLIENT_LONG_PASSWORD uint32 = 1 << iota
	CLIENT_FOUND_ROWS
	CLIENT_LONG_FLAG
	CLIENT_CONNECT_WITH_DB
	CLIENT_NO_SCHEMA
	CLIENT_COMPRESS
	CLIENT_ODBC
	CLIENT_LOCAL_FILES
	CLIENT_IGNORE_SPACE
	CLIENT_PROTOCOL_41
	CLIENT_INTERACTIVE
	CLIENT_SSL
	CLIENT_IGNORE_SIGPIPE
	CLIENT_TRANSACTIONS
	CLIENT_RESERVED
	CLIENT_SECURE_CONNECTION
	CLIENT_MULTI_STATEMENTS
	CLIENT_MULTI_RESULTS
	CLIENT_PS_MULTI_RESULTS
	CLIENT_PLUGIN_AUTH
	CLIENT_CONNECT_ATTRS
	CLIENT_PLUGIN_AUTH_LENENC_CLIENT_DATA
)

const (
	MYSQL_TYPE_DECIMAL     uint8 = 0
	MYSQL_TYPE_TINY        uint8 = 1
	MYSQL_TYPE_SHORT       uint8 = 2
	MYSQL_TYPE_LONG        uint8 = 3
	MYSQL_TYPE_FLOAT       uint8 = 4
	MYSQL_TYPE_DOUBLE      uint8 = 5
	MYSQL_TYPE_NULL        uint8 = 6
	MYSQL_TYPE_TIMESTAMP   uint8 = 7
	MYSQL_TYPE_LONGLONG    uint8 = 8
	MYSQL_TYPE_INT24       uint8 = 9
	MYSQL_TYPE_DATE        uint8 = 10
	MYSQL_TYPE_TIME        uint8 = 11
	MYSQL_TYPE_DATETIME    uint8 = 12
	MYSQL_TYPE_YEAR        uint8 = 13
	MYSQL_TYPE_NEWDATE     uint8 = 14
	MYSQL_TYPE_VARCHAR     uint8 = 15
	MYSQL_TYPE_BIT         uint8 = 16
	MYSQL_TYPE_TIMESTAMP2  uint8 = 17
	MYSQL_TYPE_DATETIME2   uint8 = 18
	MYSQL_TYPE_TIME2       uint8 = 19
	MYSQL_TYPE_NEWDECIMAL  uint8 = 246
	MYSQL_TYPE_ENUM        uint8 = 247
	MYSQL_TYPE_SET         uint8 = 248
	MYSQL_TYPE_TINY_BLOB   uint8 = 249
	MYSQL_TYPE_MEDIUM_BLOB uint8 = 250
	MYSQL_TYPE_LONG_BLOB   uint8 = 251
	MYSQL_TYPE_BLOB        uint8 = 252
	MYSQL_TYPE_VAR_STRING  uint8 = 253
	MYSQL_TYPE_STRING      uint8 = 254
	MYSQL_TYPE_GEOMETRY    uint8 = 255
)

const (
	NOT_NULL_FLAG       = 1
	PRI_KEY_FLAG        = 2
	UNIQUE_KEY_FLAG     = 4
	BLOB_FLAG           = 16
	UNSIGNED_FLAG       = 32
	ZEROFILL_FLAG       = 64
	BINARY_FLAG         = 128
	ENUM_FLAG           = 256
	AUTO_INCREMENT_FLAG = 512
	TIMESTAMP_FLAG      = 1024
	SET_FLAG            = 2048
	NUM_FLAG            = 32768
	PART_KEY_FLAG       = 16384
	GROUP_FLAG          = 32768
	UNIQUE_FLAG         = 65536
)

const (
	AUTH_NAME = "mysql_native_password"
)
