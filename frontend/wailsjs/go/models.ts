export namespace hosts {

	export class Host {
		id: number;
		show: boolean;
		ip: string;
		hostname: string;
		groupName: string;

		static createFrom(source: any = {}) {
			return new Host(source);
		}

		constructor(source: any = {}) {
			if ('string' === typeof source) source = JSON.parse(source);
			this.id = source["id"];
			this.show = source["show"];
			this.ip = source["ip"];
			this.hostname = source["hostname"];
			this.groupName = source["groupName"];
		}
	}

}

