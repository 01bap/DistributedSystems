<script lang="ts">
    // let endpoint = "http://localhost:8080";
    let endpoint = "https://vigilant-doodle-xqvwj7j6wxjhvjw4-8080.app.github.dev";

    let itemId: number | null | undefined;

    // Add specific result and error variable here 
    let itemResult = '', itemError = '';
    let itemsResult = '', itemsError = '';

    async function handleFormSubmition(event: SubmitEvent, path: string) {
        event.preventDefault();
        const method = event.target?.method;
        const formData = new FormData(event.target);
        const data = Object.fromEntries(formData);

        let setResult: Function, setError: Function;
        let uri: string;
        if({itemId}.itemId == undefined || {itemId}.itemId == null) {
            uri = endpoint + path;
        } else {
            uri = endpoint + path + `/${{itemId}.itemId}`;
        }
        const regex = new RegExp("get", "i")
        // Add new api pathes here
        switch(path) {
            case "/item":
                setResult = (val: string) => itemResult = val;
                setError = (val: string) => itemError = val;
                break;
            case "/items":
                setResult = (val: string) => itemsResult = val;
                setError = (val: string) => itemsError = val;
                break;
            default:
                throw new ReferenceError(path + " not defined!");
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
            const response = ((json != null) ? JSON.stringify(json, null, 2) : "Nothing recieved!");

            setResult(response);
            setError(null);
        } catch (err: any) {
            setError(err.message);
            setResult(null);
        }
        itemId = null;
        setTimeout(() => {
            setError(null);
            setResult(null);
        }, 3000);
    }
</script>

<fieldset class="fieldset bg-base-200 border border-base-300 p-4 rounded-box">
    <legend class="fieldset-legend">Test Endpoints</legend>

    <form method="GET" onsubmit="{(e) => handleFormSubmition(e, "/items")}" class="fieldset border border-neutral p-4 rounded-box">
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

    <form method="GET" onsubmit={(e) => handleFormSubmition(e, '/item')} class="fieldset border border-neutral p-4 rounded-box">
        <h2 class="font-bold">Get item</h2>
        <div class="flex gap-2">
            <label for="itemID">Item id:</label>
            <input bind:value={itemId} id="itemID" name="id" type="text" placeholder="Id" required/>
        </div>
        <div class="flex gap-2">
            <h3>Count:</h3>
            {#if itemError != null}
                <p class="text-error">{itemError}</p>
            {:else}
                <p class="text-success">{itemResult}</p>
            {/if}
        </div>
        <button type="submit" class="btn btn-primary">Get</button>
    </form>
</fieldset>