package helperchecker

import (
	"log/slog"
	helperexception "panel-ektensi/helper/exception"
)

func AssertFoundData[T interface{ PK() int64 }](log *slog.Logger, caller, where, notFoundMsg string, e T) *helperexception.Exception {
	if e.PK() == 0 {
		log.Info(caller, where, "Error", notFoundMsg)
		return helperexception.NotFound(notFoundMsg)
	}
	return nil
}

func AssertTenantMatch[T interface{ TenantCode() string }](
	log *slog.Logger,
	caller, where, notFoundMsg string,
	expectedTenant string,
	e T,
) *helperexception.Exception {
	if e.TenantCode() != expectedTenant {
		log.Info(caller, where, "Error", notFoundMsg)
		return helperexception.NotFound(notFoundMsg)
	}
	return nil
}
