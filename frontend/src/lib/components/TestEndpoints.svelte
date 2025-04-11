<script lang="ts">
    import {ItemCollection} from "$lib/entities/ItemCollection";
    import { itemStore } from "$lib/entities/ItemStore.svelte";
    import { handleFormSubmition } from "$lib/fetchScripts";

    let derivedItemId: number = $derived(itemStore.selectedItem?.id ?? -1);
    
    let itemId: number | null | undefined = $state(null);
    let itemName: string | null | undefined = $state(null);
    let itemQuantity: number | null | undefined = $state(null);

    // Add specific result and error variables here 
    let itemsResult: ItemCollection | null = $state(null), itemsError: string | null = $state(null);
    let itemResult: ItemCollection | null = $state(null), itemError: string | null = $state(null);
    let item2Result: ItemCollection | null = $state(null), item2Error: string | null = $state(null);
    let item3Result: ItemCollection | null = $state(null), item3Error: string | null = $state(null);
    let item4Result: ItemCollection | null = $state(null), item4Error: string | null = $state(null);
</script>

<fieldset class="fieldset bg-base-200 border border-base-300 p-4 rounded-box">
    <legend class="fieldset-legend">Test Endpoints</legend>

    <form method="GET" onsubmit={(e) => handleFormSubmition(e, "/items",
        (val) => itemsResult = val, 
        (val) => itemsError = val)} 
        class="fieldset border border-neutral p-4 rounded-box">
        <h2 class="font-bold">Get items</h2>
        <div class="flex gap-2">
            <h3>Items:</h3>
            {#if itemsError != null}
                <p class="text-error">{itemsError}</p>
            {:else}
                <p class="text-success">{itemsResult}</p>
            {/if}
        </div>
        <button type="submit" class="btn btn-primary">Get</button>
    </form>

    <form method="GET" onsubmit={(e) => handleFormSubmition(e, '/items',
        (val) => itemResult = val, 
        (val) => itemError = val)}
        class="fieldset border border-neutral p-4 rounded-box">
        <h2 class="font-bold">Get item</h2>
        <div class="flex gap-2">
            <label for="itemID">Item id:</label>
            <input class="border px-1" bind:value={itemId} id="itemID" name="id" type="text" placeholder="Id" required/>
        </div>
        <div class="flex gap-2">
            <h3>Item:</h3>
            {#if itemError != null}
                <p class="text-error">{itemError}</p>
            {:else}
                <p class="text-success">{itemResult}</p>
            {/if}
        </div>
        <button type="submit" class="btn btn-primary">Get</button>
    </form>

    <form method="POST" onsubmit={(e) => handleFormSubmition(e, '/items',
        (val) => item2Result = val, 
        (val) => item2Error = val)}
        class="fieldset border border-neutral p-4 rounded-box">
        <h2 class="font-bold">Store item</h2>
        <div class="flex gap-2">
            <label for="itemName">Item name:</label>
            <input class="border px-1" bind:value={itemStore.inputItemName} id="itemName" name="name" type="text" placeholder="Name" required/>
        </div>
        <div class="flex gap-2">
            <label for="itemQuantity">Item quantity:</label>
            <input class="border px-1" bind:value={itemQuantity} id="itemQuantity" name="quantity" type="number" placeholder="Quantity" required/>
        </div>
        <div class="flex gap-2">
            <h3>Stored item:</h3>
            {#if item2Error != null}
                <p class="text-error">{item2Error}</p>
            {:else}
                <p class="text-success">{item2Result}</p>
            {/if}
        </div>
        <button type="submit" class="btn btn-primary">Get</button>
    </form>

    <form method="POST" onsubmit={(e) => handleFormSubmition(e, '/items',
        (val) => item3Result = val, 
        (val) => item3Error = val)}
        class="fieldset border border-neutral p-4 rounded-box">
        <input type="hidden" name="_method" value="PUT"/>
        <input type="hidden" name="id" bind:value={derivedItemId}/>
        <h2 class="font-bold">Store item</h2>
        <div class="flex gap-2">
            <label for="itemName">Item name:</label>
            <input class="border px-1" bind:value={itemStore.inputItemName} id="itemName" name="name" type="text" placeholder="Name" required/>
        </div>
        <div class="flex gap-2">
            <label for="itemQuantity">Item quantity:</label>
            <input class="border px-1" bind:value={itemQuantity} id="itemQuantity" name="quantity" type="number" placeholder="Quantity" required/>
        </div>
        <div class="flex gap-2">
            <h3>Stored item:</h3>
            {#if item3Error != null}
                <p class="text-error">{item3Error}</p>
            {:else}
                <p class="text-success">{item3Result}</p>
            {/if}
        </div>
        <button type="submit" class="btn btn-primary">Update</button>
    </form>

    <form method="POST" onsubmit={(e) => handleFormSubmition(e, '/items',
        (val) => item4Result = val, 
        (val) => item4Error = val)}
        class="fieldset border border-neutral p-4 rounded-box">
        <input type="hidden" name="_method" value="DELETE"/>
        <h2 class="font-bold">Delete item</h2>
        <div class="flex gap-2">
            <label for="itemID">Item id:</label>
            <input class="border px-1" bind:value={itemId} id="itemID" name="id" type="text" placeholder="Id" required/>
        </div>
        <div class="flex gap-2">
            <h3>Stored item:</h3>
            {#if item4Error != null}
                <p class="text-error">{item4Error}</p>
            {:else}
                <p class="text-success">{item4Result}</p>
            {/if}
        </div>
        <button type="submit" class="btn btn-primary">Delete</button>
    </form>
</fieldset>