#include <stdlib.h>
#include <os/log.h>

const os_log_t os_log_default = OS_LOG_DEFAULT;

void ul_log(unsigned char level, os_log_t log, const char* const s) {
	switch (level) {
	case 0: // PanicLevel
		os_log_fault(log, "%{public}s", s);
		break;
	case 1: // FatalLevel
	case 2: // ErrorLevel
		os_log_error(log, "%{public}s", s);
		break;
	case 3: // WarnLevel
	case 4: // InfoLevel
		os_log_info(log, "%{public}s", s);
		break;
	case 5: // DebugLevel
	case 7: // TraceLevel
		os_log_debug(log, "%{public}s", s);
		break;
	default:
		os_log(log, "%{public}s", s);
	}

}

void release(os_log_t t) {
	os_release(t);
}
