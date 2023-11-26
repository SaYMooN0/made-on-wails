<template>
    <h1>
        New theme:
    </h1>
    <h2>Theme name: <input type="text" /></h2>
    <div class="color-labels-container">
        <NewThemeColorInput v-for="(value, key) in themeValues" :key="key" :label="key" :color="value" ref="colorLabels" />
    </div>
</template>
<script>
import NewThemeColorInput from '../tabsComponents/NewThemeColorInput.vue';
export default {
    props: {
        newTheme: {},
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
    }
};
</script>