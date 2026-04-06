package event

import (
	"reflect"
	"testing"
)

func TestSubscribedEventTypesForCatchAllReturnsNil(t *testing.T) {
	got := subscribedEventTypesFor("")
	if got != nil {
		t.Fatalf("subscribedEventTypesFor(\"\") = %v, want nil for catch-all", got)
	}
}

func TestSubscribedEventTypesForExplicitListSortsValues(t *testing.T) {
	got := subscribedEventTypesFor("calendar.event.updated_v1,im.message.receive_v1")
	want := []string{"calendar.event.updated_v1", "im.message.receive_v1"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("subscribedEventTypesFor(explicit) = %v, want %v", got, want)
	}
}
