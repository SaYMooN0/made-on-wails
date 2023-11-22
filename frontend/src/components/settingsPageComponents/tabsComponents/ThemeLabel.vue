<template>
    <label class="theme-label">
        {{ theme.Name }}
        <div class="right-side-container">
            <div class="colors-container">
                <div v-for="color in ThemeColorFields" class="color-div" :style="`background-color: ${color}`"></div>
            </div>
            <button class="edit-theme-button">edt</button>
            <button class="delete-theme-button">del</button>
        </div>
    </label>
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
.theme-label:hover{
    background-color: var(--back-3);
    color: var(--front-3);
}
.colors-container {
    display: flex;
    justify-content: space-between;
    height: calc(0.3vh + 0.5vw + 8px);
    max-width: calc(80px + 25vw);
    background-color: white;
    align-items: center;
    padding: calc(2px + 0.1vw + 0.1vh);
    gap: 5px;
    border-radius: 2px;
}

.right-side-container {
    display: flex;
    flex-direction: row;
}

@media (max-width: 920px) {
    .right-side-container .delete-theme-button {
        display: none;
    }

    .right-side-container .edit-theme-button {
        display: none;
    }
}

.color-div {
    height: 100%;
    aspect-ratio: 1/1;
    border-radius: 10%;
}
</style>
    