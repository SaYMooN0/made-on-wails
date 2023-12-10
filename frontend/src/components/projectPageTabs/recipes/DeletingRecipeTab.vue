<template>
    <label class="header-label">Delete recipes for </label>
    <div class="filters-container">
        <DefLine labelText="item:">
            <InputWithSuggestions :value="initialItem" @updateValue="this.itemValue = $event" suggestion-type="item" />
        </DefLine>
        <DefLine>
            <div class="checkbox-container">
                <DefCheckBox labelText="as input" v-model="fromValue" class="non-selectable" />
                <DefCheckBox labelText="as output" v-model="toValue" class="non-selectable" />
            </div>
        </DefLine>
        <DefLine labelText="types:" />
        <div class="types-container scrollbar">
            <DefLine v-for="(typeVal, index) in types" :key="`type-${types.length}`" :labelText="`${index + 1})`"
                class="non-selectable">
                <div style="display: flex; flex-direction: row; gap: 10px; align-items: center;">
                    <InputWithSuggestions :value="typeVal" @updateValue="updateTypeValue($event, index)"
                        suggestion-type="type" ref="inputComponents" />
                    <svg class='delete-type-icon' viewBox='0 0 24 24' fill='none' @click="deleteType(index)">
                        <path stroke='#1C274C' d='M20.5001 6H3.5' stroke-width='1.5' stroke-linecap='round' />
                        <path
                            d='M18.8332 8.5L18.3732 15.3991C18.1962 18.054 18.1077 19.3815 17.2427 20.1907C16.3777 21 15.0473 21 12.3865 21H11.6132C8.95235 21 7.62195 21 6.75694 20.1907C5.89194 19.3815 5.80344 18.054 5.62644 15.3991L5.1665 8.5'
                            stroke='#1C274C' stroke-width='1.5' stroke-linecap='round' />
                        <path d='M9.5 11L10 16' stroke='#1C274C' stroke-width='1.5' stroke-linecap='round' />
                        <path d='M14.5 11L14 16' stroke='#1C274C' stroke-width='1.5' stroke-linecap='round' />
                        <path
                            d='M6.5 6C6.55588 6 6.58382 6 6.60915 5.99936C7.43259 5.97849 8.15902 5.45491 8.43922 4.68032C8.44784 4.65649 8.45667 4.62999 8.47434 4.57697L8.57143 4.28571C8.65431 4.03708 8.69575 3.91276 8.75071 3.8072C8.97001 3.38607 9.37574 3.09364 9.84461 3.01877C9.96213 3 10.0932 3 10.3553 3H13.6447C13.9068 3 14.0379 3 14.1554 3.01877C14.6243 3.09364 15.03 3.38607 15.2493 3.8072C15.3043 3.91276 15.3457 4.03708 15.4286 4.28571L15.5257 4.57697C15.5433 4.62992 15.5522 4.65651 15.5608 4.68032C15.841 5.45491 16.5674 5.97849 17.3909 5.99936C17.4162 6 17.4441 6 17.5 6'
                            stroke-width='1.5' />
                    </svg>

                </div>

            </DefLine>
        </div>
        <div @click="addNewType" class="add-new-type-button">Add New Type</div>
    </div>
    <DefSave @click="submit" :submitText="submitText" />
</template>
<script>
import DefCheckBox from '../../default/DefCheckBox.vue';
import InputWithSuggestions from '../../default/InputWithSuggestions.vue';
import DefSave from '../../default/DefSave.vue';
import DefLine from '../../default/DefLine.vue';
export default {
    components:
    {
        DefCheckBox,
        InputWithSuggestions,
        DefSave,
        DefLine
    },
    props:
    {
        from: {
            type: Boolean,
            default: false
        },
        to: {
            type: Boolean,
            default: false
        },
        initialItem:
        {
            type: String,
            default: ''
        },
        initialTypes:
        {
            type: Object,
            default: () => [],
        },
        isNew: {
            type: Boolean,
            default: true
        },
        filePath: {
            type: String,
            default: ''
        },
        actionId: {
            type: Number,
            default: 0
        },
    },
    data() {
        return {
            fromValue: this.from,
            toValue: this.from,
            itemValue: this.initialItem,
            types: this.initTypes()
        }
    },
    methods: {
        updateTypeValue(newValue, index) { 
            this.types[index] = newValue;
        },
        setFocusToInput(index) {
            if (this.$refs.inputComponents && this.$refs.inputComponents[index]) {
                this.$refs.inputComponents[index].focusInput();
            }
        },
        initTypes() {
            return this.initialTypes.flat();
        },
        addNewType() {
            this.types.push('');
        },
        deleteType(index) {
            this.types.splice(index, 1);
        },
        submit() {
            let emptyTypeIndexes = [];

            this.types.forEach((type, index) => {
                console.log(type, index);
                if (!type || type.trim() === '') {
                    emptyTypeIndexes.push(index + 1);
                }
            });

            if (emptyTypeIndexes.length > 0) {
                this.showNotification(`There are empty values in types ${emptyTypeIndexes.join(', ')}. Delete or fill them`, 1);
                return;
            }
            if ((!this.itemValue || this.itemValue.trim() === '') && (this.fromValue || this.toValue)) {
                this.showNotification("The 'item' field could not be empty if 'input' or 'output' is selected", 1);
                return;
            }
        }
    },
    computed: {
        submitText() { return this.isNew ? 'Save' : 'Save changes'; }
    },
    inject: ["showNotification"],

}
</script>
<style scoped>
.header-label {
    position: relative;
    top: 3vh;
    margin-left: calc(10px + 2vw);

    color: var(--front);
    font-family: 'Figtree';
    font-size: calc(1.4vh + 1.4vw + 11px);
}

.filters-container {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    padding-top: calc(10vh - 20px);
}

.checkbox-container {
    width: 100%;
    gap: calc(40px + 5%);
    display: flex;
    flex-direction: row;
}

.types-container {
    width: 60%;
    max-height: calc(66vh - 160px);
    overflow-y: auto;
}

.delete-type-icon {
    height: calc(17px + 1vh + 0.4vw);
    padding: 2px;
    aspect-ratio: 1/1;
    border: 1px solid transparent;
    border-radius: 30%;
}

.delete-type-icon path {
    stroke: var(--front);
}

.delete-type-icon:hover {
    background-color: var(--back-3);
}

.delete-type-icon:hover path {
    stroke: var(--front-2);
}

.add-new-type-button {
    justify-self: center;
    background-color: var(--bright-2);
    border: 1px solid transparent;
    border-radius: calc(3px + 0.09vw + 0.1vh);
    font-family: 'Figtree';
    font-size: calc(0.5vh + 0.5vw + 10px);
    font-weight: 300;
    color: var(--front);
    padding: 9px;
    cursor: pointer;
    transition: 0.06s;

}
</style>