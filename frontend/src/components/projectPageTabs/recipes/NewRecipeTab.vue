<template>
     <div>
       <select id="vanilla-type-selection" onchange="handleRecipeTypeChange(this)" class='default-select vanila-type-select'>
        <option value="crafting-table">Crafting Table</option>
        <option value="furnace">Furnace</option>
        <option value="stonecutter">Stonecutter</option>
        </select>
        <div id="crafting-table-recipe-content" class='recpe-content' > ${getCraftingTableRecipeForm()} </div>
        <div id="furnace-recipe-content" class='recpe-content' style="display:none"> ${getFurnaceRecipeForm()} </div>
        <div id="stonecutter-recipe-content" class='recpe-content' style="display:none"> ${getStonecutterRecipeForm()} </div>
        </div>
</template>