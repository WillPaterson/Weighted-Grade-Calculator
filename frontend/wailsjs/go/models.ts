export namespace main {
	
	export class Grade {
	    grade: number;
	    totalPossible: number;
	    percentage: number;
	    weight: number;
	
	    static createFrom(source: any = {}) {
	        return new Grade(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.grade = source["grade"];
	        this.totalPossible = source["totalPossible"];
	        this.percentage = source["percentage"];
	        this.weight = source["weight"];
	    }
	}
	export class WeightedClass {
	    classCode: string;
	    gradeList: Grade[];
	
	    static createFrom(source: any = {}) {
	        return new WeightedClass(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.classCode = source["classCode"];
	        this.gradeList = this.convertValues(source["gradeList"], Grade);
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

