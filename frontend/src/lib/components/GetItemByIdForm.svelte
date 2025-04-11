<script lang="ts">
    import { ItemCollection } from "$lib/entities/ItemCollection";
    import { handleFormSubmition } from "$lib/fetchScripts";

    let itemId: number | null | undefined = $state(null);

    let itemResult: ItemCollection | null = $state(null), itemError: string | null = $state(null);
</script>


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