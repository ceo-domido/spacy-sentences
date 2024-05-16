#include "spacy.h"
#include <spacy/spacy>
#include <vector>
#include <string>
#include <cstring>  // for strncpy

extern "C" {

void* load_model(const char* model_name) {
    auto* spacy = new Spacy::Spacy();
    auto nlp = spacy->load(model_name);
    return reinterpret_cast<void*>(new Spacy::Nlp(nlp));
}

void free_model(void* nlp) {
    delete reinterpret_cast<Spacy::Nlp*>(nlp);
}

void* parse(void* nlp, const char* text) {
    auto doc = reinterpret_cast<Spacy::Nlp*>(nlp)->parse(text);
    return new Spacy::Doc(doc);
}

void free_doc(void* doc) {
    delete reinterpret_cast<Spacy::Doc*>(doc);
}

size_t get_sentence_count(void* doc) {
    auto sents = reinterpret_cast<Spacy::Doc*>(doc)->sents();
    return sents.size();
}

void get_sentence(void* doc, size_t index, char* buffer, size_t buffer_size) {
    auto sents = reinterpret_cast<Spacy::Doc*>(doc)->sents();
    if (index < sents.size()) {
        std::string sentence = sents[index].text();
        strncpy(buffer, sentence.c_str(), buffer_size - 1);
        buffer[buffer_size - 1] = '\0';
    }
}

} // extern "C"