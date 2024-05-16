#ifndef SPACY_H
#define SPACY_H

#ifdef __cplusplus
extern "C" {
#endif

#include <stddef.h>

void* load_model(const char* model_name);
void free_model(void* nlp);

void* parse(void* nlp, const char* text);
void free_doc(void* doc);

size_t get_sentence_count(void* doc);
void get_sentence(void* doc, size_t index, char* buffer, size_t buffer_size);

#ifdef __cplusplus
}
#endif

#endif // SPACY_H
