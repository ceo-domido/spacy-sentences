package spacy

/*
#cgo CXXFLAGS: -std=c++11
#cgo LDFLAGS: -L/home/sokolov/GolandProjects/spacy/spacy-cpp/build -lspacy
#include "spacy.h"
#include <stdlib.h>

void* load_model(const char* model_name);
void free_model(void* nlp);

void* parse(void* nlp, const char* text);
void free_doc(void* doc);

size_t get_sentence_count(void* doc);
void get_sentence(void* doc, size_t index, char* buffer, size_t buffer_size);
*/
import "C"
import (
	"errors"
	"unsafe"
)

type NLP struct {
	ptr unsafe.Pointer
}

type Doc struct {
	ptr unsafe.Pointer
}

func LoadModel(modelName string) (*NLP, error) {
	cModelName := C.CString(modelName)
	defer C.free(unsafe.Pointer(cModelName))

	ptr := C.load_model(cModelName)
	if ptr == nil {
		return nil, errors.New("failed to load model")
	}
	return &NLP{ptr: ptr}, nil
}

func (nlp *NLP) Free() {
	C.free_model(nlp.ptr)
}

func (nlp *NLP) Parse(text string) (*Doc, error) {
	cText := C.CString(text)
	defer C.free(unsafe.Pointer(cText))

	ptr := C.parse(nlp.ptr, cText)
	if ptr == nil {
		return nil, errors.New("failed to parse text")
	}
	return &Doc{ptr: ptr}, nil
}

func (doc *Doc) Free() {
	C.free_doc(doc.ptr)
}

func (doc *Doc) GetSentences() ([]string, error) {
	count := C.get_sentence_count(doc.ptr)
	sentences := make([]string, count)

	for i := 0; i < int(count); i++ {
		buffer := make([]byte, 1024)
		C.get_sentence(doc.ptr, C.size_t(i), (*C.char)(unsafe.Pointer(&buffer[0])), C.size_t(len(buffer)))
		sentences[i] = C.GoString((*C.char)(unsafe.Pointer(&buffer[0])))
	}

	return sentences, nil
}
