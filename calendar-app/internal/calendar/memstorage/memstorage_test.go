package memstorage

import (
	"testing"
	"time"

	"github.com/Cornpop456/otus-go/calendar-app/internal/calendar/models"
	"github.com/stretchr/testify/assert"
)

var (
	storage *EventsLocalStorage
	t1      time.Time
	t2      time.Time
	t3      time.Time
)

func setUp() {
	storage = NewEventsLocalStorage()

	t1, _ = time.Parse(time.RFC822, "27 Jan 20 19:00 MSK")
	t2, _ = time.Parse(time.RFC822, "28 Jan 20 20:00 MSK")
	t3, _ = time.Parse(time.RFC822, "10 Jan 21 11:00 MSK")
}

func TestEventsLocalStorage_AddItem(t *testing.T) {
	setUp()

	type args struct {
		item interface{}
	}
	tests := []struct {
		s            *EventsLocalStorage
		args         args
		wantErr      bool
		eventsLength int
	}{
		{storage, args{1}, true, 0},
		{storage, args{&models.Event{RawDate: &t1, ID: "0"}}, false, 1},
		{storage, args{&models.Event{RawDate: &t2, ID: "1"}}, false, 2},
		{storage, args{&models.Event{RawDate: &t2, ID: "2"}}, true, 2},
		{storage, args{&models.Event{RawDate: &t3, ID: "1"}}, true, 2},
	}

	for _, tc := range tests {
		err := storage.AddItem(tc.args.item)
		if tc.wantErr {
			assert.NotNil(t, err)
		} else {
			assert.Nil(t, err)
		}
		assert.Equal(t, tc.eventsLength, len(storage.events), "should be equal")
	}
}

func TestEventsLocalStorage_DeleteItem(t *testing.T) {
	setUp()

	storage.AddItem(&models.Event{ID: "0", RawDate: &t1})
	storage.AddItem(&models.Event{ID: "1", RawDate: &t2})
	storage.AddItem(&models.Event{ID: "2", RawDate: &t3})

	type args struct {
		id string
	}
	tests := []struct {
		s            *EventsLocalStorage
		args         args
		wantErr      bool
		eventsLength int
	}{
		{storage, args{"0"}, false, 2},
		{storage, args{"0"}, true, 2},
		{storage, args{"1"}, false, 1},
		{storage, args{"1"}, true, 1},
	}
	for _, tc := range tests {
		err := storage.DeleteItem(tc.args.id)
		if tc.wantErr {
			assert.NotNil(t, err)
		} else {
			assert.Nil(t, err)
		}
		assert.Equal(t, tc.eventsLength, len(storage.events), "should be equal")
	}
}

func TestEventsLocalStorage_ChangeItem(t *testing.T) {
	setUp()

	storage.AddItem(&models.Event{ID: "0", RawDate: &t1})
	storage.AddItem(&models.Event{ID: "1", RawDate: &t3})

	type args struct {
		id   string
		args map[string]string
	}
	tests := []struct {
		s              *EventsLocalStorage
		args           args
		wantErr        bool
		wantName       string
		wantDesc       string
		wantDateString string
	}{
		{storage, args{"0", map[string]string{"Name": "name1"}}, false, "name1", "", t1.String()},
		{storage, args{"0", map[string]string{"Description": "desc1"}}, false, "name1", "desc1", t1.String()},
		{storage, args{"0", map[string]string{"Date": t2.Format(time.RFC822)}}, false, "name1", "desc1", t2.String()},
		{storage, args{"0", map[string]string{"Date": t3.Format(time.RFC822)}}, true, "name1", "desc1", t2.String()},
	}
	for _, tc := range tests {
		err := storage.ChangeItem(tc.args.id, tc.args.args)
		item, _ := storage.GetItem(tc.args.id)
		event := item.(*models.Event)

		if tc.wantErr {
			assert.NotNil(t, err)
		} else {
			assert.Nil(t, err)
		}

		assert.Equal(t, tc.wantName, event.Name, "Names should be equal")
		assert.Equal(t, tc.wantDesc, event.Description, "Descriptions should be equal")
		assert.Equal(t, tc.wantDateString, event.RawDate.String(), "RawDates should be equal")
	}
}
