<template>
    <div class="vanila-recipe-container">
        <select id="vanilla-type-selection" onchange="handleRecipeTypeChange(this)" class='default-select vanila-type-select'>
        <option value="crafting-table">Crafting Table</option>
        <option value="furnace">Furnace</option>
        <option value="stonecutter">Stonecutter</option>
        </select>
        <div class='recpe-content' > ${getCraftingTableRecipeForm()} </div>
        <div class='recpe-content' style="display:none"> ${getFurnaceRecipeForm()} </div>
        <div class='recpe-content' style="display:none"> ${getStonecutterRecipeForm()} </div>
        </div>
</template>
<script>

</script>