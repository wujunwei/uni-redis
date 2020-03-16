package protocol


//const(
//	write - command may result in modifications
//readonly - command will never modify keys
//denyoom - reject command if currently OOM
//admin - server admin command
//pubsub - pubsub-related command
//noscript - deny this command from scripts
//random - command has random results, dangerous for scripts
//sort_for_script - if called from script, sort output
//loading - allow command while database is loading
//stale - allow command while replica has stale data
//skip_monitor - do not show this command in MONITOR
//asking - cluster related - accept even if importing
//fast - command operates in constant or log(N) time. Used for latency monitoring.
//movablekeys - keys have no pre-determined position. You must discover keys yourself.
//)
type Command struct {
	Name string
	FixArity bool
	MovableKeys bool
}