import { ItemCollection } from "./entities/ItemCollection";
import { itemStore } from "./entities/ItemStore.svelte";

// let endpoint = "http://localhost:8080";
let endpoint = "https://vigilant-doodle-xqvwj7j6wxjhvjw4-8080.app.github.dev";

/**
 * Doesnt escape types in forms
 * @param formData Data that will be send in a form
 * @returns formatted data with numbers not being stringified
 */
export function formDataToTypedObject(formData: FormData): Record<string, string | number> {
    const result: Record<string, string | number> = {};

    for (const [key, value] of formData.entries()) {
        const num = Number(value);
        result[key] = (!isNaN(num)) ? num : value as string;
    }
    return result;
}

/**
 * Handles the request to the backend
 * @Important Item ID has to be named as 'id' to check the value inside this function
 * @Important When fetching a message other than 'POST' or 'GET' an input with name '_method' needs to be defined with the corresponding value
 * @param event Holds the form to be submitted
 * @param path Path at the endpoint
 * @param setResult Returns an object as the response
 * @param setError Sets the error when fetch failed
 */
export async function handleFormSubmition(
    event: SubmitEvent, 
    path: string,
    setResult: (val: ItemCollection | null) => void,
    setError: (val: string | null) => void
) {
    event.preventDefault();
    const NAME_FOR_METHODCALLBACK = "_method";
    const regex = new RegExp("(get|delete)", "i");
    const form = event.target as HTMLFormElement;
    const formData = new FormData(form);
    let method: string = form.method;
    if(formData.has(NAME_FOR_METHODCALLBACK)) {
        method = formData.get(NAME_FOR_METHODCALLBACK) as string;
        formData.delete(NAME_FOR_METHODCALLBACK);
    }
    const data = formDataToTypedObject(formData);

    try {
        let uri: string;
        let res: any;
        let response: ItemCollection | null = new ItemCollection(null);

        if(isNaN(data.id as number) || data.id == undefined || data.id == null) {
            uri = endpoint + path;
        } else {
            if(data.id as number < 1) throw new Error(`No item for id ${data.id} found!`);
            uri = endpoint + path + `/${data.id}`;
        }

        if(regex.test(method)) {
            res = await fetch(uri, {
                method: method,
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
        try {
            const json = await res.json();
            response = new ItemCollection(json);

            // Custom hook for storage of items
            if(res.ok && regex.test(method) && (isNaN(data.id as number) || data.id == undefined || data.id == null)) {
                itemStore.setItems(response);
            }
        } catch(ex: any) {
            // Exceptions
            // Ok response without body
            if(res.status == 204) {
                response = new ItemCollection("Successfull request");
            }
        }

        setResult(response);
        setError(null);
    } catch (err: any) {
        setError(err.message);
        setResult(null);
    }
    // Reseting inputs
    form.reset();
    setTimeout(() => {
        setError(null);
        setResult(null);
    }, 5000);
}