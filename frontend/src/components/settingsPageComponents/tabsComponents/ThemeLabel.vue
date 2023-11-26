<template>
    <label class="theme-label">
        {{ theme.Name }}
        <div class="right-side-container">
            <div class="colors-container">
                <div v-for="color in ThemeColorFields" class="color-div" :style="`background-color: ${color}`"></div>
            </div>
            <button class="delete-theme-button" @click="deleteTheme">del</button>
        </div>
    </label>
</template>
    
<script>
export default {
    props: {
        theme: Object,
    },
    methods:
    {
        deleteTheme() {
            console.log(this.theme);
        },
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
    font-size: calc(0.5vh + 0.5vw + 10px);
    font-weight: 600;
    font-family: 'Figtree';
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding: calc(2px + 0.2vw + 0.2vh);
    color: var(--front);
    background-color: var(--back-2);
    border-radius: 2px;
    margin-top: calc(5px + 0.5vh);
    transition: 0.1s all;
    cursor: pointer;
}

.theme-label:hover {
    background-color: var(--back-3);
    color: var(--front-3);
}

.colors-container {
    display: flex;
    justify-content: space-between;
    height: calc(0.15vh + 0.7vw + 7px);
    max-width: calc(80px + 25vw);
    background-color: white;
    align-items: center;
    padding: calc(2px + 0.1vw + 0.1vh);
    gap: calc(2px + 0.25vw);
    border-radius: 2px;
}

.right-side-container {
    display: flex;
    flex-direction: row;
    gap: calc(3% - 4px);
}

.color-div {
    height: 100%;
    aspect-ratio: 1/1;
    border-radius: 10%;
}
</style>
    