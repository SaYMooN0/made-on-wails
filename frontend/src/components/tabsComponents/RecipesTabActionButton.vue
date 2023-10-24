<template>
    <div class="action-button" @click="handleClick">
        <label class="action-button-main-label">{{ mainLabel }}</label>
        <label class="action-button-hint-label">{{ hintLabel }}</label>
    </div>
</template>

<script>
export default {
    props: {
        mainLabel: String,
        hintLabel: String,
        actionOnClick: Function
    },
    methods: {
        handleClick(event) {
            if (this.actionOnClick) {
                this.actionOnClick(event);
            }
        }
    }
}
</script>
<style scoped>
.action-button {
    box-sizing: border-box;
    height: calc(5vh + 2.1vw + 60px);
    width: calc(10vw + 2vh + 100px);
    display: grid;
    grid-template-rows: 4fr 7fr;
    border: calc(0.12vw + 0.12vh) solid var(--front);
    border-radius: calc(0.25vw + 0.25vh + 8px);
    padding: 1vh;
    cursor: pointer;
    transition: 0.1s;
}

.action-button:hover {
    transform: scale(1.03);
    background-color: var(--back-2);
}

.action-button-main-label {
    color: var(--front-3);
    font-family: 'Figtree';
    font-size: calc(0.7vh + 0.82vw + 6px);
    text-align: center;
    display: flex;
    align-items: center;
    justify-content: center;
    cursor: pointer;
}

.action-button-hint-label {
    color: var(--front);
    font-family: 'Bahnschrift';
    font-size: calc(0.5vh + 0.5vw + 8px);
    text-align: center;
    cursor: pointer;
}
</style>