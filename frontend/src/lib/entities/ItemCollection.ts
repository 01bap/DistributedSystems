import { Item } from "./Item";

export class ItemCollection {
    items!: Item[];
    errorMessage: string;

    constructor(obj: any){
        this.items = new Array<Item>();
        this.errorMessage = "";
        try{
            if(obj.length > 0) {
                obj.forEach((item: any) => {
                    this.items.push(new Item(item));
                });
            } else {
                this.items.push(new Item(obj))
            }
        } catch(ex: any) {
            if(obj == null) {
                this.errorMessage = "No item recieved!";
            } else {
                this.errorMessage = obj.toString();
            }
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
        if(this.items.length > 0) {
            return this.items.join(", ");
        }
        return this.errorMessage;
    }
}