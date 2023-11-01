<template>
    <label class="default-radio-container">
    <input class="default-radio" type="radio" :name="name" :value="value" :checked="isChecked"> 
        <span class="default-radio-label">{{ spanText }}</span>
    </label>
</template>

<script>
export default {
    props: {
        name: {
            type: String,
            required: true
        },
        value: {
            type: String,
            required: true
        },
        isChecked: {
            type: Boolean,
            default: false
        },
        spanText: {
            type: String,
            default: ''
        }
    }
}
</script>

<style scoped>
.default-radio {
    display: none;
}

.default-radio-container {
    padding-left: calc(24px + 0.55vw + 0.55vh);
    position: relative;
    cursor: pointer;
    width:100%;
    display: flex;
    align-items: center;
}

.default-radio-label::before {
    content: "";
    position: absolute;
    left: 0;
    top: 50%;
    transform: translateY(-50%);
    width: calc(10px + 0.5vw + 0.5vh);
    height: calc(10px + 0.5vw + 0.5vh);
    border: calc(2px + 0.05vw + 0.05vh) solid transparent;
    border-radius: 50%;
    background-color: var(--back-3);
}
.default-radio-label {
    color: var(--front);
    font-family: 'Figtree';
    font-weight: 600;
    font-size: calc(1vw + 7px);
    white-space: nowrap;
}

.default-radio-container:hover .default-radio-label::before {
    border-color: var(--front-2);
}

.default-radio-container:hover .default-radio-label {
    color: var(--front-2);
}

.default-radio:checked + .default-radio-label {
    color: var(--front-3);
}

    .default-radio:checked + .default-radio-label::before {
        border-color: var(--front-3);
    }

</style>