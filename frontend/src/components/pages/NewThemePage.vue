<template>
    <h1>
        New theme:
    </h1>
    <h2>Theme name: <input type="text" /></h2>
    <div class="color-labels-container">
        <NewThemeColorInput v-for="(value, key) in themeValues" :key="key" :label="key" :color="value" ref="colorLabels" />
    </div>
    <button @click="this.$emit('goBack');" class="go-back-button">Go back</button>
</template>
<script>
import NewThemeColorInput from '../settingsPageComponents/tabsComponents/NewThemeColorInput.vue';
import { GetHexDefaultDark } from '../../../wailsjs/go/themerelated/ThemeCollection'
export default {
    data() {
        return {
            newTheme: {},
        }

    },
    computed: {
        themeValues() {
            return Object.fromEntries(
                Object.entries(this.newTheme).filter(([key]) => key.endsWith("Color"))
            );
        },
    },
    components:
    {
        NewThemeColorInput
    },
    created() {
        GetHexDefaultDark().then((theme) => {
            this.newTheme=theme;
        });
    }
};
</script>
<style>
.go-back-button {
    position: absolute;
    right: calc(8px + 1vw);
    bottom: calc(6px + 2%);
    height: calc(2vh + 22px + 0.2vw);
    width: calc(4vw + 60px + 2vh);
    background-color: var(--bright-2);
    font-family: 'Figtree';
    font-size: calc(0.55vh + 0.55vw + 10px);
    font-weight: 300;
    color: var(--front);
    border: 1px solid transparent;
    border-radius: calc(3px + 0.09vw + 0.1vh);
    cursor: pointer;
    transition: 0.08s;
}

.go-back-button:hover {
    background-color: var(--bright-3);
}
</style>