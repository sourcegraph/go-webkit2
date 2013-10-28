#include <gio/gio.h>

/* Wrapper that runs the Go closure for a given context */
extern void _go_gasyncreadycallback_call(gpointer user_data, void *cresult);

void _gasyncreadycallback_call(GObject *source_object, GAsyncResult *res, gpointer user_data);
