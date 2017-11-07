package redisBus

func CreateFirstCounter() error {
	err := client.Set("counter", 0, 0).Err()
	if err != nil {
		panic(err)
	}
	return err
}
