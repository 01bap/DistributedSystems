<script lang="ts">
    import { ItemCollection } from "$lib/entities/ItemCollection";
    import { handleFormSubmition } from "$lib/fetchScripts";

    let itemsResult: ItemCollection | null = $state(null), itemsError: string | null = $state(null);
</script>

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