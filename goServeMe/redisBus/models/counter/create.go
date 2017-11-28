package counter

import bus ".."

//CreateFirstCounter ...
// Creates the first Counter
func CreateFirstCounter() error {
	err := bus.Client.Set("counter", 0, 0).Err()
	if err != nil {
		panic(err)
	}
	return err
}
