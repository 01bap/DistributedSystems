<script lang="ts">
    import type { Item } from "$lib/entities/Item";
    import { itemStore } from "$lib/entities/ItemStore.svelte";

    let cachedItems = $derived(itemStore.items?.items ?? new Array<Item>());
</script>

<!-- Maybe implement banner to show fetch success -->
<!-- Maybe also an interval check to fetch all database items -->

<fieldset class="fieldset bg-base-200 border border-base-300 p-4 rounded-box">
    <legend>List of all cached items</legend>
    <ul class="list bg-base-100 rounded-box shadow-md">
        {#each cachedItems as item}
            <li class="list-row m-1 border">
                <div class="text-4xl opacity-30 tabular-nums">{item.id}</div>
                
                <div class="list-col-grow">
                    <div class="text-lg font-bold">{item.name}</div>
                    <!-- When clicked on then change to input to create a put request which overwrites the quantity -->
                    <div class="text-xs uppercase font-semibold opacity-60">Count: {item.quantity}</div>
                </div>

                <!-- Should open the post dialog to update/create items -->
                <button class="btn btn-ghost">Update</button>
                <button class="btn btn-ghost">Delete</button>
            </li>
        {/each}
    </ul>
</fieldset>