package Maps

import "testing"

func TestSearch(t *testing.T) {
	testDict := Dictionary{"Test": "this is just a test"}

	t.Run("Dictionary Search happy path", func(t *testing.T) {
		got, err := testDict.Search("Test")
		if err != nil {
			t.Errorf("We found an error when we did not expect to")
		}
		want := "this is just a test"
		if got != want {
			t.Errorf("We wanted to find %s, but we found %s", want, got)
		}
	})

	t.Run("Dictionary sad not found path", func(t *testing.T) {
		_, err := testDict.Search("yam")
		if err == nil {
			t.Errorf("this should have returned an error")
		}
		if err.Error() != "this key value was not in the map" {
			t.Errorf("We got the wrong error message wanted this key value was not in the map got %s", err.Error())
		}
	})
}

func TestAdd(t *testing.T) {
	testDict := Dictionary{"Test": "this is just a test"}
	t.Run("Add happy path", func(t *testing.T) {
		key := "test2"
		definition := "This is the definition of key 2"
		err := testDict.Add(key, definition)
		if err != nil {
			t.Errorf("Unexpected error")
		}
		val, ok := testDict[key]
		if !ok {
			t.Errorf("The key never got added to the map")
		}
		if val != definition {
			t.Errorf("The key is present, but the value is incorrect")
		}
	})
	t.Run("Cannot use Add to over write existing value", func(t *testing.T) {
		key := "Test"
		value := "New Value"
		err := testDict.Add(key, value)
		if err == nil {
			t.Errorf("We expected an error adding a duplicate value")
		}
		if err.Error() != "this key already exists, please use the Update method to change the value" {
			t.Errorf("Wrong error message returned %s", err.Error())
		}
	})
}

func TestRemove(t *testing.T) {
	testDict := Dictionary{"test": "testing", "test2": "still Testing"}
	t.Run("Happy path removing key", func(t *testing.T) {
		keyToRemove := "test2"
		err := testDict.Remove(keyToRemove)
		if err != nil {
			t.Errorf("Expected no error but receieved one")
		}
		_, ok := testDict[keyToRemove]
		if ok {
			t.Errorf("The key was not removed as expected")
		}
	})
	t.Run("Sad path removing nonexistent key", func(t *testing.T) {
		keyToRemove := "test3"
		err := testDict.Remove(keyToRemove)
		if err == nil {
			t.Errorf("We expected an error and did not recieve one")
		}
	})
}

func TestUpdate(t *testing.T) {
	testDict := Dictionary{"test": "yam"}
	updatedValue := "testing"
	key := "test"
	t.Run("Happy path testing update", func(t *testing.T) {
		err := testDict.Update(key, updatedValue)
		if err != nil {
			t.Errorf("We found an error when we were not expecting to")
		}
		if testDict[key] != updatedValue {
			t.Errorf("The value was not updated")
		}
	})
	t.Run("Sad path Updating nonexistent key", func(t *testing.T) {
		keyToRemove := "test3"
		err := testDict.Update(keyToRemove, "fake")
		if err == nil {
			t.Errorf("We expected an error and did not recieve one")
		}
	})
}
