import { Item } from "./Item";

export class ItemCollection {
    items!: Item[];

    constructor(obj: any){
        console.log(typeof(obj), obj)
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

    public getItemIdByName(name: string): number {
        const item = this.items.find(item => item.name === name);
        return item ? item.id : -1;
    }

    public getItemByName(name: string): Item | undefined {
        return this.items.find(item => item.name === name);
    }
    public getItemById(id: number): Item | undefined {
        return this.items.find(item => item.id === id);
    }

    public sortById(): ItemCollection {
        this.items.sort((a,b) => a.id - b.id);
        return this;
    }

    public toString = (): string => {
        return this.items.join(", ");
    }
}