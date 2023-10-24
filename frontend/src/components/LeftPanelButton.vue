<template>
    <div class="left-panel-btn non-selectable" @click="leftPanelButtonClick()">
        <slot></slot>
        <label class="left-panel-label">{{ buttonLabel }}</label>
    </div>
</template>
    
<script>
export default {
    props: ['buttonLabel'],
    methods: {
        leftPanelButtonClick() {
            this.$emit('createTab');
        }
    }
}
</script>
<style>
.left-panel-btn {
    margin-top: calc(0.1vw);
    width: 94%;
    height: calc(32px + 2.2vh);
    display: flex;
    cursor: pointer;
    margin-left: 2.5%;
    border: 1px solid transparent;
}

.left-panel-btn:hover {
    background-color: var(--back-3);
    border-radius: calc(0.2vw + 0.2vh + 2px);
}

.left-panel-btn:hover .left-panel-label {
    color: var(--front-3);
}

.left-panel-btn:hover .left-panel-btn-icon-fill path {
    fill: var(--front-3);
}

.left-panel-btn:hover .left-panel-btn-icon-stroke path {
    stroke: var(--front-3);
}

.left-panel-btn-icon {
    margin-left: calc(1px + 0.2vw + 0.1vh);
    height: calc(90% - 7px);
    align-self: center;
    aspect-ratio: 1/1;
}

.left-panel-btn-icon-big {
    height: calc(72% - 7px);
    margin-left: calc(1px + 0.2vw + 0.1vh);
    padding-right: calc(2% + 5px);
    align-self: center;
    aspect-ratio: 1/1;
}

.left-panel-btn-icon-fill path {
    fill: var(--front);
}

.left-panel-btn-icon-stroke path {
    stroke: var(--front);
}


.left-panel-label {
    font-family: 'Bahnschrift';
    font-size: calc(1vh + 0.5vw + 7px);
    align-self: center;
    font-weight: 400;
    color: var(--front);
    margin-left: calc(3px + 0.2vw + 0.1vh);
    padding-top: calc(0.24vh + 1px);
    cursor: pointer;
}
</style>