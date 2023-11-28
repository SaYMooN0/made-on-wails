<template>
    <label class="theme-label">
        {{ theme.Name }}
        <div class="right-side-container">
            <div class="colors-container">
                <div v-for="color in ThemeColorFields" class="color-div" :style="`background-color: ${color}`"></div>
            </div>
            <div class="delete-theme-button" @click="deleteTheme">
                <svg class='delete-theme-button-icon' viewBox="0 0 448 512">
                    <path
                        d="M135.2 17.7C140.6 6.8 151.7 0 163.8 0H284.2c12.1 0 23.2 6.8 28.6 17.7L320 32h96c17.7 0 32 14.3 32 32s-14.3 32-32 32H32C14.3 96 0 81.7 0 64S14.3 32 32 32h96l7.2-14.3zM32 128H416V448c0 35.3-28.7 64-64 64H96c-35.3 0-64-28.7-64-64V128zm96 64c-8.8 0-16 7.2-16 16V432c0 8.8 7.2 16 16 16s16-7.2 16-16V208c0-8.8-7.2-16-16-16zm96 0c-8.8 0-16 7.2-16 16V432c0 8.8 7.2 16 16 16s16-7.2 16-16V208c0-8.8-7.2-16-16-16zm96 0c-8.8 0-16 7.2-16 16V432c0 8.8 7.2 16 16 16s16-7.2 16-16V208c0-8.8-7.2-16-16-16z" />
                </svg>
            </div>
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
        deleteTheme(event) {
            event.stopPropagation();
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
    border-radius: calc(1px + 0.2vw);
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

.delete-theme-button {
    margin-right: 0.2vw;
    align-self: center;
    height: calc(0.2vh + 0.7vw + 13px);
    aspect-ratio: 1/1;
    display: flex;
    align-items: center;
    justify-content: center;
    padding: 0.1vw;
    border-radius: 10%;
}

.delete-theme-button-icon {
    height: 80%;
    aspect-ratio: 1/1;
    fill: var(--back-2);

}
.theme-label:hover svg path{
    fill: var(--front);
    transition: 0.12s;
    
}
.delete-theme-button:hover path {
    fill: var(--warning-main) !important;
    height: 82%;
}
</style>
    