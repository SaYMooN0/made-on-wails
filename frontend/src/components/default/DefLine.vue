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
            required: false
        }
    }
}
</script>
<style scoped>
.line
{
    width: calc(80px + 69%);
    display: flex;
    align-items: center;
    justify-content: space-between;
    
}
.label-text {
    color: var(--front);
    font-family: 'Figtree';
    font-weight: 600;
    font-size: calc(0.9vh + 0.71vw + 8px);
}
</style>
