<template>
    <div class="color-inputs-block">
        <label class="color-block-label">{{ title }}</label>
        <NewThemeColorInput v-for="item in colorList" :key="item.label" :label="item.label" :color="item.color"
            @color-updated="updateColor(item.label, $event)" />
    </div>
</template>
  
<script>
import NewThemeColorInput from './NewThemeColorInput.vue';
export default {
    components: {
        NewThemeColorInput
    },
    props: {
        type: String,
        title: String,
        colorList: {
            type: Array,
            default: () => []
        }
    },
    methods: {
        updateColor(label, newColor) {
            const item = this.colorList.find(item => item.label === label);
            if (item) { item.color = newColor; }
        },
        getColors() {
            let colors = {};
            this.colorList.forEach(p => {
                colors[`${this.type}${p.label}Color`] = p.color;
            });
            return colors;
        },
    }

};
</script>

  
<style scoped>
.color-block-label {
    width: 100%;
    text-align: center;
    color: #121212;
    font-family: 'Figtree';
    font-size: calc(0.5vh + 1.1vw + 8px);
}

.color-inputs-block {
    height: calc(100px + 30% + 8vw);
    width: 100%;
    background-color: #f5f5f5;
    border-radius: calc(0.3vh + 0.3vw + 2px);
    display: grid;
    grid-template-rows: auto 1fr 1fr 1fr;
    padding-top: calc(2px + 2vh);
    align-content: center;
}
</style>
  