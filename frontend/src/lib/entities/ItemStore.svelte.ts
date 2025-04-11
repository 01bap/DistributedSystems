import type { Item } from "./Item";
import type { ItemCollection } from "./ItemCollection";

class ItemStore {
    items: ItemCollection | null = $state(null);
    inputItemId: number = $state(-1);
    inputItemName: string = $state("");

    selectedItem: Item | undefined = $derived.by(() => {
        if(this.inputItemId > 0) {
            return this.items?.getItemById(this.inputItemId);
        }
        if(this.inputItemName.length > 0) {
            return this.items?.getItemByName(this.inputItemName);
        }
        return undefined;
    });

    // Singleton pattern
    static instance: ItemStore | null = null;
    private constructor() {}
    static getInstance(): ItemStore {
        if (!this.instance) {
            this.instance = new ItemStore();
        }
        return this.instance;
    }

    public setItems = (items: ItemCollection): void => {
        this.items = items.sortById();
        this.inputItemId = -1;          // Reset id and selected item
        // Remove item pointers when updated list doesnt contain the item anymore
        if(items?.getItemByName(this.inputItemName) == undefined) {
            this.inputItemName = "";
        }
    }
}
export const itemStore = ItemStore.getInstance();