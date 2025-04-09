import { Item } from "./item";

export class FetchResponse {
    items!: Item[];

    constructor(obj: any){
        this.items = new Array<Item>();
        try{
            if(obj.length > 0) {
                obj.forEach((item: any) => {
                    this.items.push(new Item(item));
                });
            } else {
                this.items.push(new Item(obj))
            }
        } catch(ex: any) {
            console.log("Failed to create response!", ex);
        }
    }

    public toString = (): string => {
        return this.items.join(", ");
    }
}