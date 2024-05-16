# Spacy Wrapper for Go

This project is a Go wrapper for the `spacy-cpp` library, allowing you to use spaCy's NLP capabilities in your Go applications.

## Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/yourusername/spacy-wrapper.git
   cd spacy-wrapper
   ```
   
2. Build the spacy-cpp library:

    ```sh
    cd spacy-cpp
    mkdir -p build && cd build
    cmake ..
    make
    sudo make install
   ```
3. Set the LD_LIBRARY_PATH environment variable:

    ```sh
      export LD_LIBRARY_PATH=/usr/local/lib:$LD_LIBRARY_PATH
    ```

4. Build the Go project:
    ```sh 
      go build -o spacy_app
      cd ../..
    ```

5.  Run the Go application:
    ```sh
    ./spacy_app
    ```

## Usage
```go
package main

import (
	"fmt"
	"log"
	"spacy/spacy"
)

func main() {
	nlp, err := spacy.LoadModel("en_core_web_sm")
	if err != nil {
		log.Fatalf("Error loading model: %v", err)
	}
	defer nlp.Free()

	text := "This is a test. This is another sentence."
	doc, err := nlp.Parse(text)
	if err != nil {
		log.Fatalf("Error parsing text: %v", err)
	}
	defer doc.Free()

	sentences, err := doc.GetSentences()
	if err != nil {
		log.Fatalf("Error getting sentences: %v", err)
	}

	for i, sentence := range sentences {
		fmt.Printf("Sentence %d: %s\n", i+1, sentence)
	}
}
```

## License
This project is licensed under the MIT License.


