<template>
    <input type="submit" :value="submitText" class="default-save" />
</template>

<script>
export default {
    props: {
        submitText: {
            type: String,
            required: true
        }
    }
}
</script>
<style scoped>
.default-save {
    position: absolute;
    right: calc(8px + 3vw);
    bottom: calc(6px + 2%);
    height: calc(2vh + 24px + 0.2vw);
    width: calc(4.2vw + 88px + 2.3vh);
    background-color: var(--bright-2);
    border: 1px solid transparent;
    border-radius: calc(3px + 0.09vw + 0.1vh);
    font-family: 'Figtree';
    font-size: calc(0.55vh + 0.55vw + 11px);
    font-weight: 300;
    color: var(--front);
    cursor: pointer;
    transition: 0.08s all;
}

.default-save:hover {
    background-color: var(--bright-3);
}
</style>
