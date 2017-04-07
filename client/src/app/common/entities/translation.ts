export class Translation {
    lv: TranslationValue;
    ru: TranslationValue;
    en: TranslationValue;
}

class TranslationValue {
    "default"?: string;
    placeholder?: string;
    name?: string;
}