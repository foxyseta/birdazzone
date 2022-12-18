export interface WordCloudOptions {
  text: string;
  width: number;
  height: number;
  fontFamily: string;
  //loadGoogleFonts: string
  fontScale: number;
  padding: number;
  maxNumWords: number;
  case: string;
  rotation: number;
  colors: string[];
  useWordList: boolean;
  scale: string;
}
