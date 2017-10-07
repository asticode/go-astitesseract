#include <tesseract/baseapi.h>
#include <leptonica/allheaders.h>

extern "C" {
    class TesseractWrapper {
        private:
            tesseract::TessBaseAPI *api;
        public:
            TesseractWrapper()
            {
                api = new tesseract::TessBaseAPI();
            }
            int Init(char* datapath, char* languages)
            {
                return api->Init(datapath, languages);
            }
            void End()
            {
                api->End();
            }
            char* GetUTF8Text(char* path)
            {
                Pix *image = pixRead(path);
                api->SetImage(image);
                char *out = api->GetUTF8Text();
                pixDestroy(&image);
                return out;
            }
    };

    TesseractWrapper* New()
    {
        return new TesseractWrapper();
    }

    int Init(TesseractWrapper* w, char* datapath, char* languages)
    {
        if (strlen(datapath) == 0) {
            return w->Init(NULL, languages);
        } else {
            return w->Init(datapath, languages);
        }
    }

    void End(TesseractWrapper* w)
    {
        w->End();
        free(w);
    }

    char* GetUTF8Text(TesseractWrapper* w, char* path)
    {
        return w->GetUTF8Text(path);
    }
}

