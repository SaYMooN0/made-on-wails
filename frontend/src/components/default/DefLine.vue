<template>
    <p class="line">
        <label class="label-text">
            {{labelText}}
        </label>
        <slot></slot>
    </p>
</template>

<script>
export default {
    props: {
        labelText: {
            type: String,
            required: true
        }
    }
}
</script>
<style scoped>
.line
{
    width: calc(100px + 64%);
    display: flex;
    align-items: center;
    justify-content: space-between;
    
}
.label-text {
    color: var(--front);
    font-family: 'Figtree';
    font-weight: 600;
    font-size: calc(0.9vh + 0.71vw + 8px);
    padding-right: calc(1px + 0.25vw + 0.25vh);
}
</style>
