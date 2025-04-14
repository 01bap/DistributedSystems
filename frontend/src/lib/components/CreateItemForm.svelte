<script lang="ts">
    import { ItemCollection } from "$lib/entities/ItemCollection";
    import { handleFormSubmition } from "$lib/fetchScripts";
    import {itemStore} from "$lib/entities/ItemStore.svelte"

    let itemQuantity: number | null | undefined = $state(null);

    let itemResult: ItemCollection | null = $state(null), itemError: string | null = $state(null);
</script>

<form method="POST" onsubmit={(e) => handleFormSubmition(e, '/items',
    (val) => itemResult = val, 
    (val) => itemError = val)}
    class="fieldset border border-neutral p-4 rounded-box">
    <h2 class="font-bold">Create item</h2>
    <div class="flex gap-2">
        <label class="floating-label">
            <input bind:value={itemStore.inputItemName} id="itemName" name="name" type="text" placeholder="Item name" required class="input input-md" />
            <span>Item name</span>
        </label>
    </div>
    <div class="flex gap-2">
        <label class="floating-label">
            <input bind:value={itemQuantity} id="itemQuantity" name="quantity" type="number" placeholder="Item quantity" required class="input input-md" />
            <span>Item quantity</span>
        </label>
    </div>
    
    <div class="flex gap-2">
        <h3>Stored item:</h3>
        {#if itemError != null}
            <p class="text-error">{itemError}</p>
        {:else}
            <p class="text-success">{itemResult}</p>
        {/if}
    </div>
    <button type="submit" class="btn btn-primary">Create or update</button>
</form>