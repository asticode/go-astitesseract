#ifdef __cplusplus
extern "C" {
#endif

    typedef void* TesseractWrapper;
    TesseractWrapper* New();
    int Init(TesseractWrapper* w, char* datapath, char* languages);
    void End(TesseractWrapper* w);
    char* GetUTF8Text(TesseractWrapper* w, char* path);

#ifdef __cplusplus
}
#endif