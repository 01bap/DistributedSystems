<script>
    import CreateItemForm from "./CreateItemForm.svelte";
    import GetAllItemsForm from "./GetAllItemsForm.svelte";
    import GetItemByIdForm from "./GetItemByIdForm.svelte";
    import UpdateItemForm from "./UpdateItemForm.svelte";

</script>


<fieldset class="fieldset bg-base-200 border border-base-300 p-4 rounded-box">
    <legend class="fieldset-legend">Unified Endpoints</legend>

    <GetAllItemsForm/>
    <CreateItemForm/>
    
    <GetItemByIdForm/>
    <UpdateItemForm/>

</fieldset>