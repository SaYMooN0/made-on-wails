<template>
    <div class="tab-content">
        <select v-model="selectedType" class='type-select'>
            <option value="crafting-table">Crafting Table</option>
            <option value="furnace">Furnace</option>
            <option value="stonecutter">Stonecutter</option>
        </select>
        <div class='recpe-content'>
            <CrafringTable v-if="selectedType === 'crafting-table'" />
            <Furnace v-if="selectedType === 'furnace'" />
            <Stonecutter v-if="selectedType === 'stonecutter'" />
        </div>
    </div>
</template>

<script>
import CrafringTable from './components/CrafringTable.vue';
import Furnace from './components/Furnace.vue';
import Stonecutter from './components/Stonecutter.vue';

export default {
    components: {
        CrafringTable,
        Furnace,
        Stonecutter
    },
    data() {
        return {
            selectedType: 'crafting-table'
        }
    }
}
</script>
<style scoped>
.tab-content{
    height: 100%;
}
.recpe-content{
    display: flex;
    padding-top: 2%;
    height: calc(90vh - 4vw - 22px);
}
.type-select {
    height: calc(1vh + 25px + 0.6vw);
    width: calc(80px + 7vw + 2vh);
    font-size: calc(0.78vh + 0.58vw + 7px);
    margin-left: calc(10px + 1vh);
    margin-top: calc(10px + 1vh);
    color: var(--front);
    background-color: var(--back-2);
    font-family: 'Figtree';
    outline: none;
    border: 1px solid transparent;
    border-radius: calc(2px + 0.12vw + 0.12vh);
    padding-left: calc(2px + 0.08vw + 0.12vh);
}
</style>