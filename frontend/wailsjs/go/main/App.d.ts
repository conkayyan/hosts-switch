// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {hosts} from '../models';

export function AddHost(arg1:string,arg2:string,arg3:string):Promise<string>;

export function DeleteHost(arg1:string,arg2:number):Promise<string>;

export function DeleteHosts(arg1:Array<hosts.Host>):Promise<string>;

export function GetAllGroupNames():Promise<Array<string>>;

export function GetHostsText():Promise<string>;

export function GetInUseHostsText():Promise<string>;

export function GetMyHosts():Promise<hosts.MyHosts>;

export function SaveAddHostsText(arg1:string):Promise<string>;

export function SaveAllGroupHosts(arg1:string,arg2:string):Promise<string>;

export function SaveAllHosts(arg1:string):Promise<string>;

export function SaveAllInUseHosts(arg1:string):Promise<string>;

export function SetGroupName(arg1:string,arg2:string):Promise<string>;

export function SetGroupNameByHostnameId(arg1:number,arg2:string):Promise<string>;

export function SetGroupNameByHosts(arg1:Array<hosts.Host>,arg2:string):Promise<string>;

export function SwitchByGroupName(arg1:string,arg2:boolean):Promise<string>;

export function SwitchByHostnameId(arg1:string,arg2:number,arg3:boolean):Promise<string>;
