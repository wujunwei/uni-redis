package redis

const (
	Write         = "write"           // command may result in modifications
	Readonly      = "readonly"        // command will never modify keys
	DenyOOM       = "denyoom"         // reject command if currently OOM
	Admin         = "admin"           // server admin command
	Pubsub        = "pubsub"          // pubsub-related command
	Noscript      = "noscript"        // deny this command from scripts
	Random        = "random"          // command has random results, dangerous for scripts
	SortForScript = "sort_for_script" // if called from script, sort output
	Loading       = "loading"         // allow command while database is loading
	Stale         = "stale"           // allow command while replica has stale data
	skipMonitor   = "skip_monitor"    // do not show this command in MONITOR
	Asking        = "asking"          // cluster related - accept even if importing
	Fast          = "fast"            // command operates in constant or log(N) time. Used for latency monitoring.
	MovableKeys   = "movablekeys"     // keys have no pre-determined position. You must discover keys yourself.
)

type Command struct {
	Name        string
	FixArity    bool
	MovableKeys bool
}
