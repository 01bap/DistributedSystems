<script lang="ts">
    import {FetchResponse} from "$lib/entities/fetchResponse";
    
    // let endpoint = "http://localhost:8080";
    let endpoint = "https://vigilant-doodle-xqvwj7j6wxjhvjw4-8080.app.github.dev";

    let itemId: number | null | undefined;
    let itemName: string | null | undefined;
    let itemQuantity: number | null | undefined;

    // Add specific result and error variables here 
    let itemsResult: FetchResponse | null, itemsError: string | null;
    let itemResult: FetchResponse | null, itemError: string | null;
    let item2Result: FetchResponse | null, item2Error: string | null;

    function formDataToTypedObject(formData: FormData): Record<string, string | number> {
        const result: Record<string, string | number> = {};

        for (const [key, value] of formData.entries()) {
            const num = Number(value);
            result[key] = (!isNaN(num)) ? num : value as string;
        }
        return result;
    }

    async function handleFormSubmition(
        event: SubmitEvent, 
        path: string,
        setResult: (val: FetchResponse | null) => void,
        setError: (val: string | null) => void
    ) {
        event.preventDefault();
        const form = event.target as HTMLFormElement;
        const method = form.method;
        const formData = new FormData(form);
        const data = formDataToTypedObject(formData);
        const regex = new RegExp("get", "i");

        let uri: string;
        if({itemId}.itemId == undefined || {itemId}.itemId == null) {
            uri = endpoint + path;
        } else {
            uri = endpoint + path + `/${{itemId}.itemId}`;
        }

        try {
            let res: any;
            if(regex.test(method)) {
                res = await fetch(uri, {
                    method: "GET",
                    headers: { 'Content-Type': 'application/json' },
                });
            } else {
                res = await fetch(uri, {
                    method: method,
                    headers: { 'Content-Type': 'application/json' },
                    body: JSON.stringify(data)
                });
            }

            if (!res.ok) throw new Error('Request failed');
            const json = await res.json();
            const response = new FetchResponse(json);

            setResult(response);
            setError(null);
        } catch (err: any) {
            setError(err.message);
            setResult(null);
        }
        itemId = null;
        itemName = null;
        itemQuantity = null;
        setTimeout(() => {
            setError(null);
        }, 3000);
    }
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
            <input class="border px-1" bind:value={itemName} id="itemName" name="name" type="text" placeholder="Name" required/>
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
</fieldset>