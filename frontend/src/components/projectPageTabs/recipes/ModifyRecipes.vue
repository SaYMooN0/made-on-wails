<template>
    <div class="action-buttons-container">
        <RecipesTabActionButton mainLabel="Add vanilla recipe"
            hintLabel="Create recipes for crafting stations from vanilla minecraft" :actionOnClick="createActionOnClick" />
        <RecipesTabActionButton mainLabel="Delete recipe" hintLabel="Delete creation recipes for and from items and tags"
            :actionOnClick="deleteActionOnClick" />
        <RecipesTabActionButton mainLabel="Change recipe" hintLabel="Change recipes by replacing the items used in them"
            :actionOnClick="changeActionOnClick" />
        <RecipesTabActionButton mainLabel="Add modded recipe" hintLabel="Create recipes for mods that are supported by Made"
            :actionOnClick="moddedActionOnClick" />
        <RecipesTabActionButton mainLabel="Add recipe group" hintLabel="Create recipes for machines from mods at a time"
            :actionOnClick="groupActionOnClick" />
    </div>
</template>

<script>
import RecipesTabActionButton from './RecipesTabActionButton.vue';

export default {
    components: {
        RecipesTabActionButton
    },
    inject: ['createNewRecipeTab'],
    methods: {
        createActionOnClick() {
            this.createNewRecipeTab();
        },
        deleteActionOnClick() {
        },
        changeActionOnClick() {
        },
        moddedActionOnClick() {
        },
        groupActionOnClick() {
        }
    }
}

</script>
<style scoped>
.action-buttons-container {
    position: absolute;
    top: 52%;
    transform: translateY(-50%);
    display: flex;
    flex-wrap: wrap;
    justify-content: center;
    align-content: center;
    gap: 2.5vw;
    padding-left: calc(10vw - 25px - 2vh);
    padding-right: calc(10vw - 25px - 2vh);
}
</style>