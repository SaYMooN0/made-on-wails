export namespace minecraftrelated {
	
	export class Block {
	    id: string;
	    inGameName: string;
	    maxStackSize: number;
	    burnTime?: number;
	    fireResistant: boolean;
	    hardness: number;
	    resistance: number;
	
	    static createFrom(source: any = {}) {
	        return new Block(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.inGameName = source["inGameName"];
	        this.maxStackSize = source["maxStackSize"];
	        this.burnTime = source["burnTime"];
	        this.fireResistant = source["fireResistant"];
	        this.hardness = source["hardness"];
	        this.resistance = source["resistance"];
	    }
	}
	export class Item {
	    Id: string;
	    InGameName: string;
	    MaxStackSize: number;
	    BurnTime?: number;
	    FireResistant: boolean;
	
	    static createFrom(source: any = {}) {
	        return new Item(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Id = source["Id"];
	        this.InGameName = source["InGameName"];
	        this.MaxStackSize = source["MaxStackSize"];
	        this.BurnTime = source["BurnTime"];
	        this.FireResistant = source["FireResistant"];
	    }
	}
	export class ProcessingType {
	    id: string;
	    inGameName: string;
	    isSupportedByMade: boolean;
	
	    static createFrom(source: any = {}) {
	        return new ProcessingType(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.inGameName = source["inGameName"];
	        this.isSupportedByMade = source["isSupportedByMade"];
	    }
	}
	export class Tag {
	    id: string;
	    inGameName: string;
	
	    static createFrom(source: any = {}) {
	        return new Tag(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.inGameName = source["inGameName"];
	    }
	}
	export class Mod {
	    id: string;
	    inGameName: string;
	    items: Item[];
	    blocks: Block[];
	    tags: Tag[];
	    supportedTypes: ProcessingType[];
	
	    static createFrom(source: any = {}) {
	        return new Mod(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.inGameName = source["inGameName"];
	        this.items = this.convertValues(source["items"], Item);
	        this.blocks = this.convertValues(source["blocks"], Block);
	        this.tags = this.convertValues(source["tags"], Tag);
	        this.supportedTypes = this.convertValues(source["supportedTypes"], ProcessingType);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	

}

export namespace projectrelated {
	
	
	export class ProjectSettings {
	    showWarningWhenDeletingAction: boolean;
	
	    static createFrom(source: any = {}) {
	        return new ProjectSettings(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.showWarningWhenDeletingAction = source["showWarningWhenDeletingAction"];
	    }
	}
	export class MadeProject {
	    Name: string;
	    FullPath: string;
	    PathToFolder: string;
	    Version: string;
	    Loader: number;
	    // Go type: time
	    CreationDate: any;
	    // Go type: time
	    LastUpdated: any;
	    Settings: ProjectSettings;
	    History: HistoryItem[];
	    Mods: minecraftrelated.Mod[];
	
	    static createFrom(source: any = {}) {
	        return new MadeProject(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Name = source["Name"];
	        this.FullPath = source["FullPath"];
	        this.PathToFolder = source["PathToFolder"];
	        this.Version = source["Version"];
	        this.Loader = source["Loader"];
	        this.CreationDate = this.convertValues(source["CreationDate"], null);
	        this.LastUpdated = this.convertValues(source["LastUpdated"], null);
	        this.Settings = this.convertValues(source["Settings"], ProjectSettings);
	        this.History = this.convertValues(source["History"], HistoryItem);
	        this.Mods = this.convertValues(source["Mods"], minecraftrelated.Mod);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	

}

export namespace themerelated {
	
	export class Theme {
	    Name: string;
	    // Go type: Color
	    MainBackColor: any;
	    // Go type: Color
	    SecondBackColor: any;
	    // Go type: Color
	    ThirdBackColor: any;
	    // Go type: Color
	    MainFrontColor: any;
	    // Go type: Color
	    SecondFrontColor: any;
	    // Go type: Color
	    ThirdFrontColor: any;
	    // Go type: Color
	    MainBrightColor: any;
	    // Go type: Color
	    SecondBrightColor: any;
	    // Go type: Color
	    ThirdBrightColor: any;
	    // Go type: Color
	    WarningMainColor: any;
	    // Go type: Color
	    WarningBrightColor: any;
	
	    static createFrom(source: any = {}) {
	        return new Theme(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Name = source["Name"];
	        this.MainBackColor = this.convertValues(source["MainBackColor"], null);
	        this.SecondBackColor = this.convertValues(source["SecondBackColor"], null);
	        this.ThirdBackColor = this.convertValues(source["ThirdBackColor"], null);
	        this.MainFrontColor = this.convertValues(source["MainFrontColor"], null);
	        this.SecondFrontColor = this.convertValues(source["SecondFrontColor"], null);
	        this.ThirdFrontColor = this.convertValues(source["ThirdFrontColor"], null);
	        this.MainBrightColor = this.convertValues(source["MainBrightColor"], null);
	        this.SecondBrightColor = this.convertValues(source["SecondBrightColor"], null);
	        this.ThirdBrightColor = this.convertValues(source["ThirdBrightColor"], null);
	        this.WarningMainColor = this.convertValues(source["WarningMainColor"], null);
	        this.WarningBrightColor = this.convertValues(source["WarningBrightColor"], null);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class ThemeHex {
	    Name: string;
	    MainBackColor: string;
	    SecondBackColor: string;
	    ThirdBackColor: string;
	    MainFrontColor: string;
	    SecondFrontColor: string;
	    ThirdFrontColor: string;
	    MainBrightColor: string;
	    SecondBrightColor: string;
	    ThirdBrightColor: string;
	    WarningMainColor: string;
	    WarningBrightColor: string;
	
	    static createFrom(source: any = {}) {
	        return new ThemeHex(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.Name = source["Name"];
	        this.MainBackColor = source["MainBackColor"];
	        this.SecondBackColor = source["SecondBackColor"];
	        this.ThirdBackColor = source["ThirdBackColor"];
	        this.MainFrontColor = source["MainFrontColor"];
	        this.SecondFrontColor = source["SecondFrontColor"];
	        this.ThirdFrontColor = source["ThirdFrontColor"];
	        this.MainBrightColor = source["MainBrightColor"];
	        this.SecondBrightColor = source["SecondBrightColor"];
	        this.ThirdBrightColor = source["ThirdBrightColor"];
	        this.WarningMainColor = source["WarningMainColor"];
	        this.WarningBrightColor = source["WarningBrightColor"];
	    }
	}

}

