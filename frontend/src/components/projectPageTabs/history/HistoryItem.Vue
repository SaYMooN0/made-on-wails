<template>
    <div class='history-item' @click="openHistoryItem">
        <label class='history-item-type'>{{ actionTypeToString(historyItem.ActionType) }}</label>
        <svg class='history-item-delete-button' @click.stop="deleteHistoryItem" viewBox='0 0 24 24' fill='none'>
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
        <label class='history-item-id'>Id: {{ historyItem.ActionID }}</label>
        <label class='history-item-input-output'>{{ historyItemLabel(historyItem) }}</label>
    </div>
</template>
  
<script>
import { actionTypeToString, historyItemLabel } from "../../../HistoryItemManager";
export default {
    props: {
        historyItem: {
            type: Object,
            required: true
        }
    },
    methods: {
        actionTypeToString,
        historyItemLabel,
        openHistoryItem() {
            this.$emit('open', this.historyItem);
        },
        deleteHistoryItem(event) {
            this.$emit('delete', this.historyItem);
        }
    }
};
</script>
<style scoped>
.history-item {
    min-height: 52px;
    height: calc(25px + 5.5vh + 0.9vw);
    width: 97%;
    background-color: var(--back);
    border-radius: calc(0.14vw + 0.12vh + 3px);
    cursor: pointer;
    display: grid;
    grid-template-rows: 1.2fr 1fr;
    grid-template-columns: 8fr 4fr calc(10px + 3vh + 2vw);
    grid-template-areas:
        "type date dlt"
        "text text text";
    padding-top: 1%;
    padding-left: 2%;
}


.history-item:hover {
    background-color: var(--back-2);
}


.history-item-id {
    color: var(--front-2);
    grid-area: date;
    font-family: 'Bahnschrift';
    cursor: inherit;
    justify-content: center;
}

.history-item-type {
    color: var(--front-3);
    grid-area: type;
    font-family: 'Figtree';
    cursor: inherit;
}

.history-item-delete-button {
    grid-area: dlt;
    height: calc(15px + 50% + 0.4vw);
    padding: 2px;
    aspect-ratio: 1/1;
    border: 1px solid transparent;
    border-radius: 30%;
}

.history-item-delete-button path {
    stroke: var(--front);
}

.history-item-delete-button:hover {
    background-color: var(--back-3);
}

.history-item-delete-button:hover path {
    stroke: var(--front-2);
}

.history-item-input-output {
    color: var(--front);
    font-family: 'Bahnschrift';
    cursor: inherit;
}
</style>