
all: run

run:
	@echo "Running Docker containers..."
	docker-compose up -d
	go run app/cmd/main.go

stop:
	@echo "Stopping Docker containers..."
	docker-compose down

clean:
	rm -rf images 

test:
	go test -v ./app/tests

example_async:
	go run app/client/client_runner.go --async https://www.youtube.com/watch\?v\=9MK0CvfZWkg https://www.youtube.com/watch\?v\=3v05NlypecM https://www.youtube.com/watch?v=9EbS6w3RSG0

example_single:
	go run app/client/client_runner.go https://www.youtube.com/watch?v=4jds773UlWE
