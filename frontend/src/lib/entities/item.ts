export class Item {
    id!: number;
    name!: string;
    quantity!: number;

    constructor(params: any) {
        try {
            if(params.id == undefined || params.name == undefined || params.quantity == undefined)
                throw "'params' does not match {id, name, quantity}!";
            this.id = params.id;
            this.name = params.name;
            this.quantity = params.quantity;
        } catch(ex: any) {
            this.id = -1;
            this.name = "No item";
            this.quantity = 0;
            console.log("Failed to create item!", ex);
        }
    }

    public toString = (): string => {
        return `${this.name} (ID: ${this.id}):  ${this.quantity}`;
    }
}