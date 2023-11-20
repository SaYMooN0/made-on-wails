<template>
    <p  class="theme-label">
        {{ theme.Name }}
        <div class="colors-container">
            <div v-for="color in ThemeColorFields" class="color-div" :style="`background-color: ${color}`"></div>
        </div>
    </p>
</template>
    
<script>
export default {
    props: {
        theme: Object,
    },
    computed: {
        ThemeColorFields() {
            return Object.fromEntries(
                Object.entries(this.theme).filter(([key]) => key.endsWith("Color"))
            );
        },
    },
};
</script>
    
<style scoped>
.theme-label {
    color: black;
    font-size: calc(0.5vh + 0.3vw + 11px);
    font-weight: 600;
    font-family: 'Figtree';
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding-right: 1%;
}

.colors-container {
    display: flex;
    justify-content: space-between;
    height: calc(0.3vh + 0.35vw + 11px);
    max-width: calc(80px + 25vw );
    background-color: white;
    align-items: center;
    padding: 1%;
    gap: 5px;
    

}

.color-div {
    height: 100%;
    aspect-ratio: 1/1;
    border-radius: 10%;
}
</style>
    