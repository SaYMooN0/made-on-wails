export namespace made {
	
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

}

