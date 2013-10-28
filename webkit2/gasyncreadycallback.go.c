#include <stdio.h>
#include <gio/gio.h>
#include "gasyncreadycallback.go.h"

void _gasyncreadycallback_call(GObject *source_object, GAsyncResult *res, gpointer user_data) {
  _go_gasyncreadycallback_call(user_data, res);
}
