export namespace hosts {
	
	export class Host {
	    id: number;
	    show: boolean;
	    ip: string;
	    hostName: string;
	    groupName: string;
	
	    static createFrom(source: any = {}) {
	        return new Host(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.show = source["show"];
	        this.ip = source["ip"];
	        this.hostName = source["hostName"];
	        this.groupName = source["groupName"];
	    }
	}
	export class Group {
	    hostsText: string;
	    groupName: string;
	    showNum: number;
	    hideNum: number;
	    show: boolean;
	    list: {[key: number]: Host};
	
	    static createFrom(source: any = {}) {
	        return new Group(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.hostsText = source["hostsText"];
	        this.groupName = source["groupName"];
	        this.showNum = source["showNum"];
	        this.hideNum = source["hideNum"];
	        this.show = source["show"];
	        this.list = source["list"];
	    }
	}
	
	export class MyHosts {
	    path: string;
	    hostsText: string;
	    inUseHostsText: string;
	    noInUseHostsText: string;
	    totalNum: number;
	    list: {[key: number]: Host};
	    listByGroup: {[key: string]: Group};
	
	    static createFrom(source: any = {}) {
	        return new MyHosts(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.path = source["path"];
	        this.hostsText = source["hostsText"];
	        this.inUseHostsText = source["inUseHostsText"];
	        this.noInUseHostsText = source["noInUseHostsText"];
	        this.totalNum = source["totalNum"];
	        this.list = source["list"];
	        this.listByGroup = this.convertValues(source["listByGroup"], Group, true);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
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

