// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {themerelated} from '../models';

export function CurrentTheme():Promise<themerelated.Theme>;

export function CurrentThemeColors():Promise<themerelated.ThemeHex>;

export function GetAllThemesValues():Promise<Array<themerelated.ThemeHex>>;

export function SaveToFile():Promise<void>;

export function SetCurrentTheme(arg1:string):Promise<string>;

export function ThemeFromHexTheme(arg1:themerelated.ThemeHex):Promise<themerelated.Theme>;

export function UpdateTheme(arg1:string,arg2:themerelated.Theme):Promise<string>;
