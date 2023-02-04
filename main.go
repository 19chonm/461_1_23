package main

func main() {
	// commands.Execute()

	// testing of sorting given a channel of ratings
	// urls output should be: F, C, B, E, A, D
	ratings := make(chan rating)
	go func() {
		ratings <- rating{NetScore: 5.0, Url: "url_for_module_A"}
		ratings <- rating{NetScore: 10.5, Url: "url_for_module_B"}
		ratings <- rating{NetScore: 42.5, Url: "url_for_module_C"}
		ratings <- rating{NetScore: -1.0, Url: "url_for_module_D"}
		ratings <- rating{NetScore: 6.5, Url: "url_for_module_E"}
		ratings <- rating{NetScore: 10000.0, Url: "url_for_module_F"}
		close(ratings)
	}()
	Sort_modules(ratings)
	NDJSON_test()
}
